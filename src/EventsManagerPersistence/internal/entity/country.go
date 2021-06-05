package entity

type Country struct {
	Id   string
	Code string
	Name string
}

type CountryRepository interface {
	GetAll() *[]Country
}
