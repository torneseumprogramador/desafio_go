package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Pagina struct {
	Id       primitive.ObjectID `bson:"_id,omitempty"` // Modificado para usar ObjectID
	Titulo   string             `bson:"titulo"`
	Conteudo string             `bson:"conteudo"`
	Status   uint32             `bson:"status"`
	Data     time.Time          `bson:"data"`
}

func gravarNoMongoDB(ctx context.Context, client *mongo.Client, pagina Pagina) (*mongo.InsertOneResult, error) {
	collection := client.Database("desafio_go").Collection("paginas")
	return collection.InsertOne(ctx, pagina)
}

func atualizarPaginaNoMongoDB(ctx context.Context, client *mongo.Client, id primitive.ObjectID, pagina Pagina) (*mongo.UpdateResult, error) {
	collection := client.Database("desafio_go").Collection("paginas")
	filter := bson.M{"_id": id}
	update := bson.M{"$set": pagina}
	return collection.UpdateOne(ctx, filter, update)
}

func apagarPaginaNoMongoDB(ctx context.Context, client *mongo.Client, id primitive.ObjectID) (*mongo.DeleteResult, error) {
	collection := client.Database("desafio_go").Collection("paginas")
	filter := bson.M{"_id": id}
	return collection.DeleteOne(ctx, filter)
}

func buscarMongoDb(ctx context.Context, client *mongo.Client) ([]Pagina, error) {
	var paginas []Pagina
	collection := client.Database("desafio_go").Collection("paginas")
	cursor, err := collection.Find(ctx, bson.D{{}})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var pagina Pagina
		if err = cursor.Decode(&pagina); err != nil {
			return nil, err
		}
		paginas = append(paginas, pagina)
	}
	return paginas, nil
}

// Buscar páginas por título usando uma busca semelhante a "like"
func buscarPorTitulo(ctx context.Context, client *mongo.Client, titulo string) ([]Pagina, error) {
	var paginas []Pagina
	collection := client.Database("desafio_go").Collection("paginas")
	filter := bson.M{"titulo": bson.M{"$regex": titulo, "$options": "i"}} // "i" para case-insensitive
	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var pagina Pagina
		if err := cursor.Decode(&pagina); err != nil {
			return nil, err
		}
		paginas = append(paginas, pagina)
	}
	return paginas, nil
}

// Buscar páginas por status com um intervalo específico
func buscarPorStatus(ctx context.Context, client *mongo.Client, minStatus, maxStatus uint32) ([]Pagina, error) {
	var paginas []Pagina
	collection := client.Database("desafio_go").Collection("paginas")
	filter := bson.M{"status": bson.M{"$gte": minStatus, "$lte": maxStatus}}
	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var pagina Pagina
		if err := cursor.Decode(&pagina); err != nil {
			return nil, err
		}
		paginas = append(paginas, pagina)
	}
	return paginas, nil
}

// Buscar páginas por data com um intervalo específico
func buscarPorData(ctx context.Context, client *mongo.Client, dataInicio, dataFim time.Time) ([]Pagina, error) {
	var paginas []Pagina
	collection := client.Database("desafio_go").Collection("paginas")
	filter := bson.M{"data": bson.M{"$gte": dataInicio, "$lte": dataFim}}
	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var pagina Pagina
		if err := cursor.Decode(&pagina); err != nil {
			return nil, err
		}
		paginas = append(paginas, pagina)
	}
	return paginas, nil
}

func main() {
	ctx := context.Background()

	// Configurando a conexão com o MongoDB
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	// Insira uma página como exemplo
	pagina1 := Pagina{Titulo: "Home", Conteudo: "Conteúdo atualizado da Home", Status: 2, Data: time.Now()}
	insertResult, err := gravarNoMongoDB(ctx, client, pagina1)
	if err != nil {
		log.Fatalf("Falha ao gravar no MongoDB: %v", err)
	}
	fmt.Printf("Página inserida com ID: %v\n", insertResult.InsertedID)

	pagina2 := Pagina{Titulo: "Dashboard", Conteudo: "Conteúdo atualizado da Dashboard", Status: 1, Data: time.Now()}
	gravarNoMongoDB(ctx, client, pagina2)

	id, _ := insertResult.InsertedID.(primitive.ObjectID) // aproveita o ID cadastrado
	// id, _ := primitive.ObjectIDFromHex("65e1a0ce9b5d335eabc53fd1") // coloca o ID fixo

	_, err = atualizarPaginaNoMongoDB(ctx, client, id, Pagina{Titulo: "Home Atualizada", Conteudo: "Conteúdo novo da Home"})
	if err != nil {
		log.Fatalf("Falha ao atualizar no MongoDB: %v", err)
	}

	_, err = apagarPaginaNoMongoDB(ctx, client, id)
	if err != nil {
		log.Fatalf("Falha ao apagar no MongoDB: %v", err)
	}

	paginas, err := buscarMongoDb(ctx, client)
	if err != nil {
		log.Fatalf("Falha ao buscar no MongoDB: %v", err)
	}

	for _, pagina := range paginas {
		fmt.Printf("ID: %s\nTítulo: %s\nConteúdo: %s\nStatus: %d\nData: %s\n", pagina.Id.Hex(), pagina.Titulo, pagina.Conteudo, pagina.Status, pagina.Data.Format("02/01/2006 15:04"))
		fmt.Println("--------------------------------------------------")
	}

	paginasFiltroTitulo, err := buscarPorTitulo(ctx, client, "Home")
	if err != nil {
		log.Fatalf("Falha ao buscar no MongoDB: %v", err)
	}

	for _, pagina := range paginasFiltroTitulo {
		fmt.Printf("ID: %s\nTítulo: %s\nConteúdo: %s\nStatus: %d\nData: %s\n", pagina.Id.Hex(), pagina.Titulo, pagina.Conteudo, pagina.Status, pagina.Data.Format("02/01/2006 15:04"))
		fmt.Println("--------------------------------------------------")
	}

	paginasFiltroStatus, err := buscarPorStatus(ctx, client, 1, 2)
	if err != nil {
		log.Fatalf("Falha ao buscar no MongoDB: %v", err)
	}

	for _, pagina := range paginasFiltroStatus {
		fmt.Printf("ID: %s\nTítulo: %s\nConteúdo: %s\nStatus: %d\nData: %s\n", pagina.Id.Hex(), pagina.Titulo, pagina.Conteudo, pagina.Status, pagina.Data.Format("02/01/2006 15:04"))
		fmt.Println("--------------------------------------------------")
	}

	layout := "02/01/2006 15:04" // "dd/MM/yyyy HH:mm" forma como o go utiliza
	dataIniStr := "01/03/2024 06:30"
	dataFimStr := "02/03/2024 23:30"

	dataIni, _ := time.Parse(layout, dataIniStr)
	dataFim, _ := time.Parse(layout, dataFimStr)

	paginasFiltroData, err := buscarPorData(ctx, client, dataIni, dataFim)
	if err != nil {
		log.Fatalf("Falha ao buscar no MongoDB: %v", err)
	}

	for _, pagina := range paginasFiltroData {
		fmt.Printf("ID: %s\nTítulo: %s\nConteúdo: %s\nStatus: %d\nData: %s\n", pagina.Id.Hex(), pagina.Titulo, pagina.Conteudo, pagina.Status, pagina.Data.Format("02/01/2006 15:04"))
		fmt.Println("--------------------------------------------------")
	}
}

