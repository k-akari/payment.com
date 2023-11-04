package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/k-akari/payment.com/internal/handler"
	"github.com/k-akari/payment.com/internal/infrastructure/database"
	"github.com/k-akari/payment.com/internal/infrastructure/repository"
	"github.com/k-akari/payment.com/internal/router"
	"github.com/k-akari/payment.com/internal/usecase"
)

func main() {
	if err := run(context.Background()); err != nil {
		log.Printf("failed to terminate server: %v", err)
		os.Exit(1)
	}
}

func run(ctx context.Context) error {
	cfg, err := newConfig()
	if err != nil {
		return err
	}

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.Port))
	if err != nil {
		log.Fatalf("failed to listen port %d: %v", cfg.Port, err)
	}
	url := fmt.Sprintf("http://%s", l.Addr().String())
	log.Printf("start with: %v", url)

	db, cleanup, err := database.New(cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBName, cfg.DBPort)
	if err != nil {
		return fmt.Errorf("failed to connect database: %w", err)
	}
	defer cleanup()

	dbc := database.NewClient(db)
	cor := repository.NewCompanyRepository(dbc)
	clr := repository.NewClientRepository(dbc)
	ir := repository.NewInvoiceRepository(dbc)
	cou := usecase.NewCompanyUsecase(cor)
	clu := usecase.NewClientUsecase(clr)
	iu := usecase.NewInvoiceUsecase(ir)
	coh := handler.NewCompanyHandler(cou)
	clh := handler.NewClientHandler(clu)
	ih := handler.NewInvoiceHandler(iu)

	mux := router.NewMux(coh, clh, ih)
	s := newServer(l, mux)

	return s.run(ctx)
}
