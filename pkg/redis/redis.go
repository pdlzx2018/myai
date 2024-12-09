package redis

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"

	"github.com/pdlzx2018/myai/config"
)

var Client *redis.Client

// Init 初始化Redis连接
func Init() error {
	conf := config.GlobalConfig.Redis

	// 创建Redis客户端
	Client = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", conf.Host, conf.Port),
		Password: conf.Password,
		DB:       conf.DB,

		// 连接池配置
		PoolSize:     10, // 连接池大小
		MinIdleConns: 5,  // 最小空闲连接数

		// 超时配置
		DialTimeout:  5 * time.Second, // 建立连接超时
		ReadTimeout:  3 * time.Second, // 读取超时
		WriteTimeout: 3 * time.Second, // 写入超时
		PoolTimeout:  4 * time.Second, // 当连接池繁忙时，等待连接的超时时间

		// 心跳检测
		IdleCheckFrequency: 60 * time.Second, // 空闲连接检查的频率
		IdleTimeout:        5 * time.Minute,  // 空闲连接超时时间
	})

	// 测试连接
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := Client.Ping(ctx).Err(); err != nil {
		return fmt.Errorf("连接Redis失败: %v", err)
	}

	return nil
}

// Close 关闭Redis连接
func Close() {
	if Client != nil {
		Client.Close()
	}
}

// GetKey 获取键值
func GetKey(ctx context.Context, key string) (string, error) {
	return Client.Get(ctx, key).Result()
}

// SetKey 设置键值
func SetKey(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	return Client.Set(ctx, key, value, expiration).Err()
}

// DelKey 删除键
func DelKey(ctx context.Context, key string) error {
	return Client.Del(ctx, key).Err()
}
