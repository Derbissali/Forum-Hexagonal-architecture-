package handler

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"tidy/pkg/model"

	"time"
)

func (h *Handlers) search(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("./templates/home_page.html", "./templates/header.html")

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	var M model.Forum
	selcection := r.FormValue("search")
	M.Post.Rows, err = h.services.PostService.GetSearch(selcection)
	if err != nil {
		log.Printf("ERROR post handler PostCreate method GetAll function:--> %v\n", err)

		return
	}
	M.Category, err = h.services.PostService.GetCategory()
	if err != nil {
		log.Printf("ERROR post handler PostCreate method GetCategory function:--> %v\n", err)

		return
	}
	if h.CheckSession(w, r) {
		c, _ := r.Cookie("session")
		M.User, err = h.services.SessionService.ReadByUUID(c.Value)

		temp.Execute(w, M)
	} else {
		temp.Execute(w, M)
	}
}

func (h *Handlers) home_page(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	temp, err := template.ParseFiles("./templates/home_page.html", "./templates/header.html")

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	var M model.Forum
	M.Post.Rows, err = h.services.PostService.GetAll()
	if err != nil {
		log.Printf("ERROR post handler PostCreate method GetAll function:--> %v\n", err)

		return
	}
	M.Category, err = h.services.PostService.GetCategory()
	if err != nil {
		log.Printf("ERROR post handler PostCreate method GetCategory function:--> %v\n", err)

		return
	}
	if h.CheckSession(w, r) {
		c, _ := r.Cookie("session")
		M.User, err = h.services.SessionService.ReadByUUID(c.Value)

		temp.Execute(w, M)
	} else {
		temp.Execute(w, M)
	}

}

