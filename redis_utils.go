package main

import (
	"github.com/go-redis/redis"
	"strconv"
	"strings"
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

func Client(host string, port int, pwd string) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     host + ":" + strconv.Itoa(port),
		Password: pwd,
	})
}

func Info(host string, port int, pwd, section string) map[string]string {
	var infoMap = map[string]string{}
	client := Client(host, port, pwd)
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

func ModeInt(host string, port int, pwd string) int {
	modeString := ModeStr(host, port, pwd)
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

func ModeStr(host string, port int, pwd string) string {
	infoMap := Info(host, port, pwd, SectionServer)
	return infoMap[RedisMode]
}

func (client *RedisClient) GetClusterNodes() []ClusterNodes {
	var result = []ClusterNodes{}
	switch client.RedisModeID {
	case 1:
		result = nil
	case 2:
		str, _ := client.Cluster.ClusterNodes().Result()
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
	case -1:
		result = nil
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

func (client *RedisClient) GetScanNodesAddresses() []string {
	nodes := client.GetClusterNodes()
	if nodes == nil {
		return []string{
			client.Addr,
		}
	} else {
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
}

func (c *RedisClient) SlowLog() {

}

func (c *RedisClient) ConfigCheck() {

}

func (c *RedisClient) MemoryStat() {

}

func DeleteKeys(client interface{redis.Cmdable}, keys ...string) int64 {
	return client.Del(keys...).Val()
}

func ExpireKey(client interface{redis.Cmdable}, keys string, duration time.Duration) bool {
	return client.Expire(keys, duration).Val()
}

