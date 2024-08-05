package helperes

import (
	"errors"

	"github.com/gin-gonic/gin"
)

func CheckUsertype(c *gin.Context, role string) (err error) {

	usertype := c.GetString("user_type")

	err = nil

	if usertype != role {

		err = errors.New("Unauthorized to aceess this resource")
		return
	}

	return err
}

func MatchuserTypetoUid(c *gin.Context, UserID string) (err error) {

	userType := c.GetString("user_type")
	uid := c.GetString("uid")
	err = nil

	if userType == "USER" && uid != UserID {

		err = errors.New("Unauthorized to aceess this resource")
	}

	err = CheckUsertype(c, userType)
	return err
}
