package middleware

import (
	"errors"
	"net/http"
	"strings"

	"github.com/tunardev/auth-user/pkg/entity"
	"github.com/tunardev/auth-user/pkg/repository"
	"github.com/tunardev/auth-user/pkg/utils"
)

func IsAuth(function func(w http.ResponseWriter, r *http.Request, user entity.User), repo repository.UserRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Authorization") == "" {
			utils.Response(w, http.StatusUnauthorized, errors.New("Wrong Authorization"), nil)
			return 
		}
	
		if strings.Split(r.Header.Get("Authorization"), " ")[1] == "" {
			utils.Response(w, http.StatusUnauthorized, errors.New("Wrong Authorization"), nil)
			return 
		}

		id, err := utils.VerifyJWT(strings.Split(r.Header.Get("Authorization"), " ")[1])
		if err != nil {
			utils.Response(w, http.StatusUnauthorized, errors.New("Wrong Authorization"), nil)
			return
		}

		user, err := repo.GetById(id)
		if err != nil {
			if err.Error() == "sql: no rows in result set" {
				utils.Response(w, http.StatusUnauthorized, errors.New("User not found"), nil)
				return
			}

			utils.Response(w, http.StatusInternalServerError, err, nil)
			return 
		}
		
		function(w, r, user)
	}
}