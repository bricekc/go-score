package mongodb

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ouvre une connection au serveur MongoDB et retourne la base de données de l'application.
// Les informations (username, password, db) seront lues dans les variables d'environnement correspondantes
func connectToDB() *mongo.Database {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	password := os.Getenv("password")
	username := os.Getenv("username")
	db := os.Getenv("db")
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	client, _ := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://"+username+":"+password+"@localhost:27017"))
	database := client.Database(db)
	return database
}

// singleton
var db *mongo.Database = connectToDB()

// retourne une collection identifiée par son nom
func Collection(name string) *mongo.Collection {
	return db.Collection(name)
}
