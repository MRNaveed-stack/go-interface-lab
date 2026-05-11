package main

import "fmt"

type Storer interface {
	Save(data string) error
}

type Database struct {
	Connection string
}

func (db Database) Save(data string) error {
	fmt.Printf("SQL:INSERT INTO logs ('%s') via %s\n", data, db.Connection)
	return nil
}

type FileSystem struct {
	Path string
}

func (fs FileSystem) Save(data string) error {
	fmt.Printf("I/O: writing %s to the file %s\n", data, fs.Path)
	return nil
}

func ProcessData(s Storer, payload string) {
	fmt.Println("Architectural logic: starting to process")
	err := s.Save(payload)
	if err != nil {
		fmt.Println("Error saving:", err)
		return
	}
	fmt.Println("Architectural logic: success")

}

func main() {
	db := Database{Connection: "postgres://localhost:5432"}
	fs := FileSystem{Path: "/var/log/app.log"}

	ProcessData(db, "User logged in")
	ProcessData(fs, "Kernel panic trace")
}
