package handlers

import (
	"encoding/json"
	"github.com/Igrok95Ronin/todolistca.drpetproject.ru-api.git/internal/models"
	"github.com/Igrok95Ronin/todolistca.drpetproject.ru-api.git/internal/service"
	"github.com/Igrok95Ronin/todolistca.drpetproject.ru-api.git/pkg/httperror"
	"github.com/Igrok95Ronin/todolistca.drpetproject.ru-api.git/pkg/logging"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

// NoteHandler обрабатывает запросы, связанные с заметками
type NoteHandler struct {
	service service.NoteService
	logger  *logging.Logger
}

// NewNoteHandler создаёт новый обработчик заметок
func NewNoteHandler(service service.NoteService, logger *logging.Logger) *NoteHandler {
	return &NoteHandler{
		service: service,
		logger:  logger,
	}
}

// GetAllNotes Получить все посты
func (h *NoteHandler) GetAllNotes(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	ctx := r.Context()

	notes, err := h.service.GetAllNotes(ctx)
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
	ctx := r.Context()
	var note models.AllNotes

	if err := json.NewDecoder(r.Body).Decode(&note); err != nil {
		// Если произошла ошибка декодирования, возвращаем клиенту ошибку с кодом 400
		httperror.WriteJSONError(w, "Ошибка декодирования в JSON", err, http.StatusBadRequest)
		// Логируем ошибку
		h.logger.Errorf("Ошибка декодирования в JSON: %s", err)
		return
	}

	if err := h.service.CreateNote(ctx, &note); err != nil {
		h.logger.Errorf("Ошибка при добавления записи в БД: %s", err)
		httperror.WriteJSONError(w, "Ошибка при добавления записи в БД", err, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// EditEntry обновить заметку
func (h *NoteHandler) EditEntry(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := r.Context()
	var modifiedEntry *models.ModifiedEntry

	if err := json.NewDecoder(r.Body).Decode(&modifiedEntry); err != nil {
		// Если произошла ошибка декодирования, возвращаем клиенту ошибку с кодом 400
		httperror.WriteJSONError(w, "Ошибка декодирования в JSON", err, http.StatusBadRequest)
		// Логируем ошибку
		h.logger.Errorf("Ошибка декодирования в JSON: %s", err)
		return
	}

	id, _ := strconv.Atoi(ps.ByName("id"))

	if err := h.service.EditEntry(ctx, modifiedEntry, int64(id)); err != nil {
		httperror.WriteJSONError(w, "Ошибка при обновления записи в БД", err, http.StatusInternalServerError)
		h.logger.Errorf("Ошибка при обновления записи в БД: %s", err)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// DeleteEntry // Удалить запись
func (h *NoteHandler) DeleteEntry(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := r.Context()
	ID := ps.ByName("id")

	id, err := strconv.Atoi(ID)
	if err != nil {
		h.logger.Errorf("Некорректный ID: %s", err)
		return
	}

	if err = h.service.DeleteEntry(ctx, int64(id)); err != nil {
		h.logger.Errorf("Ошибка удаления записи с ID %d: %s", id, err)
		httperror.WriteJSONError(w, "Ошибка удаления записи", err, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// MarkCompleteEntry Отметить выполненную запись
func (h *NoteHandler) MarkCompleteEntry(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := r.Context()
	ID := ps.ByName("id")

	var check models.Check

	if err := json.NewDecoder(r.Body).Decode(&check); err != nil {
		// Если произошла ошибка декодирования, возвращаем клиенту ошибку с кодом 400
		httperror.WriteJSONError(w, "Ошибка декодирования в JSON", err, http.StatusBadRequest)
		// Логируем ошибку
		h.logger.Errorf("Ошибка декодирования в JSON: %s", err)
		return
	}

	id, err := strconv.Atoi(ID)
	if err != nil {
		h.logger.Errorf("Некорректный ID: %s", err)
		return
	}

	if err = h.service.MarkCompleteEntry(ctx, check, int64(id)); err != nil {
		h.logger.Errorf("Ошибка обновления записи с ID %d: %s", id, err)
		httperror.WriteJSONError(w, "Ошибка обновления записи", err, http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// DeleteAllEntries Удалить все записи
func (h *NoteHandler) DeleteAllEntries(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	ctx := r.Context()
	if err := h.service.DeleteAllEntries(ctx); err != nil {
		h.logger.Errorf("Ошибка при удалении всех записей: %s", err)
		httperror.WriteJSONError(w, "Ошибка при удалении всех записей", err, http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// DeleteAllMarkedEntries Удалить все отмеченные записи
func (h *NoteHandler) DeleteAllMarkedEntries(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	ctx := r.Context()
	if err := h.service.DeleteAllMarkedEntries(ctx); err != nil {
		h.logger.Errorf("Ошибка при удалении всех отмеченных записей: %s", err)
		httperror.WriteJSONError(w, "Ошибка при удалении всех отмеченных записей", err, http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}
