package main

import (
	"github.com/go-pg/pg"
	_ "github.com/lib/pq"
	"github.com/thitemalhar/GoLang/Postgres-DB/DB"
	"time"
)

func main() {
	//DB.Connect()
	pg_db := DB.Connect()
	SaveProduct(pg_db)

}

func SaveProduct(dbRef *pg.DB) {
	newPI := DB.ProductItem{
		ID:    1,
		Name:  "Product 1",
		Desc:  "Product 1 description",
		Image: "Product 1 image",
		Price: 64.32,
		Feature: struct {
			Name       string
			Desc       string
			Importance int
		}{
			Name:       "Feature1",
			Desc:       "Feature 1 description",
			Importance: 1,
		},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		IsActive:  false,
	}
	newPI.Save(dbRef)
}
