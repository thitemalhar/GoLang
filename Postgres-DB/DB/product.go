package DB

import (
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
	"log"
	"time"
)

type ProductItem struct {
	tableName struct{} `sql:"product_item_collection"`
	ID        int      `sql:"id,pk"`
	Name      string   `sql:"name,unique"`
	Desc      string   `sql:"desc"`
	Image     string   `sql:"image"`
	Price     float64  `sql:"price,type:real"`
	Feature   struct {
		Name       string
		Desc       string
		Importance int
	} `sql:"feature"`
	CreatedAt time.Time `sql:"created_at"`
	UpdatedAt time.Time `sql:"updated_at"`
	IsActive  bool      `sql:"is_active"`
}

func CreateProdItemsTable(db *pg.DB) error {
	opts := &orm.CreateTableOptions{
		IfNotExists:   true,
	}

	createErr := db.CreateTable(&ProductItem{}, opts)
	if createErr != nil {
		log.Printf("Error while createing the table: %v\n", createErr)
		return createErr
	}
	log.Printf("Table ProductItems created successfully")
	return nil
}

func (pi *ProductItem) Save(db *pg.DB) error {
	insertErr := db.Insert(pi)
	if insertErr != nil {
		log.Printf("Error while inserting the table : %v \n", insertErr)
		return insertErr
	}
	log.Printf("Table %s inserted successfully!!", pi.Name)
	return nil
}