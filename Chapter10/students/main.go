package main

import (
	"database/sql"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/PacktPublishing/Podman-for-DevOps/Chapter10/students/models"

	_ "github.com/lib/pq"
)

func main() {

	var (
		username string
		password string
		host     string
		port     string
		database string
		usage    string
	)

	flag.StringVar(&username, "username", "admin", "Default database username")
	flag.StringVar(&password, "password", "password", "Default database password")
	flag.StringVar(&host, "host", "localhost", "Default host running the database")
	flag.StringVar(&port, "port", "5432", "Default database port")
	flag.StringVar(&database, "database", "students", "Default application database")
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), usage, os.Args[0])
		flag.PrintDefaults()
	}
	flag.Parse()

	connString := "postgres://" + username + ":" + password + "@" + host + ":" + port + "/" + database + "?sslmode=disable"
	log.Printf("Connecting to host %s:%s, database %s", host, port, database)

	var err error
	models.DB, err = sql.Open("postgres", connString)
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/students", getStudents)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}

// getStudents is a HTTP GET method to print a full list of the students in JSON encoding
func getStudents(w http.ResponseWriter, r *http.Request) {
	studentsInfo, err := models.GetStudents()
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	for _, std := range studentsInfo {
		s, err := json.Marshal(std)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Fprintf(w, "%s", string(s))
	}
}
