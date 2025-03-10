package repository

import (
	"context"
	"github.com/Igrok95Ronin/todolistca.drpetproject.ru-api.git/internal/models"
	"gorm.io/gorm"
)

// NoteRepository - интерфейс для работы с заметками
type NoteRepository interface {
	GetAllNotes(ctx context.Context) ([]models.AllNotes, error)
	CreateNote(ctx context.Context, note *models.AllNotes) error
	EditEntry(ctx context.Context, updatedEntry *models.ModifiedEntry, id int64) error
	DeleteEntry(ctx context.Context, id int64) error
	MarkCompleteEntry(ctx context.Context, check models.Check, id int64) error
	DeleteAllEntries(ctx context.Context) error
	DeleteAllMarkedEntries(ctx context.Context) error
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
func (r *noteRepository) GetAllNotes(ctx context.Context) ([]models.AllNotes, error) {
	var notes []models.AllNotes
	err := r.db.WithContext(ctx).Find(&notes).Error
	return notes, err
}

// CreateNote сохраняет новую заметку в БД
func (r *noteRepository) CreateNote(ctx context.Context, note *models.AllNotes) error {
	return r.db.WithContext(ctx).Create(&note).Error
}

// EditEntry обновить заметку в БД
func (r *noteRepository) EditEntry(ctx context.Context, updatedEntry *models.ModifiedEntry, id int64) error {
	return r.db.WithContext(ctx).Model(&models.AllNotes{}).Where("id = ?", id).Update("note", updatedEntry.ModEntry).Error
}

// DeleteEntry Удалить запись из БД
func (r *noteRepository) DeleteEntry(ctx context.Context, id int64) error {
	return r.db.WithContext(ctx).Where("id = ?", id).Unscoped().Delete(&models.AllNotes{}).Error
}

// MarkCompleteEntry Отметить выполненную запись в БД
func (r *noteRepository) MarkCompleteEntry(ctx context.Context, check models.Check, id int64) error {
	return r.db.WithContext(ctx).Model(&models.AllNotes{}).Where("id = ?", id).Update("completed", check.Check).Error
}

// DeleteAllEntries Удалить все записи из БД
func (r *noteRepository) DeleteAllEntries(ctx context.Context) error {
	return r.db.WithContext(ctx).Unscoped().Where("1 = 1").Delete(&models.AllNotes{}).Error
}

// DeleteAllMarkedEntries Удалить все отмеченные записи из БД
func (r *noteRepository) DeleteAllMarkedEntries(ctx context.Context) error {
	return r.db.WithContext(ctx).Unscoped().Where("completed = ?", true).Delete(&models.AllNotes{}).Error
}
