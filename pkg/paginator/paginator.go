package paginator

import (
	"fmt"
	"gohub/pkg/config"
	"gohub/pkg/logger"
	"math"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Paging struct {
	CurrentPage int
	PerPage     int
	TotalPage   int
	TotalCount  int64
	NextPageUrl string
	PrePageUrl  string
}

type Paginator struct {
	BaseUrl    string
	PerPage    int
	Page       int
	Offset     int
	TotalPage  int
	TotalCount int64
	Sort       string
	Order      string

	query *gorm.DB
	ctx   *gin.Context
}

func (p Paginator) getPageLink(page int) string {
	return fmt.Sprintf("%v%v&%s=%s&%s=%s&%s=%v",
		p.BaseUrl,
		page,
		config.Get("paging.url_query_sort"),
		p.Sort,
		config.Get("paging.url_query_order"),
		p.Order,
		config.Get("paging.url_query_per_page"),
		p.PerPage,
	)
}

func (p *Paginator) getNextPageUrl() string {
	if p.TotalPage > p.Page {
		return p.getPageLink(p.Page + 1)
	}
	return ""
}

func (p *Paginator) getPrePageUrl() string {
	if p.Page <= 1 || p.Page > p.TotalPage {
		return ""
	}
	return p.getPageLink(p.Page - 1)
}

func (p *Paginator) formatBaseUrl(baseUrl string) string {
	if strings.Contains(baseUrl, "?") {
		baseUrl = baseUrl + "&" + config.Get("paging.url_query_page") + "="
	} else {
		baseUrl = baseUrl + "?" + config.Get("paging.url_query_page") + "="
	}
	return baseUrl
}

func (p *Paginator) getTotalPage() int {
	if p.TotalCount == 0 {
		return 0
	}
	nums := int64(math.Ceil(float64(p.TotalCount) / float64(p.PerPage)))
	if nums == 0 {
		nums = 1
	}
	return int(nums)
}

func (p *Paginator) getTotalCount() int64 {
	var count int64
	if err := p.query.Count(&count).Error; err != nil {
		return 0
	}
	return count
}

func (p *Paginator) getCurrentPage() int {
	page := cast.ToInt(p.ctx.Query(config.Get("paging.url_query_page")))
	if page < 0 {
		page = 1
	}
	if p.TotalPage == 0 {
		return 0
	}
	if page > p.TotalPage {
		return p.TotalPage
	}
	return page
}

func (p *Paginator) getPerPage(perPage int) int {
	queryPerPage := p.ctx.Query(config.Get("paging.url_query_per_page"))
	if len(queryPerPage) > 0 {
		perPage = cast.ToInt(queryPerPage)
	}

	if perPage <= 0 {
		perPage = config.GetInt("paging.perpage")
	}

	return perPage
}

func (p *Paginator) initProperties(perPage int, baseUrl string) {
	p.BaseUrl = p.formatBaseUrl(baseUrl)
	p.PerPage = p.getPerPage(perPage)

	p.Order = p.ctx.DefaultQuery(config.Get("paging.url_query_order"), "asc")
	p.Sort = p.ctx.DefaultQuery(config.Get("paging.url_query_sort"), "id")

	p.TotalCount = p.getTotalCount()
	p.TotalPage = p.getTotalPage()
	p.Page = p.getCurrentPage()
	p.Offset = (p.Page - 1) * p.PerPage
}

func Paginate(c *gin.Context, db *gorm.DB, data interface{}, baseUrl string, perPage int) Paging {
	p := &Paginator{
		query: db,
		ctx:   c,
	}
	p.initProperties(perPage, baseUrl)

	err := p.query.Preload(clause.Associations).
		Order(p.Sort + " " + p.Order).
		Limit(p.PerPage).
		Offset(p.Offset).
		Find(data).
		Error

	if err != nil {
		logger.LogIf(err)
		return Paging{}
	}
	return Paging{
		CurrentPage: p.Page,
		PerPage:     p.PerPage,
		TotalPage:   p.TotalPage,
		TotalCount:  p.TotalCount,
		NextPageUrl: p.getNextPageUrl(),
		PrePageUrl:  p.getPrePageUrl(),
	}
}
