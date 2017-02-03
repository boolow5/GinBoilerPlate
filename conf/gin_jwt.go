package conf

import (
	"strings"
	"time"

	"gopkg.in/appleboy/gin-jwt.v2"

	"github.com/boolow5/GinBoilerPlate/models"
	"github.com/gin-gonic/gin"
)

func NewAuthMiddleware(min_level int) *jwt.GinJWTMiddleware {
	return &jwt.GinJWTMiddleware{
		Realm:         "test zone",
		Key:           []byte("secret key"),
		Timeout:       time.Hour,
		MaxRefresh:    time.Hour * 24,
		Authenticator: AuthenticatorFunc,
		Authorizator: func(username string, c *gin.Context) bool {
			user, err := models.GetUserByUsername(username)
			if err != nil {
				return false
			}
			if user.Role.Level < (min_level + 1) {
				return true
			}

			return false
		},
		Unauthorized: UnauthorizedFunc,
		TokenLookup:  "header:Authorization",
	}
}

func AuthenticatorFunc(userId string, password string, c *gin.Context) (string, bool) {
	// params: username, password string
	// operation: 1. get user from database, 2. check user's password app_id match
	// returns: username string, authenticated bool

	ids := strings.Split(userId, "@")
	var username, app_id string

	if len(ids) > 1 {
		username, app_id = ids[0], ids[1]
	}

	user, err := models.AuthenticateUser(username, password, app_id)
	if err != nil {

		return "", false
	}
	return user.Username, true
	/*
		if (username == "admin" && password == "admin") || (username == "test" && password == "test") {
			return username, true
		}

		return username, false*/
}

func UnauthorizedFunc(c *gin.Context, code int, message string) {
	c.JSON(code, gin.H{
		"code":    code,
		"message": message,
	})
}
