package repository

import "errors"

var ErrUserIDCollision = errors.New("user id collision")
var ErrUserNotFound = errors.New("user not found")
