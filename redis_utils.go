package main

import (
	"github.com/go-redis/redis"
	"strconv"
	"strings"
	"sync"
	"time"
)

type RedisClient struct {
	Addr        string
	StandAlone  *redis.Client
	Cluster     *redis.ClusterClient
	RedisModeID int
}

type ClusterNodes struct {
	Role     string
	NodeId   string
	SlaveIds []string
	MasterId string
	Addr     string
}

type RedisKeyDetail struct {
	Database     int
	Key          string
	TheType      string
	SizeByte     int64
	SizeMB       int32
	SizeGB       int
	ElementCount int64
	TTL          int64
	Expire       string
}

func Client(host string, port int, pwd string, db int) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     host + ":" + strconv.Itoa(port),
		Password: pwd,
		DB:       db,
	})
}

func Cluster(host string, port int, pwd string) *redis.ClusterClient {
	return redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:    []string{host + ":" + strconv.Itoa(port)},
		Password: pwd,
	})
}

func Info(host string, port int, pwd string, db int, section string) map[string]string {
	var infoMap = map[string]string{}
	client := Client(host, port, pwd, db)
	defer client.Close()
	result, _ := client.Info(section).Result()
	result = strings.TrimSpace(result)
	infoArr := strings.Split(result, "\n")
	for _, line := range infoArr {
		if strings.HasPrefix(line, "#") {
			continue
		}
		line = strings.TrimSpace(line)
		tmp := strings.Split(line, ":")
		if len(tmp) != 2 {
			continue
		}
		infoMap[tmp[0]] = tmp[1]
	}
	return infoMap
}

func ModeInt(host string, port int, pwd string, db int) int {
	modeString := ModeStr(host, port, pwd, db)
	var result int
	switch modeString {
	case ModeStandalone:
		result = 1
	case ModeCluster:
		result = 2
	default:
		result = -1
	}
	return result

}

func ModeStr(host string, port int, pwd string, db int) string {
	infoMap := Info(host, port, pwd, db, SectionServer)
	return infoMap[RedisMode]
}

func GetClusterNodes(host string, port int, pwd string) []ClusterNodes {
	var result = []ClusterNodes{}
	cluster := Cluster(host, port, pwd)
	defer cluster.Close()

	str, _ := cluster.ClusterNodes().Result()
	str = strings.TrimSpace(str)
	for _, line := range strings.Split(str, "\n") {
		tmp := strings.Fields(line)
		result = append(result, ClusterNodes{
			NodeId:   tmp[0],
			Addr:     strings.Split(tmp[1], "@")[0],
			Role:     strings.Split(tmp[2], ",")[0],
			MasterId: strings.Replace(tmp[3], "-", "", 1),
		})
	}

	for _, ni := range result {
		for _, nj := range result {
			if len(nj.MasterId) > 0 && nj.MasterId == ni.NodeId {
				ni.SlaveIds = append(ni.SlaveIds, nj.NodeId)
			}
		}
	}

	return result
}

func GetScanNodesAddresses(host string, port int, pwd string) []string {
	nodes := GetClusterNodes(host, port, pwd)

	var result = []string{}
	for _, n := range nodes {
		if len(n.SlaveIds) > 0 {
			result = append(result, n.SlaveIds[0])
		} else {
			result = append(result, n.NodeId)
		}
	}
	return result
}

func (c *RedisClient) SlowLog() {

}

func (c *RedisClient) ConfigCheck() {

}

func (c *RedisClient) MemoryStat() {

}

func DeleteKeys(client interface{ redis.Cmdable }, keys ...string) int64 {
	return client.Del(keys...).Val()
}

func ExpireKey(client interface{ redis.Cmdable }, keys string, duration time.Duration) bool {
	return client.Expire(keys, duration).Val()
}

func Scan(client *redis.Client, keys chan string, wg *sync.WaitGroup, elementBatch int, elementInterval int, pattern string) error {
	defer wg.Done()
	//defer client.Close()
	var (
		cursor     uint64 = 0
		resultKeys []string
		err        error
		count      = 0
	)

	for {
		resultKeys, cursor, err = client.Scan(cursor, pattern, GlobalScanBatch).Result()
		if err != nil {
			return err
		}
		if cursor == 0 {
			break
		}

		count += len(resultKeys)
		for _, k := range resultKeys {
			keys <- k
		}

		if elementInterval > 0 && count >= elementBatch {
			time.Sleep(time.Duration(elementInterval) * time.Millisecond)
		}
	}
	return nil
}

func KeyType(c interface{ redis.Cmdable }, key string) string {
	return c.Type(key).Val()
}

func GetRedisKeyDetail(c interface{ redis.Cmdable }, scanResultKeys chan string, NeedPrintKeys chan RedisKeyDetail, wg *sync.WaitGroup) {
	defer wg.Done()
	for key := range scanResultKeys {
		switch KeyType(c, key) {
		case "string":
			GetStringDetail(c, key, NeedPrintKeys)
		case "hash":
			GetHashDetail(c, key, NeedPrintKeys)
		case "list":
			GetListDetail(c, key, NeedPrintKeys)
		case "set":
			GetSetDetail(c, key, NeedPrintKeys)
		case "zset":
			GetZSetDetail(c, key, NeedPrintKeys)
		default:
			return
		}
	}
}

func GetStringDetail(c interface{ redis.Cmdable }, key string, oKey chan RedisKeyDetail) {

}

func GetHashDetail(c interface{ redis.Cmdable }, key string, oKey chan RedisKeyDetail) {

}

func GetListDetail(c interface{ redis.Cmdable }, key string, oKey chan RedisKeyDetail) {
}

func GetSetDetail(c interface{ redis.Cmdable }, key string, oKey chan RedisKeyDetail) {

}

func GetZSetDetail(c interface{ redis.Cmdable }, key string, oKey chan RedisKeyDetail) {

}
