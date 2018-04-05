package server

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

const (
	addBuku                = `insert into Buku(Kd_Buku, Nama_Buku, Status, Keterangan, CreateBy, CreateOn) values (?,?,?,?,?,?)`
	selectBukuByKode       = `select Kd_Buku, Nama_Buku, Status, Keterangan from Buku where Kd_Buku = ?`
	selectBuku             = `select Kd_Buku, Nama_Buku, Status, Keterangan from Buku`
	updateBuku             = `update Buku set Kd_Buku =?, Nama_Buku = ?, Status = ?, Keterangan = ? where Kd_Buku = ?`
	selectBukuByKeterangan = `select Kd_Buku, Nama_Buku, Status, Keterangan from Buku Where Keterangan like ?`
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

func (rw *dbReadWriter) AddBuku(buku Buku) error {
	fmt.Println("add")
	tx, err := rw.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec(addBuku, buku.KodeBuku, buku.NamaBuku, OnAdd, buku.Keterangan, buku.CreateBy, time.Now())
	if err != nil {
		return err
	}
	return tx.Commit()
}

func (rw *dbReadWriter) ReadBukuByKode(kode string) (Buku, error) {
	buku := Buku{KodeBuku: kode}
	err := rw.db.QueryRow(selectBukuByKode, kode).Scan(&buku.KodeBuku, &buku.NamaBuku,
		&buku.Status)
	if err != nil {
		return Buku{}, err
	}
	return buku, nil
}

func (rw *dbReadWriter) ReadBuku() (Bukus, error) {
	buku := Bukus{}
	rows, _ := rw.db.Query(selectBuku)
	defer rows.Close()
	for rows.Next() {
		var c Buku
		err := rows.Scan(&c.KodeBuku, &c.NamaBuku, &c.Status, &c.Keterangan)
		if err != nil {
			fmt.Println("error query:", err)
			return buku, err
		}
		buku = append(buku, c)
	}
	return buku, nil
}

func (rw *dbReadWriter) UpdateBuku(buku Buku) error {
	tx, err := rw.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec(updateBuku, buku.KodeBuku, buku.NamaBuku, buku.Status, buku.Keterangan, buku.UpdateBy)

	if err != nil {
		return err
	}

	return tx.Commit()
}

func (rw *dbReadWriter) ReadBukuByKeterangan(keterangan string) (Bukus, error) {
	buku := Bukus{}
	rows, _ := rw.db.Query(selectBukuByKeterangan, keterangan)
	defer rows.Close()
	for rows.Next() {
		var c Buku
		err := rows.Scan(&c.KodeBuku, &c.NamaBuku, &c.Status, &c.Keterangan)
		if err != nil {
			fmt.Println("error query:", err)
			return buku, err
		}
		buku = append(buku, c)
	}
	return buku, nil
}
