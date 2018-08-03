package main

const (
	RedisMode      = "redis_mode"
	SectionServer  = "server"
	ModeCluster    = "cluster"
	ModeStandalone = "standalone"

	RoleMaster = "master"
	RoleSlave  = "slave"

	GlobalScanBatch int64 = 200

	BigKeysFile    = "bigkeys.csv"
	DeleteKeysFile = "deleteKeys.txt"
	DumpKeysFile   = "dumpkeys.json"

	TTLLessThan   = "<"
	TTLLessEqual  = "<="
	TTLGreatThan  = ">"
	TTLGreatEqual = ">="
	TTLBetween    = "<>"

	RedisTypeString = "string"
	RedisTypeHash   = "hash"
	RedisTypeSet    = "set"
	RedisTypeZSet   = "zset"
	RedisTypeList   = "list"
)

var (
	BigKeysHeader = []string{
		"db",
		"key",
		"type",
		"size(Byte)",
		"size(MB)",
		"size(GB)",
		"element_count",
		"expire",
	}
)
