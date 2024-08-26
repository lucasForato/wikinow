package cmd

import (
	"wikinow/internal/config"
	"wikinow/internal/handler"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Initializes wikinow by adding a /wiki directory and configuration file",
	Long: `Initialization command:

  This command will create a config file at the current directory and a /wiki
  directory with a main.md file in it.

  You should run this command before using any other command. 
  `,
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		config.SetupConfig()

		e := echo.New()
		e.Static("/static", "static")
		e.GET("/wiki/*", handler.Wiki)
		e.GET("favicon.ico", handler.Favicon)
		e.GET("/api/search", handler.GETSearch)
		e.POST("/api/search", handler.POSTSearch)

		port, err := config.GetPort()
		if err != nil {
			log.Fatal("Error reading port from config file.")
		}

		e.Logger.Fatal(e.Start(port))
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}
