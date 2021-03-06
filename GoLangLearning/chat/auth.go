package main

import (
	"fmt"
	"github.com/stretchr/gomniauth"
	"github.com/stretchr/objx"
	"log"
	"net/http"
	"strings"
	"crypto/md5"
	"io"
)

type authHandler struct {
	next http.Handler
}

func (h *authHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if cookie, err := r.Cookie("auth"); err == http.ErrNoCookie || cookie.Value == "" {
		// No Authenticated
		w.Header().Set("Location", "/login")
		w.WriteHeader(http.StatusTemporaryRedirect)
	} else if err != nil {
		panic(err.Error())
	} else {
		h.next.ServeHTTP(w, r)
	}
}

func MustAuth(handler http.Handler) http.Handler {
	return &authHandler{next: handler}
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	segs := strings.Split(r.URL.Path, "/")
	action := segs[2]
	provider := segs[3]
	switch action {
	case "login":

		provider, err := gomniauth.Provider(provider)
		if err != nil {
			log.Fatalln("Error when trying to get provider", provider, "-", err)
		}

		loginUrl, err := provider.GetBeginAuthURL(nil, nil)
		if err != nil {
			log.Fatalln("Error when trying to GetBeginAuthURL for", provider, "-", err)
		}

		w.Header()["Location"] = []string{loginUrl}
		w.WriteHeader(http.StatusTemporaryRedirect)

	case "callback":

		provider, err := gomniauth.Provider(provider)
		if err != nil {
			log.Fatalln("Error when trying to get provider", provider, "-", err)
		}

		// get the credentials
		creds, err := provider.CompleteAuth(objx.MustFromURLQuery(r.URL.RawQuery))
		if err != nil {
			log.Fatalln("Error when trying to complete auth for", provider, "-", err)
		}
		user, err := provider.GetUser(creds)
		
		if err != nil {
			log.Fatalln("Error when trying to get user from", provider, "-", err)
		}
		m := md5.New()

		io.WriteString(m, strings.ToLower(user.Name()))
		userId := fmt.Sprintf("%x", m.Sum(nil))

		// get the user
		
		
		authCookieValue := objx.New(map[string]interface{}{
			"userid"		: 		userId,
			"name"			: 		user.Name(),
			"avatar_url" 	:		user.AvatarURL(),
			"email"			:		user.Email(),
		}).MustBase64()

		// save some data
		/*authCookieValue := objx.New(map[string]interface{}{
			"name"			: 		user.Name(),
			"avatar_url" 	:		user.AvatarURL(),
			"email"			:		user.Email(),
		}).MustBase64()*/

		http.SetCookie(w, &http.Cookie{
			Name:  "auth",
			Value: authCookieValue,
			Path:  "/"})

		w.Header().Set("Location", "/chat")
		w.WriteHeader(http.StatusTemporaryRedirect)

	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Auth action %s not supported", action)
	}
}
