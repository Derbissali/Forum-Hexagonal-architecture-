package handler

import (
	"net/http"
	"tidy/pkg/service"
)

type Handlers struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handlers {
	return &Handlers{services: services}
}

func (h *Handlers) Register(router *http.ServeMux) {
	router.HandleFunc("/", h.home_page)
	router.HandleFunc("/signup", h.Signup)
	router.HandleFunc("/signin", h.Signin)
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
