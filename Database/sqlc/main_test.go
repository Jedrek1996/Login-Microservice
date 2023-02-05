package db

import (
	"Microservice-Login/util"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

//If login connect to database
//Authenticsation - https://www.sohamkamani.com/golang/password-authentication-and-storage/
//Constant check for session cookies

//Circuit breaker - https://levelup.gitconnected.com/circuit-breaker-example-in-golang-e6459c87eaeb

//Seperate database connection? Queries seperate
//Sqlc to generate query codes

//check if connected

var testQueries *Queries

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "1234"
	dbname   = "first_db"
)

type userCreds struct {
	Username string `json:"password", db:"password"`
	Password string `json:"username", db:"username"`
}

func init() {
	//Test template for checking login
	//tpl = template.Must(template.ParseGlob("./TestTemplates/*"))
}

func main() {
	//http.ListenAndServe(":8000", nil) //Start port on 8080

	// http.HandleFunc("/", MainPageFunc())

	fmt.Println("Login Page")

	fmt.Println(util.RandomFirstNameGenerator())
	fmt.Println(util.RandomLastNameGenerator())
	fmt.Println(util.RandomEmailGenerator())
	fmt.Println(util.RandomMobileGenerator())
	fmt.Println(util.RandomUsernameGenerator())

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
