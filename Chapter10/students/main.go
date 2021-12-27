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

func getStudents(w http.ResponseWriter, r *http.Request) {
	studentsInfo, err := models.GetStudents()
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	for _, std := range studentsInfo {
		//s, err := json.Marshal(std.FirstName, std.MiddleName.std.LastName, std.Class, std.Course)
		s, err := json.Marshal(std)
		if err != nil {
			log.Fatal(err)
		}
		//fmt.Fprintf(w, "%s\n%s\n%s\n%s\n%s\n", std.FirstName, std.MiddleName, std.LastName, std.Class, std.Course)
		fmt.Fprintf(w, "%s", string(s))
	}
}
