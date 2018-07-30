package main

import (
	"github.com/urfave/cli"
	"fmt"
	"os"
	"log"
)

func main() {

	app := cli.NewApp()
	app.Name = ">> Redis Utils <<"
	app.Author = "caiqyxyx"
	app.Email = "cqyisbug@163.com"
	app.Version = "0.0.1"
	app.Usage = "Delete Keys(delete) | Find Big Keys(bigkeys) | Show SlowLog(slowlog) |Check Redis Config(check) | Monitor Redis Status(monitor)"
	app.UsageText = "redis-utils [GLOBAL OPTIONS] COMMAND [COMMAND OPTIONS] [ARGS...]"

	app.EnableBashCompletion = true

	//global options
	app.Flags = appFlags
	//command
	app.Commands = appCommands
	//before
	app.Before = beforeCommand
	//deal handler
	//app.Action = appAction
	//command not found
	app.CommandNotFound = func(ctx *cli.Context, command string) {
		err := fmt.Errorf("invalid command %q", command)
		log.Fatal(err)
	}

	err := app.Run(os.Args)
	if err != nil {
		//log.Fatal(err)
	}
}
