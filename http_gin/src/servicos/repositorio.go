package servicos

type Repositorio[T any] struct {
	Servico CRUDService[T]
}

func NovoRepositorio[T any](servico CRUDService[T]) *Repositorio[T] {
	return &Repositorio[T]{
		Servico: servico,
	}
}
