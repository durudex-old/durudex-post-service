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

package service

import (
	"context"

	"github.com/durudex/durudex-post-service/internal/domain"
	"github.com/durudex/durudex-post-service/internal/repository/postgres"

	"github.com/segmentio/ksuid"
)

// Post interface.
type Post interface {
	// Creating a new post.
	Create(ctx context.Context, post domain.Post) (ksuid.KSUID, error)
	// Getting a post.
	GetBy(ctx context.Context, id ksuid.KSUID) (domain.Post, error)
	// Getting author posts.
	GetPosts(ctx context.Context, authorId ksuid.KSUID, sort domain.SortOptions) ([]domain.Post, error)
	// Deleting a post.
	Delete(ctx context.Context, id, authorId ksuid.KSUID) error
	// Updating a post.
	Update(ctx context.Context, post domain.Post) error
}

// Post service structure.
type PostService struct{ repos postgres.Post }

// Creating a new post service.
func NewPostService(repos postgres.Post) *PostService {
	return &PostService{repos: repos}
}

// Creating a new post.
func (s *PostService) Create(ctx context.Context, post domain.Post) (ksuid.KSUID, error) {
	var err error

	// Validate a post.
	if err := post.Validate(); err != nil {
		return ksuid.Nil, err
	}

	// Generating a new user id.
	if post.Id.IsNil() {
		post.Id, err = ksuid.NewRandom()
		if err != nil {
			return ksuid.Nil, err
		}
	}

	// Create a new post.
	if err := s.repos.Create(ctx, post); err != nil {
		return ksuid.Nil, err
	}

	return post.Id, nil
}

// Getting a post.
func (s *PostService) GetBy(ctx context.Context, id ksuid.KSUID) (domain.Post, error) {
	// Get post by id.
	post, err := s.repos.GetById(ctx, id)
	if err != nil {
		return domain.Post{}, err
	}

	return post, nil
}

// Getting author posts.
func (s *PostService) GetPosts(ctx context.Context, authorId ksuid.KSUID, sort domain.SortOptions) ([]domain.Post, error) {
	// Check is first and last are set.
	if sort.First == nil && sort.Last == nil {
		return nil, &domain.Error{
			Message: "Must be `first` or `last`",
			Code:    domain.CodeInvalidArgument,
		}
	}

	// Getting author posts.
	posts, err := s.repos.GetPosts(ctx, authorId, domain.SortOptions{
		First: sort.First,
		Last:  sort.Last,
	})
	if err != nil {
		return nil, err
	}

	return posts, err
}

// Deleting a post.
func (s *PostService) Delete(ctx context.Context, id, authorId ksuid.KSUID) error {
	return s.repos.Delete(ctx, id, authorId)
}

// Updating a post.
func (s *PostService) Update(ctx context.Context, post domain.Post) error {
	// Validate a post.
	if err := post.Validate(); err != nil {
		return err
	}

	return s.repos.Update(ctx, post)
}
