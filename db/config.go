package db

import (
	"fmt"

	"taskManager/db/model"
	"taskManager/envconf"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connection() (DB *gorm.DB) {
	env, err := envconf.EnvSetting()
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
        panic(fmt.Sprintf("Gagal koneksi awal GORM: %v", err))
    }

	pool, err := db.DB()
    if err != nil {
        panic(fmt.Sprintf("Gagal mengambil SQL DB Pool: %v", err))
    }

	pool.SetMaxIdleConns(10)
	pool.SetMaxOpenConns(10)
	pool.SetConnMaxLifetime(10)
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&model.LoginUser{}, &model.Task{}, &model.User{})
	fmt.Println("Connected", db)
	return db
}
