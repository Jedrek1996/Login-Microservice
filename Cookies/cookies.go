package cookies

import (
	"fmt"
	"net/http"

	uuid "github.com/satori/go.uuid"
)

type cookieDetails struct {
	Name  string
	Value string
}

func setCookie(res http.ResponseWriter, req *http.Request, currentUser string) {
	fmt.Println("Setting Cookie!")
	id := uuid.NewV4()

	http.SetCookie(res, &http.Cookie{
		Name:  currentUser,
		Value: id.String(),
	})
}
