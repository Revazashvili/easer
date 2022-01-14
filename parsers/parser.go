package parsers

type Parser interface {
	Parse(name string, templateToParse string, data interface{}) (string, bool)
}
