package models

import "errors"

var (
	userErr  error = errors.New("Invalid user slice pointer ")
	videoErr error = errors.New("Invalid video slice pointer ")
)
