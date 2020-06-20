package example

type Handler struct {
	Service *Service `inject:"service"`
}
