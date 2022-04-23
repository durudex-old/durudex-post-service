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

package service

import (
	"context"

	"github.com/durudex/durudex-post-service/internal/domain"
	"github.com/durudex/durudex-post-service/internal/repository"

	"github.com/gofrs/uuid"
)

// Post interface.
type Post interface {
	Create(ctx context.Context, authorID uuid.UUID, text string) (uuid.UUID, error)
	GetByID(ctx context.Context, id uuid.UUID) (domain.Post, error)
	Delete(ctx context.Context, id uuid.UUID) error
}

// Post service structure.
type PostService struct{ repos repository.Post }

// Creating a new post service.
func NewPostService(repos repository.Post) *PostService {
	return &PostService{repos: repos}
}

// Creating a new post.
func (s *PostService) Create(ctx context.Context, authorID uuid.UUID, text string) (uuid.UUID, error) {
	// TODO: Check length of text.

	// Create a new post.
	id, err := s.repos.Create(ctx, authorID, text)
	if err != nil {
		return uuid.Nil, err
	}

	return id, nil
}

// Getting a post by id.
func (s *PostService) GetByID(ctx context.Context, id uuid.UUID) (domain.Post, error) {
	// Get post by id.
	post, err := s.repos.GetByID(ctx, id)
	if err != nil {
		return domain.Post{}, err
	}

	return post, nil
}

// Deleting a post.
func (s *PostService) Delete(ctx context.Context, id uuid.UUID) error {
	return s.repos.Delete(ctx, id)
}
