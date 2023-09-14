package main

import (
	"database/sql"
	"errors"
	"fmt"
	_ "go-sqlite3"
	"hw11"
	"log"
	"math/rand"
	"os"
	"strconv"
)

func main() {
	var students []hw11.Student

	//check cmd line options
	if len(os.Args) != 3 {
		log.Fatal("usage: ./hw08-main N seed")
	}

	//get options, seed random
	N, seed := getCmdOptions(os.Args)
	rand.Seed(seed)

	//delete all rows from table
	deleteAllStudents()

	//insert N random students
	insertStudents(N)

	//get all students, print
	fmt.Println("\n=== ALL ===")
	printAllStudents()

	//print all students with age > 20
	fmt.Println("\n=== age > 20 ===")
	students = _select("age", '>', 20)
	printStudents(students)

	//print all students with gpa < 2.1
	fmt.Println("\n=== gpa < 2.1 ===")
	students = _select("gpa", '<', 2.1)
	printStudents(students)

	//randomly ignore, delete, or update
	students = selectAllStudents()
	for _, s := range students {
		switch rand.Intn(3) {
		case 0:
			continue
		case 1:
			_delete(s)
		case 2:
			update(s)
		}
	}

	//print all records in students table
	fmt.Println("\n=== ALL: AFTER MODIFYING ===")
	printAllStudents()

}

//error checking
func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

//helper function to connect to local sqlite database
func connect() *sql.DB {
	db, err := sql.Open("sqlite3", "./hw11.sqlite")
	check(err)
	return db
}

//wrapper to delete a specific student
func _delete(s hw11.Student) {
	db := connect()
	s.Delete(db)
	db.Close()
}

//delete all students from database table
func deleteAllStudents() {
	db := connect()
	students := selectAllStudents()
	for _, s := range students {
		s.Delete(db)
	}
	db.Close()
}

func getCmdOptions(args []string) (int, int64) {
	N, err := strconv.Atoi(args[1])
	check(err)
	seed, err := strconv.ParseInt(args[2], 10, 64)
	check(err)
	if N <= 0 {
		check(errors.New("N must be greater than 0"))
	}
	return N, seed
}

//insert N random students into table
func insertStudents(N int) {
	names := []string{"Olivia", "Emma", "Amelia", "Ava", "Sophia",
		"Isabella", "Mia", "Luna", "Charlotte", "Harper",
		"Noah", "Liam", "Oliver", "Elijah", "Lucas",
		"Mateo", "Gold", "Levi", "Ethan", "Asher"}

	db := connect()
	for i := 0; i < N; i++ {
		index := rand.Intn(len(names))
		name := names[index]
		age := rand.Intn(25-17) + 17
		gpa := float64(rand.Intn(3)) + rand.Float64()
		hw11.Insert(db, name, age, gpa)
	}
	db.Close()
}

//print all students
func printAllStudents() {
	students := selectAllStudents()
	printStudents(students)
}

//helper function to print slice of students
func printStudents(students []hw11.Student) {
	for _, s := range students {
		s.Print()
	}
}

//wrapper function to select specific students
func _select(column string, op rune, value interface{}) []hw11.Student {
	db := connect()
	defer db.Close()
	return hw11.Select(db, column, op, value)
}

//get all students from database table
func selectAllStudents() []hw11.Student {
	db := connect()
	defer db.Close()
	return hw11.SelectAll(db)
}

//update random column in Student
func update(s hw11.Student) {
	column := []string{"name", "age", "gpa"}[rand.Intn(3)]
	var value interface{}
	switch column {
	case "name":
		value = "FFFFFFFF"
	case "age":
		value = 8675309
	case "gpa":
		value = 867.5309
	}
	db := connect()
	s.Update(db, column, value)
	db.Close()
}
