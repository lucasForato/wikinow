package cmd

import (
	"fmt"
	"wikinow/infra/logger"
	"wikinow/infra/config"
	"wikinow/internal/handler"

	"github.com/labstack/echo/v4"
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
		e.HideBanner = true
		e.HidePort = true

		e.GET("/wiki/*", handler.Wiki)
		e.GET("favicon.ico", handler.Favicon)
		e.GET("/api/search", handler.GETSearch)
		e.POST("/api/search", handler.POSTSearch)

		e.Static("/images", "images")

		port, err := config.GetPort()
		if err != nil {
			logger.Error("Error reading port from config file.")
		}

		fmt.Println(`
/\ \  _ \ \   /\ \   /\ \/ /    /\ \   /\ "-.\ \   /\  __ \   /\ \  _ \ \   
\ \ \/ ".\ \  \ \ \  \ \  _"-.  \ \ \  \ \ \-.  \  \ \ \/\ \  \ \ \/ ".\ \  
 \ \__/".~\_\  \ \_\  \ \_\ \_\  \ \_\  \ \_\\"\_\  \ \_____\  \ \__/".~\_\ 
  \/_/   \/_/   \/_/   \/_/\/_/   \/_/   \/_/ \/_/   \/_____/   \/_/   \/_/
      `)

		if port == ":" {
			port = ":4000"
		}
		docUrl := fmt.Sprintf("http://localhost%s/wiki/", port)
		fmt.Println("You can access the documentation at: ", docUrl)
		e.Logger.Fatal(e.Start(port))
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}
