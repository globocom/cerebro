package modules

import (
	"fmt"

	logger "github.com/bakatz/echo-logrusmiddleware"
	"github.com/labstack/echo"
	log "github.com/sirupsen/logrus"
)

// Init initializes current app.
func Init(fn func(settings Settings) PersistenceClient) {
	settings, err := LoadSettings()
	log.Debug("Current configuration:", settings)

	if err != nil {
		log.Fatal(err)
	}

	log.SetLevel(log.Level(settings.LogLevel))
	fmt.Println("Setting Log Level to", settings.LogLevel)

	server := echo.New()
	server.Debug = settings.Debug
	server.Logger = logger.Logger{Logger: log.StandardLogger()}

	client := fn(settings)
	defer client.Close()

	createEndpoints(server, settings, client)
	bindMiddlewares(server)

	bind := fmt.Sprintf(":%d", settings.Port)
	log.Fatal(server.Start(bind))
}

func createEndpoints(server *echo.Echo, settings Settings, client PersistenceClient) {
	handler := NewHTTPHandler(settings, client)

	server.GET("/", handler.Index)

	server.GET("/healthcheck", handler.Healthcheck)

	server.GET("/attribute/:name", handler.GetAttribute)
	server.POST("/attribute", handler.PostAttribute)
	server.PUT("/attribute/:name", handler.UpdateAttribute)
	server.DELETE("/attribute/:name", handler.DeleteAttribute)

	server.File("/favicon.ico", "/dev/null")
}

func bindMiddlewares(server *echo.Echo) {
	server.Use(logger.Hook())
}
