package handler

import (
	"CRUD/internal/domain"
	"CRUD/internal/usecase"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

type PostHandler struct {
	postUsecase usecase.PostUsecase
	logger      *zap.Logger
}

func NewPostHandler(postUsecase usecase.PostUsecase, logger *zap.Logger) *PostHandler {
	return &PostHandler{
		postUsecase: postUsecase,
		logger:      logger,
	}
}

func (h *PostHandler) Get(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		http.Error(w, "Invalid ID parameter", http.StatusBadRequest)
		h.logger.Error("Invalid ID parameter", zap.Error(err))
		return
	}

	post, err := h.postUsecase.GetPostByID(r.Context(), id)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to fetch post with ID %s", id.String()), http.StatusInternalServerError)
		h.logger.Error("Failed to fetch post", zap.Error(err))
		return
	}

	if post == nil {
		http.Error(w, "Post not found", http.StatusNotFound)
		h.logger.Info("Post not found", zap.String("ID", id.String()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(post)
}

func (h *PostHandler) Create(w http.ResponseWriter, r *http.Request) {
	var newPost domain.Post
	if err := json.NewDecoder(r.Body).Decode(&newPost); err != nil {
		http.Error(w, "Failed to parse request body", http.StatusBadRequest)
		h.logger.Error("Failed to parse request body", zap.Error(err))
		return
	}

	newPost.ID = uuid.New()

	if err := h.postUsecase.CreatePost(r.Context(), &newPost); err != nil {
		http.Error(w, "Failed to create post", http.StatusInternalServerError)
		h.logger.Error("Failed to create post", zap.Error(err), zap.Any("Post", newPost))
		return
	}

	w.Write([]byte("Post created successfully"))
	h.logger.Info("Post created successfully", zap.Any("Post", newPost))
}

func (h *PostHandler) Update(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		http.Error(w, "Invalid ID parameter", http.StatusBadRequest)
		h.logger.Error("Invalid ID parameter", zap.Error(err))
		return
	}

	var updatedPost domain.Post
	if err := json.NewDecoder(r.Body).Decode(&updatedPost); err != nil {
		http.Error(w, "Failed to parse request body", http.StatusBadRequest)
		h.logger.Error("Failed to parse request body", zap.Error(err))
		return
	}

	if err := h.postUsecase.UpdatePost(r.Context(), id, &updatedPost); err != nil {
		http.Error(w, "Failed to update post", http.StatusInternalServerError)
		h.logger.Error("Failed to update post", zap.Error(err), zap.Any("Post", updatedPost))
		return
	}

	w.Write([]byte("Post updated successfully"))
	h.logger.Info("Post updated successfully", zap.Any("Post", updatedPost))
}

func (h *PostHandler) Delete(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		http.Error(w, "Invalid ID parameter", http.StatusBadRequest)
		h.logger.Error("Invalid ID parameter", zap.Error(err))
		return
	}

	if err := h.postUsecase.DeletePost(r.Context(), id); err != nil {
		http.Error(w, "Failed to delete post", http.StatusInternalServerError)
		h.logger.Error("Failed to delete post", zap.Error(err), zap.String("ID", id.String()))
		return
	}

	w.Write([]byte("Post deleted successfully"))
	h.logger.Info("Post deleted successfully", zap.String("ID", id.String()))
}
