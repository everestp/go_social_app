package controller

import (
	"context"
	"os"
	"server/database"
	"server/models"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

// Register
// @Summary Gegister a new user
// @Description Register an ew user by providing email, password , first name , last name
// @Tags Authentication
// @Accept json
// @Produce json
// @Param user body models.CreateUser true "user register deatils"
// @Success 201 {object} models.UserModel
// @Failure 400 {object} map[string]interface{}
// @Router /user/signup [post]
func Register(c *fiber.Ctx) error {

	UserSchema := database.DB.Collection("users")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	var body models.CreateUser

	// Parse Body
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   "Invalid request body",
			"details": err.Error(),
		})
	}

	// Check if user already exists
	var existingUser models.UserModel
	err := UserSchema.FindOne(ctx, bson.M{"email": body.Email}).Decode(&existingUser)

	if err == nil {
		// User found -> already exists
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"message": "User with email " + body.Email + " already exists",
		})
	}

	// If error is something other than "no documents found"
	if err != mongo.ErrNoDocuments {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Hash password
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to hash password",
		})
	}

	newUser := models.UserModel{
		Name:      body.FirstName + " " + body.LastName,
		Email:     body.Email,
		Password:  string(hashPassword),
		Followers: make([]string, 0),
		Following: make([]string, 0),
	}

	result, err := UserSchema.InsertOne(ctx, newUser)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// get the  new User
	 var createdUser *models.UserModel
	 query := bson.M{"_id":result.InsertedID}
	 UserSchema.FindOne(ctx, query).Decode(&createdUser)

	 // Create the token
	 claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
             Issuer: createdUser.ID.Hex(),
			 ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),


	 })

	 jwtSecret := os.Getenv("JWT_SECRET")
	 
	 token , _ := claims.SignedString([]byte(jwtSecret))




	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"sucess":    true,
		"result": createdUser,
		"token":token,
		"message": "User registered successfully",
	})
}
