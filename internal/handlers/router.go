package handlers

import (
	"github.com/Igrok95Ronin/todolistca.drpetproject.ru-api.git/internal/config"
	"github.com/Igrok95Ronin/todolistca.drpetproject.ru-api.git/internal/repository"
	"github.com/Igrok95Ronin/todolistca.drpetproject.ru-api.git/pkg/logging"
	"github.com/julienschmidt/httprouter"
	"gorm.io/gorm"
)

// Handler управляет роутами
type Handler struct {
	cfg      *config.Config
	logger   *logging.Logger
	noteRepo repository.NoteRepository
}

// NewHandler создаёт новый обработчик
func NewHandler(cfg *config.Config, logger *logging.Logger, db *gorm.DB) *Handler {
	return &Handler{
		cfg:      cfg,
		logger:   logger,
		noteRepo: repository.NewNoteRepository(db, logger),
	}
}

// RegisterRoutes регистрирует маршруты
func (h *Handler) RegisterRoutes(router *httprouter.Router) {
	noteHandler := NewNoteHandler(h.noteRepo, h.logger)

	router.GET("/", noteHandler.GetAllNotes)     // Получения всех записей
	router.POST("/addpost", noteHandler.AddPost) // Добавить пост
}
