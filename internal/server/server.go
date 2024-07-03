package server

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	_ "github.com/joho/godotenv/autoload"

	"github.com/faisalhardin/sawitpro/internal/database"
	estateHandler "github.com/faisalhardin/sawitpro/internal/entity/interfaces"
)

type Server struct {
	port int

	db database.Service
	EstateHandler estateHandler.EstateHandler
}

func NewServer(handler estateHandler.EstateHandler) *http.Server {
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	NewServer := &Server{
		port: port,
		EstateHandler: handler,
	}

	// Declare Server config
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", NewServer.port),
		Handler:      NewServer.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}
