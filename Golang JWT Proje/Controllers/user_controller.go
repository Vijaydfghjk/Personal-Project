package controllers

import (
	database "golang-jwt-project/Database"
	helperes "golang-jwt-project/Helperes"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/mongo"
)

var usercollection *mongo.Collection = database.Opencollection(database.Client, "user")

var validate = validator.New()

func Hashpassword() {

}

func Verifypassword() {

}

func Login() {

}

func Signup() {

}

func GetUsers() {

}

func GetUser() gin.HandlerFunc{

   
return func (c *gin.Context) {
	
	  userID:= c.Param("user_id")

	  if err := helperes.MatchuserTypetoUid(c, userId) err != nil{


		  c.json(http.StatusBadRequest,gin.H{"eooro":err.Error()}) 
		  return
	  }
  }
}
