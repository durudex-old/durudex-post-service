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

package psql

import (
	"context"

	"github.com/durudex/dugopg"
	"github.com/durudex/durudex-post-service/internal/domain"

	"github.com/gofrs/uuid"
)

// Post postgres repository.
type PostRepository struct{ psql dugopg.Native }

// Creating a new post repository.
func NewPostRepository(psql dugopg.Native) *PostRepository {
	return &PostRepository{psql: psql}
}

// Creating a new post in postgres database.
func (r *PostRepository) Create(ctx context.Context, text string) (uuid.UUID, error) {
	return uuid.Nil, nil
}

// Getting a post by id in postgres database.
func (r *PostRepository) GetByID(ctx context.Context, id uuid.UUID) (domain.Post, error) {
	return domain.Post{}, nil
}

// Deleting a post in postgres database.
func (r *PostRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return nil
}
