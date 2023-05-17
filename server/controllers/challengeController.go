package controllers

import (
	"context"
	"net/http"

	"github.com/shangsuru/ctf-website/db"
	middlewares "github.com/shangsuru/ctf-website/handlers"
	"github.com/shangsuru/ctf-website/models"
	"go.mongodb.org/mongo-driver/bson"
)

var client = db.Dbconnect()

var GetChallengesEndpoint = http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
	var challenges []*models.Challenge

	collection := client.Database("ctf").Collection("challenges")
	cursor, err := collection.Find(context.TODO(), bson.D{{}})
	if err != nil {
		middlewares.ServerErrResponse(err.Error(), response)
		return
	}

	if err = cursor.All(context.TODO(), &challenges); err != nil {
		middlewares.ServerErrResponse(err.Error(), response)
		return
	}
	
	middlewares.SuccessArrRespond(challenges, response)
})
