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

package v1

import (
	"context"

	"github.com/durudex/dugopb/type/timestamp"
	"github.com/durudex/durudex-post-service/internal/domain"
	"github.com/durudex/durudex-post-service/internal/service"
	v1 "github.com/durudex/durudex-post-service/pkg/pb/durudex/v1"

	"github.com/segmentio/ksuid"
)

// Sample gRPC server handler.
type PostHandler struct {
	service service.Post
	v1.UnimplementedPostServiceServer
}

// Creating a new post gRPC handler.
func NewPostHandler(service service.Post) *PostHandler {
	return &PostHandler{service: service}
}

// Creating a new post handler.
func (h *PostHandler) CreatePost(ctx context.Context, input *v1.CreatePostRequest) (*v1.CreatePostResponse, error) {
	// Create a new post.
	id, err := h.service.Create(ctx, domain.Post{
		AuthorId: ksuid.FromBytesOrNil(input.AuthorId),
		Text:     input.Text,
	})
	if err != nil {
		return &v1.CreatePostResponse{}, err
	}

	return &v1.CreatePostResponse{Id: id.Bytes()}, nil
}

// Getting a post handler.
func (h *PostHandler) GetPost(ctx context.Context, input *v1.GetPostRequest) (*v1.GetPostResponse, error) {
	// Getting post by id.
	post, err := h.service.GetBy(ctx, ksuid.FromBytesOrNil(input.Id))
	if err != nil {
		return &v1.GetPostResponse{}, err
	}

	return &v1.GetPostResponse{
		AuthorId:  post.AuthorId.Bytes(),
		Text:      post.Text,
		UpdatedAt: timestamp.NewOptional(post.UpdatedAt),
	}, nil
}

// Getting posts handler.
func (h *Handler) GetPosts(ctx context.Context, input *v1.GetPostsRequest) (*v1.GetPostsResponse, error) {
	// Getting posts.
	posts, err := h.service.GetPosts(ctx, ksuid.FromBytesOrNil(input.AuthorId),
		domain.SortOptions{
			First:  input.SortOptions.First,
			Last:   input.SortOptions.Last,
			Before: ksuid.FromBytesOrNil(input.SortOptions.Before),
			After:  ksuid.FromBytesOrNil(input.SortOptions.After),
		})
	if err != nil {
		return &v1.GetPostsResponse{}, err
	}

	responsePosts := make([]*v1.Post, len(posts))

	for i, post := range posts {
		responsePosts[i] = &v1.Post{
			Id:        post.Id.Bytes(),
			Text:      post.Text,
			UpdatedAt: timestamp.NewOptional(post.UpdatedAt),
		}
	}

	return &v1.GetPostsResponse{Posts: responsePosts}, nil
}

// Deleting a post handler.
func (h *PostHandler) DeletePost(ctx context.Context, input *v1.DeletePostRequest) (*v1.DeletePostResponse, error) {
	// Deleting post.
	if err := h.service.Delete(ctx, ksuid.FromBytesOrNil(input.Id), ksuid.FromBytesOrNil(input.AuthorId)); err != nil {
		return &v1.DeletePostResponse{}, err
	}

	return &v1.DeletePostResponse{}, nil
}

// Updating a post handler.
func (h *PostHandler) UpdatePost(ctx context.Context, input *v1.UpdatePostRequest) (*v1.UpdatePostResponse, error) {
	// Updating post.
	if err := h.service.Update(ctx, domain.Post{
		Id:       ksuid.FromBytesOrNil(input.Id),
		AuthorId: ksuid.FromBytesOrNil(input.AuthorId),
		Text:     input.Text,
	}); err != nil {
		return &v1.UpdatePostResponse{}, err
	}

	return &v1.UpdatePostResponse{}, nil
}
