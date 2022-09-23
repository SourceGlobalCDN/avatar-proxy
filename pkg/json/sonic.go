//go:build sonic && avx && (linux || windows || darwin) && amd64
// +build sonic
// +build avx
// +build linux windows darwin
// +build amd64

package json

import "github.com/bytedance/sonic"

var (
	json = sonic.ConfigStd
	// Marshal is exported by gin/json package.
	Marshal = json.Marshal
	// Unmarshal is exported by gin/json package.
	Unmarshal = json.Unmarshal
	// MarshalIndent is exported by gin/json package.
	MarshalIndent = json.MarshalIndent
	// NewDecoder is exported by gin/json package.
	NewDecoder = json.NewDecoder
	// NewEncoder is exported by gin/json package.
	NewEncoder = json.NewEncoder
)
