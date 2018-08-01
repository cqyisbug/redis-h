package main

import (
	"os"
	"fmt"
)

var (
	FormatBigKeysBody   string = "%-2d, %-10d, %-10d, %-6s, %-19s, %-10d, %-20s, %-20s\n"
	FormatBigKeysHeader string = "%-2s, %-20s, %-8s, %-20s ,%-15s, %-10s, %-10s, %-10s, %-20s\n"
)

func FileExists(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}

func GetBigKeysHeader() string {
	return fmt.Sprintf(FormatBigKeysHeader, "database", "key", "type", "size(Byte)", "size(MB)", "size(GB)", "element_count", "ttl", "expire")
}
