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

package service_test

import (
	"context"
	"reflect"
	"testing"

	"github.com/durudex/durudex-post-service/internal/domain"
	mock_postgres "github.com/durudex/durudex-post-service/internal/repository/postgres/mock"
	"github.com/durudex/durudex-post-service/internal/service"

	"github.com/golang/mock/gomock"
	"github.com/segmentio/ksuid"
)

// Testing creating a new post.
func TestPostService_Create(t *testing.T) {
	// Creating a new mock controller.
	c := gomock.NewController(t)
	defer c.Finish()

	// Creating a new mock repository.
	psql := mock_postgres.NewMockPost(c)

	// Testing args.
	type args struct{ post domain.Post }

	// Test behavior.
	type mockBehavior func(r *mock_postgres.MockPost, args args)

	// Tests structures.
	tests := []struct {
		name         string
		args         args
		want         ksuid.KSUID
		wantErr      bool
		mockBehavior mockBehavior
	}{
		{
			name: "OK",
			args: args{domain.Post{
				Id:       ksuid.New(),
				AuthorId: ksuid.New(),
				Text:     "This is a test post.",
			}},
			mockBehavior: func(r *mock_postgres.MockPost, args args) {
				r.EXPECT().Create(context.Background(), args.post).Return(nil)
			},
		},
	}

	// Conducting tests in various structures.
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setting a mock behavior.
			tt.mockBehavior(psql, tt.args)

			// Creating a new post service.
			service := service.NewPostService(psql)

			// Creating a new post.
			id, err := service.Create(context.Background(), tt.args.post)
			if err != nil {
				t.Errorf("error creating post: %s", err.Error())
			}

			// Check id is nil.
			if id.IsNil() {
				t.Error("post id is nil")
			}
		})
	}
}

// Testing getting a post.
func TestPostService_Get(t *testing.T) {
	// Creating a new mock controller.
	c := gomock.NewController(t)
	defer c.Finish()

	// Creating a new mock repository.
	psql := mock_postgres.NewMockPost(c)

	// Testing args.
	type args struct{ id ksuid.KSUID }

	// Test behavior.
	type mockBehavior func(r *mock_postgres.MockPost, args args, want domain.Post)

	// Tests structures.
	tests := []struct {
		name         string
		args         args
		want         domain.Post
		wantErr      bool
		mockBehavior mockBehavior
	}{
		{
			name: "OK",
			args: args{id: ksuid.New()},
			want: domain.Post{
				AuthorId: ksuid.New(),
				Text:     "This is a test post.",
			},
			mockBehavior: func(r *mock_postgres.MockPost, args args, want domain.Post) {
				r.EXPECT().Get(context.Background(), args.id).Return(want, nil)
			},
		},
	}

	// Conducting tests in various structures.
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setting a mock behavior.
			tt.mockBehavior(psql, tt.args, tt.want)

			// Creating a new post service.
			service := service.NewPostService(psql)

			// Getting a post by id.
			got, err := service.Get(context.Background(), tt.args.id)
			if err != nil {
				t.Errorf("error getting post by id: %s", err.Error())
			}

			// Check for similarity of post.
			if !reflect.DeepEqual(got, tt.want) {
				t.Error("error post are not similar")
			}
		})
	}
}

// Testing getting author posts.
func TestPostService_GetPosts(t *testing.T) {
	// Creating a new mock controller.
	c := gomock.NewController(t)
	defer c.Finish()

	// Creating a new mock repository.
	psql := mock_postgres.NewMockPost(c)

	// Testing args.
	type args struct {
		authorId ksuid.KSUID
		sort     domain.SortOptions
	}

	// Test behavior.
	type mockBehavior func(r *mock_postgres.MockPost, args args, want []domain.Post)

	// Query filter,
	var filer int32 = 12

	// Tests structures.
	tests := []struct {
		name         string
		args         args
		want         []domain.Post
		wantErr      bool
		mockBehavior mockBehavior
	}{
		{
			name: "OK",
			args: args{
				authorId: ksuid.New(),
				sort: domain.SortOptions{
					First: &filer,
				},
			},
			want: []domain.Post{
				{
					Id:   ksuid.New(),
					Text: "This is a test post.",
				},
			},
			mockBehavior: func(r *mock_postgres.MockPost, args args, want []domain.Post) {
				r.EXPECT().GetPosts(context.Background(), args.authorId, args.sort).Return(want, nil)
			},
		},
	}

	// Conducting tests in various structures.
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setting a mock behavior.
			tt.mockBehavior(psql, tt.args, tt.want)

			// Creating a new post service.
			service := service.NewPostService(psql)

			// Getting a post by id.
			got, err := service.GetPosts(context.Background(), tt.args.authorId, tt.args.sort)
			if err != nil {
				t.Errorf("error getting posts: %s", err.Error())
			}

			// Check for similarity of post.
			if !reflect.DeepEqual(got, tt.want) {
				t.Error("error posts are not similar")
			}
		})
	}
}

// Testing deleting a post.
func TestPostService_Delete(t *testing.T) {
	// Creating a new mock controller.
	c := gomock.NewController(t)
	defer c.Finish()

	// Creating a new mock repository.
	psql := mock_postgres.NewMockPost(c)

	// Testing args.
	type args struct{ id, authorId ksuid.KSUID }

	// Test behavior.
	type mockBehavior func(r *mock_postgres.MockPost, args args)

	// Tests structures.
	tests := []struct {
		name         string
		args         args
		wantErr      bool
		mockBehavior mockBehavior
	}{
		{
			name: "OK",
			args: args{
				id:       ksuid.New(),
				authorId: ksuid.New(),
			},
			wantErr: false,
			mockBehavior: func(r *mock_postgres.MockPost, args args) {
				r.EXPECT().Delete(context.Background(), args.id, args.authorId).Return(nil)
			},
		},
	}

	// Conducting tests in various structures.
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setting a mock behavior.
			tt.mockBehavior(psql, tt.args)

			// Creating a new post service.
			service := service.NewPostService(psql)

			// Deleting a post.
			if err := service.Delete(context.Background(), tt.args.id, tt.args.authorId); err != nil {
				if !tt.wantErr {
					t.Errorf("error deleting a post: %s", err.Error())
				}
			}
		})
	}
}

// Testing updating a post.
func TestPostService_Update(t *testing.T) {
	// Creating a new mock controller.
	c := gomock.NewController(t)
	defer c.Finish()

	// Creating a new mock repository.
	psql := mock_postgres.NewMockPost(c)

	// Testing args.
	type args struct{ post domain.Post }

	// Test behavior.
	type mockBehavior func(r *mock_postgres.MockPost, args args)

	// Tests structures.
	tests := []struct {
		name         string
		args         args
		wantErr      bool
		mockBehavior mockBehavior
	}{
		{
			name: "OK",
			args: args{domain.Post{
				Id:       ksuid.New(),
				AuthorId: ksuid.New(),
				Text:     "This is a test post.",
			}},
			wantErr: false,
			mockBehavior: func(r *mock_postgres.MockPost, args args) {
				r.EXPECT().Update(context.Background(), args.post).Return(nil)
			},
		},
	}

	// Conducting tests in various structures.
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setting a mock behavior.
			tt.mockBehavior(psql, tt.args)

			// Creating a new post service.
			service := service.NewPostService(psql)

			// Updating a post.
			if err := service.Update(context.Background(), tt.args.post); err != nil {
				if !tt.wantErr {
					t.Errorf("error updating a post: %s", err.Error())
				}
			}
		})
	}
}
