package blacklist

import "github.com/SourceGlobalCDN/avatar-proxy/pkg/json"

type Blacklist struct {
	Gravatar []string `json:"gravatar"`
	GitHub   []string `json:"github"`
}

func (b Blacklist) Marshal() ([]byte, error) {
	return json.Marshal(b)
}

func UnmarshalBlacklist(data []byte) (Blacklist, error) {
	var r Blacklist
	err := json.Unmarshal(data, &r)
	return r, err
}

var global = &Blacklist{
	Gravatar: []string{},
	GitHub:   []string{},
}
