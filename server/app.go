package server

import (
	"context"
	delivery "github.com/Revazashvili/easer/delivery/http"
	"github.com/Revazashvili/easer/parsers"
	"github.com/Revazashvili/easer/renderers"
	"github.com/Revazashvili/easer/storage"
	"github.com/Revazashvili/easer/storage/mongo"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

type App struct {
	httpServer *http.Server
	storage    storage.Storage
	parser     parsers.Parser
	renderer   renderers.Renderer
}

func NewApp() *App {
	dbOptions := mongo.DbOptions{
		Uri:              viper.GetString("mongo.uri"),
		DbName:           viper.GetString("mongo.name"),
		TemplateCollName: viper.GetString("mongo.template_collection"),
	}
	templateStorage := mongo.NewTemplateStorage(dbOptions)
	htmlParser := parsers.NewHtmlParser()
	pdfRenderer := renderers.NewPdfRenderer(htmlParser)
	return &App{
		storage:  templateStorage,
		parser:   htmlParser,
		renderer: pdfRenderer,
	}
}

func (a *App) Run(port string) error {
	router := gin.Default()
	router.Use(
		gin.Recovery(),
		gin.Logger(),
	)
	api := router.Group("api")
	delivery.RegisterPdfHTTPEndpoints(api, a.storage, a.renderer)
	delivery.RegisterTemplateHTTPEndpoints(api, a.storage, a.parser)
	a.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if err := a.httpServer.ListenAndServe(); err != nil {
			log.Fatalf("Failed to listen and serve: %+v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Interrupt)

	s := <-quit
	log.Println("Got signal:", s)
	ctx, shutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdown()
	return a.httpServer.Shutdown(ctx)
}
