package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"

	"github.com/pdlzx2018/myai/config"
	"github.com/pdlzx2018/myai/internal/api/router"
	"github.com/pdlzx2018/myai/internal/model"
	"github.com/pdlzx2018/myai/pkg/database"
	"github.com/pdlzx2018/myai/pkg/redis"
)

var (
	g errgroup.Group
)

func main() {
	// 设置生产模式
	gin.SetMode(gin.ReleaseMode)

	// 加载配置
	if err := config.Load(); err != nil {
		log.Fatalf("加载配置失败: %v", err)
	}

	// 初始化数据库连接池
	if err := database.Init(); err != nil {
		log.Fatalf("初始化数据库失败: %v", err)
	}
	defer database.Close()

	// 配置数据库连接池
	sqlDB, err := database.DB.DB()
	if err != nil {
		log.Fatalf("获取数据库连接池失败: %v", err)
	}
	// 设置最大连接数
	sqlDB.SetMaxOpenConns(100)
	// 设置最大空闲连接数
	sqlDB.SetMaxIdleConns(10)
	// 设置连接最大存活时间
	sqlDB.SetConnMaxLifetime(time.Hour)

	// 自动迁移数据库表
	if err := database.DB.AutoMigrate(&model.User{}, &model.Chat{}); err != nil {
		log.Fatalf("数据库迁移失败: %v", err)
	}

	// 初始化Redis连接池
	if err := redis.Init(); err != nil {
		log.Fatalf("初始化Redis失败: %v", err)
	}
	defer redis.Close()

	// 创建 gin 引擎
	r := router.SetupRouter()

	// 创建 HTTP 服务器
	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
		// 读取请求头的超时时间
		ReadHeaderTimeout: 20 * time.Second,
		// 读取请求体的超时时间
		ReadTimeout: 60 * time.Second,
		// 写入响应的超时时间
		WriteTimeout: 60 * time.Second,
		// 空闲连接超时时间
		IdleTimeout: 120 * time.Second,
		// 最大请求头大小
		MaxHeaderBytes: 1 << 20, // 1MB
	}

	// 在错误组中启动服务器
	g.Go(func() error {
		return srv.ListenAndServe()
	})

	// 优雅关闭
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit
	log.Println("正在关闭服务器...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("服务器关闭失败: %v", err)
	}

	if err := g.Wait(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("服务器运行错误: %v", err)
	}

	log.Println("服务器已成功关闭")
}
