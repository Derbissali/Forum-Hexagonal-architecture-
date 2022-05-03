package handler

import (
	"html/template"
	"net/http"
	"tidy/pkg/model"
)

func (h *Handlers) LikeDis(w http.ResponseWriter, r *http.Request) {
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
		M.User, err = h.services.SessionService.ReadByUUID(c.Value)
	} else {
		http.Redirect(w, r, "/signin", 301)
		return
	}
	n := M.User.ID
	l := r.FormValue("like")
	d := r.FormValue("dislike")

	h.services.LikeService.PostLike(l, d, idPost, n)

	http.Redirect(w, r, r.Header.Get("Referer"), 301)
	return
}
func (h *Handlers) CommentLikeDis(w http.ResponseWriter, r *http.Request) {
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

	if h.CheckSession(w, r) {
		idPost := r.FormValue("postId")

		idComment := r.FormValue("comId")
		c, _ := r.Cookie("session")
		M.User, err = h.services.SessionService.ReadByUUID(c.Value)
		n := M.User.ID
		l := r.FormValue("commnetLike")
		d := r.FormValue("commentDislike")

		h.services.LikeService.CommentLike(l, d, idPost, idComment, n)

		http.Redirect(w, r, r.Header.Get("Referer"), 301)
	} else {
		http.Redirect(w, r, "/signin", 301)
		return
	}

}
