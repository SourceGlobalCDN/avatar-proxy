package avatar

import (
	"fmt"
	"io"

	"github.com/SourceGlobalCDN/avatar-proxy/pkg/json"
)

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
	GetAvatarInfo(hash string) (*Info, error)
}

func NewFactory() Factory {
	return NewImpl()
}

var (
	ErrNotFound = fmt.Errorf("avatar not found")
)

func UnmarshalInfo(data []byte) (*Info, error) {
	var info Info
	if err := json.Unmarshal(data, &info); err != nil {
		return nil, err
	}

	return &info, nil
}

type Info struct {
	Entry []Entry `json:"entry" xml:"entry" yaml:"entry"`
}

func (i Info) Marshal() ([]byte, error) {
	return json.Marshal(i)
}

type Entry struct {
	ID                string  `json:"id" xml:"id" yaml:"id"`
	Hash              string  `json:"hash" xml:"hash" yaml:"hash"`
	RequestHash       string  `json:"requestHash" xml:"request_hash" yaml:"request_hash"`
	ProfileURL        string  `json:"profileUrl" xml:"profileUrl" yaml:"profileUrl"`
	PreferredUsername string  `json:"preferredUsername" xml:"preferredUsername" yaml:"preferredUsername"`
	ThumbnailURL      string  `json:"thumbnailUrl" xml:"thumbnailUrl" yaml:"thumbnailUrl"`
	Photos            []Pho   `json:"photos" xml:"photos" yaml:"photos"`
	Name              *Name   `json:"name,omitempty" xml:"name,omitempty" yaml:"name,omitempty"`
	DisplayName       string  `json:"displayName" xml:"display_name" yaml:"display_name"`
	CurrentLocation   *string `json:"currentLocation,omitempty" xml:"current_location,omitempty" yaml:"current_location,omitempty"`
	PhoneNumbers      []Pho   `json:"phoneNumbers,omitempty" xml:"phone_numbers,omitempty" yaml:"phone_numbers,omitempty"`
	Emails            []Email `json:"emails,omitempty" xml:"emails,omitempty" yaml:"emails,omitempty"`
	Urls              []URL   `json:"urls,omitempty" xml:"urls,omitempty" yaml:"urls,omitempty"`
}

type Email struct {
	Primary string `json:"primary" xml:"primary" yaml:"primary"`
	Value   string `json:"value" xml:"value" yaml:"value"`
}

type Name struct {
	GivenName  string `json:"givenName" xml:"givenName" yaml:"givenName"`
	FamilyName string `json:"familyName" xml:"familyName" yaml:"familyName"`
	Formatted  string `json:"formatted" xml:"formatted" yaml:"formatted"`
}

type Pho struct {
	Type  string `json:"type" xml:"type" yaml:"type"`
	Value string `json:"value" xml:"value" yaml:"value"`
}

type URL struct {
	Value string `json:"value" xml:"value" yaml:"value"`
	Title string `json:"title" xml:"title" yaml:"title"`
}
