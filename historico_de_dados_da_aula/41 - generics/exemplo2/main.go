package main

import "fmt"

// Entity é uma interface vazia que serve como um marcador para tipos de entidade.
// Assumimos que todos os tipos de entidade implementam esta interface implicitamente.
type Entity interface {
	GetId() int
}

// Repository é uma interface genérica para um repositório que opera sobre um tipo específico de entidade.
type Repository[T Entity] interface {
	Save(entity T) error
	All() ([]T, error)
	FindById(id int) (T, error)
}

// InMemoryRepository é uma implementação concreta de Repository que armazena entidades em memória.
type InMemoryRepository[T Entity] struct {
	entities []T
}

func NewInMemoryRepository[T Entity]() *InMemoryRepository[T] {
	return &InMemoryRepository[T]{}
}

func (r *InMemoryRepository[T]) Save(entity T) error {
	r.entities = append(r.entities, entity)
	return nil // Simplificação: ignorando erros
}

func (r *InMemoryRepository[T]) All() ([]T, error) {
	return r.entities, nil // Simplificação: ignorando erros
}

func (r *InMemoryRepository[T]) FindById(id int) (T, error) {
	// Simplificação: Esta implementação assume que há um campo ID ou um método GetID nas entidades,
	// o que não é representável diretamente em Go sem reflexão ou interfaces adicionais.
	// Portanto, estamos retornando o primeiro item e ignorando erros para simplificar.
	if len(r.entities) > 0 {
		for i := 0; i < len(r.entities); i++ {
			if r.entities[i].GetId() == id {
				return r.entities[i], nil
			}
		}
	}
	var zero T
	return zero, fmt.Errorf("id não encontrado")
}

type User struct {
	ID   int
	Name string
}

func (u User) GetId() int {
	return u.ID
}

type Product struct {
	Code  int
	Name  string
	Price float64
}

func (p Product) GetId() int {
	return p.Code
}

func main() {
	userRepo := NewInMemoryRepository[User]()
	userRepo.Save(User{ID: 1, Name: "John Doe"})
	users, _ := userRepo.All()
	fmt.Println("Users:", users)

	user, _ := userRepo.FindById(1)
	fmt.Println("Nome do usuario:", user.Name)

	productRepo := NewInMemoryRepository[Product]()
	productRepo.Save(Product{Code: 1, Name: "Laptop", Price: 999.99})
	products, _ := productRepo.All()
	fmt.Println("Products:", products)

	product, _ := productRepo.FindById(1)
	fmt.Println("Nome do produto:", product.Name)

	userNotFound, error := userRepo.FindById(10)
	fmt.Println(userNotFound)
	fmt.Println(error)

}
