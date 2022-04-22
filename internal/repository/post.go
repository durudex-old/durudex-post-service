/*
 * Copyright © 2022 Durudex

 * This file is part of Durudex: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as
 * published by the Free Software Foundation, either version 3 of the
 * License, or (at your option) any later version.

 * Durudex is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
 * GNU Affero General Public License for more details.

 * You should have received a copy of the GNU Affero General Public License
 * along with Durudex. If not, see <https://www.gnu.org/licenses/>.
 */

package repository

import (
	"context"

	"github.com/durudex/dugopg"
	"github.com/durudex/durudex-post-service/internal/domain"
	"github.com/durudex/durudex-post-service/internal/repository/psql"

	"github.com/gofrs/uuid"
)

// Post repository interface.
type Post interface {
	Create(ctx context.Context, text string) (uuid.UUID, error)
	GetByID(ctx context.Context, id uuid.UUID) (domain.Post, error)
	Delete(ctx context.Context, id uuid.UUID) error
}

// Post repository structure.
type PostRepository struct{ psql *psql.PostRepository }

// Creating a new post repository.
func NewPostRepository(conn dugopg.Native) *PostRepository {
	return &PostRepository{psql: psql.NewPostRepository(conn)}
}

// Creating a new post in database.
func (r *PostRepository) Create(ctx context.Context, text string) (uuid.UUID, error) {
	return r.psql.Create(ctx, text)
}

// Getting a post by id in database.
func (r *PostRepository) GetByID(ctx context.Context, id uuid.UUID) (domain.Post, error) {
	return r.psql.GetByID(ctx, id)
}

// Deleting a post in database.
func (r *PostRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.psql.Delete(ctx, id)
}
