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

package postgres_test

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	"github.com/durudex/durudex-post-service/internal/domain"
	"github.com/durudex/durudex-post-service/internal/repository/postgres"

	"github.com/pashagolub/pgxmock"
	"github.com/segmentio/ksuid"
)

// Testing creating a new post in postgres database.
func TestPostRepository_Create(t *testing.T) {
	// Creating a new mock connection.
	mock, err := pgxmock.NewConn()
	if err != nil {
		t.Fatalf("error creating a new mock connection: %s", err.Error())
	}
	defer mock.Close(context.Background())

	// Testing args.
	type args struct{ post domain.Post }

	// Test behavior.
	type mockBehavior func(args args)

	// Creating a new repository.
	repos := postgres.NewPostRepository(mock)

	// Tests structures.
	tests := []struct {
		name         string
		args         args
		wantErr      bool
		mockBehavior mockBehavior
	}{
		{
			name: "OK",
			args: args{post: domain.Post{Id: ksuid.New(), AuthorId: ksuid.New(), Text: "text"}},
			mockBehavior: func(args args) {
				mock.ExpectExec(fmt.Sprintf(`INSERT INTO "%s"`, postgres.PostTable)).
					WithArgs(args.post.Id, args.post.AuthorId, args.post.Text).
					WillReturnResult(pgxmock.NewResult("", 1))
			},
		},
	}

	// Conducting tests in various structures.
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockBehavior(tt.args)

			// Creating a new post in postgres database.
			err := repos.Create(context.Background(), tt.args.post)
			if (err != nil) != tt.wantErr {
				t.Errorf("error creating post: %s", err.Error())
			}
		})
	}
}

// Testing getting a post by id in postgres database.
func TestPostRepository_GetByID(t *testing.T) {
	// Creating a new mock connection.
	mock, err := pgxmock.NewConn()
	if err != nil {
		t.Fatalf("error creating a new mock connection: %s", err.Error())
	}
	defer mock.Close(context.Background())

	// Testing args.
	type args struct{ id ksuid.KSUID }

	// Test behavior.
	type mockBehavior func(args args, post domain.Post)

	// Creating a new repository.
	repos := postgres.NewPostRepository(mock)

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
				AuthorId:  ksuid.New(),
				Text:      "text",
				UpdatedAt: nil,
			},
			mockBehavior: func(args args, post domain.Post) {
				rows := mock.NewRows([]string{"author_id", "text", "updated_at"}).AddRow(
					post.AuthorId, post.Text, post.UpdatedAt)

				mock.ExpectQuery(fmt.Sprintf(`SELECT (.+) FROM "%s"`, postgres.PostTable)).
					WithArgs(args.id).
					WillReturnRows(rows)
			},
		},
	}

	// Conducting tests in various structures.
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockBehavior(tt.args, tt.want)

			// Getting a post by id in postgres database.
			got, err := repos.GetById(context.Background(), tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("error getting post by id: %s", err.Error())
			}

			// Check for similarity of post.
			if !reflect.DeepEqual(got, tt.want) {
				t.Error("error post are not similar")
			}
		})
	}
}

