package repository

import (
	"github.com/Igrok95Ronin/todolistca.drpetproject.ru-api.git/internal/models"
	"gorm.io/gorm"
)

// NoteRepository - интерфейс для работы с заметками
type NoteRepository interface {
	GetAllNotes() ([]models.AllNotes, error)
	CreateNote(note *models.AllNotes) error
	EditEntry(updatedEntry *models.ModifiedEntry, id int64) error
	DeleteEntry(id int64) error
	MarkCompleteEntry(check models.Check, id int64) error
	DeleteAllEntries() error
	DeleteAllMarkedEntries() error
}

type noteRepository struct {
	db *gorm.DB
}

func NewNoteRepository(db *gorm.DB) NoteRepository {
	return &noteRepository{
		db: db,
	}
}

// GetAllNotes Получить все посты
func (r *noteRepository) GetAllNotes() ([]models.AllNotes, error) {
	var notes []models.AllNotes
	err := r.db.Find(&notes).Error
	return notes, err
}

// CreateNote сохраняет новую заметку в БД
func (r *noteRepository) CreateNote(note *models.AllNotes) error {
	return r.db.Create(&note).Error
}

// EditEntry обновить заметку в БД
func (r *noteRepository) EditEntry(updatedEntry *models.ModifiedEntry, id int64) error {
	return r.db.Model(&models.AllNotes{}).Where("id = ?", id).Update("note", updatedEntry.ModEntry).Error
}

// DeleteEntry Удалить запись из БД
func (r *noteRepository) DeleteEntry(id int64) error {
	return r.db.Where("id = ?", id).Unscoped().Delete(&models.AllNotes{}).Error
}

// MarkCompleteEntry Отметить выполненную запись в БД
func (r *noteRepository) MarkCompleteEntry(check models.Check, id int64) error {
	return r.db.Model(&models.AllNotes{}).Where("id = ?", id).Update("completed", check.Check).Error
}

// DeleteAllEntries Удалить все записи из БД
func (r *noteRepository) DeleteAllEntries() error {
	return r.db.Unscoped().Where("1 = 1").Delete(&models.AllNotes{}).Error
}

// DeleteAllMarkedEntries Удалить все отмеченные записи из БД
func (r *noteRepository) DeleteAllMarkedEntries() error {
	return r.db.Unscoped().Where("completed = ?", true).Delete(&models.AllNotes{}).Error
}
