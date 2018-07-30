package main

import (
	"github.com/urfave/cli"
	"fmt"
)

// scan delete slowlog config memory_stat

type BigKeysConfig struct {
	keyBatch        int
	keyInterval     int
	elementBatch    int
	elementInterval int
	sizeLimit       int
	pattern         string
	dump            bool
}

type DeleteKeysConfig struct {
	keyInterval  int
	keyBatch     int
	pattern      string
	directDelete bool
	expireTime   int
}

type SlowLogConfig struct {
}

type ConfigCheckConfig struct {
}

type ServerConfig struct {
}

type CmdConfig struct {
	host              string
	port              int
	db                int
	password          string
	commands          []string
	command           string
	scanThreads       int
	processThreads    int
	bigKeysConfig     *BigKeysConfig
	deleteKeysConfig  *DeleteKeysConfig
	slowLogConfig     *SlowLogConfig
	configCheckConfig *ConfigCheckConfig
	serverConfig      *ServerConfig
}

var DefaultConfig = &CmdConfig{
	host:           "127.0.0.1",
	port:           6379,
	password:       "",
	db:             0,
	scanThreads:    1,
	processThreads: 1,
	bigKeysConfig: &BigKeysConfig{
		keyBatch:    400,
		keyInterval: 10,
		pattern:     "",
		sizeLimit:   10,
	},
	deleteKeysConfig: &DeleteKeysConfig{
		keyBatch:     400,
		keyInterval:  10,
		pattern:      "",
		directDelete: false,
		expireTime:   100,
	},
}

var InputConfig *CmdConfig

var appFlags = []cli.Flag{
	cli.StringFlag{
		Name:        "host",
		Value:       DefaultConfig.host,
		Usage:       "redis `hostname`",
	},
	cli.IntFlag{
		Name:        "port",
		Value:       DefaultConfig.port,
		Usage:       "redis `port`",
	},
	cli.StringFlag{
		Name:        "pwd",
		Value:       DefaultConfig.password,
		Usage:       "redis `password`",
	},
	cli.IntFlag{
		Name:        "db",
		Value:       DefaultConfig.db,
		Usage:       "redis `database`",
	},
	cli.IntFlag{
		Name:  "scan-thread",
		Value: 1,
		Usage: "threads concurrently to scan keys",
	},
	cli.IntFlag{
		Name:  "process-thread",
		Value: 3,
		Usage: "threads concurrently to process keys",
	},
	cli.StringFlag{
		Name:        "config,c",
		Value:       "config.ini",
		Usage:       "config `file`, if not present, this program will create it!",
	},
}

var appCommands = []cli.Command{
	bigKeysCommand,
	delKeysCommand,
	checkCommand,
	monitorCommand,
	slowLogCommand,
}

func beforeCommand(ctx *cli.Context) error {

	InputConfig = &CmdConfig{
		host:           ctx.GlobalString("host"),
		port:           ctx.GlobalInt("port"),
		password:       ctx.GlobalString("pwd"),
		db:             ctx.GlobalInt("db"),
		scanThreads:    ctx.GlobalInt("scan-thread"),
		processThreads: ctx.GlobalInt("process-thread"),
	}

	fmt.Printf("*****************************************************************\n")
	fmt.Printf("\t\tWelcome to use the Redis-Util\n")
	fmt.Printf(
		"\t\tredis host : %s\n\t\tredis port : %d\n\t\tredis database : %d\n",
		ctx.GlobalString("host"),
		ctx.GlobalInt("port"),
		ctx.GlobalInt("db"))
	fmt.Printf("*****************************************************************\n\n")
	return nil
}
