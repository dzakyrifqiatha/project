package data

import (
	"database/sql"
	"day15/Selling/models"
	"fmt"
)

type SellingRepository struct {
	DB *sql.DB
}

func GetAll(db *SellingRepository) []models.Selling {
	fmt.Println(db.DB)

	result, err := db.DB.Query("Select Item, Total, Paid, OfficerCode From tblSelling")
	if err != nil {
		return nil
	}

	defer result.Close()
	fmt.Println(result)
	var selling []models.Selling
	for result.Next() {
		var s models.Selling
		if err := result.Scan(&s.Item, &s.Total, &s.Paid, &s.OfficerCode); err != nil {
			return nil
		}
		selling = append(selling, s)
	}
	return selling
}
