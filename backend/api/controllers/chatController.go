package controllers

import (
	"context"
	"strconv"

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

// GetMsgsByNums
// @Summary get message by pagenation
// @Description GetMsgsByNumbetween two users by pagenation
// @Tags Chat
// @Accept json
// @Produce json
// @Param from query int true "Staring point page num"
// @Param firstuid query string true "first user id"
// @Param seconduid query string true "second user id"
// @Success 201 {object} []models.Message
// @Failure 400 {object} map[string]interface{}
// @Router /chat/getmsgsbynums [get]
func GetMsgsByNums(c *fiber.Ctx) error {

	var MessageSchema = database.DB.Collection("messages")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	from, err := strconv.Atoi(c.Query("from"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "invalid value form from",
			"err":     err.Error(),
		})
	}

	firstuid := c.Query("firstuid")
	seconduid := c.Query("seconduid")

	// construct the filer
	senderFilter := bson.M{"sender": firstuid, "recever": seconduid}
	receiverFilter := bson.M{"sender": seconduid, "recever": firstuid}
	filter := bson.M{"$or": []bson.M{senderFilter, receiverFilter}}

	// pagenation options
	options := options.Find()
	options.SetSort(bson.D{{Key: "_id", Value: -1}})
	options.SetSkip(int64(from * 2))
	options.SetLimit(2)

	// query the db
	cursor, err := MessageSchema.Find(ctx, filter, options)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Faild to retrieve messages",
			"error":   err.Error(),
		})
	}
	defer cursor.Close(ctx)

	// iterate over the cursor and build the res array
	var messages []models.Message
	for cursor.Next(ctx) {
		var msg models.Message
		err := cursor.Decode(&msg)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Faild to decode messages",
				"error":   err.Error(),
			})
		}
		messages = append(messages, msg)
	}

	// reverce the message array

	for i := 0; i < len(messages)/2; i++ {
		j := len(messages) - 1 - i
		messages[i], messages[j] = messages[j], messages[i]
	}

	if len(messages) == 0 {
		messages = []models.Message{}
	}

	// Return the created message
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"msgs": messages,
	})

}


