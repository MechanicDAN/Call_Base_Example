package dbDirectory

import (
	"fmt"
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	"log"
	"test/structures"
)

func CreateSchema(name string) {
	err := Db.RunInTransaction(func(tx *pg.Tx) error {
		_, err := tx.Exec("CREATE SCHEMA IF NOT EXISTS ?", pg.SafeQuery(name))
		if err != nil {
			log.Printf("dbDirectory.dbServiceCommand.CreateSchema: cant create schema: +%v\n",err)
			return err
		}

		_,err = tx.Exec("SET LOCAL search_path TO ?", pg.SafeQuery(name))
		if err != nil {
			fmt.Printf("dbDirectory.dbServiceCommand.CreateSchema: cant set local: +%v\n", err)
			return err
		}

		err = tx.Model(&structures.User{}).CreateTable(&orm.CreateTableOptions{})
		if err != nil {
			log.Printf("dbDirectory.dbServiceCommand.CreateSchema: cant create table: +%v\n",err)
			return err
		}
		err = tx.Model(&structures.Call{}).CreateTable(&orm.CreateTableOptions{})
		if err != nil {
			log.Printf("dbDirectory.dbServiceCommand.CreateSchema: cant create table: +%v\n",err)
			return err
		}
		return nil
	})
	if err != nil {
		log.Printf("dbDirectory.dbServiceCommand.CreateSchema: Transaction faild: +%v\n",err)
	}
}

func GetVersion() (string,error) {
	var versionStr string
	_, err := Db.Query(pg.Scan(&versionStr), "SELECT version()")
	if err != nil {
		log.Printf("dbDirectory.dbServerCommand.GetVersion: +%v\n",err)
		return "",err
	}
	return versionStr,nil
}

func GetAllSchemaFromDb() (schema []structures.Schemata, err error) {
	_, err = Db.Query(&schema, "SELECT * FROM information_schema.schemata")
	if err != nil {
		log.Printf("dbDirectory.dbServerCommand.GetAllSchemaFromDb: +%v\n",err)
		return schema,err
	}
	return schema,nil
}
