package config

import (
	"database/sql"
	"net/http"
	post3 "tezt/hexagonal/internal/adapters/api/post"
	user3 "tezt/hexagonal/internal/adapters/api/user"
	"tezt/hexagonal/internal/adapters/db/comment"
	"tezt/hexagonal/internal/adapters/db/like"
	"tezt/hexagonal/internal/adapters/db/post"
	"tezt/hexagonal/internal/adapters/db/session"
	"tezt/hexagonal/internal/adapters/db/user"
	comment2 "tezt/hexagonal/internal/domain/comment"
	like2 "tezt/hexagonal/internal/domain/like"
	post2 "tezt/hexagonal/internal/domain/post"
	session2 "tezt/hexagonal/internal/domain/session"
	user2 "tezt/hexagonal/internal/domain/user"
)

func Config(db *sql.DB) *http.ServeMux {

	router := http.NewServeMux()
	postStorage := post.NewStorage(db)
	likeStorage := like.NewLikeStorage(db)
	likeService := like2.NewService(likeStorage)
	sessionStorage := session.NewStorage(db)
	sessionService := session2.NewService(sessionStorage)
	commentStorage := comment.NewStorage(db)
	commentService := comment2.NewService(commentStorage)
	postService := post2.NewService(postStorage, likeService)
	postHandler := post3.NewHandler(postService, sessionService, commentService)

	postHandler.Register(router)
	userStorage := user.NewStorage(db)
	userService := user2.NewService(userStorage)
	userHandler := user3.NewHandler(userService, sessionService)
	userHandler.Register(router)
	return router
}
