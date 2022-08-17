package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"github.com/wilcatarino/labs-full-cycle-bank/domain"
	"github.com/wilcatarino/labs-full-cycle-bank/infrastructure/repository"
	"github.com/wilcatarino/labs-full-cycle-bank/usecase"
)

// func init() {
// 	err := godotenv.Load()
// 	if err != nil {
// 		log.Fatal("error loading .env file")
// 	}
// }

func main() {
	db := setupDb()
	defer db.Close()

	cc := domain.NewCreditCard()
	cc.Number = "5340571799979248"
	cc.Name = "Wilson C. Tavares"
	cc.ExpirationYear = 2023
	cc.ExpirationMonth = 1
	cc.CVV = 509
	cc.Limit = 10000
	cc.Balance = 0

	repo := repository.NewTransactionRepositoryDatabase(db)
	err := repo.CreateCreditCard(*cc)
	if err != nil {
		fmt.Println(err)
	}

	// producer := setupKafkaProducer()
	// processTransactionUseCase := setupTransactionUseCase(db)
	// serveGrpc(processTransactionUseCase)
}

func setupTransactionUseCase(db *sql.DB) usecase.UseCaseTransaction {
	transactionRepository := repository.NewTransactionRepositoryDatabase(db)
	useCase := usecase.NewUseCaseTransaction(transactionRepository)
	// useCase.KafkaProducer = producer
	return useCase
}

// func setupKafkaProducer() kafka.KafkaProducer {
// 	producer := kafka.NewKafkaProducer()
// 	producer.SetupProducer(os.Getenv("KafkaBootstrapServers"))
// 	return producer
// }

func setupDb() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		"db",           // os.Getenv("host"),
		"5432",         // os.Getenv("port"),
		"postgres",     // os.Getenv("user"),
		"root",         // os.Getenv("password"),
		"catarinobank", // os.Getenv("dbname"),
	)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal("error connection to database")
	}
	return db
}

// func serveGrpc(processTransactionUseCase usecase.UseCaseTransaction) {
// 	grpcServer := server.NewGRPCServer()
// 	grpcServer.ProcessTransactionUseCase = processTransactionUseCase
// 	fmt.Println("Rodando gRPC Server")
// 	grpcServer.Serve()
// }
