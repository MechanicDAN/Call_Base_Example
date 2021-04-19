package dbDirectory

import (
	"crypto/sha1"
	"encoding/hex"
	"github.com/go-pg/pg/v10"
	"log"
	"test/structures"
)

func (dao UserDaoImpl) GetAllUsersFromDb(params Params) (users []structures.User,err error) {
	if params.SortBy == ""{
		params.SortBy = "Id"
	}
	if params.SortBy != "Id"{
		params.SortBy = params.SortBy + "_p"
	}

	_ ,err = dao.tx.Query(&users,
		"SELECT * FROM user_t WHERE login_p LIKE ? AND name_p LIKE ? AND phone_p LIKE ? ORDER BY ? ?",
		"%" + params.Login + "%", "%" + params. Name + "%", "%" + params.Phone + "%", pg.SafeQuery(params.SortBy), pg.SafeQuery(params.SortType))
	if err != nil {
		log.Printf("dbDirectory.dbServerCommand.GetAllFromDb: +%v\n", err)
		return nil,err
	}
	return users,nil
}

func (dao UserDaoImpl) GetUserFromDbById(id string) (users structures.User,err error) {
	_ ,err = dao.tx.Query(&users,"SELECT * FROM user_t WHERE id = ?", id)
	if err != nil {
		log.Printf("dbDirectory.dbServerCommand.GetAllFromDb: +%v\n", err)
		return structures.User{} ,err
	}
	return users,nil
}

func (dao UserDaoImpl) GetAuthorizationUsers(login string,password string) (users []structures.User,err error) {
	hasher := sha1.New()
	hasher.Write([]byte(password))

	_ ,err = dao.tx.Query(&users,"SELECT * FROM user_t WHERE login_p = ? AND password_p = ?",
		login ,hex.EncodeToString(hasher.Sum(nil)))

	if err != nil {
		log.Printf("dbDirectory.dbServerCommand.GetAuthorizationUsers: +%v\n", err)
		return nil,err
	}
	return users,err
}

func (dao UserDaoImpl) CreateUser(params Params) (err error){
	user := structures.User{Name: params.Name, Phone:params.Phone ,Login:params.Login ,Language:params.Language}

	_ ,err = dao.tx.Query(&user,"INSERT INTO user_t (name_p, login_p, phone_p, language_p, password_p) VALUES (?, ?, ?, ?, ?)",
		params.Name, params.Login, params.Phone, params.Language, "356a192b7913b04c54574d18c28d46e6395428ab")
	if err != nil {
		log.Printf("dbDirectory.dbServerCommand.CreateUser: +%v\n", err)
		return err
	}
	return nil
}

func (dao UserDaoImpl) UpdateUser(params Params) (err error){
	user := structures.User{Name: params.Name, Phone:params.Phone ,Login:params.Login ,Language:params.Language}
	_ ,err = dao.tx.Query(&user,"UPDATE user_t SET name_p = ?, login_p = ?, phone_p = ?, language_p = ? WHERE id = ?" ,
		params.Name, params.Login, params.Phone, params.Language, params.Id)
	if err != nil {
		log.Printf("dbDirectory.dbServerCommand.CreateUser: +%v\n", err)
		return err
	}
	return nil
}

func (dao UserDaoImpl) DeleteUser(params Params) (err error){
	user := structures.User{Name: params.Name, Phone:params.Phone ,Login:params.Login ,Language:params.Language}
	_ ,err = dao.tx.Query(&user,"DELETE FROM user_t WHERE id = ?" ,
		params.Id)
	if err != nil {
		log.Printf("dbDirectory.dbServerCommand.CreateUser: +%v\n", err)
		return err
	}
	return nil
}