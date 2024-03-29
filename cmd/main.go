package main

import (
	"context"
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"

	api "Microservice-Login/api"
	db "Microservice-Login/database/sqlc"
	util "Microservice-Login/util"

	_ "github.com/lib/pq"
)

var tpl *template.Template
var queries *db.Queries
var newDB *sql.DB

const (
	dbDriver      = "postgres"
	dbSource      = "postgresql://root:secret@localhost:5430/loginMicroservice9?sslmode=disable"
	serverAddress = "0.0.0.0:8080"
)

func init() {
	tpl = template.Must(template.ParseGlob("./templates/testTemplates/*"))

	var err error
	newDB, err = sql.Open(dbDriver, dbSource)

	if err != nil {
		log.Fatal("Error connecting to db:", err)
	}

	store := db.NewStore(newDB)
	server := api.NewServer(store)

	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("Error connecting to server:", server)
	}

	db.New(newDB)

	fmt.Println("Database conncted")
}

func main() {

	// r := gin.Default()
	// config := cors.DefaultConfig()
	// config.AllowOrigins = []string{"http://localhost:3000"}
	// r.Use(cors.New(config))
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

func getConfig() *api.AppConfiguration {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	expireSec, err := strconv.Atoi(os.Getenv("TOKEN_EXPIRE_SECS"))
	if err != nil {
		log.Fatal(err)
	}

	isRunOnHost, err := strconv.ParseBool(os.Getenv("Run_ON_HOST"))
	if err != nil {
		log.Fatal(err)
	}

	return &api.AppConfiguration{
		RunOnHost:                        isRunOnHost,
		TokenExpireSecs:                  expireSec,
		ServicePort:                      os.Getenv("SERVICE_PORT"),
		EmailServiceContainerName:        os.Getenv("MAIL_SERVICE"),
		LoginServiceContainerName:        os.Getenv("LOGIN_SERVICE"),
		PlaylistServiceContainerName:     os.Getenv("PLATLIST_SERVICE"),
		SubscriptionServiceContainerName: os.Getenv("SUBSCRIPTION_SERVICE"),
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
