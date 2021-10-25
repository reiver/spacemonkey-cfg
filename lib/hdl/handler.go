package hdl

type Handler interface {
	Value() (string, error)
}

type HandlerFunc func()(string, error)

func (fn HandlerFunc) Value() (string, error) {
	return fn()
}
