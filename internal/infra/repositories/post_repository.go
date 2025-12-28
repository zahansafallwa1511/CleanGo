package repositories

import (
	"context"
	"errors"
	"sync"
	"time"

	"cleanandclean/internal/core/domain"
	"cleanandclean/internal/core/usecases/post"
)

var _ post.PostRepository = (*InMemoryPostRepository)(nil)

type InMemoryPostRepository struct {
	mu     sync.RWMutex
	posts  map[uint64]*domain.Post
	nextID uint64
}

func NewInMemoryPostRepository() *InMemoryPostRepository {
	return &InMemoryPostRepository{
		posts:  make(map[uint64]*domain.Post),
		nextID: 1,
	}
}

func (r *InMemoryPostRepository) Create(ctx context.Context, p *domain.Post) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	p.ID = r.nextID
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()
	r.posts[p.ID] = p
	r.nextID++

	return nil
}

func (r *InMemoryPostRepository) FindByID(ctx context.Context, id uint64) (*domain.Post, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	p, exists := r.posts[id]
	if !exists {
		return nil, errors.New("post not found")
	}

	return p, nil
}

func (r *InMemoryPostRepository) FindAll(ctx context.Context, limit, offset int) ([]*domain.Post, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	posts := make([]*domain.Post, 0, len(r.posts))
	for _, p := range r.posts {
		posts = append(posts, p)
	}

	if offset >= len(posts) {
		return []*domain.Post{}, nil
	}

	end := offset + limit
	if end > len(posts) {
		end = len(posts)
	}

	return posts[offset:end], nil
}

func (r *InMemoryPostRepository) Update(ctx context.Context, p *domain.Post) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.posts[p.ID]; !exists {
		return errors.New("post not found")
	}

	p.UpdatedAt = time.Now()
	r.posts[p.ID] = p

	return nil
}

func (r *InMemoryPostRepository) Delete(ctx context.Context, id uint64) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.posts[id]; !exists {
		return errors.New("post not found")
	}

	delete(r.posts, id)
	return nil
}
