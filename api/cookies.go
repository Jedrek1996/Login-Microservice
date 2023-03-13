package api

import (
	"context"
	"fmt"
	"math/big"
	"net/http"
	"strconv"
	"time"

	db "Microservice-Login/database/sqlc"

	uuid "github.com/satori/go.uuid"
)

func (s *Server) SetCookie(w http.ResponseWriter, userDetail db.UserDetail, duration time.Duration) {

	expires := time.Now().Add(duration)
	cookieID := getUUIDInt32()
	cookie := http.Cookie{
		Name:    userDetail.UserName,
		Value:   strconv.Itoa(int(cookieID)),
		Expires: expires,
	}
	http.SetCookie(w, &cookie)

	arg := db.InsertCookieParams{
		UserName:  userDetail.UserName,
		CookieID:  cookieID,
		ExpiresAt: expires,
	}

	userCookie, err := s.store.InsertCookie(context.Background(), arg)
	if err != nil {
		fmt.Println("Cannot set cookie")
		fmt.Println(arg)
		fmt.Println(userCookie)
		fmt.Println(err)
		return
	}

	// JED TO CONTINUE WIT COOKIES WORKFLOW. NOW shld be okay
}

func GetCookie(r *http.Request, name string) (*http.Cookie, error) {
	cookie, err := r.Cookie(name)
	if err != nil {
		return nil, err
	}
	return cookie, nil
}

func ClearCookie(w http.ResponseWriter, name string) {
	expires := time.Unix(0, 0)
	cookie := http.Cookie{
		Name:    name,
		Value:   "",
		Expires: expires,
	}
	http.SetCookie(w, &cookie)
}

func getUUIDInt32() int32 {
	cookieValue := uuid.NewV4()

	// Convert UUID to int32
	idBytes := cookieValue.Bytes()
	idInt := big.NewInt(0)
	idInt.SetBytes(idBytes)
	idInt32 := int32(idInt.Uint64())

	return idInt32
}
