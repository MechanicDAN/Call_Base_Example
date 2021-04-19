package server

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/jwtauth"
	"log"
	"net/http"
	"strconv"
	"strings"
	"test/dbDirectory"
	"test/structures"
	"time"
)

func GetRouter() (r *chi.Mux) {
	r = chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))
	r.Use(middleware.Compress(5))

	r.Handle("/*", http.FileServer(http.Dir("server/static")))


	r.Route("/login", func(r chi.Router) {
		r.Post("/auth", getAuthorizationCookie)
	})

	r.Group(func(r chi.Router) {
		r.Use(jwtauth.Verify(tokenAuth, findToken))
		r.Use(myAuthenticator)

			r.Route("/pg", func(r chi.Router) {
				r.Get("/version", getVersion)
				r.Get("/schemata", getSchemata)
			})
			r.Route("/api", func(r chi.Router) {
				r.Get("/checkAuth",checkAuth)
				r.Get("/users", createHandler(getUsers))
				r.Get("/calls",createHandler(getCalls))
				r.Get("/users/{id}", createHandler(getUserById))
				r.Post("/users/{id}/update", createHandler(updateUser))
				r.Post("/users/{id}/delete", createHandler(deleteUser))
				r.Post("/users/create",createHandler(createUser))
			})
	})
	r.Group(func(r chi.Router) {
		r.Get("/login", func(w http.ResponseWriter, r *http.Request) {
			http.ServeFile(w, r, "server/static/index.html")
		})
		r.Get("/authorized", func(w http.ResponseWriter, r *http.Request) {
			http.ServeFile(w, r, "server/static/index.html")
		})
		r.Get("/users", func(w http.ResponseWriter, r *http.Request) {
			http.ServeFile(w, r, "server/static/index.html")
		})
		r.Get("/users/edit", func(w http.ResponseWriter, r *http.Request) {
			http.ServeFile(w, r, "server/static/index.html")
		})
		r.Get("/calls", func(w http.ResponseWriter, r *http.Request) {
			http.ServeFile(w, r, "server/static/index.html")
		})
	})
	return r
}

func createHandler(fn func(Api)(interface{}, int, error)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		_, claims, _ := jwtauth.FromContext(r.Context())
		authId ,err := strconv.Atoi(fmt.Sprintf("%v",claims["id"]))
		if err != nil{
			w.WriteHeader(http.StatusInternalServerError)
			log.Printf("server.command.createHandler: strconv+%v\n", err)
			return
		}
		if fmt.Sprintf("%v",claims["schemaName"]) == ""{
			log.Printf("server.router.handler: wrong schema name")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		api := Api{	authId: authId,
					authSchema: fmt.Sprintf("%v",claims["schemaName"]),
					urlPath: 	r.URL.Path,
					request:     r,
					}

		users,request,err := fn(api)
		if err != nil{
			w.WriteHeader(http.StatusInternalServerError)
			log.Printf("server.command.createHandler: fn+%v\n", err)
			return
		}

		if users != nil{
			data, err := json.Marshal(users)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				log.Printf("server.command.createHandler: cant parse json+%v\n", err)
				return
			}
			_,err =	w.Write(data)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				log.Printf("server.command.createHandler: print fail: +%v\n", err)
				return
			}
		}else{
			_,err =	w.Write([]byte("[]"))
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				log.Printf("server.command.createHandler: print fail: +%v\n", err)
				return
			}
		}
		if request != 0{
			w.WriteHeader(request)
		}else {
			w.WriteHeader(http.StatusOK)
		}
	}
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	http.FileServer(http.Dir("server/static"))
}

func getVersion(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	str,err := dbDirectory.GetVersion()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("server.command.getVersion: connection lost: +%v\n", err)
		return
	}else {
		_,err = w.Write([]byte(str))
		if err != nil {
			log.Printf("server.command.getVersion: write fail: +%v\n", err)
			return
		}
	}
}

func getSchemata(w http.ResponseWriter, r *http.Request)  {
	schemata,err := dbDirectory.GetAllSchemaFromDb()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("server.command.getShemata: print fail: +%v\n", err)
		return
	}
	for i := range schemata {
		_,err = w.Write([]byte(schemata[i].SchemaName + " "))
		if err != nil {
			log.Printf("server.command.getSchemata: write fail: +%v\n", err)
			return
		}
	}
}

func checkAuth(w http.ResponseWriter, r *http.Request)  {
	w.WriteHeader(http.StatusOK)
}

func getAuthorizationCookie(w http.ResponseWriter, r *http.Request) {
	var count, id int
	var schemaName string
	user := structures.User{}
	schemata,err := dbDirectory.GetAllSchemaFromDb()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Printf("server.command.authorization: print fail: +%v\n", err)
		return
	}

	for i := range schemata {
		if !strings.HasPrefix(schemata[i].SchemaName,"pg_") && schemata[i].SchemaName != "information_schema"{
			dao := dbDirectory.GetUserDao(schemata[i].SchemaName)
			users,_ := dao.GetAuthorizationUsers(r.FormValue("login"),r.FormValue("password"))

			if users != nil{
				count =+ len(users)
			}
			if count > 1{
				w.WriteHeader(http.StatusBadRequest)
				log.Printf("server.command.authorization: wrong data:" + r.FormValue("login") +":"+ r.FormValue("password"))
				return
			}
			if len(users) == 1{
				id = users[0].Id
				schemaName = schemata[i].SchemaName
				user = users[0]
			}
		}
	}
	if count == 1 {
		cookie, err := createJwtCookie(schemaName, id)
		if err != nil{
			w.WriteHeader(http.StatusBadRequest)
			log.Printf("server.command.getAuthorizationUsers: createToken +%v\n", err)
			return
		}

		http.SetCookie(w,cookie)

		data, err := json.Marshal(user)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Printf("server.command.getUsers: cant parse json+%v\n", err)
			return
		}

		_,err =	w.Write(data)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Printf("server.command.getUsers: print fail: +%v\n", err)
			return
		}

		w.WriteHeader(http.StatusOK)
	}else {
		w.WriteHeader(http.StatusBadRequest)
		log.Printf("server.command.authorization: wrong data:" + r.FormValue("login") +":"+ r.FormValue("password"))
		return
	}
}