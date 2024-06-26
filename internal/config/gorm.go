package config

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

type GormConfig struct {
}

func NewDatabase(viper *viper.Viper, log *logrus.Logger) *gorm.DB {
	username := viper.GetString("database.username")
	password := viper.GetString("database.password")
	host := viper.GetString("database.host")
	port := viper.GetInt("database.port")
	database := viper.GetString("database.name")
	idleConnection := viper.GetInt("database.pool.idle")
	maxConnection := viper.GetInt("database.pool.max")
	maxLifeTimeConnection := viper.GetInt("database.pool.lifetime")

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Jakarta", host, port, username, password, database)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.New(&logrusWriter{Logger: log}, logger.Config{
			SlowThreshold:             time.Second * 5,
			Colorful:                  false,
			IgnoreRecordNotFoundError: true,
			ParameterizedQueries:      true,
			LogLevel:                  logger.Info,
		}),
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
	})

	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
		panic("Error : " + err.Error())
	}

	//err = db.Migrator().DropTable(
	//	&entity.Admin{},
	//	&entity.Book{},
	//	&entity.Class{},
	//	&entity.Family{},
	//	&entity.Schedule{},
	//	&entity.StatusActive{},
	//	&entity.Student{},
	//	&entity.Teacher{},
	//)
	//
	//if err != nil {
	//	log.Fatalf("Error migrating database: %v", err)
	//	panic("Error : " + err.Error())
	//}
	//
	//log.Info("HAPUS TABEL SUKSES")
	//
	//err = db.AutoMigrate(
	//	&entity.Admin{},
	//	&entity.Book{},
	//	&entity.Class{},
	//	&entity.Family{},
	//	&entity.Schedule{},
	//	&entity.StatusActive{},
	//	&entity.Student{},
	//	&entity.Teacher{},
	//)
	//
	//if err != nil {
	//	log.Fatalf("Error migrating database: %v", err)
	//	panic("Error : " + err.Error())
	//}

	//log.Info("MIGRATE TABEL SUKSES")

	connection, err := db.DB()
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
		panic("Error : " + err.Error())
	}

	connection.SetMaxIdleConns(idleConnection)
	connection.SetMaxOpenConns(maxConnection)
	connection.SetConnMaxLifetime(time.Second * time.Duration(maxLifeTimeConnection))

	return db

}

type logrusWriter struct {
	Logger *logrus.Logger
}

func (l *logrusWriter) Printf(message string, args ...interface{}) {
	l.Logger.Tracef(message, args...)
}
