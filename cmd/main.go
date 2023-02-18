package main

import (
	"database/sql"
	"fmt"
	"net"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/tortuga-softworks/hestia/internal/registration"
	"github.com/tortuga-softworks/hestia/pkg/account"
	"github.com/tortuga-softworks/hestia/proto"
)

func main() {
	fmt.Println("<== HESTIA ==>")

	accountStore := initAccountStore()
	registrationService := initRegistrationService(accountStore)
	registrationServer := initRegistrationServer(registrationService)
	listener := initListener()

	server := grpc.NewServer()
	reflection.Register(server) // added for services discovery
	proto.RegisterRegistrationServer(server, registrationServer)
	if err := server.Serve(listener); err != nil {
		panic(err)
	}
}

func initAccountStore() account.AccountStore {
	dbConnectionString := os.Getenv("HESTIA_DB")

	db, err := sql.Open("postgres", dbConnectionString)
	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to PostgreSQL.")

	store, err := account.NewSqlAccountStore(db)

	if err != nil {
		panic(err)
	}

	fmt.Println("Account store ready.")
	return store
}

func initRegistrationService(accountStore account.AccountStore) *registration.RegistrationService {
	service, err := registration.NewRegistrationService(accountStore)

	if err != nil {
		panic(err)
	}

	fmt.Println("Registration service ready.")
	return service
}

func initRegistrationServer(registrationService *registration.RegistrationService) *registration.RegistrationServer {
	server, err := registration.NewRegistrationServer(registrationService)

	if err != nil {
		panic(err)
	}

	fmt.Println("Registration server ready.")
	return server
}

func initListener() net.Listener {
	port := os.Getenv("HESTIA_PORT")

	if port == "" {
		port = "9000"
		fmt.Println("No port configration found. Using default: 9000.")
	}

	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Listening on port " + port + ".")
	}

	return listener
}
