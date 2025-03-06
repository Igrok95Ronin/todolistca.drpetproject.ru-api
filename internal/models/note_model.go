package models

import "time"

// AllNotes представляет таблицу заметок
type AllNotes struct {
	ID        int64     `json:"id" gorm:"primaryKey"`
	Note      string    `json:"note"`
	Completed bool      `json:"completed"`
	CreatedAt time.Time `json:"created_at"`
}

// TableName переопределяет имя таблицы для GORM
func (AllNotes) TableName() string {
	return "all_notes"
}

// Обновить заметку
type ModifiedEntry struct {
	ModEntry string `json:"modEntry"`
}

type Check struct {
	Check bool `json:"check"`
}
