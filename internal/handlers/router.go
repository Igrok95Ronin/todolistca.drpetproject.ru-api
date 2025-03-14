package handlers

import (
	"github.com/Igrok95Ronin/todolistca.drpetproject.ru-api.git/internal/config"
	"github.com/Igrok95Ronin/todolistca.drpetproject.ru-api.git/internal/repository"
	"github.com/Igrok95Ronin/todolistca.drpetproject.ru-api.git/internal/service"
	"github.com/Igrok95Ronin/todolistca.drpetproject.ru-api.git/pkg/logging"
	"github.com/julienschmidt/httprouter"
	"gorm.io/gorm"
)

// Handler управляет роутами
type Handler struct {
	cfg      *config.Config
	logger   *logging.Logger
	noteRepo repository.NoteRepository
	noteSvc  service.NoteService
}

// NewHandler создаёт новый обработчик
func NewHandler(cfg *config.Config, logger *logging.Logger, db *gorm.DB) *Handler {
	noteRepo := repository.NewNoteRepository(db)
	noteSvc := service.NewNoteService(noteRepo)

	return &Handler{
		cfg:      cfg,
		logger:   logger,
		noteRepo: noteRepo,
		noteSvc:  noteSvc,
	}
}

// RegisterRoutes регистрирует маршруты
func (h *Handler) RegisterRoutes(router *httprouter.Router) {
	noteHandler := NewNoteHandler(h.noteSvc, h.logger)

	router.GET("/", noteHandler.GetAllNotes)                              // Получения всех записей
	router.POST("/notes", noteHandler.AddPost)                            // Добавить пост
	router.PUT("/notes/:id", noteHandler.EditEntry)                       // Редактировать запись
	router.DELETE("/note/:id", noteHandler.DeleteEntry)                   // Удалить запись
	router.PUT("/notes/:id/complete", noteHandler.MarkCompleteEntry)      // Отметить выполненную запись
	router.DELETE("/notes", noteHandler.DeleteAllEntries)                 // Удалить все записи
	router.DELETE("/notes/completed", noteHandler.DeleteAllMarkedEntries) // Удалить все отмеченные записи
}
