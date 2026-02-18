package controllers

import (
	"context"
	
	"server/database"
	"server/models"
	"time"

	
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	
)

// SendMessage
// @Summary send message to friend user
// @Description SendMessage form one user to another
// @Tags Chat
// @Accept json
// @Produce json
// @Param message body models.SendMessageM true "user SendMessage deatils"
// @Success 201 {object} models.Message
// @Failure 400 {object} map[string]interface{}
func SendMessage(c *fiber.Ctx) error {

	var MessageSchema = database.DB.Collection("messages")
	var UnReadedMsgSchema = database.DB.Collection("unReadedmessages")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	var body models.SendMessageM
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"error":   "Invalid request body",
			"deatils": err.Error(),
		})
	}

  var msg models.Message
  c.BodyParser(&msg)

  // save the message to db
  result , err := MessageSchema.InsertOne(ctx, &msg)
  if  err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"error":   "I failed to save message",
			"deatils": err.Error(),
		})
	}

	//update or create the unreaded messages count and is readed
	var unReadedMsg models.UnReadedMsg
	filter := bson.M{"mainUserId":msg.Receiver ,"otherUserId":msg.Sender}
	update := bson.M{"$inc":bson.M{"numOfUnreadedMessages":1},"$set":bson.M{"isReaded":false}}
	opts := options.FindOneAndUpdate().SetUpsert(true)
	err = UnReadedMsgSchema.FindOneAndUpdate(ctx ,filter ,update,opts).Decode(&unReadedMsg)
	if err !=nil && err != mongo.ErrNoDocuments{
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{
			"error":   "Failed to update unreaded  message count",
			"deatils": err.Error(),
		})

	}

     
    // return the created message
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message":"Message send Sucessfully",
		"result":result.InsertedID,
	})
}

