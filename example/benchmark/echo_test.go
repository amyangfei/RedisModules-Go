package benchmark

import (
	"github.com/garyburd/redigo/redis"
	"math/rand"
	"testing"
	"time"
)

func newPool(server, password string) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", server)
			if err != nil {
				return nil, err
			}
			if password != "" {
				if _, err := c.Do("AUTH", password); err != nil {
					c.Close()
					return nil, err
				}
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}

var (
	Pool        *redis.Pool
	RedisServer string = "127.0.0.1:6379"
	EchoMsg     []byte
)

func init() {
	Pool = newPool(RedisServer, "")

	letters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	// 42KB string
	strlen := 42 * 1024
	EchoMsg = make([]byte, strlen)
	for i := 0; i < strlen; i++ {
		EchoMsg[i] = letters[rand.Intn(len(letters))]
	}
}

func RedisEcho(pool *redis.Pool, cmd string) {
	conn := pool.Get()
	if _, err := conn.Do(cmd, EchoMsg); err != nil {
		panic(err)
	}
}

func Benchmark_Echo1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RedisEcho(Pool, "hello.echo1")
	}
}

func Benchmark_Echo2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RedisEcho(Pool, "hello.echo2")
	}
}

func Benchmark_Echo3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RedisEcho(Pool, "hello.echo3")
	}
}

// TODO: hello.echo4 has unexpected memory problem
// func Benchmark_Echo4(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		RedisEcho(Pool, "hello.echo4")
// 	}
// }

func Benchmark_Echo6(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RedisEcho(Pool, "hello.echo6")
	}
}

func Benchmark_Echo7(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RedisEcho(Pool, "hello.echo7")
	}
}
