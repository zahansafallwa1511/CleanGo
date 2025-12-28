package repositories

import (
	"context"
	"database/sql"

	"cleanandclean/internal/core/domain"
	"cleanandclean/internal/core/usecases/post"
)

var _ post.PostRepository = (*PostgresPostRepository)(nil)

type PostgresPostRepository struct {
	db *sql.DB
}

func NewPostgresPostRepository(db *sql.DB) *PostgresPostRepository {
	return &PostgresPostRepository{db: db}
}

func (r *PostgresPostRepository) Create(ctx context.Context, p *domain.Post) error {
	query := `
		INSERT INTO posts (title, content, author_id, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id
	`
	return r.db.QueryRowContext(
		ctx,
		query,
		p.Title,
		p.Content,
		p.AuthorID,
		p.CreatedAt,
		p.UpdatedAt,
	).Scan(&p.ID)
}

func (r *PostgresPostRepository) FindByID(ctx context.Context, id uint64) (*domain.Post, error) {
	query := `
		SELECT id, title, content, author_id, created_at, updated_at
		FROM posts
		WHERE id = $1
	`
	p := &domain.Post{}
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&p.ID,
		&p.Title,
		&p.Content,
		&p.AuthorID,
		&p.CreatedAt,
		&p.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return p, nil
}

func (r *PostgresPostRepository) FindAll(ctx context.Context, limit, offset int) ([]*domain.Post, error) {
	query := `
		SELECT id, title, content, author_id, created_at, updated_at
		FROM posts
		ORDER BY created_at DESC
		LIMIT $1 OFFSET $2
	`
	rows, err := r.db.QueryContext(ctx, query, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []*domain.Post
	for rows.Next() {
		p := &domain.Post{}
		if err := rows.Scan(
			&p.ID,
			&p.Title,
			&p.Content,
			&p.AuthorID,
			&p.CreatedAt,
			&p.UpdatedAt,
		); err != nil {
			return nil, err
		}
		posts = append(posts, p)
	}
	return posts, rows.Err()
}

func (r *PostgresPostRepository) Update(ctx context.Context, p *domain.Post) error {
	query := `
		UPDATE posts
		SET title = $1, content = $2, updated_at = $3
		WHERE id = $4
	`
	_, err := r.db.ExecContext(ctx, query, p.Title, p.Content, p.UpdatedAt, p.ID)
	return err
}

func (r *PostgresPostRepository) Delete(ctx context.Context, id uint64) error {
	query := `DELETE FROM posts WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}
