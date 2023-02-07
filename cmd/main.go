package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("./templates/testTemplates/*"))
}

func main() {
	//CSS Files for testing
	fs := http.FileServer(http.Dir("./templates"))
	http.Handle("/templates/", http.StripPrefix("/templates/", fs))

	http.HandleFunc("/", loginTest)
	// http.HandleFunc("/welcome", cookies.Welcome)
	// http.HandleFunc("/welcome", cookies.Refresh)
	// http.HandleFunc("/welcome", cookies.Logout)
	// start the server on port 8080
	log.Fatal(http.ListenAndServe(":8080", nil))

}

func loginTest(res http.ResponseWriter, req *http.Request) {
	//Login Template
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
