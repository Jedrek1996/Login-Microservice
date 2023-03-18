package main

import (
	"context"
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	"Microservice-Login/api"
	db "Microservice-Login/database/sqlc"
	util "Microservice-Login/util"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var tpl *template.Template
var queries *db.Queries
var newDB *sql.DB

const (
	dbDriver      = "postgres"
	serverAddress = ":8080"
)

func main() {

	tpl = template.Must(template.ParseGlob("./templates/testTemplates/*"))

	dbSource := ""
	if isRunOnHost() {
		dbSource = "postgresql://root:secret@localhost:5430/loginMicroservice9?sslmode=disable"
	} else {
		dbSource = "postgresql://root:secret@postgres12:5432/loginMicroservice9?sslmode=disable"
	}
	var err error
	newDB, err = sql.Open(dbDriver, dbSource)

	if err != nil {
		log.Fatal("Error connecting to db:", err)
	}

	store := db.NewStore(newDB)
	server := api.NewServer(store)

	err = server.Start(serverAddress)
	if err != nil {
		fmt.Println("hit here")
		log.Fatal("Error connecting to server:", server)
	}

	store.Queries = db.New(newDB)

	//CSS Files for testing
	queries = db.New(newDB)
	fs := http.FileServer(http.Dir("./templates"))
	http.Handle("/templates/", http.StripPrefix("/templates/", fs))

	http.HandleFunc("/", loginTest)
	http.HandleFunc("/signup", signUpTest)
	// http.HandleFunc("/welcome", cookies.Welcome)
	// http.HandleFunc("/welcome", cookies.Refresh)
	// http.HandleFunc("/welcome", cookies.Logout)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func isRunOnHost() bool {
	fmt.Println(goDotEnvVariable("RUN_ON_HOST"))
	if goDotEnvVariable("Run_ON_HOST") == "true" {
		return true
	} else {
		return false
	}
}

// Login
func loginTest(res http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		fmt.Println("Loggin test")

		err := tpl.ExecuteTemplate(res, "Login.gohtml", "")

		if err != nil {
			log.Fatalln(err)
		}
	} else if req.Method == "POST" {
		userName := req.FormValue("loginUsername")
		userPassword := req.FormValue("loginPassword")

		fmt.Println(userName)
		fmt.Println(userPassword)

	}
}

// Signup
func signUpTest(res http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		fmt.Println("Signup test")

		err := tpl.ExecuteTemplate(res, "Signup.gohtml", "")

		if err != nil {
			log.Fatalln(err)
		}
	} else if req.Method == "POST" {
		signupFirstName := req.FormValue("signupFirstName")
		signupLastName := req.FormValue("signupLastName")
		signupUsername := req.FormValue("signupUsername")
		signuPassword := req.FormValue("signuPassword")
		signupEmail := req.FormValue("signupEmail")
		signupMobile := req.FormValue("signupMobile")
		//Validate credentials and store them in database
		fmt.Println(signupFirstName, signupLastName, signupUsername, signuPassword, signupEmail, signupMobile)
		//Change this
		arg := db.CreateUserParams{
			FirstName: util.RandomFirstNameGenerator(),
			LastName:  util.RandomLastNameGenerator(),
			UserName:  util.RandomUsernameGenerator(),
			Email:     util.RandomEmailGenerator(),
			Mobile:    util.RandomMobileGenerator(),
		}

		fmt.Println(arg)

		userDetails, err := queries.CreateUser(context.Background(), arg)

		if err != nil {
			fmt.Println(err, userDetails)
		}
	}
}

// use godot package to load/read the .env file and
// return the value of the key
func goDotEnvVariable(key string) string {

	// load .env file which is located at the root path
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}
