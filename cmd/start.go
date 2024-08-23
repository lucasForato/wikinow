package cmd

import (
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
		e := echo.New()

		e.Static("/static", "static")
		e.GET("/wiki/*", handler.Wiki)
		e.GET("favicon.ico", handler.Favicon)

		e.Logger.Fatal(e.Start(":4000"))
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}
