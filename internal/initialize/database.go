package initialize

import (
	"gorm.io/driver/sqlite" // Sqlite driver based on CGO
	"gorm.io/gorm"
)

func ConnectToDB() {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
}