// Testing getting author posts by author id in postgres database.
func TestPostRepository_GetPosts(t *testing.T) {
	// Creating a new mock connection.
	mock, err := pgxmock.NewConn()
	if err != nil {
		t.Fatalf("error creating a new mock connection: %s", err.Error())
	}
	defer mock.Close(context.Background())

	// Testing args.
	type args struct {
		authorId ksuid.KSUID
		first    *int32
		last     *int32
	}

	// Test behavior.
	type mockBehavior func(args args, want []domain.Post)

	// Creating a new repository.
	repos := postgres.NewPostRepository(mock)

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
				first:    &filer,
				last:     nil,
			},
			want: []domain.Post{
				{
					Id:        ksuid.New(),
					AuthorId:  ksuid.New(),
					Text:      "text",
					UpdatedAt: nil,
				},
			},
			mockBehavior: func(args args, want []domain.Post) {
				rows := mock.NewRows([]string{"id", "author_id", "text", "updated_at"}).AddRow(
					want[0].Id, want[0].AuthorId, want[0].Text, want[0].UpdatedAt,
				)

				mock.ExpectQuery(fmt.Sprintf(`SELECT (.+) FROM "%s"`, postgres.PostTable)).
					WithArgs(args.authorId, *args.first).
					WillReturnRows(rows)
			},
		},
	}

	// Conducting tests in various structures.
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockBehavior(tt.args, tt.want)

			// Getting a post by id in postgres database.
			got, err := repos.GetPosts(context.Background(), tt.args.authorId, tt.args.first, tt.args.last)
			if (err != nil) != tt.wantErr {
				t.Errorf("error getting author posts: %s", err.Error())
			}

			// Check for similarity of posts.
			if !reflect.DeepEqual(got[0], tt.want[0]) {
				t.Error("error posts are not similar")
			}
		})
	}
}

// Testing deleting a post in postgres database.
func TestPostRepository_Delete(t *testing.T) {
	// Creating a new mock connection.
	mock, err := pgxmock.NewConn()
	if err != nil {
		t.Fatalf("error creating a new mock connection: %s", err.Error())
	}
	defer mock.Close(context.Background())

	// Testing args.
	type args struct{ id, authorId ksuid.KSUID }

	// Test behavior.
	type mockBehavior func(args args)

	// Creating a new repository.
	repos := postgres.NewPostRepository(mock)

	// Tests structures.
	tests := []struct {
		name         string
		args         args
		wantErr      bool
		mockBehavior mockBehavior
	}{
		{
			name:    "OK",
			args:    args{id: ksuid.New(), authorId: ksuid.New()},
			wantErr: false,
			mockBehavior: func(args args) {
				mock.ExpectExec(fmt.Sprintf(`DELETE FROM "%s"`, postgres.PostTable)).
					WithArgs(args.id, args.authorId).
					WillReturnResult(pgxmock.NewResult("", 1))
			},
		},
	}

	// Conducting tests in various structures.
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockBehavior(tt.args)

			// Deleting a post in postgres database.
			err := repos.Delete(context.Background(), tt.args.id, tt.args.authorId)
			if (err != nil) != tt.wantErr {
				t.Errorf("error deleting post by id: %s", err.Error())
			}
		})
	}
}

// Testing updating a post in postgres database.
func TestPostRepository_Update(t *testing.T) {
	// Creating a new mock connection.
	mock, err := pgxmock.NewConn()
	if err != nil {
		t.Fatalf("error creating a new mock connection: %s", err.Error())
	}
	defer mock.Close(context.Background())

	// Testing args.
	type args struct{ post domain.Post }

	// Test behavior.
	type mockBehavior func(args args)

	// Creating a new repository.
	repos := postgres.NewPostRepository(mock)

	// Tests structures.
	tests := []struct {
		name         string
		args         args
		wantErr      bool
		mockBehavior mockBehavior
	}{
		{
			name:    "OK",
			args:    args{post: domain.Post{Id: ksuid.New(), AuthorId: ksuid.New(), Text: "text"}},
			wantErr: false,
			mockBehavior: func(args args) {
				mock.ExpectExec(fmt.Sprintf(`UPDATE "%s"`, postgres.PostTable)).
					WithArgs(args.post.Text, args.post.Id, args.post.AuthorId).
					WillReturnResult(pgxmock.NewResult("", 1))
			},
		},
	}

	// Conducting tests in various structures.
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockBehavior(tt.args)

			// Updating a post in postgres database.
			err := repos.Update(context.Background(), tt.args.post)
			if (err != nil) != tt.wantErr {
				t.Errorf("error updating post by id: %s", err.Error())
			}
		})
	}
}
