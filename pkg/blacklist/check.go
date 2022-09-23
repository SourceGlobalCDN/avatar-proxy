package blacklist

import "github.com/samber/lo"

// CheckGitHub checks if the given username is blacklisted
func CheckGitHub(hash string) bool {
	_, exist := lo.Find(global.GitHub, func(item string) bool {
		return item == hash
	})

	return exist
}

// CheckGravatar checks if the given email is blacklisted
func CheckGravatar(hash string) bool {
	_, exist := lo.Find(global.Gravatar, func(item string) bool {
		return item == hash
	})

	return exist
}
