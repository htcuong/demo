package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/htcuong/demo/config"
	"github.com/htcuong/demo/handler"
	"github.com/htcuong/demo/pkg/database"
	"github.com/htcuong/demo/pkg/log"
	"github.com/spf13/cobra"
)

type Handlers struct {
	wager handler.IWagerHandler
	buy   handler.IBuyHandler
}

func NewServerCmd(configs *config.Configurations, logger *log.Logger) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "api",
		Short: "run api server",
		Long:  "run api server with graphql",
		Run: func(cmd *cobra.Command, args []string) {
			// generate new logger with package name
			logger = log.FromLogger(logger, log.PackageName("cmd/server"))

			db, err := database.Connection(*configs.Database)
			if err != nil {
				logger.WithError(err).Fatal("Connect to db has failed")
				defer db.Close()
			}
			handlers := Handlers{
				wager: handler.NewWagerHandler(db, logger),
				buy:   handler.NewBuyHandler(db, logger),
			}

			r := setupRouter(handlers)
			r.Run(fmt.Sprintf("%s:%s", configs.Service.Host, configs.Service.Port))
		},
	}

	return cmd
}

func setupRouter(handlers Handlers) *gin.Engine {
	r := gin.Default()
	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	r.POST("/wagers", handlers.wager.CreateWager)
	r.GET("/wagers", handlers.wager.GetListWager)

	r.POST("/buy/:wager_id", handlers.buy.BuyWager)

	return r
}
