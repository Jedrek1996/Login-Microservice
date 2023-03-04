package cookies

import (
	"fmt"
	"net/http"

	uuid "github.com/satori/go.uuid"
)

func setCookie(res http.ResponseWriter, req *http.Request, currentUser string) {
	fmt.Println("Setting Cookie!")
	id := uuid.NewV4()

	http.SetCookie(res, &http.Cookie{
		Name:  currentUser,
		Value: id.String(),
	})
}
