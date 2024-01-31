package app

import (
	"fmt"
	"net/http"
	routes "orgmanagementapp/pkg/api/routes"
	mongodb "orgmanagementapp/pkg/database/mongodb/repository"

	"github.com/gorilla/mux"

	"go.mongodb.org/mongo-driver/mongo"
)

// App struct holds the application dependencies
type App struct {
	Port        int
	MongoClient *mongo.Client
}

// InitializeApp sets up the application
func InitializeApp(port int) *App {
	// Connect to MongoDB
	// To run locally change this to "mongodb://localhost:27017"
	err := mongodb.InitMongoDB("mongodb://db:27017", "new-db")
	if err != nil {
		fmt.Printf("Error connecting to MongoDB: %v\n", err)
		return nil
	}

	// Create a new App instance with MongoDB connection
	app := &App{
		Port:        port,
		MongoClient: mongodb.GetMongoClient(),
	}
	router := mux.NewRouter()
	routes.SetUserRoutes(router)
	routes.SetOrganizationRoutes(router)
	app.setupHTTPServer(router)
	return app
}

func (app *App) setupHTTPServer(router *mux.Router) {
	// Define the port you want to listen on
	port := app.Port
	addr := fmt.Sprintf(":%d", port)
	fmt.Printf("Listening on http://localhost%s\n", addr)
	http.Handle("/", router)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}
