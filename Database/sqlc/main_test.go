package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

var testQueries *Queries
var testDB *sql.DB

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "1234"
	dbname   = "first_db"
	dbDriver = "postgres"
	dbSource = "postgresql://root:secret@localhost:5430/FoodPanda9?sslmode=disable"
)

type userCreds struct {
	Username string `json:"password", db:"password"`
	Password string `json:"username", db:"username"`
}

func init() {

}

func TestMain(m *testing.M) {

	var err error

	testDB, err = sql.Open(dbDriver, dbSource)

	if err != nil {
		log.Fatal("Error connecting to db:", err)
	}

	testQueries = New(testDB)

	os.Exit(m.Run())

	fmt.Println(testQueries)

	//dbConnection()
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func dbConnection() {

	// connection string
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// open database
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)

	// close database
	defer db.Close()

	// check db
	err = db.Ping()
	CheckError(err)

	fmt.Println("Connected!")

	insertTest := `insert into "Employee" ("Name", "EmpId") values('Jed',123)`

	_, insertErr := db.Exec(insertTest)

	CheckError(insertErr)
}
