package db

import (
	"fmt"

	"taskManager/envconf"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connection() {
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
	db, _ := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	pool, err := db.DB()
	pool.SetMaxIdleConns(10)
	pool.SetMaxOpenConns(10)
	pool.SetConnMaxLifetime(10)
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected", db)
}
