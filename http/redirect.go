package userHttp

import "net/http"

func redirectLogin(res http.ResponseWriter, req *http.Request) {
	http.Redirect(res, req, "/", http.StatusSeeOther)
}
