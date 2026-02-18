package controllers

import (
	"context"
	"fmt"
	"server/database"
	"server/models"
	"time"


	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	
)

// GetUserBy ID
// @Summary Get User By ID
// @Description Get User Detail By ID
// @Tags  Users
// @Accept json
// @Produce json
// @Param   id path string true "User ID"
// @Success 201 {object} models.UserModel
// @Failure 400 {object} map[string]interface{}
// @Router /user/getUser/{id} [get]
func GetUserByID(c *fiber.Ctx) error {

	var UserSchema = database.DB.Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	var user models.UserModel

	objID, _ :=  primitive.ObjectIDFromHex(c.Params("id"))
	strID :=c.Params("id")
	fmt.Println(strID)
	//TODO GET and Return user posts

	//get user data
	userResult := UserSchema.FindOne(ctx , bson.M{"_id":objID})
	if userResult.Err() != nil{
			return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"sucess":   false,
			"message": "user not found",
		})
	}
 userResult.Decode(&user)
	
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"result": user,
		"post":  []string{},
	})
}
