package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
)

type artist struct {
	gorm.Model
	id          int
	name        string
	description string
	image_url   string
	rating      float64
}

/*
func initDBGorm() {
	var err error
	db, err := gorm.Open("postgres", "host=grandline.clclpmbitplj.ap-southeast-1.rds.amazonaws.com user=kaidou dbname=grandline_app sslmode=disable password=Seizetheday24$ connect_timeout=5")
	if err != nil {
		panic(err)
	}
	//db.LogMode(true)
	log.Printf("Connected")
	defer db.Close()

	db.AutoMigrate(&artist{})

	var art artist
	rows := db.First(&art, 1)
	//fmt.Fprintln(rows)
	log.Printf("DB ok!")
	//handleRows(rows, err)
}
*/
func initDB() {
	var err error
	now := time.Now()
	defer func() {
		log.Printf("Exited after %v", time.Since(now))
	}()
	log.Printf("Opening...")
	db, err := sql.Open("postgres", "host=grandline.clclpmbitplj.ap-southeast-1.rds.amazonaws.com user=kaidou dbname=grandline_app sslmode=disable password=grandline connect_timeout=5")
	defer db.Close()
	if err != nil {
		log.Print(err)
		return
	}
	log.Printf("Opened, pinging...")
	err = db.Ping()
	if err != nil {
		log.Print(err)
		return
	}
	log.Printf("Ping ok!")

	rows, err := db.Query("select * from artist where rating > $1", 4)
	handleRows(rows, err)
}
func handleRows(rows *sql.Rows, err error) {
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	artists := []artist{}
	for rows.Next() {
		a := artist{}
		err := rows.Scan(&a.id, &a.name, &a.description, &a.image_url, &a.rating)
		if err != nil {
			log.Println(err)
			continue
		}
		artists = append(artists, a)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(artists)
}
func main() {
	//db, err := sql.Open("postgres", "postgres://kaidou:Seizetheday24$@grandline.clclpmbitplj.ap-southeast-1.rds.amazonaws.com:5432/grandline_app?sslmode=verify-full")
	initDB()
	//initDBGorm()
	/*
	  http.HandleFunc("/",func(w http.ResponseWriter, r *http.Request){

	      fmt.Fprintf(w, "Hello Web Development!")
	  })
	  fmt.Println(http.ListenAndServe(":8000",nil))
	*/
}
