package middleware

import (
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)


func AuthMiddleware(c *fiber.Ctx )  error{
	tok := c.Get("Authorization")
	if tok == ""{
		return  c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message":"unauthenticate",
		})
	}

	//Check if token start with bearer
	if !strings.HasPrefix(tok , "Bearer "){
		tok = strings.TrimSpace(tok)
	} else{
		splited := strings.Split(tok , "Bearer ")
		if len(splited) != 2{
			return  c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message":"unauthenticate",
		})
		}
		tok =splited[1]
	}

 SecretKey := os.Getenv("JWT_SECRET")
 token , err := jwt.ParseWithClaims(tok , &jwt.StandardClaims{}, func (t *jwt.Token) (interface{} , error)  {
	 return  []byte(SecretKey),nil
 })
 if err != nil{
	return  c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message":"unauthenticate",
		})
 }

claims , ok := token.Claims.(*jwt.StandardClaims)
if !ok || !token.Valid {
	return  c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message":"unauthenticate",
		})
}
c.Locals("userId",claims.Issuer)
 return c.Next()

}