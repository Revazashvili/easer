package pdf

type UseCase interface {
	Render(id string,data interface{}) ([]byte,error)
}
