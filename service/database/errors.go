package database

import "errors"

var (
	ErrLikeNotFound    = errors.New("like not found")
	ErrPhotoNotFound   = errors.New("photo not found")
	ErrFollowNotFound  = errors.New("follow not found")
	ErrCommentNotFound = errors.New("comment not found")
	ErrBanNotFound     = errors.New("ban not found")
)
