package database

import (
	"fmt"
	"log"
	"unila-helpdesk-backend/config"
	"unila-helpdesk-backend/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// InitDatabase - Inisialisasi koneksi database
func InitDatabase(cfg *config.Config) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Jakarta",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Gagal menghubungkan ke database: %v", err)
	}

	log.Println("Berhasil terhubung ke database")

	// Migrasi model ke database
	AutoMigrate()

}

// AutoMigrate - Melakukan migrasi model ke database
func AutoMigrate() {
	err := DB.AutoMigrate(
		&models.User{},
		&models.ServiceCategory{},
		&models.Ticket{},
		&models.Questionnaire{},
		&models.Question{},
		&models.QuestionOption{},
		&models.SurverResponse{},
		&models.SurveyAnswer{},
		&models.Notification{},
		&models.CohortAnalysis{},
	)

	if err != nil {
		log.Fatalf("Gagal melakukan migrasi model: %v", err)
	}

	log.Println("Migrasi model berhasil")
}

// SeedServiceCategories - Melakukan seeding data kategori layanan
func SeedServiceCategories() {
	categories := []models.ServiceCategory{
		// kategori layanan yang memerlukan login
		{Name: "Website", RequiresLogin: true},
		{Name: "Jaringan Internet", RequiresLogin: true},
		{Name: "Siakadu", RequiresLogin: true},
		{Name: "Sistem Informasi", RequiresLogin: true},
		{Name: "Lainnya", RequiresLogin: true},

		// kategori layanan yang tidak memerlukan login
		{Name: "Lupa Password", RequiresLogin: false},
		{Name: "Buat Email unila.ac.id", RequiresLogin: false},
		{Name: "Buat SSO Unila", RequiresLogin: false},
	}

	for _, category := range categories {
		var existingCategory models.ServiceCategory
		result := DB.Where("name = ?", category.Name).First(&existingCategory)

		if result.Error != nil {
			// Kategori belum ada, buat baru
			if err := DB.Create(&category).Error; err != nil {
				log.Printf("Gagal membuat kategori %s: %v", category.Name, err)
			} else {
				log.Printf("Kategori '%s' berhasil dibuat", category.Name)
			}
		}
	}
}

func GetDB() *gorm.DB {
	return DB
}
