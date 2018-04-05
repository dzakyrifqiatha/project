package server

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

const (
	addDonaturBk                    = `insert into donaturBuku(Kd_DonaturBuku, Nama_DonaturBuku, Status) values (?,?)`
	selectDonaturBkByKd_DonaturBuku = `select Kd_DonaturBuku, Nama_DonaturBuku, Status from donaturBuku where Kd_DonaturBuku = ?`
	selectDonaturBk                 = `select Kd_DonaturBuku, Nama_DonaturBuku, Status from donaturBuku`
	updateDonaturBk                 = `update donaturBuku set Kd_DonaturBuku =?, Nama_DonaturBuku = ?, Status = ?, where Kd_DonaturBuku = ?`
	selectMahasiswaByNama           = `select Kd_DonaturBuku, Nama_DonaturBuku, Status from donaturBuku Where Nama_DonaturBuku = ?`
)

//langkah 4
type dbReadWriter struct {
	db *sql.DB
}

func NewDBReadWriter(url string, schema string, user string, password string) ReadWriter {
	schemaURL := fmt.Sprintf("%s:%s@tcp(%s)/%s", user, password, url, schema)
	db, err := sql.Open("mysql", schemaURL)
	if err != nil {
		panic(err)
	}
	return &dbReadWriter{db: db}
}

func (rw *dbReadWriter) addDonaturBk(donaturBk DonaturBuku) error {
	fmt.Println("add")
	tx, err := rw.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec(addDonaturBk, donaturBk.Kd_DonaturBuku, donaturBk.Nama_DonaturBuku, OnAdd, time.Now())
	if err != nil {
		return err
	}
	return tx.Commit()
}

func (rw *dbReadWriter) ReadDonaturBkByKd_DonaturBuku(kd_donaturBk string) (DonaturBuku error) {
	donaturBk := DonaturBuku{Kd_DonaturBuku: kd_donaturBk}
	err := rw.db.QueryRow(selectDonaturBkByKd_DonaturBuku, kd_donaturBk).Scan(&donaturBk.Kd_DonaturBuku, &donaturBk.Nama_DonaturBuku,
		&donaturBk.Status)
	if err != nil {
		return DonaturBuku{}, err
	}
	return donaturBk, nil
}

func (rw *dbReadWriter) ReadDonaturBk() (DonaturBuku, error) {
	donaturBk := DonaturBuku{}
	rows, _ := rw.db.Query(selectDonaturBk)
	defer rows.Close()
	for rows.Next() {
		var c DonaturBuku
		err := rows.Scan(&c.Kd_DonaturBuku, &c.Nama_DonaturBuku, &c.Status)
		if err != nil {
			fmt.Println("error query:", err)
			return donaturBk, err
		}
		mahasiswa = append(mahasiswa, c)
	}
	return mahasiswa, nil
}

func (rw *dbReadWriter) updateDonaturBk(dbk DonaturBuku) error {
	tx, err := rw.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec(updateDonaturBk, dbk.Nama_DonaturBuku, dbk.Status)

	if err != nil {
		return err
	}

	return tx.Commit()
}

func (rw *dbReadWriter) ReadDonaturBkByNama_DonaturBuku(nama string) (DonaturBuku error) {
	donaturBk := DonaturBuku{Nama_DonaturBuku: nama}
	err := rw.db.QueryRow(selectDonaturBkByNama_DonaturBuku, nama).Scan(&donaturBk.Kd_DonaturBuku,
		&donaturBk.Nama_DonaturBuku, &donaturBk.Status)
	if err != nil {
		return DonaturBuku{}, err
	}
	return donaturBk, nil
}
