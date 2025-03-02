package handlers

import (
	"encoding/json"
	"github.com/Igrok95Ronin/todolistca.drpetproject.ru-api.git/internal/repository"
	"github.com/Igrok95Ronin/todolistca.drpetproject.ru-api.git/pkg/httperror"
	"github.com/Igrok95Ronin/todolistca.drpetproject.ru-api.git/pkg/logging"
	"github.com/julienschmidt/httprouter"
	"net/http"
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
