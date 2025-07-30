package main

import (
	"log"
	"unila-helpdesk-backend/config"
	"unila-helpdesk-backend/database"
	"unila-helpdesk-backend/routes"
	"unila-helpdesk-backend/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load konfigurasi
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	log.Println("Configuration loaded successfully")

	// initialize database connection
	database.InitDatabase(cfg)
	log.Println("Database berhasil diinisialisasi")

	// Inisialisasi Firebase FCM (opsional)
	if err := utils.InitFCM(cfg.FCMKeyPath); err != nil {
		log.Printf("Warning: Gagal menginisialisasi FCM: %v", err)
		log.Println("Aplikasi akan berjalan tanpa notifikasi push")
	}

	// Setup Gin mode
	gin.SetMode(gin.ReleaseMode)

	// Setup routes
	router := routes.SetupRouter()

	// Start the server
	log.Printf("Server berjalan di port %s", cfg.Port)
	log.Printf("API Documentation: http://localhost:%s/health", cfg.Port)
	log.Printf("Network Access: http://192.168.32.125:%s/health", cfg.Port)

	if err := router.Run("0.0.0.0:" + cfg.Port); err != nil {
		log.Fatal("Gagal menjalankan server:", err)
	}
}
