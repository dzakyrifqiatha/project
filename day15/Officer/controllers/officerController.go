package controllers

import (
	"day15/Officer/data"
	"encoding/json"
	"net/http"
)

func GetOfficer(w http.ResponseWriter, r *http.Request) {

	context := Context{}
	d := DBInitial(context.DB)
	defer d.Close()

	repo := &data.OfficerRepository{d}

	officer := data.GetAll(repo)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	mdl, _ := json.Marshal(officer)
	w.Write(mdl)
}
