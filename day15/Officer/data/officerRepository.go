package data

import (
	"database/sql"
	"day15/Officer/models"
	"fmt"
)

type OfficerRepository struct {
	DB *sql.DB
}

func GetAll(db *OfficerRepository) []models.Officer {
	fmt.Println(db.DB)

	result, err := db.DB.Query("Select OfficerName, OfficerCode From tblOfficer")
	if err != nil {
		return nil
	}

	defer result.Close()
	fmt.Println(result)
	var officer []models.Officer
	for result.Next() {
		var o models.Officer
		if err := result.Scan(&o.OfficerName, &o.OfficerCode); err != nil {
			return nil
		}
		officer = append(officer, o)
	}
	return officer
}
