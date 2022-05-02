package post

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
	"tezt/hexagonal/internal/adapters/api"
	"tezt/hexagonal/internal/model"
	"time"
)

type handlerPost struct {
	postService    api.PostService
	sessionService api.SessionService
	commentService api.CommentService
}

func NewHandler(service api.PostService, sessionService api.SessionService, commentService api.CommentService) api.Handler {
	return &handlerPost{
		postService:    service,
		sessionService: sessionService,
		commentService: commentService,
	}
}
func (h *handlerPost) Register(router *http.ServeMux) {
	router.HandleFunc("/", h.home_page)
	router.HandleFunc("/Category/", h.postsByCategory)
	router.HandleFunc("/likedPosts", h.likedPosts)
	router.HandleFunc("/createdPosts", h.createdPosts)
	router.HandleFunc("/addpost", h.addPost)
	router.HandleFunc("/post/", h.postPage)
	router.HandleFunc("/search", h.search)
	router.HandleFunc("/likeNdis", h.LikeDis)
	router.HandleFunc("/commenting", h.Comment)
	router.HandleFunc("/commentLike", h.CommentLikeDis)
	router.HandleFunc("/logout", h.logout)

}
func (h *handlerPost) search(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("./templates/home_page.html", "./templates/header.html")

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	var M model.Forum
	selcection := r.FormValue("search")
	M.Post.Rows, err = h.postService.GetSearch(selcection)
	if err != nil {
		log.Printf("ERROR post handler PostCreate method GetAll function:--> %v\n", err)

		return
	}
	M.Category, err = h.postService.GetCategory()
	if err != nil {
		log.Printf("ERROR post handler PostCreate method GetCategory function:--> %v\n", err)

		return
	}
	if h.CheckSession(w, r) {
		c, _ := r.Cookie("session")
		M.User, err = h.sessionService.ReadByUUID(c.Value)

		temp.Execute(w, M)
	} else {
		temp.Execute(w, M)
	}
}

func (h *handlerPost) home_page(w http.ResponseWriter, r *http.Request) {
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
	M.Post.Rows, err = h.postService.GetAll()
	if err != nil {
		log.Printf("ERROR post handler PostCreate method GetAll function:--> %v\n", err)

		return
	}
	M.Category, err = h.postService.GetCategory()
	if err != nil {
		log.Printf("ERROR post handler PostCreate method GetCategory function:--> %v\n", err)

		return
	}
	if h.CheckSession(w, r) {
		c, _ := r.Cookie("session")
		M.User, err = h.sessionService.ReadByUUID(c.Value)

		temp.Execute(w, M)
	} else {
		temp.Execute(w, M)
	}

}

func (h *handlerPost) postPage(w http.ResponseWriter, r *http.Request) {
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
	M.Category, err = h.postService.GetCategory()
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

	M.Post.Rows, err = h.postService.PostPage(i)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	if h.CheckSession(w, r) {
		c, _ := r.Cookie("session")
		M.User, err = h.sessionService.ReadByUUID(c.Value)
		tmpl.Execute(w, M)
	} else {
		tmpl.Execute(w, M)
	}

}
func (h *handlerPost) postsByCategory(w http.ResponseWriter, r *http.Request) {
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
	M.Category, err = h.postService.GetCategory()
	if err != nil {
		log.Printf("ERROR post handler PostCreate method GetById function:--> %v\n", err)

		return
	}
	title := r.RequestURI[10:]
	M.Post.Rows, err = h.postService.SortedByCategory(title)
	if h.CheckSession(w, r) {
		c, _ := r.Cookie("session")
		M.User, err = h.sessionService.ReadByUUID(c.Value)
		temp.Execute(w, M)
	} else {
		temp.Execute(w, M)
	}

}
func (h *handlerPost) addPost(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("./templates/addpost.html", "./templates/header.html")
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	var M model.Forum
	M.Category, err = h.postService.GetCategory()
	if err != nil {
		log.Printf("ERROR post handler PostCreate method GetById function:--> %v\n", err)

		return
	}
	if h.CheckSession(w, r) {
		c, _ := r.Cookie("session")
		M.User, err = h.sessionService.ReadByUUID(c.Value)
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

		M.Post = h.postService.CreatePost(argument, Caty, n)
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
func (h *handlerPost) likedPosts(w http.ResponseWriter, r *http.Request) {
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
		M.User, err = h.sessionService.ReadByUUID(c.Value)
		M.Category, err = h.postService.GetCategory()
		if err != nil {
			log.Printf("ERROR post handler likedPost method GetCategoryfunction:--> %v\n", err)

			return
		}
		M.Post.Rows, err = h.postService.LikedPosts(M.User.ID)
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
func (h *handlerPost) createdPosts(w http.ResponseWriter, r *http.Request) {
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
		M.User, err = h.sessionService.ReadByUUID(c.Value)

		M.Category, err = h.postService.GetCategory()
		if err != nil {
			log.Printf("ERROR post handler PostCreate method GetById function:--> %v\n", err)

			return
		}
		M.Post.Rows, err = h.postService.CreatedPosts(M.User.ID)
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

func (h *handlerPost) logout(w http.ResponseWriter, r *http.Request) {
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

	m, err := h.sessionService.ReadByUUID(c.Value)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = h.sessionService.Delete(strconv.Itoa(m.ID))
	if err != nil {
		fmt.Println(err)
		return
	}
	http.SetCookie(w, c)
	http.Redirect(w, r, r.Header.Get("Referer"), 303)
	return
}
func (h *handlerPost) CheckSession(w http.ResponseWriter, r *http.Request) bool {
	_, err := r.Cookie("session")
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}