// ====== Gravando com ID controlado ======
// package main

// import (
// 	"context"
// 	"fmt"
// 	"log"

// 	"go.mongodb.org/mongo-driver/bson"
// 	"go.mongodb.org/mongo-driver/mongo"
// 	"go.mongodb.org/mongo-driver/mongo/options"
// )

// type Pagina struct {
// 	Id       int    `bson:"id"`
// 	Titulo   string `bson:"titulo"`
// 	Conteudo string `bson:"conteudo"`
// }

// func gravarNoMongoDB(ctx context.Context, client *mongo.Client, pagina Pagina) (*mongo.InsertOneResult, error) {
// 	collection := client.Database("desafio_go").Collection("paginas")
// 	return collection.InsertOne(ctx, pagina)
// }

// func atualizarPaginaNoMongoDB(ctx context.Context, client *mongo.Client, id int, pagina Pagina) (*mongo.UpdateResult, error) {
// 	collection := client.Database("desafio_go").Collection("paginas")
// 	filter := bson.M{"id": id}
// 	update := bson.M{"$set": pagina}
// 	return collection.UpdateOne(ctx, filter, update)
// }

// func apagarPaginaNoMongoDB(ctx context.Context, client *mongo.Client, id int) (*mongo.DeleteResult, error) {
// 	collection := client.Database("desafio_go").Collection("paginas")
// 	filter := bson.M{"id": id}
// 	return collection.DeleteOne(ctx, filter)
// }

// func buscarMongoDb(ctx context.Context, client *mongo.Client) ([]Pagina, error) {
// 	var paginas []Pagina
// 	collection := client.Database("desafio_go").Collection("paginas")
// 	cursor, err := collection.Find(ctx, bson.D{{}})
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer cursor.Close(ctx)

// 	for cursor.Next(ctx) {
// 		var pagina Pagina
// 		if err = cursor.Decode(&pagina); err != nil {
// 			return nil, err
// 		}
// 		paginas = append(paginas, pagina)
// 	}
// 	return paginas, nil
// }

// func main() {
// 	ctx := context.Background()

// 	// Configurando a conexão com o MongoDB
// 	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
// 	client, err := mongo.Connect(ctx, clientOptions)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer client.Disconnect(ctx)

// 	// Insira, atualize e apague uma página como exemplo
// 	pagina1 := Pagina{Id: 1, Titulo: "Home", Conteudo: "Conteúdo atualizado da Home"}
// 	_, err = gravarNoMongoDB(ctx, client, pagina1)
// 	if err != nil {
// 		log.Fatalf("Falha ao gravar no MongoDB: %v", err)
// 	}

// 	pagina2 := Pagina{Id: 2, Titulo: "Dashboard", Conteudo: "Conteúdo atualizado da Dashboard"}
// 	gravarNoMongoDB(ctx, client, pagina2)

// 	_, err = atualizarPaginaNoMongoDB(ctx, client, 1, Pagina{Id: 1, Titulo: "Home Atualizada", Conteudo: "Conteúdo novo da Home"})
// 	if err != nil {
// 		log.Fatalf("Falha ao atualizar no MongoDB: %v", err)
// 	}

// 	_, err = apagarPaginaNoMongoDB(ctx, client, 1)
// 	if err != nil {
// 		log.Fatalf("Falha ao apagar no MongoDB: %v", err)
// 	}

// 	paginas, err := buscarMongoDb(ctx, client)
// 	if err != nil {
// 		log.Fatalf("Falha ao buscar no MongoDB: %v", err)
// 	}

// 	for _, pagina := range paginas {
// 		fmt.Printf("ID: %s\nTítulo: %s\nConteúdo: %s\nStatus: %d\nData: %s\n", pagina.Id.Hex(), pagina.Titulo, pagina.Conteudo, pagina.Status, pagina.Data.Format("02/01/2006 15:04"))
// 		fmt.Println("--------------------------------------------------")
// 	}
// }
