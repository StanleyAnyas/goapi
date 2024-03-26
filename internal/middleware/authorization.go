package middleware

import (
	"errors"
	"net/http"

	"github.com/StanleyAnyas/goapi/api"
	"github.com/StanleyAnyas/goapi/internal/tools"
	log "github.com/sirupsen/logrus"
)

var UnAuthorized = errors.New("invalid username or token")

func Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var username string = r.URL.Query().Get("username")
		var token string = r.Header.Get("Authorization")
		var err error

		if username == "" || token == "" {
			log.Error(UnAuthorized)
			api.RequestErrorHandler(w, UnAuthorized)
			return
		}

		var database *tools.DatabaseInterface
		database, err = tools.NewDatabase()
		if err != nil {
			api.InternalErrorHandler(w)
			return
		}

		var loginDetails *tools.LoginDetails
		loginDetails = (*database).GetUserLoginDetails(username)

		if loginDetails == nil || (token != (*loginDetails).AuthToken) {
			log.Error(UnAuthorized)
			api.RequestErrorHandler(w, UnAuthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
