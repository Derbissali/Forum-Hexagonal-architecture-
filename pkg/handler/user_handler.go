package handler

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"tidy/pkg/model"

	uuid "github.com/satori/go.uuid"
)

func (h *Handlers) Signup(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("./templates/signup.html")

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
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
		cred, err := h.services.UserService.Create(creds)
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

func (h *Handlers) Signin(w http.ResponseWriter, r *http.Request) {

	temp, err := template.ParseFiles("./templates/login.html")

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
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
		cred, err := h.services.UserService.Check(creds)
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
func (h *Handlers) CreateSession(w http.ResponseWriter, r *http.Request) {

	email := r.FormValue("email")

	// c, err := r.Cookie("session")

	sID := uuid.NewV4()
	c := &http.Cookie{
		Name:     "session",
		Value:    sID.String(),
		HttpOnly: true,
	}
	http.SetCookie(w, c)
	id, err := h.services.UserService.GetIDbyName(email)

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(c.Value)
	err = h.services.SessionService.Create(c.Value, id)

	if err != nil {
		h.services.SessionService.Delete(id)
		h.services.SessionService.Create(c.Value, id)
		return
	}

	return
}
