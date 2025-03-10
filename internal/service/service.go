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
