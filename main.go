package main

import (
	"drakk-bk/dynamo"
	"drakk-bk/handlers"
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic("Error loading .env file")
	}

	tableName := os.Getenv("DYNAMODB_TABLE_NAME")
	if tableName == "" {
		panic("DYNAMODB_TABLE_NAME is not set")
	}
	
    dynamoClient, err := dynamo.New()
    if err != nil {
        panic("Failed to initialize DynamoDB client: " + err.Error())
    }

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello World")
    })

    http.HandleFunc("/menu-item/add", func(w http.ResponseWriter, r *http.Request) {
        handlers.PostMenuItemHandler(dynamoClient, w, r)
    })

	http.HandleFunc("/menu-item/all", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetAllMenuItemsHandler(dynamoClient, w, r)
	})

	http.HandleFunc("/menu-item/get", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetOneMenuItemHandler(dynamoClient, w, r)
	})

	http.HandleFunc("/menu-item/delete", func(w http.ResponseWriter, r *http.Request) {
		handlers.DeleteMenuItemHandler(dynamoClient, w, r)
	})

	http.HandleFunc("/menu-item/update", func(w http.ResponseWriter, r *http.Request) {
		handlers.UpdateMenuItemHandler(dynamoClient, w, r)
	})

    fmt.Println("Server is running on port 8080...")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        panic(err)
    }
}
