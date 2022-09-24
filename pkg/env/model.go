package env

type system struct {
	Listen string
	Debug  bool
}

type proxy struct {
	Remote    string
	UserAgent string
	Timeout   int
}

type cache struct {
	Enabled  bool
	Mode     string `binding:"eq=redis|eq=memory"`
	Host     string
	Port     int
	Password string
	Database int
}
