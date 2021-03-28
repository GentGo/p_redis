package pedis

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
	"sync"
	"time"
)
var redisClient  *redis.Client
var reidsClient_Once sync.Once
var addr = "127.0.0.1:6379"

func Redis() *redis.Client{
	reidsClient_Once.Do(func() {
		redisClient = redis.NewClient(&redis.Options{
			Network:"tcp",
			Addr:addr,
			Password:"",
			DB:0,

			//连接池容量及闲置连接数
			PoolSize:15,//连接池数量
			MinIdleConns:10,//好比最少连接数

			//超时
			DialTimeout:5*time.Second,//连接建立超时时间
			ReadTimeout:5*time.Second,//读超时，默认3秒，-1表示取消读超时
			WriteTimeout:5*time.Second,//写超时，默认等于读超时
			PoolTimeout:5*time.Second,//当所有连接都处于繁忙状态时，客户端等待可用连接的最大等待时长

			//闲置连接检查包括IdleTimeOut,MaxConnAge
			IdleCheckFrequency:60*time.Second,//闲置连接检查的周期，默认1分钟，-1表示不做周期性检查
			IdleTimeout:5*time.Second,//闲置超时
			MaxConnAge:0*time.Second,//连接存活时长，从创建开始计时，超过指定时长则关闭连接

			//命令执行失败时的重试策略
			MaxRetries:0,//命令执行失败时，最多重试多少次，默认为0即不重试
			MinRetryBackoff:8*time.Millisecond,//每次计算重试间隔时间的下限，默认8毫秒，-1表示取消
			MaxRetryBackoff:512*time.Millisecond,//每次计算重试间隔时间的上限，默认512毫秒，-1表示取消
		})

		Pong,err:=redisClient.Ping(context.Background()).Result()
		if err!=nil{
			log.Fatal(fmt.Errorf("connect error:%s",err))
		}
		log.Println(Pong)
	})

	return redisClient
}