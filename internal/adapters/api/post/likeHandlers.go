package post

import (
	"html/template"
	"net/http"
	"tezt/hexagonal/internal/model"
)

func (h *handlerPost) LikeDis(w http.ResponseWriter, r *http.Request) {
	_, err := template.ParseFiles("templates/post_page.html", "./templates/header.html")
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	var M model.Forum
	if r.Method != http.MethodPost {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return

	}
	idPost := r.FormValue("postId")

	if h.CheckSession(w, r) {
		c, _ := r.Cookie("session")
		M.User, err = h.sessionService.ReadByUUID(c.Value)
	} else {
		http.Redirect(w, r, "/signin", 301)
		return
	}
	n := M.User.ID
	l := r.FormValue("like")
	d := r.FormValue("dislike")

	h.postService.PostLike(l, d, idPost, n)

	http.Redirect(w, r, r.Header.Get("Referer"), 301)
	return
}
func (h *handlerPost) CommentLikeDis(w http.ResponseWriter, r *http.Request) {
	_, err := template.ParseFiles("templates/post_page.html", "./templates/header.html")
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	var M model.Forum
	if r.Method != http.MethodPost {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return

	}
	idPost := r.FormValue("postId")

	idComment := r.FormValue("comId")

	if h.CheckSession(w, r) {
		c, _ := r.Cookie("session")
		M.User, err = h.sessionService.ReadByUUID(c.Value)
	} else {
		http.Redirect(w, r, "/signin", 301)
		return
	}
	n := M.User.ID
	l := r.FormValue("commnetLike")
	d := r.FormValue("commentDislike")

	h.postService.CommentLike(l, d, idPost, idComment, n)

	http.Redirect(w, r, r.Header.Get("Referer"), 301)
	return
}
