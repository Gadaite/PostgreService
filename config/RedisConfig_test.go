package config

import (
	"github.com/gomodule/redigo/redis"
	"reflect"
	"testing"
)

func TestAddRedisKeyStringValue(t *testing.T) {
	type args struct {
		key   string
		value string
		conn  redis.Conn
	}
	conn, _ := redis.Dial("tcp", "1.15.94.16:6380", redis.DialDatabase(0))
	conn1, _ := redis.Dial("tcp", "1.15.94.16:6380", redis.DialDatabase(1))
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{"case 1", args{"visit", "visitValue", conn}, false},
		{"case 2", args{"key", "value", conn1}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AddRedisKeyStringValue(tt.args.key, tt.args.value, tt.args.conn); got != tt.want {
				t.Errorf("AddRedisKeyStringValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetRedisValueByKey(t *testing.T) {
	type args struct {
		key  string
		conn redis.Conn
	}
	conn, _ := redis.Dial("tcp", "1.15.94.16:6380", redis.DialDatabase(0))
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{"case 1", args{"visit", conn}, "12"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetRedisValueByKey(tt.args.key, tt.args.conn); got != tt.want {
				t.Errorf("GetRedisValueByKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetRedisAllKeys(t *testing.T) {
	type args struct {
		conn redis.Conn
	}
	conn, _ := redis.Dial("tcp", "1.15.94.16:6380", redis.DialDatabase(0))
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
		{"case 1", args{conn}, GetRedisAllKeys(conn)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetRedisAllKeys(tt.args.conn); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetRedisAllKeys() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAddRedisKeyFloatMapValue(t *testing.T) {
	type args struct {
		key   string
		value map[string]float64
		conn  redis.Conn
	}
	conn, _ := redis.Dial("tcp", "1.15.94.16:6380", redis.DialDatabase(1))
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{"case 1", args{"keyFloatMap", map[string]float64{
			"key1": 1.01,
			"key2": 2.02,
		}, conn}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AddRedisKeyFloatMapValue(tt.args.key, tt.args.value, tt.args.conn); got != tt.want {
				t.Errorf("AddRedisKeyFloatMapValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAddRedisKeyIntMapValue(t *testing.T) {
	type args struct {
		key   string
		value map[string]int
		conn  redis.Conn
	}
	conn, _ := redis.Dial("tcp", "1.15.94.16:6380", redis.DialDatabase(1))
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{"case 1", args{"keyIntMap", map[string]int{
			"key1": 1,
			"key2": 2,
		}, conn}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AddRedisKeyIntMapValue(tt.args.key, tt.args.value, tt.args.conn); got != tt.want {
				t.Errorf("AddRedisKeyIntMapValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAddRedisKeyStringListMapValue(t *testing.T) {
	type args struct {
		key   string
		value []string
		conn  redis.Conn
	}
	conn, _ := redis.Dial("tcp", "1.15.94.16:6380", redis.DialDatabase(1))
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{"case 1", args{"keyListMap", []string{"1", "2"}, conn}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AddRedisKeyStringListMapValue(tt.args.key, tt.args.value, tt.args.conn); got != tt.want {
				t.Errorf("AddRedisKeyStringListMapValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAddRedisKeyStringMapValue(t *testing.T) {
	type args struct {
		key   string
		value map[string]string
		conn  redis.Conn
	}
	conn, _ := redis.Dial("tcp", "1.15.94.16:6380", redis.DialDatabase(1))
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{"case 1", args{"keyStringMap", map[string]string{
			"key1": "value1",
			"key2": "value2",
		}, conn}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AddRedisKeyStringMapValue(tt.args.key, tt.args.value, tt.args.conn); got != tt.want {
				t.Errorf("AddRedisKeyStringMapValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpdateRedisKeyFloatMapValue(t *testing.T) {
	type args struct {
		key   string
		value map[string]float64
		conn  redis.Conn
	}
	conn, _ := redis.Dial("tcp", "1.15.94.16:6380", redis.DialDatabase(1))
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{"case 1", args{"keyFloatMap", map[string]float64{
			"key3": 3.03,
			"key4": 4.04,
		}, conn}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UpdateRedisKeyFloatMapValue(tt.args.key, tt.args.value, tt.args.conn); got != tt.want {
				t.Errorf("UpdateRedisKeyFloatMapValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpdateRedisKeyIntMapValue(t *testing.T) {
	type args struct {
		key   string
		value map[string]int
		conn  redis.Conn
	}
	conn, _ := redis.Dial("tcp", "1.15.94.16:6380", redis.DialDatabase(1))
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{"case 1", args{"keyIntMap", map[string]int{
			"key3": 3,
			"key4": 4,
		}, conn}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UpdateRedisKeyIntMapValue(tt.args.key, tt.args.value, tt.args.conn); got != tt.want {
				t.Errorf("UpdateRedisKeyIntMapValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpdateRedisKeyStringListMapValue(t *testing.T) {
	type args struct {
		key   string
		value []string
		conn  redis.Conn
	}
	conn, _ := redis.Dial("tcp", "1.15.94.16:6380", redis.DialDatabase(1))
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{"case 1", args{"keyListMap", []string{"3", "4"}, conn}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UpdateRedisKeyStringListMapValue(tt.args.key, tt.args.value, tt.args.conn); got != tt.want {
				t.Errorf("UpdateRedisKeyStringListMapValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpdateRedisKeyStringMapValue(t *testing.T) {
	type args struct {
		key   string
		value map[string]string
		conn  redis.Conn
	}
	conn, _ := redis.Dial("tcp", "1.15.94.16:6380", redis.DialDatabase(1))
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{"case 1", args{"keyStringMap", map[string]string{
			"key3": "value3",
			"key4": "value4",
		}, conn}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UpdateRedisKeyStringMapValue(tt.args.key, tt.args.value, tt.args.conn); got != tt.want {
				t.Errorf("UpdateRedisKeyStringMapValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpdateRedisKeyStringValue(t *testing.T) {
	type args struct {
		key   string
		value string
		conn  redis.Conn
	}
	conn, _ := redis.Dial("tcp", "1.15.94.16:6380", redis.DialDatabase(1))
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{"case 2", args{"key", "valueUpdate", conn}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UpdateRedisKeyStringValue(tt.args.key, tt.args.value, tt.args.conn); got != tt.want {
				t.Errorf("UpdateRedisKeyStringValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
