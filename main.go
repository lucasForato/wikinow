/*
Copyright © 2024 Lucas Forato lucas.e.forato@gmail.com
*/
package main

import (
	"wikinow/cmd"
	"wikinow/infra/logger"
)

func main() {
  logger.Init()
	cmd.Execute()
}