func (h *Handlers) postPage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./templates/post_page.html", "./templates/header.html")
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return

	}
	if !strings.HasPrefix(r.URL.Path, "/post/") {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	var M model.Forum
	M.Category, err = h.services.PostService.GetCategory()
	if err != nil {
		log.Printf("ERROR post handler PostCreate method GetById function:--> %v\n", err)

		return
	}

	id := r.RequestURI[6:]
	i, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	// CommentId

	M.Post.Rows, err = h.services.PostService.PostPage(i)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	if h.CheckSession(w, r) {
		c, _ := r.Cookie("session")
		M.User, err = h.services.SessionService.ReadByUUID(c.Value)
		tmpl.Execute(w, M)
	} else {
		tmpl.Execute(w, M)
	}

}
func (h *Handlers) postsByCategory(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("./templates/home_page.html", "./templates/header.html")
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	if !strings.HasPrefix(r.URL.Path, "/Category/") {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	var M model.Forum
	M.Category, err = h.services.PostService.GetCategory()
	if err != nil {
		log.Printf("ERROR post handler PostCreate method GetById function:--> %v\n", err)

		return
	}
	title := r.RequestURI[10:]
	M.Post.Rows, err = h.services.PostService.SortedByCategory(title)
	if h.CheckSession(w, r) {
		c, _ := r.Cookie("session")
		M.User, err = h.services.SessionService.ReadByUUID(c.Value)
		temp.Execute(w, M)
	} else {
		temp.Execute(w, M)
	}

}
func (h *Handlers) addPost(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("./templates/addpost.html", "./templates/header.html")
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	var M model.Forum
	M.Category, err = h.services.PostService.GetCategory()
	if err != nil {
		log.Printf("ERROR post handler PostCreate method GetById function:--> %v\n", err)

		return
	}
	if h.CheckSession(w, r) {
		c, _ := r.Cookie("session")
		M.User, err = h.services.SessionService.ReadByUUID(c.Value)
	} else {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	switch r.Method {
	case "GET":

		if err != nil {
			log.Printf("ERROR post handler PostCreate method SelectCategory function:--> %v\n", err)

			return
		}
		temp.Execute(w, M)
	case "POST":
		r.ParseMultipartForm(0)
		r.ParseForm()
		n := M.User.ID
		fmt.Printf("%+v\n", r.Form)
		Caty := r.Form["category"]
		fmt.Println(Caty, "dddd")
		file, handler, err := r.FormFile("myFile")
		var fileName string
		if err == nil {
			dst, err := os.Create(fmt.Sprintf("assets/temp-images/%d%s", time.Now().UnixNano(), filepath.Ext(handler.Filename)))
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			defer dst.Close()
			fileName = strings.TrimPrefix(dst.Name(), "assets/temp-images/")
			if len(fileName) == 0 {
				fileName = "1"
			}
			fmt.Println(fileName) // Copy the uploaded file to the filesystem
			// at the specified destination
			_, err = io.Copy(dst, file)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			//func() {}
		}
		argument := model.Post{
			Name:  r.FormValue("postN"),
			Body:  r.FormValue("postB"),
			Image: fileName,
		}

		M.Post = h.services.PostService.CreatePost(argument, Caty, n)
		if M.Post.TitBodOver {
			temp.Execute(w, M)
			return
			// http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		}
		if M.Post.TitBodNull {
			temp.Execute(w, M)
			return
		}
		if M.Post.CategoryNull {
			temp.Execute(w, M)
			return
		}
		http.Redirect(w, r, "/", 301)
		return
	}
}
func (h *Handlers) likedPosts(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("./templates/home_page.html", "./templates/header.html")
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	if h.CheckSession(w, r) {
		var M model.Forum
		c, _ := r.Cookie("session")
		M.User, err = h.services.SessionService.ReadByUUID(c.Value)
		M.Category, err = h.services.PostService.GetCategory()
		if err != nil {
			log.Printf("ERROR post handler likedPost method GetCategoryfunction:--> %v\n", err)

			return
		}
		M.Post.Rows, err = h.services.PostService.LikedPosts(M.User.ID)
		if err != nil {
			log.Printf("ERROR post handler PostCreate method ReadByUUID function:--> %v\n", err)

			return
		}
		temp.Execute(w, M)
	} else {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

}
func (h *Handlers) createdPosts(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("./templates/home_page.html", "./templates/header.html")
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	if h.CheckSession(w, r) {
		var M model.Forum
		c, _ := r.Cookie("session")
		M.User, err = h.services.SessionService.ReadByUUID(c.Value)

		M.Category, err = h.services.PostService.GetCategory()
		if err != nil {
			log.Printf("ERROR post handler PostCreate method GetById function:--> %v\n", err)

			return
		}
		M.Post.Rows, err = h.services.PostService.CreatedPosts(M.User.ID)
		if err != nil {
			log.Printf("ERROR post handler PostCreate method GetById function:--> %v\n", err)

			return
		}
		temp.Execute(w, M)
	} else {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

}

func (h *Handlers) logout(w http.ResponseWriter, r *http.Request) {
	_, err := template.ParseFiles("./templates/home_page.html")
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	c, err := r.Cookie("session")
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	c.Name = "session"
	c.MaxAge = -1
	c.HttpOnly = true

	m, err := h.services.SessionService.ReadByUUID(c.Value)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = h.services.SessionService.Delete(strconv.Itoa(m.ID))
	if err != nil {
		fmt.Println(err)
		return
	}
	http.SetCookie(w, c)
	http.Redirect(w, r, r.Header.Get("Referer"), 303)
	return
}
func (h *Handlers) CheckSession(w http.ResponseWriter, r *http.Request) bool {
	c, err := r.Cookie("session")
	if err != nil {
		return false
	}
	var M model.User
	M, _ = h.services.SessionService.ReadByUUID(c.Value)
	if !M.Session {
		fmt.Println("kirdi")
		c.Name = "session"
		c.MaxAge = -1
		c.HttpOnly = true
		http.SetCookie(w, c)
		return false
	}
	return true
}
