package main

import (
	"context"
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const dbname = "PersonDb"
const collectionName = "Person"

func indexRoute(res *fiber.Ctx) error {
	collection, err := getMongoDbCollection(dbname, collectionName)
	if err != nil {
		return res.Status(400).SendString("There is some problem! Please Try Again !")
	}
	var filter bson.M = bson.M{}
	curr, err := collection.Find(context.Background(), filter)
	if err != nil {
		return res.Status(400).SendString("There is some problem! Please Try Again !")
	}
	defer curr.Close(context.Background())
	var result []bson.M
	curr.All(context.Background(), &result)

	json, _ := json.Marshal(result)
	return res.Status(200).Send(json)
}

func addPerson(res *fiber.Ctx) error {
	collection, err := getMongoDbCollection(dbname, collectionName)
	if err != nil {
		return res.Status(400).SendString("There is some problem! Please Try Again !")
	}
	var newPerson Person
	json.Unmarshal([]byte(res.Body()), &newPerson)
	curr, err := collection.InsertOne(context.Background(), newPerson)

	if err != nil {
		return res.Status(400).SendString("There is some problem! Please Try Again !")
	}

	response, _ := json.Marshal(curr)
	return res.Status(200).Send(response)
}

func updatePerson(res *fiber.Ctx) error {
	collection, err := getMongoDbCollection(dbname, collectionName)
	if err != nil {
		return res.Status(400).SendString("There is some error! Try again later")
	}
	id := res.Params("id")
	objId, _ := primitive.ObjectIDFromHex(id)
	var updatePerson Person
	json.Unmarshal([]byte(res.Body()), &updatePerson)
	var filter bson.M = bson.M{
		"_id": objId,
	}
	update := bson.M{
		"$set": updatePerson,
	}
	curr, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return res.Status(400).SendString("There is some error! Try again later")
	}
	response, _ := json.Marshal(curr)
	return res.Status(200).Send(response)
}

func deletePerson(res *fiber.Ctx) error {
	collection, err := getMongoDbCollection(dbname, collectionName)
	if err != nil {
		return res.Status(400).SendString("There is some error! Try again later")
	}
	id := res.Params("id")
	objId, _ := primitive.ObjectIDFromHex(id)
	var filter bson.M = bson.M{
		"_id": objId,
	}
	curr, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		return res.Status(400).SendString("There is some error! Try again later")
	}
	response, _ := json.Marshal(curr)
	return res.Status(200).Send(response)
}

func createAssignment(res *fiber.Ctx) error {
	collection, err := getMongoDbCollection(dbname, "Assignment")
	if err != nil {
		return res.Status(400).SendString("There is some problem! Please Try Again !")
	}
	var newAssignment Assignment
	json.Unmarshal([]byte(res.Body()), &newAssignment)
	personCollection, err := getMongoDbCollection(dbname, collectionName)

	if err != nil {
		return res.Status(400).SendString("There is some error! Try again later")
	}

	id := res.Params("id")
	newAssignment.Person = id

	objId, _ := primitive.ObjectIDFromHex(id)
	curr, _ := collection.InsertOne(context.Background(), newAssignment)

	var filter bson.M = bson.M{
		"_id": objId,
	}

	var temp Person
	personCollection.FindOne(context.Background(), filter).Decode(&temp)
	temp.Assignments = append(temp.Assignments, curr.InsertedID.(primitive.ObjectID).Hex())
	update := bson.M{
		"$set": temp,
	}
	result, _ := personCollection.UpdateOne(context.Background(), filter, update)
	response, _ := json.Marshal(result)
	return res.Status(200).Send(response)
}

func getAssignments(res *fiber.Ctx) error {
	collection, err := getMongoDbCollection(dbname, "Assignment")
	if err != nil {
		return res.Status(400).SendString("There is some problem! Please Try Again !")
	}
	id := res.Params("id")
	var filter bson.M = bson.M{
		"person": id,
	}
	curr, err := collection.Find(context.Background(), filter)
	if err != nil {
		return res.Status(400).SendString("There is some problem! Please Try Again !")
	}
	defer curr.Close(context.Background())
	var result []bson.M
	curr.All(context.Background(), &result)

	json, _ := json.Marshal(result)
	return res.Status(200).Send(json)
}

func main() {
	app := fiber.New()
	app.Get("/", indexRoute)
	app.Post("/create", addPerson)
	app.Put("/update/:id", updatePerson)
	app.Delete("/delete/:id", deletePerson)
	app.Post("/assignment/:id", createAssignment)
	app.Get("/assignments/:id", getAssignments)
	app.Listen(":8000")
}
