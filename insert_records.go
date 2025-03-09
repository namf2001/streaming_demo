package main

import (
	"fmt"
	"time"

	"streaming_demo/database"
)

func main() {
	database.InitDB()
	var records []database.Record
	startDate := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
	for i := 0; i < 100000; i++ {
		// Tạo ngày ngẫu nhiên trong khoảng 2 năm
		daysOffset := time.Duration(i%730) * 24 * time.Hour
		createAt := startDate.Add(daysOffset)
		updateAt := createAt.Add(time.Duration(i%30) * 24 * time.Hour) // Cập nhật trong 30 ngày sau

		record := database.Record{
			Keyword:  fmt.Sprintf("keyword%d", i),
			Node:     fmt.Sprintf("node%d", i),
			CreateAd: createAt,
			UpdateAp: updateAt,
		}
		records = append(records, record)
		if len(records) == 1000 { // Chèn theo lô 1000 bản ghi để tăng hiệu suất
			database.DB.Create(&records)
			records = records[:0] // Reset slice
		}
	}
	if len(records) > 0 {
		database.DB.Create(&records)
	}
	fmt.Println("Inserted 100,000 records successfully")
}
