// инициализация DB и запуск gRPC
package server

import (
	"log"
	"user-service/internal/database"
	"user-service/internal/user"
)

func main() {

	database.InitDB()
	if err := database.DB.AutoMigrate(&user.User{}); err != nil {
		log.Fatalf("Migration failed: %v", err)
	}
}
