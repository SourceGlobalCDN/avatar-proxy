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

var CacheConfig = &cache{
	Enabled:  util.EnvBool("CACHE_ENABLED", false),
	Mode:     util.EnvStr("CACHE_MODE", "memory"),
	Host:     util.EnvStr("CACHE_HOST", "localhost"),
	Port:     util.EnvInt("CACHE_PORT", 6379),
	Password: util.EnvStr("CACHE_PASSWORD", ""),
	Database: util.EnvInt("CACHE_DATABASE", 0),
}
