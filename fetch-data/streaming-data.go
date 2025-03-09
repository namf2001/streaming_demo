package fetch_data

import (
	"encoding/json"
	"net/http"

	"streaming_demo/database"
	"streaming_demo/models"

	"gorm.io/gorm/clause"
)

// GetAllRecordsStreaming fetches all records from the database and sends them as a stream of responses
func GetAllRecordsStreaming(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Transfer-Encoding", "chunked")

	batchSize := 100
	last_id := 0

	for {
		var records []models.Record
		database.DB.Db.Clauses(clause.OrderBy{
			Expression: clause.Expr{SQL: "id ASC"},
		}).Where("id > ?", last_id).Limit(batchSize).Find(&records)

		if len(records) == 0 {
			break
		}

		encoder := json.NewEncoder(w)
		for _, record := range records {
			if err := encoder.Encode(record); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.(http.Flusher).Flush()
		}

		last_id = records[len(records)-1].ID
	}
}
