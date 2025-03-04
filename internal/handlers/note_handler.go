package handlers

import (
	"encoding/json"
	"github.com/Igrok95Ronin/todolistca.drpetproject.ru-api.git/internal/models"
	"github.com/Igrok95Ronin/todolistca.drpetproject.ru-api.git/internal/repository"
	"github.com/Igrok95Ronin/todolistca.drpetproject.ru-api.git/pkg/httperror"
	"github.com/Igrok95Ronin/todolistca.drpetproject.ru-api.git/pkg/logging"
	"github.com/julienschmidt/httprouter"
	"html"
	"net/http"
	"strconv"
	"strings"
)

// NoteHandler обрабатывает запросы, связанные с заметками
type NoteHandler struct {
	repo   repository.NoteRepository
	logger *logging.Logger
}

// NewNoteHandler создаёт новый обработчик заметок
func NewNoteHandler(repo repository.NoteRepository, logger *logging.Logger) *NoteHandler {
	return &NoteHandler{
		repo:   repo,
		logger: logger,
	}
}

// GetAllNotes Получить все посты
func (h *NoteHandler) GetAllNotes(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	notes, err := h.repo.GetAllNotes()
	if err != nil {
		h.logger.Error(err)
		httperror.WriteJSONError(w, "Ошибка получения заметок", err, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err = json.NewEncoder(w).Encode(notes); err != nil {
		h.logger.Error(err)
		httperror.WriteJSONError(w, "Ошибка при отправке данных", err, http.StatusInternalServerError)
	}
}

// AddPost добавляет новую заметку
func (h *NoteHandler) AddPost(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var note models.AllNotes

	if err := json.NewDecoder(r.Body).Decode(&note); err != nil {
		// Если произошла ошибка декодирования, возвращаем клиенту ошибку с кодом 400
		httperror.WriteJSONError(w, "Ошибка декодирования в JSON", err, http.StatusBadRequest)
		// Логируем ошибку
		h.logger.Errorf("Ошибка декодирования в JSON: %s", err)
		return
	}

	note.Note = html.EscapeString(strings.TrimSpace(note.Note))
	if note.Note == "" {
		httperror.WriteJSONError(w, "Заметка не может быть пустой", nil, http.StatusBadRequest)
		return
	}

	allNotes := models.AllNotes{
		Note: note.Note,
	}

	if err := h.repo.CreateNote(&allNotes); err != nil {
		h.logger.Errorf("Ошибка при добавления записи в БД: %s", err)
		httperror.WriteJSONError(w, "Ошибка при добавления записи в БД", err, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// EditEntry обновить заметку
func (h *NoteHandler) EditEntry(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var modifiedEntry *models.ModifiedEntry

	if err := json.NewDecoder(r.Body).Decode(&modifiedEntry); err != nil {
		// Если произошла ошибка декодирования, возвращаем клиенту ошибку с кодом 400
		httperror.WriteJSONError(w, "Ошибка декодирования в JSON", err, http.StatusBadRequest)
		// Логируем ошибку
		h.logger.Errorf("Ошибка декодирования в JSON: %s", err)
		return
	}

	id, _ := strconv.Atoi(ps.ByName("id"))

	modifiedEntry.ModEntry = html.EscapeString(strings.TrimSpace(modifiedEntry.ModEntry))
	if modifiedEntry.ModEntry == "" {
		return
	}

	if err := h.repo.EditEntry(modifiedEntry, int64(id)); err != nil {
		httperror.WriteJSONError(w, "Ошибка при обновления записи в БД", err, http.StatusInternalServerError)
		h.logger.Errorf("Ошибка при обновления записи в БД: %s", err)
		return
	}

	w.WriteHeader(http.StatusOK)
}
