package main

import (
	"net/http"

	"streaming_demo/database"
	"streaming_demo/fetch-data"
)

func main() {
	database.InitDB()
	http.HandleFunc("/all_records", fetch_data.GetAllRecordsStreaming)
	http.HandleFunc("/all_records_non_streaming", fetch_data.GetAllRecordsNonStreaming)
	http.ListenAndServe(":8080", nil)
}
