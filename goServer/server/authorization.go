package server

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/go-chi/jwtauth"
	"log"
	"net/http"
)

var tokenAuth = jwtauth.New("HS256", []byte("key"), nil)

func myAuthenticator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, _, err := jwtauth.FromContext(r.Context())

		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		if token == nil || !token.Valid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		// Token is authenticated, pass it through
		next.ServeHTTP(w, r)
	})
}

func findToken(r *http.Request) string {
	cookie, err := r.Cookie("token")
	if err != nil{
		log.Printf("server.autorization.findToken: cant find cookie +%v\n",err)
		return ""
	}
	return cookie.Value
}

func createJwtCookie(schemaName string, id int) (cookie *http.Cookie,err error) {
		_, tokenString, err := tokenAuth.Encode(jwt.MapClaims{
			"schemaName": schemaName,
			"id":         id,
		})
		cookie = &http.Cookie{
			Name: "token",
			Value: tokenString,
			Path: "/",
		}
		if err != nil {
			log.Printf("server.command.authorization: Encode faild: +%v\n", err)
			return nil,err
		}
	return cookie,nil
}