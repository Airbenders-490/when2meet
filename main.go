package main

import (
	"context"
	"log"

	"github.com/airbenders/when2meet/domain"
	"github.com/airbenders/when2meet/repository"
	"github.com/jackc/pgx/v4/pgxpool"
)

func main() {

	connectionstr := "postgres://postgres:password@localhost/postgres"
	pool, err := pgxpool.Connect(context.Background(), connectionstr)
	if err != nil {
		log.Fatalln(err)
	}
	repository := repository.NewAvailibilitiesRepository(pool)
	availibilities := &domain.Availibilities{
		AID:   "1994",
		SID:   "meetings",
		ADAY:  "Monday",
		Times: [2]int{10, 12},
	}
	repository.Create(context.Background(), availibilities.AID, availibilities)

}
