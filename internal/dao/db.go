package dao

import (
	"context"
	"gsn-wordle/internal/model"
	"gsn-wordle/internal/pkg/config"
	"gsn-wordle/internal/pkg/util"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB
var Redis *redis.Client

func InitDB() {
	// Build DSN string for PostgreSQL connection
	dsn := "port=" + strconv.Itoa(config.Config.Postgresql.Port) +
		" sslmode=disable" +
		" TimeZone=Asia/Shanghai"

	// 区分开发和生产环境
	if gin.Mode() == gin.ReleaseMode {
		dsn = dsn +
			" user=" + config.Config.Postgresql.User +
			" password=" + config.Config.Postgresql.Password +
			" host=" + config.Config.Postgresql.Host +
			" dbname=" + config.Config.Postgresql.Db
	} else {
		dsn = dsn +
			" user=" + config.Config.Postgresql.DevUser +
			" password=" + config.Config.Postgresql.DevPassword +
			" host=" + config.Config.Postgresql.DevHost +
			" dbname=" + config.Config.Postgresql.DevDb
	}

	// Open database connection with PostgreSQL
	writer, _ := util.GetFileWriter("log/db.log")
	DB, _ = gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy:         schema.NamingStrategy{SingularTable: true},
		SkipDefaultTransaction: true,
		DisableAutomaticPing:   true,
		Logger: logger.New(log.New(writer, "\n", log.LstdFlags), logger.Config{
			Colorful: false,
			LogLevel: logger.Warn,
		}),
	})

	// Migrate struct model to database
	err := DB.AutoMigrate(&model.Word{})
	if err != nil {
		log.Fatal(err)
	}
}

func InitRedis() {
	var opt *redis.Options

	// 区分开发和生产环境
	if gin.Mode() == gin.ReleaseMode {
		opt = &redis.Options{
			Addr:     config.Config.Redis.Host + ":" + strconv.Itoa(config.Config.Redis.Port),
			Password: config.Config.Redis.Password,
			DB:       config.Config.Redis.Db,
		}
	} else {
		opt = &redis.Options{
			Addr:     config.Config.Redis.DevHost + ":" + strconv.Itoa(config.Config.Redis.DevPort),
			Username: config.Config.Redis.DevUser,
			Password: config.Config.Redis.DevPassword,
			DB:       config.Config.Redis.DevDb,
		}
	}

	Redis = redis.NewClient(opt)

	_, err := Redis.Ping(context.Background()).Result()
	if err != nil {
		log.Fatal(err)
	}
}
