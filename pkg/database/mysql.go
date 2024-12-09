package database

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"

	"github.com/pdlzx2018/myai/config"
)

var DB *gorm.DB

// Init 初始化数据库连接
func Init() error {
	conf := config.GlobalConfig.MySQL
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local&timeout=10s",
		conf.User,
		conf.Password,
		conf.Host,
		conf.Port,
		conf.DBName,
	)

	// 配置 GORM
	gormConfig := &gorm.Config{
		// 设置日志级别
		Logger: logger.Default.LogMode(logger.Info),
		// 禁用默认事务
		SkipDefaultTransaction: true,
		// 命名策略
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 使用单数表名
		},
	}

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                      dsn,  // 数据库连接配置
		DefaultStringSize:        256,  // string 类型字段的默认长度
		DisableDatetimePrecision: true, // 禁用 datetime 精度
		DontSupportRenameIndex:   true, // 重命名索引时采用删除并新建的方式
		DontSupportRenameColumn:  true, // 用 `change` 重命名列
	}), gormConfig)

	if err != nil {
		return fmt.Errorf("连接数据库失败: %v", err)
	}

	// 获取通用数据库对象 sql.DB
	sqlDB, err := db.DB()
	if err != nil {
		return fmt.Errorf("获取数据库实例失败: %v", err)
	}

	// 设置连接池参数
	sqlDB.SetMaxIdleConns(10)           // 最大空闲连接数
	sqlDB.SetMaxOpenConns(100)          // 最大打开连接数
	sqlDB.SetConnMaxLifetime(time.Hour) // 连接最大复用时间

	DB = db
	return nil
}

// Close 关闭数据库连接
func Close() {
	if DB != nil {
		sqlDB, err := DB.DB()
		if err != nil {
			return
		}
		sqlDB.Close()
	}
}
