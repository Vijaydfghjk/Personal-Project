package main

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct {
	ID    primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name  string             `json:"name" bson:"name"`
	Email string             `json:"email" bson:"email"`
	Age   int                `json:"age" bson:"age"`
}

var userCollection *mongo.Collection

func initMongodb() {

	myclient := options.Client().ApplyURI("mongodb://localhost:27017")

	client, err := mongo.Connect(context.TODO(), myclient)

	if err != nil {

		log.Fatalf("Failed to mongodb %v", err)
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatalf("MongoDB connection error: %v", err)
	}

	log.Println("Connected to MongoDB!")

	userCollection = client.Database("products").Collection("users")
}

func Createuser(c *gin.Context) {

	var user User

	if err := c.ShouldBindJSON(&user); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": "failed ro create the user"})
		return
	}
	user.ID = primitive.NewObjectID()
	_, err := userCollection.InsertOne(context.TODO(), user)

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

func Getusers(c *gin.Context) {

	cursor, err := userCollection.Find(context.TODO(), bson.M{})

	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var users []User
	for cursor.Next(context.TODO()) {

		var user User
		if err := cursor.Decode(&user); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error decoding user"})
			return
		}

		users = append(users, user)
	}

	c.JSON(http.StatusOK, users)
}

func GetuserbyId(c *gin.Context) {

	id := c.Param("id")

	objectid, err := primitive.ObjectIDFromHex(id)

	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Id"})
		return
	}

	var user User
	err = userCollection.FindOne(context.TODO(), bson.M{"_id": objectid}).Decode(&user)

	if err != nil {
		if err == mongo.ErrNoDocuments {

			c.JSON(http.StatusBadRequest, gin.H{"error": "user not fond"})
			return
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch user"})
		}
		return
	}

	c.JSON(http.StatusOK, user)
}
func Updateuser(c *gin.Context) {

	id := c.Param("id")

	objectId, err := primitive.ObjectIDFromHex(id)

	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
		return
	}

	var user User
	if err := c.ShouldBindJSON(&user); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	update := bson.M{
		"$set": bson.M{
			"name":  user.Name,
			"email": user.Email,
			"age":   user.Age,
		},
	}

	_, err = userCollection.UpdateOne(context.TODO(), bson.M{"_id": objectId}, update)

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	//var updatedUser User
	//fmt.Printf("check", updatedUser)
	err = userCollection.FindOne(context.TODO(), bson.M{"_id": objectId}).Decode(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch updated user"})
		return
	}

	c.JSON(http.StatusOK, user)
}

func Deleteuser(c *gin.Context) {

	id := c.Param("id")

	objectId, err := primitive.ObjectIDFromHex(id)

	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
		return
	}

	_, err = userCollection.DeleteOne(context.TODO(), bson.M{"_id": objectId})

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Message": "Successfully deleted"})
}

func main() {

	initMongodb()

	router := gin.Default()

	router.POST("/users", Createuser)
	router.GET("/users", Getusers)
	router.GET("/users/:id", GetuserbyId)
	router.PUT("users/:id", Updateuser)
	router.DELETE("users/:id", Deleteuser)

	if err := router.Run(":8000"); err != nil {
		log.Fatalf("Faild to run the server :%v", err)
	}

}
