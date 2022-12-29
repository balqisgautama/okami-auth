package router

import (
	"github.com/balqisgautama/okami-auth/constanta"
	"github.com/balqisgautama/okami-auth/dao"
	"github.com/balqisgautama/okami-auth/util/jwt"
	"net/http"
	"strconv"
)

func MiddlewareCORSOrigin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		(*&w).Header().Set("Access-Control-Allow-Origin", "*")
		(*&w).Header().Set("Access-Control-Allow-Headers", "origin, content-type, accept, authorization")
		(*&w).Header().Set("Access-Control-Allow-Credentials", "true")
		(*&w).Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS, HEAD")
		(*&w).Header().Set("Access-Control-Max-Age", "1209600")
		next.ServeHTTP(w, r)
	})
}

func MiddlewareCheckUserToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get(constanta.TokenHeaderName)
		if token != "" {
			tokenData, output_ := jwt.ValidatorJWTUser(token)
			if output_.Status.Code != "" {
				// user doesn't have access
				r.Header.Set(constanta.AccessHeaderName, "")
			} else {
				userFound, _ := dao.UserDAO.GetUserByClientID(tokenData.ClientID)
				r.Header.Set(constanta.ClientIDHeaderName, userFound.ClientID.String)
				r.Header.Set(constanta.AccessHeaderName, strconv.FormatInt(int64(userFound.Sysadmin.Int16), 10))
			}
		}
		next.ServeHTTP(w, r)
	})
}
