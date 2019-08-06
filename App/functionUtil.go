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

}
