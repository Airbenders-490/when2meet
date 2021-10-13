package main

import (
	"context"
	"log"

	"github.com/airbenders/when2meet/domain"
	"github.com/airbenders/when2meet/repository"
	"github.com/jackc/pgx/v4/pgxpool"
)

func main() {

	connectionstr := "user=postgres dbname=when2meet password=Team7SOEN363 host=local sslmode=disable"
	pool, err := pgxpool.Connect(context.Background(), connectionstr)
	if err != nil {
		log.Fatalln(err)
	}
	repository := repository.NewWhen2meetRepository(pool)
	when2meet := &domain.When2meet{
		W2M_ID:   "1994",
		W2M_name: "meetings",
	}
	repository.Create(context.Background(), when2meet.W2M_ID, when2meet)

}
