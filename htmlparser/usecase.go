package htmlparser

type UseCase interface {
	Parse(name string, html string, data interface{}) (string, bool)
}
