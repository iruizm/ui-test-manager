package router

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	controller "ui-test-manager/internal/app/controllers"
	"ui-test-manager/internal/pkg/configuration"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Setup(controller *controller.Controller) *gin.Engine {
	app := gin.New()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:" + configuration.Config.FrontPort, "http://192.168.1.1:" + configuration.Config.FrontPort}
	app.Use(cors.New(config))

	f, _ := os.Create(filepath.Join(configuration.Config.DataPath, configuration.Config.LoggingPath))
	gin.DefaultWriter = io.MultiWriter(f)

	app.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - - [%s] \"%s %s %s %d %s \" \" %s\" \" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format("02/Jan/2006:15:04:05 -0700"),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))
	app.Use(gin.Recovery())

	app.GET("/api/sequences", controller.GetSequences)
	app.GET("/api/patterns", controller.GetPatterns)
	app.GET("/api/tests", controller.GetTests)
	app.POST("/api/sequences", controller.SaveSequence)
	app.POST("/api/patterns", controller.SavePattern)
	app.DELETE("/api/sequences/:id", controller.DeleteSequence)
	app.DELETE("/api/patterns/:id", controller.DeletePattern)
	app.DELETE("/api/sequences/:id/:idPrecedent", controller.DeletePrecedent)

	app.GET("/api/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})
	// ================== Docs Routes
	app.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return app
}
