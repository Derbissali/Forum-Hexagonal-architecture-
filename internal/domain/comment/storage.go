package comment

type CommentStorage interface {
	AddComment(comment string, id, n int) error
}
