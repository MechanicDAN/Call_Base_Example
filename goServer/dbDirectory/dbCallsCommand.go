package dbDirectory

import (
	"github.com/go-pg/pg/v10"
	"log"
	"test/structures"
)

func (dao UserDaoImpl) GetAllCallsFromDb(params Params) (calls []structures.Call,err error) {
	if params.SortBy == "" || params.SortBy == "Id"{
		params.SortBy = "calls.Id"
	}
	if params.SortBy != "Id" && params.SortBy != "calls.Id"{
		params.SortBy = params.SortBy + "_p"
	}
	if params.User != ""{
		_ ,err = dao.tx.Query(&calls, "SELECT * FROM calls INNER JOIN user_t ON calls.user_p = user_t.id WHERE user_p = ? ORDER BY ? ?",
			pg.SafeQuery(params.User), pg.SafeQuery(params.SortBy), pg.SafeQuery(params.SortType))
	}else{
		_ ,err = dao.tx.Query(&calls, "SELECT * FROM calls INNER JOIN user_t ON calls.user_p = user_t.id ORDER BY ? ?",
			pg.SafeQuery(params.SortBy), pg.SafeQuery(params.SortType))
	}

	if err != nil {
		log.Printf("dbDirectory.dbCallCommand.GetAllFromDb: +%v\n", err)
		return nil,err
	}
	return calls,nil
}