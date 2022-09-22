package env

import "github.com/SourceGlobalCDN/avatar-proxy/pkg/util"

var SystemConfig = &system{
	Listen: util.EnvStr("LISTEN", ":8080"),
	Debug:  util.EnvBool("DEBUG", false),
}

var ProxyConfig = &proxy{
	Remote: util.EnvStr("REMOTE", "https://www.gravatar.com/avatar/"),
}
