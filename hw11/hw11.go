package hw11

import (
  "fmt"
  "log"
  "database/sql"
  _ "github.com/mattn/go-sqlite3"
)

//CODE GOES HERE!
type Student struct {
	id   int
	name string
	age  int
	gpa  float64
}

func Insert(db *sql.DB, name string, age int, gpa float64) {
	tx, errTx := db.Begin()
	if errTx != nil {
		log.Fatal(errTx)
	}

	cmd := "INSERT INTO students" +
		"(name,age,gpa)" +
		"values" +
		"(?,?,?)"

	stmt, errStmt := tx.Prepare(cmd)

	if errStmt != nil {

		log.Fatal(errStmt)
	}
	defer stmt.Close()

	stmt.Exec(name, age, gpa)
	tx.Commit()
}

func SelectAll(db *sql.DB) []Student {
	cmd := "SELECT * FROM students"
	var students []Student

	rows, err := db.Query(cmd)
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var id int
		var name string
		var age int
		var gpa float64
		errScan := rows.Scan(&id, &name, &age, &gpa)
		if errScan != nil {
			log.Fatal(errScan)
		}
		s1 := Student{id: id, name: name, age: age, gpa: gpa}
		students = append(students, s1)
	}
	return students
}

func Select(db *sql.DB, column string, operator rune, value any) []Student {
	var students []Student
	rows, rowErr := db.Query("Select * from students Where "+column+" "+string(operator)+" $1", value)
	if rowErr != nil {
		log.Fatal(rowErr)
	}
	for rows.Next() {
		var id int
		var name string
		var age int
		var gpa float64
		errScan := rows.Scan(&id, &name, &age, &gpa)
		if errScan != nil {
			log.Fatal(errScan)
		}
		s1 := Student{id: id, name: name, age: age, gpa: gpa}
		students = append(students, s1)
	}
	return students
}

func (s *Student) Print() {
	fmt.Printf("%d %s %d %.1f\n", s.id, s.name, s.age, s.gpa)
}

func (s *Student) Delete(db *sql.DB) {
	cmd := "DELETE FROM students " +
		"WHERE id = ?"
	db.Exec(cmd, s.id)
}

func (s *Student) Update(db *sql.DB, column string, value any) {
	cmd := "UPDATE students " +
		"SET " + column + "= ?" +
		"WHERE id = ?"
	db.Exec(cmd, value, s.id)
}
