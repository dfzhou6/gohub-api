package seed

import "gorm.io/gorm"

type Seeder struct {
	Name string
	Func SeederFunc
}

type SeederFunc func(*gorm.DB)

var seeders []Seeder

var orderedSeederNames []string

func Add(name string, fn SeederFunc) {
	seeders = append(seeders, Seeder{
		Name: name,
		Func: fn,
	})
}

func SetRunOrder(names []string) {
	orderedSeederNames = names
}
