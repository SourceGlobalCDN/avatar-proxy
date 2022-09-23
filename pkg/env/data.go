package env

import "github.com/SourceGlobalCDN/avatar-proxy/pkg/util"

var SystemConfig = &system{
	Listen: util.EnvStr("LISTEN", ":9000"),
	Debug:  util.EnvBool("DEBUG", false),
}

var ProxyConfig = &proxy{
	Remote:    util.EnvStr("PROXY_REMOTE", "https://www.gravatar.com/avatar/"),
	UserAgent: util.EnvStr("PROXY_USERAGENT", "AvatarProxy/1.0"),
	Timeout:   util.EnvInt("PROXY_TIMEOUT", 5),
}
