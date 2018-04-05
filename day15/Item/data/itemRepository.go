package data

import (
	"database/sql"
	"day15/Item/models"
	"fmt"
)

type ItemRepository struct {
	DB *sql.DB
}

func GetAll(db *ItemRepository) []models.Item {
	fmt.Println(db.DB)

	result, err := db.DB.Query("Select ItemName From tblItem")
	if err != nil {
		return nil
	}

	defer result.Close()
	fmt.Println(result)
	var item []models.Item
	for result.Next() {
		var i models.Item
		if err := result.Scan(&i.ItemName); err != nil {
			return nil
		}
		item = append(item, i)
	}
	return item
}
