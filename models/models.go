package models

import (
	"time"

	"gorm.io/gorm"
)

// User - Model untuk pengguna
type User struct {
	ID           uint           `gorm:"primaryKey" json:"id"`
	Username     string         `gorm:"uniqueIndex;not null" json:"username"`
	Password     string         `gorm:"not null" json:"-"`
	Email        string         `gorm:"uniqueIndex;not null" json:"email"`
	Name         string         `gorm:"not null" json:"name"`
	EntityType   string         `gorm:"not null" json:"entity_type"` // "student", "lecturer", "staff"
	IsRegistered bool           `gorm:"default:false" json:"is_registered"`
	FCMToken     string         `gorm:"default:''" json:"fcm_token"`
	CreatedAt    time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt    time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"-"`

	// Relasi
	Tickets         []Ticket         `json:"tickets, omitempty"`
	SurverResponses []SurverResponse `json:"survey_responses, omitempty"`
	Notifications   []Notification   `json:"notifications, omitempty"`
}

// ServiceCategory - Model untuk kategori layanan
type ServiceCategory struct {
	ID            uint           `gorm:"primaryKey" json:"id"`
	Name          string         `gorm:"uniqueIndex;not null" json:"name"`
	RequiresLogin bool           `gorm:"default:true" json:"requires_login"`
	CreatedAt     time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt     time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"-"`

	// Relasi
	Tickets       []Ticket        `json:"tickets, omitempty"`
	Questionnaire []Questionnaire `json:"questionnaires, omitempty"`
}

// Ticket - Model untuk tiket layanan
type Ticket struct {
	ID                uint   `gorm:"primaryKey" json:"id"`
	TicketNumber      string `gorm:"uniqueIndex;not null" json:"ticket_number"`
	Title             string `gorm:"not null" json:"title"`
	Description       string `gorm:"not null" json:"description"`
	Status            string `gorm:"not null" json:"status"`
	Priority          string `gorm:"not null" json:"priority"`
	ServiceCategoryID uint   `gorm:"not null" json:"service_category_id"`

	// Field untuk user terdaftar
	UserID         *uint  `gorm:"index" json:"user_id"`        // Nullable untuk user yang tidak terdaftar
	EntityType     string `gorm:"not null" json:"entity_type"` // "student", "lecturer", "staff"
	AttachmentPath string `json:"attachment_path"`

	// Field untuk user tidak terdaftar
	GuestName     string `gorm:"not null" json:"guest_name"`
	GuestEmail    string `gorm:"not null" json:"guest_email"`
	GuestUserType string `gorm:"not null" json:"guest_usertype"` // "student", "lecturer", "staff"
	IDCardPath    string `json:"id_card_path"`
	SelfiePath    string `json:"selfie_path"`

	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	// Relasi
	User            *User           `json:"user, omitempty"`
	ServiceCategory ServiceCategory `json:"service_category, omitempty"`
	SurveyResponse  *SurverResponse `json:"survey_response, omitempty"`
}

// Questionnaire - Model untuk kuesioner
type Questionnaire struct {
	ID                uint           `gorm:"primaryKey" json:"id"`
	ServiceCategoryID uint           `gorm:"not null" json:"service_category_id"`
	Title             string         `gorm:"not null" json:"title"`
	Description       string         `gorm:"not null" json:"description"`
	IsActive          bool           `gorm:"default:true" json:"is_active"`
	CreatedAt         time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt         time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt         gorm.DeletedAt `gorm:"index" json:"-"`

	// Relasi
	ServiceCategory ServiceCategory  `json:"service_category"`
	Questions       []Question       `json:"questions, omitempty"`
	SurveyResponses []SurveyResponse `json:"survey_responses, omitempty"`
}

