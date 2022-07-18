package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"encoding/json"

	"github.com/joho/godotenv"
)

type RawSheetData struct {
	Range          string     `json:"range"`
	MajorDimension string     `json:"majorDimension"`
	Values         [][]string `json:"values"`
}

type SheetData struct {
	Values [][]string
}

// Get data range from Google Sheets API using key param authorization.
func sheetDataToStruct() (*RawSheetData) {
	// Format URL with environment variables
	endpoint := "https://sheets.googleapis.com/v4/spreadsheets/%s/values/%s?key=%s"
	url := fmt.Sprintf(endpoint, os.Getenv("SHEET_ID"), os.Getenv("SHEET_RANGE"), os.Getenv("API_KEY"))

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal("No response from server")
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Failed to read response body", err)
	}
	
	var data RawSheetData

	if err := json.Unmarshal(body, &data); err != nil {
		log.Fatal("Failed to unmarshal JSON")
	}
	
	return &data
}

// Create stringifed JSON from SheetData struct
func getSheetData(sheet *RawSheetData) string {
	// Copy sheets values into new values-only struct
	data := &SheetData{Values: sheet.Values}

	// Encode struct as JSON
	b, err := json.Marshal(data)
	if err != nil {
		log.Fatal("Failed to martial data", err)
	}

	return string(b)
}

// Create JSON response from sheet data
func httpResponse(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, getSheetData(sheetDataToStruct()))
}

/*
 * This is used to spin up a local web server outside
 * for testing the HTTP endpoint. Cloud Function should
 * call httpResponse() directly.
 */
func main() {
	// Load environment variables from '.env' file.
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Failed to load '.env' file")
	}

	port := ":8090"

	fmt.Printf("Listening on localhost:%s\n", port)

	http.HandleFunc("/", httpResponse)
	http.ListenAndServe(port, nil)
}
