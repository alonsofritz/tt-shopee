package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/alonsofritz/tt-shopee/config"
	"github.com/alonsofritz/tt-shopee/internal/api/router"
)

func main() {
	cfg := config.LoadConfig()

	fmt.Printf("APP_NAME: %s\n", cfg.AppName)
	fmt.Printf("APP_VERSION: %s\n", cfg.AppVersion)
	fmt.Printf("SERVER_PORT: %s\n", cfg.ServerPort)

	mux := router.SetupRouter()
	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", cfg.ServerPort),
		Handler: mux,
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		log.Printf("Servidor iniciado na porta :%s", cfg.ServerPort)
		if err := server.ListenAndServe(); err != nil {
			log.Fatal()
		}
	}()

	<-stop
	log.Println("Sinal recebido. Iniciando shutdown ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Erro durante shutdown: %v", err)
	}

	log.Println("Servidor finalizado.")
}
