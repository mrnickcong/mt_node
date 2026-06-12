// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package svc

import (
	"context"
	"fmt"
	"mt_node/internal/config"
	"mt_node/internal/model"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
)

type ServiceContext struct {
	Config config.Config
	// 数据库
	DB *pgxpool.Pool
	// Redis
	Redis *redis.Client

	// Models
	ChainInfoModel *model.ChainInfoModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	// 初始化数据库连接
	dbURL := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=%s",
		c.Database.User,
		c.Database.Password,
		c.Database.Host,
		c.Database.Port,
		c.Database.Name,
		c.Database.SSLMode,
	)

	poolConfig, err := pgxpool.ParseConfig(dbURL)
	if err != nil {
		panic(fmt.Sprintf("Failed to parse database config: %v", err))
	}

	poolConfig.MaxConns = int32(c.Database.MaxOpenConns)
	poolConfig.MinConns = int32(c.Database.MaxIdleConns)
	//SQL查询日志
	poolConfig.ConnConfig.Tracer = &config.SqlTracer{}

	db, err := pgxpool.NewWithConfig(context.Background(), poolConfig)
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to database: %v", err))
	}

	// 初始化 Redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     c.Redis.Host,
		Password: c.Redis.Pass,
		DB:       c.Redis.DB,
	})

	// 测试 Redis 连接
	if _, err := rdb.Ping(context.Background()).Result(); err != nil {
		fmt.Printf("Warning: Failed to connect to Redis: %v\n", err)
	}

	// 初始化数据库模型
	//chainInfoModel := model.NewChainInfoModel(db)

	return &ServiceContext{
		Config:         c,
		DB:             db,
		Redis:          rdb,
		ChainInfoModel: model.NewChainInfoModel(db),
	}
}
