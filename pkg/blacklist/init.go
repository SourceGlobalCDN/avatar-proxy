package blacklist

import (
	_ "embed"
	"github.com/SourceGlobalCDN/avatar-proxy/pkg/log"
)

func Init(data []byte) {
	model, err := UnmarshalBlacklist(data)
	if err != nil {
		log.Log().Panicf("Error unmarshalling blacklist: %s", err)
	}

	global = &model
}
