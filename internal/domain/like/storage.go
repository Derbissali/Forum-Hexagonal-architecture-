package like

type LikeStorage interface {
	SetLike(idPost string, n int)
	SetDislike(idPost string, n int)
	UpdateLike(idPost string, n int)
	UpdateDislike(idPost string, n int)
	PostDislike(n int, idPost string) int
	PostLike(n int, idPost string) int
	DeleteLikeNDis(idPost string, n int)
	UpdateLikeCount(idPost string)
	UpdateDislikeCount(idPost string)

	SetCommentLike(idPost, idComment string, n int)
	SetCommentDislike(idPost, idComment string, n int)
	UpdateCommentLike(idPost, idComment string, n int)
	UpdateCommentDislike(idPost, idComment string, n int)
	CommentDislike(n int, idPost, idComment string) int
	CommentLike(n int, idPost, idComment string) int
	DeleteCommentLikeNDis(idPost, idComment string, n int)
	UpdateCommentLikeCount(idPost, idComment string)
	UpdateCommentDislikeCount(idPost, idComment string)
}
