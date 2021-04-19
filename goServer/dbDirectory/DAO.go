package dbDirectory

import (
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	"log"
)

type UserDaoImpl struct {
	tx DaoTxImpl
}

type DaoTxImpl struct {
	name string
}

func (daoTx *DaoTxImpl) Query(model interface{}, query interface{}, params ...interface{}) (res orm.Result, err error) {
	err = Db.RunInTransaction(func(tx *pg.Tx) error {
		_,err = tx.Exec("SET LOCAL search_path TO ?", daoTx.name)
		if err != nil{
			log.Printf("dbDirectory.DAO.GetUserDao: set local + %v\\n", err)
			return err
		}

		_,err = tx.Query(model,query,params...)
		if err != nil{

			log.Printf("dbDirectory.DAO.GetUserDao: Query + %v\\n", err)
			return err
		}

		return nil
	})
return
}

func GetUserDao(name string) UserDaoImpl {
	var tx =  DaoTxImpl{name}
	return UserDaoImpl{tx}
}

