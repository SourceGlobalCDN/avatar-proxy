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
