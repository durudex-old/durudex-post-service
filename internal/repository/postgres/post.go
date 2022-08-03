/*
 * Copyright © 2022 Durudex
 *
 * This file is part of Durudex: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as
 * published by the Free Software Foundation, either version 3 of the
 * License, or (at your option) any later version.
 *
 * Durudex is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with Durudex. If not, see <https://www.gnu.org/licenses/>.
 */

package postgres

import (
	"context"
	"errors"
	"fmt"

	"github.com/durudex/durudex-post-service/internal/domain"
	"github.com/durudex/durudex-post-service/pkg/database/postgres"

	"github.com/jackc/pgx/v4"
	"github.com/segmentio/ksuid"
)

// Post table name.
const PostTable string = "post"

// Post repository interface.
type Post interface {
	Create(ctx context.Context, post domain.Post) error
	GetById(ctx context.Context, id ksuid.KSUID) (domain.Post, error)
	GetAuthorPosts(ctx context.Context, authorId ksuid.KSUID, first, last *int32) ([]domain.Post, error)
	Delete(ctx context.Context, id, authorId ksuid.KSUID) error
	Update(ctx context.Context, post domain.Post) error
}

// Post repository structure.
type PostRepository struct{ psql postgres.Postgres }

// Creating a new post repository.
func NewPostRepository(psql postgres.Postgres) *PostRepository {
	return &PostRepository{psql: psql}
}

// Creating a new post in postgres database.
func (r *PostRepository) Create(ctx context.Context, post domain.Post) error {
	// Query to create post.
	query := fmt.Sprintf(`INSERT INTO "%s" (id, author_id, text) VALUES ($1, $2, $3)`, PostTable)

	// Scan post id.
	if _, err := r.psql.Exec(ctx, query, post.Id, post.AuthorId, post.Text); err != nil {
		return err
	}

	return nil
}

// Getting a post by id in postgres database.
func (r *PostRepository) GetById(ctx context.Context, id ksuid.KSUID) (domain.Post, error) {
	var post domain.Post

	// Query for get post by id.
	query := fmt.Sprintf(`SELECT "author_id", "text", "updated_at" FROM "%s" WHERE "id"=$1`, PostTable)

	row := r.psql.QueryRow(ctx, query, id)

	// Scanning query row.
	if err := row.Scan(&post.AuthorId, &post.Text, &post.UpdatedAt); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return domain.Post{}, &domain.Error{Code: domain.CodeNotFound, Message: "User not found"}
		}

		return domain.Post{}, &domain.Error{Code: domain.CodeInternal, Message: "Internal Server Error"}
	}

	return post, nil
}

// Getting author posts by author id in postgres database.
func (r *PostRepository) GetAuthorPosts(ctx context.Context, authorId ksuid.KSUID, first, last *int32) ([]domain.Post, error) {
	var (
		// Posts numbers.
		n int32
		// Query post filter.
		filter string
	)

	// Set query filter.
	if first == nil {
		n = *last
		filter = "DESC"
	} else {
		n = *first
		filter = "ASC"
	}

	posts := make([]domain.Post, n)

	// Query for get author posts.
	query := fmt.Sprintf(`SELECT "id", "author_id", "text", "updated_at" FROM "%s"
		WHERE "author_id"=$1 ORDER BY "id" %s LIMIT $2`, PostTable, filter)

	rows, err := r.psql.Query(ctx, query, authorId, n)
	if err != nil {
		return nil, err
	}

	var i int

	// Scanning query rows.
	for rows.Next() {
		var post domain.Post

		// Scanning query row.
		if err := rows.Scan(&post.Id, &post.AuthorId, &post.Text, &post.UpdatedAt); err != nil {
			return nil, err
		}

		posts[i] = post

		i++
	}

	// Check is rows error.
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return posts, nil
}

// Deleting a post in postgres database.
func (r *PostRepository) Delete(ctx context.Context, id, authorId ksuid.KSUID) error {
	// Query for delete post by id.
	query := fmt.Sprintf(`DELETE FROM "%s" WHERE id=$1 AND author_id=$2`, PostTable)
	_, err := r.psql.Exec(ctx, query, id, authorId)

	return err
}

// Updating a post in postgres database.
func (r *PostRepository) Update(ctx context.Context, post domain.Post) error {
	// Query for update post by id.
	query := fmt.Sprintf(`UPDATE "%s" SET text=$1, updated_at=now() WHERE "id"=$2 AND author_id=$3`, PostTable)
	_, err := r.psql.Exec(ctx, query, post.Text, post.Id, post.AuthorId)

	return err
}
