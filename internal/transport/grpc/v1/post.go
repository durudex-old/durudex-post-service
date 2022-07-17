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

package v1

import (
	"context"

	"github.com/durudex/dugopb/type/timestamp"
	"github.com/durudex/durudex-post-service/internal/domain"
	"github.com/durudex/durudex-post-service/internal/service"
	v1 "github.com/durudex/durudex-post-service/pkg/pb/durudex/v1"

	"github.com/segmentio/ksuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
	// Getting post author ksuid from bytes.
	authorId, err := ksuid.FromBytes(input.AuthorId)
	if err != nil {
		return &v1.CreatePostResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	// Create a new post.
	id, err := h.service.Create(ctx, domain.Post{
		AuthorId: authorId,
		Text:     input.Text,
	})
	if err != nil {
		return &v1.CreatePostResponse{}, err
	}

	return &v1.CreatePostResponse{Id: id.Bytes()}, nil
}

// Getting a post by id handler.
func (h *PostHandler) GetPostById(ctx context.Context, input *v1.GetPostByIdRequest) (*v1.GetPostByIdResponse, error) {
	// Getting post ksuid from bytes.
	id, err := ksuid.FromBytes(input.Id)
	if err != nil {
		return &v1.GetPostByIdResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	// Getting post by id.
	post, err := h.service.GetByID(ctx, id)
	if err != nil {
		return &v1.GetPostByIdResponse{}, err
	}

	return &v1.GetPostByIdResponse{
		AuthorId:  post.AuthorId.Bytes(),
		Text:      post.Text,
		UpdatedAt: timestamp.NewOptional(post.UpdatedAt),
	}, nil
}

// Deleting a post handler.
func (h *PostHandler) DeletePost(ctx context.Context, input *v1.DeletePostRequest) (*v1.DeletePostResponse, error) {
	// Getting post ksuid from bytes.
	id, err := ksuid.FromBytes(input.Id)
	if err != nil {
		return &v1.DeletePostResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	// Getting post author ksuid from bytes.
	authorId, err := ksuid.FromBytes(input.AuthorId)
	if err != nil {
		return &v1.DeletePostResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	// Deleting post.
	if err := h.service.Delete(ctx, id, authorId); err != nil {
		return &v1.DeletePostResponse{}, err
	}

	return &v1.DeletePostResponse{}, nil
}

// Updating a post handler.
func (h *PostHandler) UpdatePost(ctx context.Context, input *v1.UpdatePostRequest) (*v1.UpdatePostResponse, error) {
	// Getting post ksuid from bytes.
	id, err := ksuid.FromBytes(input.Id)
	if err != nil {
		return &v1.UpdatePostResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	// Getting post author ksuid from bytes.
	authorId, err := ksuid.FromBytes(input.AuthorId)
	if err != nil {
		return &v1.UpdatePostResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	// Updating post.
	if err := h.service.Update(ctx, domain.Post{
		Id:       id,
		AuthorId: authorId,
		Text:     input.Text,
	}); err != nil {
		return &v1.UpdatePostResponse{}, err
	}

	return &v1.UpdatePostResponse{}, nil
}
