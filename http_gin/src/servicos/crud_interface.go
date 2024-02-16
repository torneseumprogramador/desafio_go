package servicos

type CRUDService[T any] interface {
	Lista() ([]T, error)
	BuscarPorId(id string) (*T, error)
	Adicionar(entidade T) error
	Alterar(entidade T) error
	Excluir(id string) error
}
