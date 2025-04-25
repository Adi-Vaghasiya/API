package main

import (
	"api_DB/database"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

// func ConnectMongo() (*mongo.Database, error) {
// 	err := godotenv.Load()
// 	if err != nil {
// 		return nil, fmt.Errorf("error loading .env file: %v", err)
// 	}

// 	mongoURI := os.Getenv("MONGO_URI")
// 	clientOptions := options.Client().ApplyURI(mongoURI)

// 	client, err := mongo.Connect(context.TODO(), clientOptions)
// 	if err != nil {
// 		return nil, fmt.Errorf("MongoDB connection failed: %v", err)
// 	}

// 	db := client.Database(os.Getenv("DATABASE_NAME"))
// 	return db, nil
// }

func main() {

	//database.ConnectMongo()
	dbInstance, err := database.ConnectMongo()
	if err != nil {
		log.Fatal("Failed to initialize database:", err)
	}

	r := chi.NewRouter()
	r.Post("/login", dbInstance.LoginHandler)

	log.Println("Server started on :8000")
	http.ListenAndServe(":8000", r)
}
