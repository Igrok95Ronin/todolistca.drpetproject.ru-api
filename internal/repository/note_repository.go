package repository

import (
	"github.com/Igrok95Ronin/todolistca.drpetproject.ru-api.git/internal/models"
	"github.com/Igrok95Ronin/todolistca.drpetproject.ru-api.git/pkg/logging"
	"gorm.io/gorm"
)

// NoteRepository - интерфейс для работы с заметками
type NoteRepository interface {
	GetAllNotes() ([]models.AllNotes, error)
	CreateNote(note *models.AllNotes) error
}

type noteRepository struct {
	db     *gorm.DB
	logger *logging.Logger
}

func NewNoteRepository(db *gorm.DB, logger *logging.Logger) NoteRepository {
	return &noteRepository{
		db:     db,
		logger: logger,
	}
}

// GetAllNotes Получить все посты
func (r *noteRepository) GetAllNotes() ([]models.AllNotes, error) {
	var note []models.AllNotes
	if err := r.db.Find(&note).Error; err != nil {
		r.logger.Error(err)
		return nil, err
	}

	return note, nil
}

// CreateNote сохраняет новую заметку в БД
func (r *noteRepository) CreateNote(note *models.AllNotes) error {
	return r.db.Create(&note).Error
}
