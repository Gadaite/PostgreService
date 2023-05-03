package config

import (
	"github.com/gomodule/redigo/redis"
	"log"
)

func GetRedisConnect(network string, address string, db int) redis.Conn {
	conn, err := redis.Dial(network, address, redis.DialDatabase(db))
	if err != nil {
		panic(err)
	}
	return conn
}

func GetRedisValueByKey(key string, conn redis.Conn) string {
	value, err := redis.String(conn.Do("Get", key))
	if err != nil {
		panic(err)
	}
	return value
}

func GetRedisAllKeys(conn redis.Conn) []string {
	keys, err := redis.Strings(conn.Do("Keys", "*"))
	if err != nil {
		panic(keys)
	}
	return keys
}

func GetRedisStringMapByKey(key string, conn redis.Conn) map[string]string {
	mapValue, err := redis.StringMap(conn.Do("Get", key))
	if err != nil {
		panic(err)
	}
	return mapValue
}

func GetRedisIntMapByKey(key string, conn redis.Conn) map[string]int {
	mapValue, err := redis.IntMap(conn.Do("Get", key))
	if err != nil {
		panic(err)
	}
	return mapValue
}

func GetRedisFloatMapByKey(key string, conn redis.Conn) map[string]float64 {
	mapValue, err := redis.Float64Map(conn.Do("Get", key))
	if err != nil {
		panic(err)
	}
	return mapValue
}

func JudgeContainKey(key string, conn redis.Conn) bool {
	res, err := redis.Bool(conn.Do("exists", key))
	if err != nil {
		panic(err)
	}
	return res
}

func AddRedisKeyStringValue(key string, value string, conn redis.Conn) bool {
	containOrNot := JudgeContainKey(key, conn)
	if containOrNot {
		log.Fatalf("%s is exists,do nothing!", key)
		return false
	} else {
		res, err := redis.String(conn.Do("Set", key, value))
		if err != nil {
			panic(err)
		}
		log.Printf("Set %s %s %s", key, value, res)
		return true
	}
}

func AddRedisKeyStringMapValue(key string, value map[string]string, conn redis.Conn) bool {
	containOrNot := JudgeContainKey(key, conn)
	if containOrNot {
		log.Fatalf("%s is exists,do nothing!", key)
		return false
	} else {
		res, err := redis.String(conn.Do("Set", key, value))
		if err != nil {
			panic(err)
		}
		log.Printf("Set %s %s %s", key, value, res)
		return true
	}
}

func AddRedisKeyStringListMapValue(key string, value []string, conn redis.Conn) bool {
	containOrNot := JudgeContainKey(key, conn)
	if containOrNot {
		log.Fatalf("%s is exists,do nothing!", key)
		return false
	} else {
		res, err := redis.String(conn.Do("Set", key, value))
		if err != nil {
			panic(err)
		}
		log.Printf("Set %s %s %s", key, value, res)
		return true
	}
}

func AddRedisKeyIntMapValue(key string, value map[string]int, conn redis.Conn) bool {
	containOrNot := JudgeContainKey(key, conn)
	if containOrNot {
		log.Fatalf("%s is exists,do nothing!", key)
		return false
	} else {
		res, err := redis.String(conn.Do("Set", key, value))
		if err != nil {
			panic(err)
		}
		log.Printf("Set %s %v %v", key, value, res)
		return true
	}
}

func AddRedisKeyFloatMapValue(key string, value map[string]float64, conn redis.Conn) bool {
	containOrNot := JudgeContainKey(key, conn)
	if containOrNot {
		log.Fatalf("%s is exists,do nothing!", key)
		return false
	} else {
		res, err := redis.String(conn.Do("Set", key, value))
		if err != nil {
			panic(err)
		}
		log.Printf("Set %s %v %v", key, value, res)
		return true
	}
}

func UpdateRedisKeyStringValue(key string, value string, conn redis.Conn) bool {
	containOrNot := JudgeContainKey(key, conn)
	if !containOrNot {
		log.Fatalf("%s is not exists,do nothing!", key)
		return false
	} else {
		res, err := redis.String(conn.Do("Set", key, value))
		if err != nil {
			panic(err)
		}
		log.Printf("Set %s %s %s", key, value, res)
		return true
	}
}

func UpdateRedisKeyStringMapValue(key string, value map[string]string, conn redis.Conn) bool {
	containOrNot := JudgeContainKey(key, conn)
	if !containOrNot {
		log.Fatalf("%s is not exists,do nothing!", key)
		return false
	} else {
		res, err := redis.String(conn.Do("set", key, value))
		if err != nil {
			panic(err)
		}
		log.Printf("Set %s %s %s", key, value, res)
		return true
	}
}

func UpdateRedisKeyStringListMapValue(key string, value []string, conn redis.Conn) bool {
	containOrNot := JudgeContainKey(key, conn)
	if !containOrNot {
		log.Fatalf("%s is not exists,do nothing!", key)
		return false
	} else {
		res, err := redis.String(conn.Do("Set", key, value))
		if err != nil {
			panic(err)
		}
		log.Printf("Set %s %s %s", key, value, res)
		return true
	}
}

func UpdateRedisKeyIntMapValue(key string, value map[string]int, conn redis.Conn) bool {
	containOrNot := JudgeContainKey(key, conn)
	if !containOrNot {
		log.Fatalf("%s is not exists,do nothing!", key)
		return false
	} else {
		res, err := redis.String(conn.Do("Set", key, value))
		if err != nil {
			panic(err)
		}
		log.Printf("Set %s %v %v", key, value, res)
		return true
	}
}

func UpdateRedisKeyFloatMapValue(key string, value map[string]float64, conn redis.Conn) bool {
	containOrNot := JudgeContainKey(key, conn)
	if !containOrNot {
		log.Fatalf("%s is not exists,do nothing!", key)
		return false
	} else {
		res, err := redis.String(conn.Do("Set", key, value))
		if err != nil {
			panic(err)
		}
		log.Printf("Set %s %v %v", key, value, res)
		return true
	}
}
