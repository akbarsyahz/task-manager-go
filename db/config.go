package db

import (
	"fmt"

	"taskManager/db/model"
	"taskManager/envconf"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Connection its for connecting to database and auto migrate for now
func Connection() (DB *gorm.DB) {
	env, err := envconf.EnvSetting()
	if err != nil {
		panic(fmt.Errorf("%s", fmt.Sprintf("ENV Gone: %v", err)))
	}
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		env.Host,
		env.User,
		env.Pass,
		env.Name,
		env.Port,
		env.SSL,
		env.Timezone,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Errorf("%s", fmt.Sprintf("Connection Failed when Instance GORM: %v", err)))
	}

	pool, err := db.DB()
	if err != nil {
		panic(fmt.Sprintf("Failed to take SQL DB Pool: %v", err))
	}

	pool.SetMaxIdleConns(10)
	pool.SetMaxOpenConns(10)
	pool.SetConnMaxLifetime(10)

	errMigrate := db.AutoMigrate(&model.LoginUser{}, &model.Task{}, &model.User{}, &model.UserRole{})
	if errMigrate != nil {
		panic(fmt.Sprintf("Error when migrate: %v", err))
	}
	return db
}
