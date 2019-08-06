package App

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/Harits514/BahasaKitaBETes/App/model"
	_ "github.com/go-sql-driver/mysql"
)

func dbConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := ""
	dbName := "dbtest"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		log.Println("connection error")
		panic(err.Error())
	}
	return db
}

func InsertData(w http.ResponseWriter, r *http.Request) {
	music := model.Music{}
	log.Println(r.Body)
	err := json.NewDecoder(r.Body).Decode(&music)
	if err != nil {
		log.Println(err)
		http.Error(w, "json decode fail", http.StatusInternalServerError)
		return
	}
	db := dbConn()
	insForm, err := db.Prepare("INSERT INTO musik(namalagu, namaartis, link, durasi) VALUES(?,?,?,?)")
	if err != nil {
		panic(err.Error())
	}
	insForm.Exec(music.NamaLagu, music.NamaArtis, music.Link, music.Durasi)
	defer db.Close()
	w.WriteHeader(http.StatusCreated)
	return
}

func UpdateData(w http.ResponseWriter, r *http.Request) {
	music := model.Music{}
	err := json.NewDecoder(r.Body).Decode(&music)
	if err != nil {
		log.Println(err)
		http.Error(w, "json decode fail", http.StatusInternalServerError)
		return
	}
	db := dbConn()
	insForm, err := db.Prepare("UPDATE musik SET namalagu=?, namaartis=? WHERE id=?")
	if err != nil {
		panic(err.Error())
	}
	insForm.Exec(music.NamaLagu, music.NamaArtis, music.ID)
	defer db.Close()
	w.WriteHeader(http.StatusCreated)
	return
}

func LoadData(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	selDB, err := db.Query("SELECT * FROM musik order by id")
	if err != nil {
		log.Println(err)
		http.Error(w, "data retrieval failed", http.StatusInternalServerError)
		return
	}
	emp := model.Music{}
	res := []model.Music{}
	for selDB.Next() {
		var id, durasi int
		var namalagu, namaartis, link string
		err = selDB.Scan(&namalagu, &namaartis, &link, &durasi, &id)
		if err != nil {
			panic(err.Error())
		}
		emp.ID = id
		emp.NamaLagu = namalagu
		emp.NamaArtis = namaartis
		emp.Link = link
		emp.Durasi = durasi
		res = append(res, emp)
		log.Println(res)
	}
	usersJSON, err := json.Marshal(res)
	w.Write(usersJSON)
}
