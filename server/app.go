package server

import (
	"context"
	"fmt"
	"github.com/Revazashvili/easer/template"
	thttp "github.com/Revazashvili/easer/template/delivery/http"
	tmongo "github.com/Revazashvili/easer/template/repository/mongo"
	tusecase "github.com/Revazashvili/easer/template/usecase"
	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

type App struct {
	httpServer *http.Server
	templateUC template.UseCase
}

func NewApp() *App {
	db := initDB()
	templateRepo :=tmongo.NewTemplateRepository(db,viper.GetString("mongo.template_collection"))
	return &App{
		templateUC: tusecase.NewTemplateUseCase(templateRepo),
	}
}

func (a *App) Run(port string) error {
	// Init gin handler
	router := gin.Default()
	router.Use(
		gin.Recovery(),
		gin.Logger(),
	)
	api := router.Group("api")
	// Set up http handlers
	thttp.RegisterHTTPEndpoints(api,a.templateUC)

	a.httpServer = &http.Server{
		Addr: ":" + port,
		Handler: router,
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if err := a.httpServer.ListenAndServe(); err != nil {
			log.Fatalf("Failed to listen and serve: %+v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Interrupt)

	s:=<-quit
	log.Println("Got signal:",s)
	ctx, shutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdown()
	return a.httpServer.Shutdown(ctx)
}

func initDB() *mgo.Database {
	uri :=viper.GetString("mongo.uri")
	fmt.Println(uri)
	session, err := mgo.Dial(uri)
	if err != nil {
		log.Fatalf("Error occured while establishing connection to mongoDB")
	}
	db := session.DB(viper.GetString("mongo.name"))
	return db
}