// Question - Model untuk pertanyaan dalam kuesioner
type Question struct {
	ID              uint           `gorm:"primaryKey" json:"id"`
	QuestionnaireID uint           `gorm:"not null" json:"questionnaire_id"`
	QuestionText    string         `gorm:"not null" json:"question_text"`
	QuestionType    string         `gorm:"not null" json:"question_type"` // "text", "single_choice", "multiple_choice"
	Isrequired      bool           `gorm:"default:true" json:"is_required"`
	OrderNumber     int            `gorm:"not null" json:"order_number"`
	CreatedAt       time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt       time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt       gorm.DeletedAt `gorm:"index" json:"-"`

	// Relasi
	Questionnaire   Questionnaire    `json:"questionnaire"`
	QuestionOptions []QuestionOption `json:"question_options, omitempty"`
	SurveyAnswers   []SurveyAnswer   `json:"survey_answers, omitempty"`
}

// QuestionOption - Model untuk opsi pertanyaan dalam kuesioner
type QuestionOption struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	QuestionID  uint           `gorm:"not null" json:"question_id"`
	OptionText  string         `gorm:"not null" json:"option_text"`
	OptionValue string         `gorm:"not null" json:"option_value"` // Nilai yang dikirimkan jika opsi ini dipilih
	OrderNumber int            `gorm:"not null" json:"order_number"`
	CreatedAt   time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`

	// Relasi
	Question      Question       `json:"question"`
	SurveyAnswers []SurveyAnswer `json:"survey_answers, omitempty"`
}

// SurverResponse - Model untuk respons survei
type SurverResponse struct {
	ID                uint           `gorm:"primaryKey" json:"id"`
	TicketID          uint           `gorm:"not null" json:"ticket_id"`
	UserID            *uint          `gorm:"index" json:"user_id"` // Nullable untuk user yang tidak terdaftar
	QuestionnaireID   uint           `gorm:"not null" json:"questionnaire_id"`
	SatisfactionScore float64        `gorm:"not null" json:"satisfaction_rating"`
	CreatedAt         time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt         time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt         gorm.DeletedAt `gorm:"index" json:"-"`

	// Relasi
	Ticket        Ticket         `json:"ticket"`
	User          *User          `json:"user, omitempty"`
	Questionnaire Questionnaire  `json:"questionnaire"`
	SurveyAnswers []SurveyAnswer `json:"survey_answers, omitempty"`
}

// SurveyAnswer - Model untuk jawaban survei
type SurveyAnswer struct {
	ID               uint           `gorm:"primaryKey" json:"id"`
	SurverResponseID uint           `gorm:"not null" json:"survey_response_id"`
	QuestionID       uint           `gorm:"not null" json:"question_id"`
	QuestionOptionID *uint          `gorm:"index" json:"question_option_id"` // Nullable jika pertanyaan tipe teks
	AnswerText       string         `gorm:"not null" json:"answer_text"`     // Jawaban teks jika pertanyaan tipe teks
	AnswerValue      string         `gorm:"not null" json:"answer_value"`    // Nilai yang dikirimkan jika opsi dipilih
	CreatedAt        time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt        time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt        gorm.DeletedAt `gorm:"index" json:"-"`

	// Relasi
	SurverResponse SurverResponse  `json:"survey_response"`
	Question       Question        `json:"question"`
	QuestionOption *QuestionOption `json:"question_option, omitempty"`
}

// Notification - Model untuk notifikasi
type Notification struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	UserID    uint           `gorm:"not null" json:"user_id"`
	Title     string         `gorm:"not null" json:"title"`
	Message   string         `gorm:"not null" json:"message"`
	Type      string         `gorm:"not null" json:"type"` // "ticket_update", "survey_response", etc.
	IsRead    bool           `gorm:"default:false" json:"is_read"`
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	// Relasi
	User User `json:"user"`
}

// CohortAnalysis - Model untuk analisis cohort
type CohortAnalysis struct {
	ID                  uint      `gorm:"primaryKey" json:"id"`
	Period              string    `gorm:"not null" json:"period"` // Format: "YYYY-MM"
	ServiceCategoryID   uint      `gorm:"not null" json:"service_category_id"`
	UserType            string    `gorm:"not null" json:"user_type"`            // "student", "lecturer", "staff"
	UsageCount          int       `gorm:"not null" json:"usage_count"`          // Jumlah penggunaan layanan
	AverageSatisfaction float64   `gorm:"not null" json:"average_satisfaction"` // Rata-rata kepuasan
	CreatedAt           time.Time `gorm:"autoCreateTime" json:"created_at"`
}
