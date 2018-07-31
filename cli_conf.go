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
	patterns        string
	patternSplit    string
	patternTest     bool
	ttlOpts         string
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
	delay             int
	cron              string
	bigKeysConfig     *BigKeysConfig
	deleteKeysConfig  *DeleteKeysConfig
	slowLogConfig     *SlowLogConfig
	configCheckConfig *ConfigCheckConfig
	serverConfig      *ServerConfig
	emailConfig       *EmailMessage
}

var DefaultConfig = &CmdConfig{
	host:           "127.0.0.1",
	port:           6379,
	password:       "",
	db:             0,
	scanThreads:    1,
	processThreads: 1,
	bigKeysConfig: &BigKeysConfig{
		keyBatch:     400,
		keyInterval:  10,
		patterns:     "",
		patternSplit: "@",
		sizeLimit:    10,
	},
	deleteKeysConfig: &DeleteKeysConfig{
		keyBatch:     400,
		keyInterval:  10,
		pattern:      "",
		directDelete: false,
		expireTime:   100,
	},
}

var InputConfig = &CmdConfig{}

var appFlags = []cli.Flag{
	cli.StringFlag{
		Name:        "host",
		Value:       DefaultConfig.host,
		Usage:       "redis `hostname`",
		Destination: &InputConfig.host,
	},
	cli.IntFlag{
		Name:        "port",
		Value:       DefaultConfig.port,
		Usage:       "redis `port`",
		Destination: &InputConfig.port,
	},
	cli.StringFlag{
		Name:        "pwd",
		Value:       DefaultConfig.password,
		Usage:       "redis `password`",
		Destination: &InputConfig.password,
	},
	cli.IntFlag{
		Name:        "db",
		Value:       DefaultConfig.db,
		Usage:       "redis `database`",
		Destination: &InputConfig.db,
	},
	//cli.IntFlag{
	//	Name:        "scan-thread",
	//	Value:       1,
	//	Usage:       "threads concurrently to scan keys",
	//	Destination: &InputConfig.scanThreads,
	//},
	cli.IntFlag{
		Name:        "process-thread",
		Value:       3,
		Usage:       "threads concurrently to process keys",
		Destination: &InputConfig.processThreads,
	},
	cli.IntFlag{
		Name:        "delay",
		Value:       -1,
		Usage:       "",
		Destination: &InputConfig.delay,
	},

	cli.StringFlag{
		Name:        "cron",
		Value:       "",
		Usage:       "",
		Destination: &InputConfig.cron,
	},
	cli.StringFlag{
		Name:        "email-user",
		Value:       "",
		Usage:       "",
		Destination: &InputConfig.emailConfig.EmailUserName,
	},
	cli.StringFlag{
		Name:        "email-pwd",
		Value:       "",
		Usage:       "",
		Destination: &InputConfig.emailConfig.EmailPassword,
	},
	cli.StringFlag{
		Name:        "email-to-users",
		Value:       "",
		Usage:       "",
		Destination: &InputConfig.emailConfig.To,
	},
	cli.StringFlag{
		Name:        "email-cc-users",
		Value:       "",
		Usage:       "",
		Destination: &InputConfig.emailConfig.CC,
	},
	cli.StringFlag{
		Name:        "email-subject",
		Value:       "",
		Usage:       "",
		Destination: &InputConfig.emailConfig.Subject,
	},
	cli.StringFlag{
		Name:        "email-body",
		Value:       "",
		Usage:       "",
		Destination: &InputConfig.emailConfig.Body,
	},
	cli.StringFlag{
		Name:        "smtp-host",
		Value:       "",
		Usage:       "",
		Destination: &InputConfig.emailConfig.EmailSMTPServerHost,
	},
	cli.IntFlag{
		Name:        "smtp-port",
		Value:       25,
		Usage:       "",
		Destination: &InputConfig.emailConfig.EmailSMTPServerPort,
	},
	cli.StringFlag{
		Name:  "config,c",
		Value: "config.ini",
		Usage: "config `file`, if not present, this program will create it!",
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
