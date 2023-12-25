package repository

import (
	"context"
	"fmt"

	"CRUD/internal/domain"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
)

// PostRepository is the interface for the post repository.
type PostRepository interface {
	GetPostByID(ctx context.Context, id uuid.UUID) (*domain.Post, error)
	CreatePost(ctx context.Context, post *domain.Post) error
	UpdatePost(ctx context.Context, id uuid.UUID, post *domain.Post) error
	DeletePost(ctx context.Context, id uuid.UUID) error
}

// PostRepositoryImpl is the implementation of the PostRepository interface.
type PostRepositoryImpl struct {
	db *pgx.Conn
}

// NewPostRepository creates a new instance of PostRepositoryImpl.
func NewPostRepository(db *pgx.Conn) *PostRepositoryImpl {
	return &PostRepositoryImpl{db: db}
}

func (r *PostRepositoryImpl) GetPostByID(ctx context.Context, id uuid.UUID) (*domain.Post, error) {
	query := "SELECT id, title, content FROM posts WHERE id = $1"
	row := r.db.QueryRow(ctx, query, id)

	var post domain.Post
	err := row.Scan(&post.ID, &post.Title, &post.Content)
	if err != nil {
		// Enhance error message
		return nil, fmt.Errorf("failed to fetch post with ID %s: %w", id.String(), err)
	}

	return &post, nil
}

func (r *PostRepositoryImpl) CreatePost(ctx context.Context, post *domain.Post) error {
	// Placeholder logic to create a post in the database
	query := "INSERT INTO posts (id, title, content) VALUES ($1, $2, $3)"
	_, err := r.db.Exec(ctx, query, post.ID, post.Title, post.Content)
	if err != nil {
		return fmt.Errorf("failed to create post: %w", err)
	}

	return nil
}

func (r *PostRepositoryImpl) UpdatePost(ctx context.Context, id uuid.UUID, post *domain.Post) error {
	// Placeholder logic to update a post in the database
	query := "UPDATE posts SET title = $1, content = $2 WHERE id = $3"
	_, err := r.db.Exec(ctx, query, post.Title, post.Content, id)
	if err != nil {
		return fmt.Errorf("failed to update post: %w", err)
	}

	return nil
}

func (r *PostRepositoryImpl) DeletePost(ctx context.Context, id uuid.UUID) error {
	// Placeholder logic to delete a post from the database
	query := "DELETE FROM posts WHERE id = $1"
	_, err := r.db.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed to delete post: %w", err)
	}

	return nil
}
