package mypkg

import (
	"fmt"

	"github.com/go-redis/redis"
)

// 定义一个RedisSingleObj结构体
type RedisSingleObj struct {
	Redis_host string
	Redis_port uint16
	Redis_auth string
	Database   int
	Db         *redis.Client
}

// 结构体InitSingleRedis方法: 用于初始化redis数据库
func (r *RedisSingleObj) InitSingleRedis() (err error) {
	// Redis连接格式拼接
	redisAddr := fmt.Sprintf("%s:%d", r.Redis_host, r.Redis_port)
	// Redis 连接对象: NewClient将客户端返回到由选项指定的Redis服务器。
	r.Db = redis.NewClient(&redis.Options{
		Addr:     redisAddr,    // redis服务ip:port
		Password: r.Redis_auth, // redis的认证密码
		DB:       r.Database,   // 连接的database库
		PoolSize: 10000,        // 连接池
	})
	fmt.Printf("Connecting Redis : %v\n", redisAddr)
	//延迟到程序结束关闭链接
	// defer r.Db.Close()

	// 验证是否连接到redis服务端
	res, err := r.Db.Ping().Result()
	if err != nil {
		fmt.Printf("Connect Failed! Err: %v\n", err)
		return err
	} else {
		fmt.Printf("Connect Successful! Ping => %v\n", res)
		return nil
	}
}
