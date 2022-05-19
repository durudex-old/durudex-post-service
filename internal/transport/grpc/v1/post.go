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

	"github.com/durudex/durudex-post-service/internal/service"
	v1 "github.com/durudex/durudex-post-service/pkg/pb/durudex/v1"
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

func (h *PostHandler) CreatePost(ctx context.Context, input *v1.CreatePostRequest) (*v1.CreatePostResponse, error) {
	return &v1.CreatePostResponse{}, nil
}

func (h *PostHandler) GetPostById(ctx context.Context, input *v1.GetPostByIdRequest) (*v1.GetPostByIdResponse, error) {
	return &v1.GetPostByIdResponse{}, nil
}

func (h *PostHandler) DeletePost(ctx context.Context, input *v1.DeletePostRequest) (*v1.DeletePostResponse, error) {
	return &v1.DeletePostResponse{}, nil
}

func (h *PostHandler) UpdatePost(ctx context.Context, input *v1.UpdatePostRequest) (*v1.UpdatePostResponse, error) {
	return &v1.UpdatePostResponse{}, nil
}
