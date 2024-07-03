package main

import (
	"fmt"

	"github.com/faisalhardin/sawitpro/internal/database"
	"github.com/faisalhardin/sawitpro/internal/handler"
	"github.com/faisalhardin/sawitpro/internal/repo"
	"github.com/faisalhardin/sawitpro/internal/server"
	"github.com/faisalhardin/sawitpro/internal/usecase"
)

func main() {

	engine, err := database.NewXormDB()
	if err != nil {
		panic(err)
	}
	
	
	repoEstate := repo.NewEstateDBRepo(&repo.Conn{
		XormEngine: engine,
	})
	
	estateUC := usecase.NewEstateUC(&usecase.EstateUC{
		EstateDBRepo: repoEstate,
	})
	
	handler := handler.NewEstateHandler(&handler.EstateHandler{
		EstateUsecase: estateUC,
	})
	
	
	server := server.NewServer(handler)

	err = server.ListenAndServe()
	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}
}
