package handlers

import (
	"context"
	"drakk-bk/dynamo"
	"drakk-bk/model"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func PostMenuItemHandler(svc *dynamo.Client, w http.ResponseWriter, r *http.Request) {
    if r.Method != "POST" {
        http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
        return
    }

    var menuItem model.MenuItem
    err := json.NewDecoder(r.Body).Decode(&menuItem)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    av, err := attributevalue.MarshalMap(menuItem)
    if err != nil {
        http.Error(w, "Could not marshal menu item", http.StatusInternalServerError)
        return
    }

	tableName := os.Getenv("DYNAMODB_TABLE_NAME")

    _, err = svc.Service.PutItem(context.TODO(), &dynamodb.PutItemInput{
        TableName: aws.String(tableName),
        Item:      av,
    })

    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, "Item inserted successfully")
}

func GetAllMenuItemsHandler(svc *dynamo.Client, w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	tableName := os.Getenv("DYNAMODB_TABLE_NAME")

	result, err := svc.Service.Scan(context.TODO(), &dynamodb.ScanInput{
		TableName: aws.String(tableName),
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var menuItems []model.MenuItem
	err = attributevalue.UnmarshalListOfMaps(result.Items, &menuItems)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(menuItems)
}

func GetOneMenuItemHandler(svc *dynamo.Client, w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	id := r.URL.Query().Get("pk")
	sk := r.URL.Query().Get("sk")

	key, err := attributevalue.MarshalMap(map[string]string{
		"PK": id,
		"SK": sk,
	})
	if err != nil {
		http.Error(w, "Failed to marshal DynamoDB key", http.StatusInternalServerError)
		return
	}

	tableName := os.Getenv("DYNAMODB_TABLE_NAME")

	result, err := svc.Service.GetItem(context.TODO(), &dynamodb.GetItemInput{
		TableName: aws.String(tableName),
		Key:       key,
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var menuItem model.MenuItem
	err = attributevalue.UnmarshalMap(result.Item, &menuItem)
	if err != nil {
		http.Error(w, "Failed to unmarshal DynamoDB item", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(menuItem)
}

func UpdateMenuItemHandler(svc *dynamo.Client, w http.ResponseWriter, r *http.Request) {
	if r.Method != "PUT" {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var menuItem model.MenuItem
	err := json.NewDecoder(r.Body).Decode(&menuItem)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	av, err := attributevalue.MarshalMap(menuItem)
	if err != nil {
		http.Error(w, "Could not marshal menu item", http.StatusInternalServerError)
		return
	}

	tableName := os.Getenv("DYNAMODB_TABLE_NAME")

	_, err = svc.Service.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName: aws.String(tableName),
		Item:      av,
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Item updated successfully")
}

func DeleteMenuItemHandler(svc *dynamo.Client, w http.ResponseWriter, r *http.Request) {
	if r.Method != "DELETE" {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	id := r.URL.Query().Get("pk")
	sk := r.URL.Query().Get("sk")

	key, err := attributevalue.MarshalMap(map[string]string{
		"PK": id,
		"SK": sk,
	})
	if err != nil {
		http.Error(w, "Failed to marshal DynamoDB key", http.StatusInternalServerError)
		return
	}

	tableName := os.Getenv("DYNAMODB_TABLE_NAME")

	_, err = svc.Service.DeleteItem(context.TODO(), &dynamodb.DeleteItemInput{
		TableName: aws.String(tableName),
		Key:       key,
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Item deleted successfully")
}