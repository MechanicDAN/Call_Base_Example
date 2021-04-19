package server

import (
	"log"
	"test/dbDirectory"
	"test/structures"
)

func getUsers(api Api) (_ interface{},request int,err error) {
	var users []structures.User

	dao := dbDirectory.GetUserDao(api.authSchema)
	users,err = dao.GetAllUsersFromDb(
		dbDirectory.Params{
			Login: 		api.request.FormValue("filterLogin"),
			Name: 		api.request.FormValue("filterName"),
			Phone: 		api.request.FormValue("filterPhone"),
			SortBy: 	api.request.FormValue("sortBy"),
			SortType:	api.request.FormValue("sortType")})
	if err  != nil {
		log.Printf("server.command.getUsers: GetAllUsersFromDb+%v\n", err)
		return nil,0,err
	}

	if len(users) == 0{
		return nil,0,nil
	}
	return users,0,nil
}

func getUserById(api Api) (_ interface{},request int,err error) {
	var user structures.User

	requireId := api.GetUrlPathParam("id")
	dao := dbDirectory.GetUserDao(api.authSchema)
	user,err = dao.GetUserFromDbById(requireId)
	if err  != nil {
		log.Printf("server.command.getUsersById: +%v\n", err)
		return nil,0,err
	}

	if user.Id == 0{
		return nil,0,nil
	}
	return user,0,nil
}

func getCalls(api Api) (_ interface{},request int,err error)  {
	var calls []structures.Call

	dao := dbDirectory.GetUserDao(api.authSchema)
	calls,err = dao.GetAllCallsFromDb(
		dbDirectory.Params{User: 		api.request.FormValue("filterUser"),
							SortBy: 	api.request.FormValue("sortBy"),
							SortType:	api.request.FormValue("sortType")})
	if err  != nil {
		log.Printf("server.command.getCalls: %v\n", err)
		return nil,0,err
	}

	if len(calls) == 0{
		return nil,0,nil
	}
	return calls,0,nil
}

func createUser(api Api) (_ interface{},request int,err error) {
	dao := dbDirectory.GetUserDao(api.authSchema)
	err = dao.CreateUser(dbDirectory.Params{
		Login: 		api.request.FormValue("Login"),
		Name: 		api.request.FormValue("Name"),
		Phone: 		api.request.FormValue("Phone"),
		Language:   api.request.FormValue("Language")})
	if err  != nil {
		log.Printf("server.command.createUser: %v\n", err)
		return nil,0,err
	}
	return nil,0,nil
}

func updateUser(api Api) (_ interface{},request int,err error) {
	dao := dbDirectory.GetUserDao(api.authSchema)
	err = dao.UpdateUser(dbDirectory.Params{
		Id: 		api.GetUrlPathParam("id"),
		Login: 		api.request.FormValue("Login"),
		Name: 		api.request.FormValue("Name"),
		Phone: 		api.request.FormValue("Phone"),
		Language:   api.request.FormValue("Language")})
	if err  != nil {
		log.Printf("server.command.createUser: %v\n", err)
		return nil,0,err
	}
	return nil,0,nil
}

func deleteUser(api Api) (_ interface{},request int,err error) {
	dao := dbDirectory.GetUserDao(api.authSchema)
	err = dao.DeleteUser(dbDirectory.Params{
		Id: 		api.GetUrlPathParam("id")})
	if err  != nil {
		log.Printf("server.command.createUser: %v\n", err)
		return nil,0,err
	}
	return nil,0,nil
}

