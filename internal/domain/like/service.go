package like

type service struct {
	storage LikeStorage
}
type LikeService interface {
	SetLikeDislike(l, d, idPost string, n int) error
	SetLikeDislikeComment(l, d, idPost, idComment string, n int) error
}

func NewService(storage LikeStorage) LikeService {
	return &service{
		storage: storage,
	}
}

func (s *service) SetLikeDislike(l, d, idPost string, n int) error {
	if l != "" {
		a := s.storage.PostDislike(n, idPost)
		liked := s.storage.PostLike(n, idPost)
		if liked != 0 {
			s.storage.DeleteLikeNDis(idPost, n)
		}
		if a != 0 {
			s.storage.UpdateLike(idPost, n)
		}
		if a == 0 && liked == 0 {
			s.storage.SetLike(idPost, n)
		}

	} else if d != "" {
		b := s.storage.PostLike(n, idPost)
		disliked := s.storage.PostDislike(n, idPost)
		if disliked != 0 {
			s.storage.DeleteLikeNDis(idPost, n)
		}
		if b != 0 {
			s.storage.UpdateDislike(idPost, n)
		}
		if b == 0 && disliked == 0 {
			s.storage.SetDislike(idPost, n)
		}
	}
	s.storage.UpdateLikeCount(idPost)
	s.storage.UpdateDislikeCount(idPost)
	return nil
}
func (s *service) SetLikeDislikeComment(l, d, idPost, idComment string, n int) error {
	if l != "" {
		a := s.storage.CommentDislike(n, idPost, idComment)
		liked := s.storage.CommentLike(n, idPost, idComment)
		if liked != 0 {
			s.storage.DeleteCommentLikeNDis(idPost, idComment, n)
		}
		if a != 0 {
			s.storage.UpdateCommentLike(idPost, idComment, n)
		}
		if a == 0 && liked == 0 {
			s.storage.SetCommentLike(idPost, idComment, n)
		}

	} else if d != "" {
		b := s.storage.CommentLike(n, idPost, idComment)
		disliked := s.storage.CommentDislike(n, idPost, idComment)
		if disliked != 0 {
			s.storage.DeleteCommentLikeNDis(idPost, idComment, n)
		}
		if b != 0 {
			s.storage.UpdateCommentDislike(idPost, idComment, n)
		}
		if b == 0 && disliked == 0 {
			s.storage.SetCommentDislike(idPost, idComment, n)
		}
	}
	s.storage.UpdateCommentLikeCount(idPost, idComment)
	s.storage.UpdateCommentDislikeCount(idPost, idComment)
	return nil
}
