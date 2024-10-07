package services

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	_ "github.com/mattn/go-sqlite3"
	"gorm.io/gorm"
)

// Structure to hold product data from SQLite
type Product struct {
	gorm.Model // handle createAt  updateAt and deleteAt
    Name        string  `json:"name" binding:"required"`
    Description string  `json:"description"`
    Price       float64 `json:"price" binding:"required"`
	Qty uint `json:"qty"`
}

// Function to replicate SQLite data to CouchDB
func ReplicateFromSQLiteToCouchDB() error {
	// Open SQLite connection
	db, err := sql.Open("sqlite3","./cmd/api/product.db")
	if err != nil {
		return fmt.Errorf("error opening SQLite DB: %v", err)
	}  
	defer db.Close()

	// Query all products from SQLite
	rows, err := db.Query("SELECT id, name, description, price FROM products")
	if err != nil {
		return fmt.Errorf("error querying SQLite DB: %v", err)
	}
	defer rows.Close()

	// Prepare the list of products to replicate
	var products []Product
	for rows.Next() {
		var product Product
		if err := rows.Scan(&product.ID, &product.Name, &product.Price); err != nil {
			return fmt.Errorf("error scanning row: %v", err)
		}
		products = append(products, product)
	}

	// Convert products to JSON format
	productsJSON, err := json.Marshal(products)
	if err != nil {
		return fmt.Errorf("error marshaling products to JSON: %v", err)
	}

	// Send the products to CouchDB
	err = sendToCouchDB(productsJSON)
	if err != nil {
		return fmt.Errorf("error sending data to CouchDB: %v", err)
	}
	return nil
}

// Function to send data to CouchDB
func sendToCouchDB(data []byte) error {

	url := "http://localhost:5984/rest_api_replica_01/_bulk_docs"

	// Prepare the payload for CouchDB bulk insert
	payload := map[string]interface{}{
		"docs": json.RawMessage(data),
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("error marshaling payload: %v", err)
	}

	// Make the POST request to CouchDB
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(payloadBytes))
	if err != nil {
		return fmt.Errorf("error sending POST request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	log.Println("Data replicated successfully to CouchDB")
	return nil
}
