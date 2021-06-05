package repository

import (
	"database/sql"
	"eventsmanagerpersistence/internal/entity"
	"log"

	_ "github.com/lib/pq"
)

type countryRepository struct {
}

func NewCountryRepository() entity.CountryRepository {
	return &countryRepository{}
}

func (*countryRepository) GetAll() *[]entity.Country {
	connStr := "user=postgres password=postgres dbname=EventsManager sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query("SELECT country_id as Id, code , name FROM public.country")
	if err != nil {
		log.Fatal(err)
	}

	result := make([]entity.Country, 10)
	for rows.Next() {

		var country entity.Country

		err = rows.Scan(&country.Id, &country.Code, &country.Name)
		if err != nil {
			log.Fatal(err)
		}
		result = append(result, country)
	}

	return &result
}
