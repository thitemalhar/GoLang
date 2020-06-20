package DB

import (
	pg "github.com/go-pg/pg"
	_ "github.com/lib/pq"
	"log"
	"os"
)

func Connect() *pg.DB {
	opts := &pg.Options{
		User:     "postgres",
		Password: "SWAdes1!",
		Addr:     "localhost:5432",
		Database: "postgres",
	}

	var db *pg.DB = pg.Connect(opts)
	if db == nil {
		log.Printf("Failed to connect to db \n")
		os.Exit(100)
	}
	log.Printf("Connection to DB successful!! \n")
	CreateProdItemsTable(db)
	//defer db.Close()
	return db

}

//const (
//	host     = "localhost"
//	port     = 5432
//	user     = "malhar.thite"
//	password = ""
//	dbname   = "testdb"
//)
//psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
//	"password=%s dbname=%s sslmode=disable",
//	host, port, user, password, dbname)
//db, err := sql.Open("postgres", psqlInfo)
//if err != nil {
//	panic(err)
//}
////CreateProdItemsTable(db)
//defer db.Close()
//
//err = db.Ping()
//if err != nil {
//	panic(err)
//}
//
//fmt.Println("Successfully connected!")
