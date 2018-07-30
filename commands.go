package main

import (
	"github.com/urfave/cli"
	"fmt"
	"errors"
)

var bigKeysCommand = cli.Command{
	Name: "big-keys",
	Aliases: []string{
		"bk",
	},
	Usage:     "",
	UsageText: "",
	Flags: []cli.Flag{
		cli.StringFlag{
			//考虑支持数组
			Name:        "pattern",
			Value:       "",
			Usage:       "",
			Destination: &InputConfig.bigKeysConfig.pattern,
		},
		cli.IntFlag{
			Name:        "key-batch",
			Value:       100,
			Usage:       "sleep (microseconds) after processing this count of keys",
			Destination: &InputConfig.bigKeysConfig.keyBatch,
		},
		cli.IntFlag{
			Name:        "key-interval",
			Value:       10,
			Usage:       "sleep for the time(microseconds), after processing --key-batch count of keys , no sleep if 0",
			Destination: &InputConfig.bigKeysConfig.keyInterval,
		},
		cli.IntFlag{
			Name:        "element-batch",
			Value:       100,
			Usage:       "sleep (microseconds) after scanning this count of keys",
			Destination: &InputConfig.bigKeysConfig.elementBatch,
		},
		cli.IntFlag{
			Name:        "element-interval",
			Value:       10,
			Usage:       "sleep for the time(microseconds), after scanning --element-batch count of keys , no sleep if 0",
			Destination: &InputConfig.bigKeysConfig.elementInterval,
		},
		cli.BoolFlag{
			Name:        "dump",
			Usage:       "dump keys from redis",
			Destination: &InputConfig.bigKeysConfig.dump,
		},
	},
	Action: func(ctx *cli.Context) error {
		infoMap := Info(ctx.GlobalString("host"), ctx.GlobalInt("port"), ctx.GlobalString("pwd"), SectionServer)
		//infoMap := GetInfo(SectionServer)
		if len(infoMap) == 0 {
			return errors.New("* could not get info from redis server")
		}
		fmt.Printf(">>> Searching big keys...\n")
		fmt.Printf("%d", ctx.Int("scan-thread"))
		return nil
	},
}

var delKeysCommand = cli.Command{}

var checkCommand = cli.Command{}

var monitorCommand = cli.Command{}

var slowLogCommand = cli.Command{}
