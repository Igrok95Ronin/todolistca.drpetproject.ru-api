package service

import (
	"context"
	"errors"
	"github.com/Igrok95Ronin/todolistca.drpetproject.ru-api.git/internal/models"
	"github.com/Igrok95Ronin/todolistca.drpetproject.ru-api.git/internal/repository"
	"html"
	"strings"
)

// NoteService - интерфейс для работы с бизнес-логикой заметок
type NoteService interface {
	GetAllNotes(ctx context.Context) ([]models.AllNotes, error)
	CreateNote(ctx context.Context, note *models.AllNotes) error
	EditEntry(ctx context.Context, updatedEntry *models.ModifiedEntry, id int64) error
	DeleteEntry(ctx context.Context, id int64) error
	MarkCompleteEntry(ctx context.Context, check models.Check, id int64) error
	DeleteAllEntries(ctx context.Context) error
	DeleteAllMarkedEntries(ctx context.Context) error
}

type noteService struct {
	repo repository.NoteRepository
}

func NewNoteService(repo repository.NoteRepository) NoteService {
	return &noteService{
		repo: repo,
	}
}

// GetAllNotes - Получить все заметки
func (s *noteService) GetAllNotes(ctx context.Context) ([]models.AllNotes, error) {
	return s.repo.GetAllNotes(ctx)
}

// CreateNote - Создаёт новую заметку с валидацией
func (s *noteService) CreateNote(ctx context.Context, note *models.AllNotes) error {

	note.Note = html.EscapeString(strings.TrimSpace(note.Note))
	if note.Note == "" {
		return errors.New("заметка не может быть пустой")
	}

	return s.repo.CreateNote(ctx, note)
}

// EditEntry обновить заметку
func (s *noteService) EditEntry(ctx context.Context, updatedEntry *models.ModifiedEntry, id int64) error {
	updatedEntry.ModEntry = html.EscapeString(strings.TrimSpace(updatedEntry.ModEntry))
	if updatedEntry.ModEntry == "" {
		return errors.New("заметка не может быть пустой")
	}

	return s.repo.EditEntry(ctx, updatedEntry, id)
}

// DeleteEntry Удалить запись
func (s *noteService) DeleteEntry(ctx context.Context, id int64) error {

	if id <= 0 {
		return errors.New("ID должен быть больше 0")
	}

	return s.repo.DeleteEntry(ctx, id)
}

// MarkCompleteEntry Отметить выполненную запись
func (s *noteService) MarkCompleteEntry(ctx context.Context, check models.Check, id int64) error {

	if id <= 0 {
		return errors.New("ID должен быть больше 0")
	}

	return s.repo.MarkCompleteEntry(ctx, check, id)
}

// DeleteAllEntries Удалить все записи
func (s *noteService) DeleteAllEntries(ctx context.Context) error {
	return s.repo.DeleteAllEntries(ctx)
}

// DeleteAllMarkedEntries Удалить все отмеченные записи
func (s *noteService) DeleteAllMarkedEntries(ctx context.Context) error {
	return s.repo.DeleteAllMarkedEntries(ctx)
}
