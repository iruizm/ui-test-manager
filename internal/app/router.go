package api

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"polonium/internal/pkg/configuration"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Setup() *gin.Engine {
	app := gin.New()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:" + configuration.Config.FrontPort}
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

	app.GET("/api/sequences", GetSequences)
	app.POST("/api/sequences", SaveSequence)
	app.DELETE("/api/sequences/:id", DeleteSequence)

	// ================== Docs Routes
	app.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return app
}
