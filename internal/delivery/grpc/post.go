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

package grpc

import (
	"context"

	"github.com/durudex/dugopb/types/timestamp"
	"github.com/durudex/durudex-post-service/internal/delivery/grpc/pb"
	"github.com/durudex/durudex-post-service/internal/service"

	"github.com/gofrs/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Post handler.
type PostHandler struct {
	service service.Post
	pb.UnimplementedPostServiceServer
}

// Creating a new post handler.
func NewPostHandler(service service.Post) *PostHandler {
	return &PostHandler{service: service}
}

// Creating a new post.
func (h *PostHandler) CreatePost(ctx context.Context, input *pb.CreatePostRequest) (*pb.CreatePostResponse, error) {
	// Get author uuid from bytes.
	authorID, err := uuid.FromBytes(input.AuthorId)
	if err != nil {
		return &pb.CreatePostResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	id, err := h.service.Create(ctx, authorID, input.Text)
	if err != nil {
		return &pb.CreatePostResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &pb.CreatePostResponse{Id: id.Bytes()}, nil
}

// Getting a post by id.
func (h *PostHandler) GetPostByID(ctx context.Context, input *pb.GetPostByIDRequest) (*pb.GetPostByIDResponse, error) {
	// Get user uuid from bytes.
	userID, err := uuid.FromBytes(input.Id)
	if err != nil {
		return &pb.GetPostByIDResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	// Get post by id.
	post, err := h.service.GetByID(ctx, userID)
	if err != nil {
		return &pb.GetPostByIDResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &pb.GetPostByIDResponse{
		AuthorId:  post.AuthorID.Bytes(),
		Text:      post.Text,
		CreatedAt: timestamp.New(post.CreatedAt),
		UpdatedAt: timestamp.NewOptional(post.UpdatedAt),
	}, nil
}

// Deleting a post.
func (h *PostHandler) DeletePost(ctx context.Context, input *pb.DeletePostRequest) (*pb.DeletePostResponse, error) {
	// Get user uuid from bytes.
	userID, err := uuid.FromBytes(input.Id)
	if err != nil {
		return &pb.DeletePostResponse{}, status.Error(codes.InvalidArgument, err.Error())
	}

	// Delete post.
	if err := h.service.Delete(ctx, userID); err != nil {
		return &pb.DeletePostResponse{}, status.Error(codes.Internal, err.Error())
	}

	return &pb.DeletePostResponse{}, nil
}
