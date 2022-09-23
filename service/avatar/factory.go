package avatar

import "io"

type Payload struct {
	Size         int    `json:"size" form:"s"`
	DefaultImg   string `json:"default_img" form:"d"`
	ForceDefault bool   `json:"force_default" form:"f"`
	Rating       Rating `json:"rating" form:"r"`
}

type Rating string

const (
	G Rating = "g"
	P Rating = "pg"
	R Rating = "r"
	X Rating = "x"
)

type Factory interface {
	GetAvatar(hash string, option Payload) (*io.ReadCloser, int, error)
}

func NewFactory() Factory {
	return NewImpl()
}
