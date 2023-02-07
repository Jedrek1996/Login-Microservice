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
	//Test template for checking login
	//tpl = template.Must(template.ParseGlob("./TestTemplates/*"))
}

func TestMain(m *testing.M) {
	//http.ListenAndServe(":8000", nil) //Start port on 8080

	// http.HandleFunc("/", MainPageFunc())

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

// func checkUserDatabase

// func MainPageFunc(res http.ResponseWriter, req *http.Request) {

// 	if req.Method == "GET" {
// 		err := tpl.ExecuteTemplate(res, "MainLoginPage.gohtml")

// 		if err != nil {
// 			log.Fatalln(err)
// 		}
// 	} else if req.Method == "POST" {
// 	}
// }
