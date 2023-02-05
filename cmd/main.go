package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("./testTemplates/*"))
}

func main() {
	http.HandleFunc("/login", loginFunc)
	// http.HandleFunc("/welcome", cookies.Welcome)
	// http.HandleFunc("/welcome", cookies.Refresh)
	// http.HandleFunc("/welcome", cookies.Logout)
	// start the server on port 8080
	log.Fatal(http.ListenAndServe(":8080", nil))

}

func loginFunc(res http.ResponseWriter, req *http.Request) {
	//Login Template
	fmt.Println(tpl)
	if req.Method == "GET" {
		fmt.Println("Loggin test")

		err := tpl.ExecuteTemplate(res, "Login.gohtml", "")

		if err != nil {
			log.Fatalln(err)
		}

		// err := userLoggedInTpl.ExecuteTemplate(res, "MainUserPage.gohtml", &appointmentsData)
		// checkLoggedIn(res, req, currentUser)

		// if err != nil {
		// 	log.Fatalln(err)
		// 	return
		// }
	}
}
