package database

import (
	"log"
	"time"

	"github.com/aidarkhanov/nanoid"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Model struct {
	ID        string    `json:"id" gorm:"primaryKey"` // nanoid
	CreatedAt time.Time `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"autoUpdateTime"`
}

func (model *Model) BeforeCreate(tx *gorm.DB) error {
	if model.ID == "" {
		model.ID = nanoid.New()
	}
	return nil
}

type Product struct {
	Model
	Name      string   `json:"name" gorm:"not null"`
	Price     int      `json:"price" gorm:"not null"` // in cents
	SectionID uint     `json:"-"`
	Section   *Section `json:"section,omitempty" gorm:"foreignKey:SectionID"`
}

type Section struct {
	Model
	Name string `gorm:"not null"`
}

func SetupDatabase() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Automatically migrate the schema
	err = db.AutoMigrate(&Product{}, &Section{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}
	return db
}
