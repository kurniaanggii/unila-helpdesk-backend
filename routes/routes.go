package routes

import (
	"unila-helpdesk-backend/handlers"

	"github.com/gin-gonic/gin"
)

// SetupRoutes - Inisialisasi rute aplikasi
func SetupRoutes() *gin.Engine {
	r := gin.Default()

	// Middleware untuk CORS
	r.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204) // No Content
			return
		}

		c.Next()
	})

	// Initialize Handlers
	authHandler := handlers.NewAuthHandler()
	ticketHandler := handlers.NewTicketHandler()
	surveyHandler := handlers.NewSurveyHandler()
	analyticsHandler := handlers.NewAnalyticsHandler()
	notificationHandler := handlers.NewNotificationHandler()

	// Health Check Route
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "OK"})
	})

	// API V1 Routes
	apiV1 := r.Group("/api/v1")

	// Authentication Routes (public)
	auth := apiV1.Group("/auth")
	{
		auth.POST("/login", authHandler.Login)
		auth.GET("/validate", authHandler.ValidateToken)
		auth.POST("/logut", authHandler.Logout)
	}

	// Authenticated Routes (protected by middleware)
	authProtected := apiV1.Group("/auth")
	authProtected.Use(handlers.AuthMiddleware())
	{
		authProtected.GET("/profile", authHandler.GetProfile)
		authProtected.POST("/fcm-token", authHandler.UpdateFCMToken)
	}

	// Admin auth Routes
	adminAuth := apiV1.Group("/admin/auth")
	adminAuth.Use(middleware.AuthMiddleware(), middleware.AdminMiddleware())
	{
		adminAuth.GET("/users", authHandler.GetAllUsers)
	}

	// Ticket Routes (public untuk guest dan pencarian)
	tickets := apiV1.Group("/tickets")
	{
		tickets.POST("", ticketHandler.CreateTicket)                    // Create ticket
		tickets.GET("/search", ticketHandler.SearchTickets)             // Search tickets
		tickets.GET("/number/:number", ticketHandler.GetTicketByNumber) // Get ticket by number
		tickets.GET("/categories", ticketHandler.GetServiceCategories)  // Get service categories
	}

	// Ticket Routes (authenticated untuk user terdaftar)
	ticketsAuth := apiV1.Group("/tickets")
	ticketsAuth.Use(middleware.AuthMiddleware())
	{
		ticketsAuth.GET("/my", ticketHandler.GetMyTickets)     // Get my tickets
		ticketsAuth.GET("/:id", ticketHandler.GetTicketByID)   // Get ticket by ID
		ticketsAuth.PUT("/:id", ticketHandler.UpdateTicket)    // Update ticket
		ticketsAuth.DELETE("/:id", ticketHandler.DeleteTicket) // Delete ticket
	}

	// Ticket Routes (admin)
	ticketsAdmin := apiV1.Group("/admin/tickets")
	ticketsAdmin.Use(middleware.AuthMiddleware(), middleware.AdminMiddleware())
	{
		ticketsAdmin.GET("", ticketHandler.GetAllTickets) // Get all tickets
	}

	// Survey routes (public - untuk melihat kuesioner)
	surveys := v1.Group("/surveys")
	{
		surveys.GET("/questionnaires/category/:categoryId", surveyHandler.GetQuestionnaireByCategory)
		surveys.GET("/questionnaires/:id", surveyHandler.GetQuestionnaireByID)
	}

	// Survey routes (protected - untuk user terdaftar)
	surveysProtected := v1.Group("/surveys")
	surveysProtected.Use(middleware.AuthMiddleware())
	{
		surveysProtected.POST("/submit", surveyHandler.SubmitSurvey)                  // Submit survei
		surveysProtected.GET("/tickets/:ticketId", surveyHandler.GetSurveyByTicketID) // Survei berdasarkan tiket
	}

	// Survey routes (admin only)
	surveysAdmin := v1.Group("/surveys")
	surveysAdmin.Use(middleware.AuthMiddleware(), middleware.AdminMiddleware())
	{
		// Kelola kuesioner
		surveysAdmin.POST("/questionnaires", surveyHandler.CreateQuestionnaire)
		surveysAdmin.GET("/questionnaires", surveyHandler.GetAllQuestionnaires)
		surveysAdmin.PUT("/questionnaires/:id", surveyHandler.UpdateQuestionnaire)
		surveysAdmin.DELETE("/questionnaires/:id", surveyHandler.DeleteQuestionnaire)

		// Kelola pertanyaan
		surveysAdmin.POST("/questions", surveyHandler.CreateQuestion)
		surveysAdmin.DELETE("/questions/:id", surveyHandler.DeleteQuestion)

		// Kelola pilihan jawaban
		surveysAdmin.POST("/question-options", surveyHandler.CreateQuestionOption)
		surveysAdmin.DELETE("/question-options/:id", surveyHandler.DeleteQuestionOption)

		// Lihat hasil survei
		surveysAdmin.GET("/responses", surveyHandler.GetAllSurveyResponses)
		surveysAdmin.GET("/responses/category/:categoryId", surveyHandler.GetSurveyResponsesByCategory)
	}

	// Analytics routes (admin only)
	analytics := v1.Group("/analytics")
	analytics.Use(middleware.AuthMiddleware(), middleware.AdminMiddleware())
	{
		analytics.GET("/cohort", analyticsHandler.GetCohortAnalysis)                // Analisis kohort
		analytics.POST("/cohort/save", analyticsHandler.SaveCohortAnalysis)         // Simpan analisis kohort
		analytics.GET("/service-trends", analyticsHandler.GetServiceTrends)         // Tren layanan
		analytics.GET("/ticket-status", analyticsHandler.GetTicketStatusStats)      // Statistik status tiket
		analytics.GET("/user-entities", analyticsHandler.GetUserEntityStats)        // Statistik entitas user
		analytics.GET("/satisfaction-trend", analyticsHandler.GetSatisfactionTrend) // Tren kepuasan
		analytics.GET("/top-issues", analyticsHandler.GetTopIssues)                 // Masalah teratas
		analytics.GET("/resolution-time", analyticsHandler.GetResolutionTimeStats)  // Waktu penyelesaian
		analytics.GET("/dashboard", analyticsHandler.GetDashboardStats)             // Statistik dashboard
	}

	// Notification routes (admin only - untuk testing)
	notifications := v1.Group("/notifications")
	notifications.Use(middleware.AuthMiddleware(), middleware.AdminMiddleware())
	{
		notifications.POST("/send", notificationHandler.SendTestNotification)        // Send test notification
		notifications.POST("/subscribe", notificationHandler.SubscribeToTopic)       // Subscribe to topic
		notifications.POST("/unsubscribe", notificationHandler.UnsubscribeFromTopic) // Unsubscribe from topic
	}

	return r
}
