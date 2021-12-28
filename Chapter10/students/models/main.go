package models

import (
	"database/sql"
	"strings"
)

var DB *sql.DB

type Student struct {
	Id         int
	FirstName  string
	MiddleName string
	LastName   string
	Class      string
	Course     string
}

func GetStudents() ([]Student, error) {
	rows, err := DB.Query("SELECT * from students")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var students []Student

	for rows.Next() {
		var student Student
		err := rows.Scan(&student.Id, &student.FirstName, &student.MiddleName, &student.LastName, &student.Class, &student.Course)
		if err != nil {
			return nil, err
		}

		// Trim uncessary spaces from string fields
		student.FirstName = strings.TrimSpace(student.FirstName)
		student.MiddleName = strings.TrimSpace(student.MiddleName)
		student.LastName = strings.TrimSpace(student.LastName)
		student.Class = strings.TrimSpace(student.Class)
		student.Course = strings.TrimSpace(student.Course)

		// Append the struct to the slice of students
		students = append(students, student)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return students, nil
}
