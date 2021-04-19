package dbDirectory

import (
	"github.com/go-pg/pg/v10"
	"log"
	"test/parsers"
)

var Db *pg.DB = nil

func ConnectToDb() {
	Db = pg.Connect(&pg.Options{
		Addr:     parsers.GetDbAddr(),
		User:     parsers.GetDbUser(),
		Password: parsers.GetDbPassword(),
		Database: parsers.GetDbDatabase(),
	})

	if err := Db.Ping(Db.Context()); err != nil {
		log.Printf("dbDirectory.dbConnection: +%v\n", err)
	}
}

func DbClose()  {
	Db.Close()
}