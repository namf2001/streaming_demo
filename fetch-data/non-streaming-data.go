package fetch_data

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"streaming_demo/database"
	"streaming_demo/models"
)

// GetAllRecordsNonStreaming fetches all records from the database and sends them as a single response
func GetAllRecordsNonStreaming(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	start := time.Now()
	var records []models.Record
	database.DB.Db.Find(&records)
	endFetch := time.Now()
	fmt.Println("Time to fetch records:", endFetch.Sub(start))

	encoder := json.NewEncoder(w)
	startEncode := time.Now()
	err := encoder.Encode(records)
	if err != nil {
		return
	}
	endEncode := time.Now()
	fmt.Println("Time to encode records:", endEncode.Sub(startEncode))

	totalTime := endEncode.Sub(start)
	fmt.Println("Total time:", totalTime)
}
