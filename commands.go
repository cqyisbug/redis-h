package main

import (
	"errors"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/urfave/cli"
	"math"
	"strings"
	"sync"
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
			Name:  "patterns",
			Value: "",
			Usage: "",
		},
		cli.StringFlag{
			Name:        "pattern-split",
			Value:       "@",
			Usage:       "",
			Destination: &InputConfig.bigKeysConfig.patternSplit,
		},
		cli.BoolFlag{
			Name:        "pattern-test",
			Usage:       "",
			Destination: &InputConfig.bigKeysConfig.patternTest,
		},
		cli.StringFlag{
			Name:        "ttl-opts",
			Value:       "none",
			Usage:       "",
			Destination: &InputConfig.bigKeysConfig.ttlOpts,
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

		var (
			scanResultKeys = make(chan string, math.MaxInt64)
			outputKeys     = make(chan RedisKeyDetail, math.MaxInt64)
			//scan key wait group
			swg *sync.WaitGroup
			//get key value wait group
			gwg *sync.WaitGroup
			//file handler wait group
			fwg               *sync.WaitGroup
			clientPool        []*redis.Client
			clusterClientPool []*redis.ClusterClient
		)

		modeInt := ModeInt(ctx.GlobalString("host"), ctx.GlobalInt("port"), ctx.GlobalString("pwd"), ctx.GlobalInt("db"))
		if modeInt == -1 {
			return errors.New("* could not get info from redis server")
		}
		fmt.Printf(">>> Searching big keys...\n")

		if modeInt == 1 {
			client := Client(ctx.GlobalString("host"), ctx.GlobalInt("port"), ctx.GlobalString("pwd"), ctx.GlobalInt("db"))
			clientPool = append(clientPool, client)
			swg.Add(1)
			patterns := ctx.String("patterns")
			for _, p := range strings.Split(patterns, ctx.String("pattern-split")) {
				go Scan(client, scanResultKeys, swg, ctx.Int("element-batch"), ctx.Int("element-interval"), p)
			}

			for i := 0; i < ctx.Int("process-thread"); i++ {
				go GetRedisKeyDetail(client, scanResultKeys, outputKeys, gwg)
			}
		} else {
			scanAddresses := GetScanNodesAddresses(ctx.GlobalString("host"), ctx.GlobalInt("port"), ctx.GlobalString("pwd"))
			for _, addr := range scanAddresses {
				client := redis.NewClient(&redis.Options{
					Addr:     addr,
					Password: ctx.GlobalString("pwd"),
				})
				clientPool = append(clientPool, client)
				patterns := ctx.String("patterns")
				for _, p := range strings.Split(patterns, ctx.String("pattern-split")) {
					swg.Add(1)
					go Scan(client, scanResultKeys, swg, ctx.Int("element-batch"), ctx.Int("element-interval"), p)
				}
			}
			for _, addr := range scanAddresses {
				client := redis.NewClusterClient(&redis.ClusterOptions{
					Addrs:    []string{addr},
					Password: ctx.GlobalString("pwd"),
				})
				clusterClientPool = append(clusterClientPool, client)
				for i := 0; i < ctx.Int("process-thread"); i++ {
					gwg.Add(1)
					go GetRedisKeyDetail(client, scanResultKeys, outputKeys, gwg)
				}
			}
		}

		swg.Wait()
		fwg.Wait()
		gwg.Wait()

		for _, c := range clientPool {
			c.Close()
		}
		for _, c := range clusterClientPool {
			c.Close()
		}

		return nil
	},
}

var delKeysCommand = cli.Command{}

var checkCommand = cli.Command{}

var monitorCommand = cli.Command{}

var slowLogCommand = cli.Command{}
