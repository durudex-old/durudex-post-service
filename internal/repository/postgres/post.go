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

	"github.com/durudex/durudex-post-service/internal/domain"
	"github.com/durudex/durudex-post-service/pkg/database/postgres"

	"github.com/jackc/pgx/v4"
	"github.com/leporo/sqlf"
	"github.com/segmentio/ksuid"
)

// Post repository interface.
type Post interface {
	// Creating a new post in postgres database.
	Create(ctx context.Context, post domain.Post) error
	// Getting a post by id in postgres database.
	Get(ctx context.Context, id ksuid.KSUID) (domain.Post, error)
	// Getting author posts by author id in postgres database.
	GetPosts(ctx context.Context, authorId ksuid.KSUID, sort domain.SortOptions) ([]domain.Post, error)
	// Deleting a post in postgres database.
	Delete(ctx context.Context, id, authorId ksuid.KSUID) error
	// Updating a post in postgres database.
	Update(ctx context.Context, post domain.Post) error
	// Getting total author posts count in postgres database.
	GetTotalCount(ctx context.Context, authorId ksuid.KSUID) (int32, error)
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
	query := "INSERT INTO post (id, author_id, text) VALUES ($1, $2, $3)"

	// Scan post id.
	if _, err := r.psql.Exec(ctx, query, post.Id, post.AuthorId, post.Text); err != nil {
		return err
	}

	return nil
}

// Getting a post by id in postgres database.
func (r *PostRepository) Get(ctx context.Context, id ksuid.KSUID) (domain.Post, error) {
	var post domain.Post

	// Query for get post by id.
	query := "SELECT author_id, text, updated_at FROM post WHERE id=$1"

	row := r.psql.QueryRow(ctx, query, id)

	// Scanning query row.
	if err := row.Scan(&post.AuthorId, &post.Text, &post.UpdatedAt); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return domain.Post{}, &domain.Error{Code: domain.CodeNotFound, Message: "Post not found"}
		}

		return domain.Post{}, &domain.Error{Code: domain.CodeInternal, Message: "Internal Server Error"}
	}

	return post, nil
}

// Getting author posts by author id in postgres database.
func (r *PostRepository) GetPosts(ctx context.Context, authorId ksuid.KSUID, sort domain.SortOptions) ([]domain.Post, error) {
	var n int32

	qb := sqlf.Select("id, text, updated_at").From("post").Where("author_id = ?", authorId)

	// Added first or last sort option.
	if sort.First != nil {
		n = *sort.First
		qb.OrderBy("created_at ASC, id ASC").Limit(*sort.First)
	} else if sort.Last != nil {
		n = *sort.Last
		qb.OrderBy("created_at DESC, id DESC").Limit(*sort.Last)
	}

	// Added before sort option.
	if sort.Before != ksuid.Nil {
		qb.Where("(created_at, id) < (?, ?)", sort.Before.Time(), sort.Before)
	}
	// Added after sort option.
	if sort.After != ksuid.Nil {
		qb.Where("(created_at, id) > (?, ?)", sort.After.Time(), sort.After)
	}

	posts := make([]domain.Post, n)

	// Query for getting author posts by author id.
	rows, err := r.psql.Query(ctx, qb.String(), qb.Args()...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var i int

	// Scanning query rows.
	for rows.Next() {
		var post domain.Post

		// Scanning query row.
		if err := rows.Scan(&post.Id, &post.Text, &post.UpdatedAt); err != nil {
			return nil, err
		}

		posts[i] = post

		i++
	}

	// Check is rows error.
	if err := rows.Err(); err != nil {
		return nil, err
	}

	// Check for fullness of the slice.
	if i == int(n) {
		return posts, nil
	}

	res := make([]domain.Post, i)
	copy(res, posts[:i])

	return res, nil
}

// Deleting a post in postgres database.
func (r *PostRepository) Delete(ctx context.Context, id, authorId ksuid.KSUID) error {
	// Query for delete post by id.
	query := "DELETE FROM post WHERE id=$1 AND author_id=$2"
	_, err := r.psql.Exec(ctx, query, id, authorId)

	return err
}

// Updating a post in postgres database.
func (r *PostRepository) Update(ctx context.Context, post domain.Post) error {
	// Query for update post by id.
	query := "UPDATE post SET text=$1, updated_at=now() WHERE id=$2 AND author_id=$3"
	_, err := r.psql.Exec(ctx, query, post.Text, post.Id, post.AuthorId)

	return err
}

// Getting total author posts count in postgres database.
func (r *PostRepository) GetTotalCount(ctx context.Context, authorId ksuid.KSUID) (int32, error) {
	var count int32

	// Query to get author posts total count.
	query := "SELECT count(*) FROM post WHERE author_id=$1"

	row := r.psql.QueryRow(ctx, query, authorId)

	// Scanning query row.
	if err := row.Scan(&count); err != nil {
		return 0, err
	}

	return count, nil
}
