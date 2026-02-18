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

// UpdateUser
// @Summary update user data
// @Description update user deatils
// @Tags Users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Param user body models.UpdateUser true "deatils "
// @Success 201 {object} models.UserModel
// @Failure 400 {object} map[string]interface{}
// @security BearerAuth
// @Router /user/Update/{id} [patch]
func UpdateUser(c *fiber.Ctx) error {

	var UserSchema = database.DB.Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	//
	extUid := c.Locals("userId").(string)

	if extUid != c.Params("id") {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"success": false,
			"message": "You Are Not Authroized to Update This Profile",
		})
	}

	userid, _ := primitive.ObjectIDFromHex(c.Params("id"))

	var user models.UpdateUser
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"error":   "Invalid request body",
			"deatils": err.Error(),
		})
	}

	update := bson.M{"name": user.Name, "imageUrl": user.ImageUrl, "bio": user.Bio}

	result, err := UserSchema.UpdateOne(ctx, bson.M{"_id": userid}, bson.M{"$set": update})

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error":   "cannot update the user data",
			"deatils": err.Error(),
		})
	}
	//
	var updateUsser models.UserModel
	if result.MatchedCount == 1 {
		err := UserSchema.FindOne(ctx, bson.M{"_id": userid}).Decode(&updateUsser)
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"deatils": err.Error(),
			})
		}
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"data": updateUsser})

}
