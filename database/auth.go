package database

import (
	security "api_DB/utils"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Database struct {
	DB *mongo.Database
}

type Credential struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (db *Database) LoginHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to LoginHandler\n"))
	creds := Credential{}

	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		http.Error(w, "Invalid JSON Format", http.StatusUnauthorized)
		return
	}
	log.Println("Received Username:", creds.Username)
	log.Println("Received Password:", creds.Password)
	hasedPassword, _ := security.HashPassword(creds.Password)

	userDocument := bson.M{"username": creds.Username, "password": hasedPassword}

	collection := db.DB.Collection("user")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := collection.InsertOne(ctx, userDocument)
	if err != nil {
		http.Error(w, "Database Error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "User Registerd Successfully!"})
}

// func ConnectMongo() (*mongo.Database, error) {
// 	err := godotenv.Load()
// 	if err != nil {
// 		log.Fatal("Error Loading .env file")
// 	}

// 	mongoURI := os.Getenv("MONGO_URI")
// 	databaseName := os.Getenv("DATABASE_NAME")

// 	clientOption := options.Client().ApplyURI(mongoURI)

// 	client, err := mongo.Connect(context.TODO(), clientOption)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	err = client.Ping(context.TODO(), nil)
// 	if err != nil {
// 		log.Fatal("Could not able to connect to MongoDB:", err)

// 	}
// 	fmt.Println("Connected to MongoDB Successfully")

// 	db := client.Database(databaseName)
// 	fmt.Println("using database", db.Name())
// }

func ConnectMongo() (*Database, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("error loading .env file: %v", err)
	}

	mongoURI := os.Getenv("MONGO_URI")
	clientOptions := options.Client().ApplyURI(mongoURI)

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, fmt.Errorf("MongoDB connection failed: %v", err)
	}

	db := client.Database(os.Getenv("DATABASE_NAME"))
	log.Println("Connected to MongoDB successfully!")
	log.Println("Using database:", db.Name())
	return &Database{DB: db}, nil
}
