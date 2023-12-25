package usecase

import (
	"CRUD/internal/domain"
	"CRUD/internal/repository"
	"context"

	"github.com/google/uuid"
)

// PostUsecase is the interface for the post use case.
type PostUsecase interface {
	GetPostByID(ctx context.Context, id uuid.UUID) (*domain.Post, error)
	CreatePost(ctx context.Context, post *domain.Post) error
	UpdatePost(ctx context.Context, id uuid.UUID, post *domain.Post) error
	DeletePost(ctx context.Context, id uuid.UUID) error
}

// PostUsecaseImpl is the implementation of the PostUsecase interface.
type PostUsecaseImpl struct {
	postRepository repository.PostRepository
}

// NewPostUsecase creates a new instance of PostUsecaseImpl.
func NewPostUsecase(postRepository repository.PostRepository) PostUsecase {
	return &PostUsecaseImpl{
		postRepository: postRepository,
	}
}

func (u *PostUsecaseImpl) GetPostByID(ctx context.Context, id uuid.UUID) (*domain.Post, error) {
	return u.postRepository.GetPostByID(ctx, id)
}

func (u *PostUsecaseImpl) CreatePost(ctx context.Context, post *domain.Post) error {
	return u.postRepository.CreatePost(ctx, post)
}

func (u *PostUsecaseImpl) UpdatePost(ctx context.Context, id uuid.UUID, post *domain.Post) error {
	return u.postRepository.UpdatePost(ctx, id, post)
}

func (u *PostUsecaseImpl) DeletePost(ctx context.Context, id uuid.UUID) error {
	return u.postRepository.DeletePost(ctx, id)
}
