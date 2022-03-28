package post

import (
	"html/template"
	"net/http"
	"tezt/hexagonal/internal/model"
)

func (h *handlerPost) Comment(w http.ResponseWriter, r *http.Request) {
	_, err := template.ParseFiles("templates/post_page.html")
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	if r.Method != http.MethodPost {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	var M model.Forum

	id := r.FormValue("idwka")

	// CommentId := r.FormValue("comIdd")
	// fmt.Println(CommentId)

	if h.CheckSession(w, r) {
		c, _ := r.Cookie("session")
		M.User, err = h.sessionService.ReadByUUID(c.Value)
	} else {
		http.Redirect(w, r, "/signin", 301)
		return
	}

	n := M.User.ID

	comment := r.FormValue("comment")
	err = h.commentService.AddComment(comment, id, n)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	http.Redirect(w, r, r.Header.Get("Referer"), 301)
	return
}
