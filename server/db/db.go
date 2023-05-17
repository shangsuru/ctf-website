package db

import (
	"context"
	"log"

	"github.com/fatih/color"
	middlewares "github.com/shangsuru/ctf-website/handlers"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"github.com/shangsuru/ctf-website/models"
)


func Dbconnect() *mongo.Client {
	clientOptions := options.Client().ApplyURI(middlewares.DotEnvVariable("MONGO_URL"))
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal("⛒ Connection Failed to Database")
		log.Fatal(err)
	}
	
	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal("⛒ Connection Failed to Database")
		log.Fatal(err)
	}
	color.Green("⛁ Connected to Database")

	// Insert challenges
	challenges := []interface{}{
		models.Challenge{ID: 1, Name: "Integer Overflow", Flag: "CTF{1N73G3r_0v3RFL0W_15_Kw173_C0mm0N_d84F97c11151809c9551C6}"},
		models.Challenge{ID: 2, Name: "Format String", Flag: "CTF{mY_5trIN9_I5_8L33DING_0a8dfeb25b78db2bcfe524}"},
		models.Challenge{ID: 3, Name: "Stack Overflow", Flag: "CTF{M4573r1n9_574CK_0v3RfL0w_83080867736FF2449dff26}"},
	}
	collection := client.Database("ctf").Collection("challenges")
	_, err = collection.InsertMany(context.TODO(), challenges)
	if err != nil {
		log.Fatal("⛒ Could not insert challenges")
	}

	color.Green("⛁ Inserted challenges")

	return client
}
