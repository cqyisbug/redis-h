package main

import "time"

func Nano2Second(d time.Duration) int64 {
	return int64(d / time.Second)
}

func GetExpireTime(d time.Duration) string{
	s:= Nano2Second(d)
	switch s {
	case -2:
		return "expired"
	case -1:
		return "not exists"
	default:
		now := time.Now()
		exp := now.Add(d)
		return exp.Format("2006-01-02 15:04:05")
	}
}
