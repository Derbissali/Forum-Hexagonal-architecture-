package user

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
	"tezt/hexagonal/internal/adapters/api"
	"tezt/hexagonal/internal/model"

	uuid "github.com/satori/go.uuid"
)

type handlerUser struct {
	userService    api.UserService
	sessionService api.SessionService
}

func NewHandler(service api.UserService, sessionService api.SessionService) api.Handler {
	return &handlerUser{
		userService:    service,
		sessionService: sessionService,
	}
}
func (h *handlerUser) Register(router *http.ServeMux) {
	router.HandleFunc("/signup", h.Signup)
	router.HandleFunc("/signin", h.Signin)

}

func (h *handlerUser) Signup(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("./templates/signup.html")

	if err != nil {
		log.Printf("Error signup html User Handler GetAll method:--> %v\n", err)

		return
	}
	creds := &model.User{}
	switch r.Method {
	case "GET":
		temp.Execute(w, creds)
	case "POST":
		creds = &model.User{
			Name:     r.FormValue("login"),
			Email:    r.FormValue("email"),
			Password: r.FormValue("password"),
		}
		cred, err := h.userService.Create(creds)
		if err != nil {
			log.Printf("ERROR post handler PostCreate method GetById function:--> %v\n", err)

			return
		}
		if cred.ErrorEmpty {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Println("empty")
			temp.Execute(w, creds)
			return
		}
		if cred.ErrorE {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Println("email")
			temp.Execute(w, creds)
			return
		}
		if cred.ErrorEm {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Println("notUniq")
			temp.Execute(w, creds)
			return
		}
		http.Redirect(w, r, "/", 301)

	default:
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
}

func (h *handlerUser) Signin(w http.ResponseWriter, r *http.Request) {

	temp, err := template.ParseFiles("./templates/login.html")

	if err != nil {
		log.Printf("Error signup html User Handler GetAll method:--> %v\n", err)

		return
	}

	creds := &model.User{}
	_, err = r.Cookie("session")
	if err == nil {
		log.Println("signin")
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	switch r.Method {
	case "GET":
		temp.Execute(w, creds)
	case "POST":
		creds = &model.User{
			Email:    r.FormValue("email"),
			Password: r.FormValue("password"),
		}
		cred, err := h.userService.Check(creds)
		if err != nil {
			log.Printf("ERROR user handler UserCreate method GetById function:--> %v\n", err)

			return
		}
		if cred.ErrorEm {
			fmt.Println("wrong login")
			temp.Execute(w, cred)
			return
		}
		h.CreateSession(w, r)

		http.Redirect(w, r, "/", 301)

	default:
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
}
func (h *handlerUser) CreateSession(w http.ResponseWriter, r *http.Request) {

	email := r.FormValue("email")

	// c, err := r.Cookie("session")

	sID := uuid.NewV4()
	c := &http.Cookie{
		Name:     "session",
		Value:    sID.String(),
		HttpOnly: true,
	}
	http.SetCookie(w, c)
	id, err := h.userService.GetIDbyName(email)

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(c.Value)
	err = h.sessionService.Create(c.Value, id)

	if err != nil {
		h.sessionService.Delete(id)
		h.sessionService.Create(c.Value, id)
		return
	}

	return
}
