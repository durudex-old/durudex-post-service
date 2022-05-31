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

package postgres_test

import (
	"context"
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/durudex/durudex-post-service/internal/domain"
	"github.com/durudex/durudex-post-service/internal/repository/postgres"

	"github.com/gofrs/uuid"
	"github.com/pashagolub/pgxmock"
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
	type mockBehavior func(args args, id uuid.UUID)

	// Creating a new repository.
	repos := postgres.NewPostRepository(mock)

	// Tests structures.
	tests := []struct {
		name         string
		args         args
		want         uuid.UUID
		wantErr      bool
		mockBehavior mockBehavior
	}{
		{
			name: "OK",
			args: args{post: domain.Post{AuthorID: uuid.UUID{}, Text: "text"}},
			want: uuid.UUID{},
			mockBehavior: func(args args, want uuid.UUID) {
				mock.ExpectQuery(fmt.Sprintf(`INSERT INTO "%s"`, postgres.PostTable)).
					WithArgs(args.post.AuthorID, args.post.Text).
					WillReturnRows(mock.NewRows([]string{"id"}).AddRow(want))
			},
		},
	}

	// Conducting tests in various structures.
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockBehavior(tt.args, tt.want)

			// Creating a new post in postgres database.
			got, err := repos.Create(context.Background(), tt.args.post)
			if (err != nil) != tt.wantErr {
				t.Errorf("error creating post: %s", err.Error())
			}

			// Check for similarity of id.
			if !reflect.DeepEqual(got, tt.want) {
				t.Error("error id are not similar")
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
	type args struct{ id uuid.UUID }

	// Test behavior.
	type mockBehavior func(args args, user domain.Post)

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
			args: args{id: uuid.UUID{}},
			want: domain.Post{
				AuthorID:  uuid.UUID{},
				Text:      "text",
				CreatedAt: time.Now(),
				UpdatedAt: nil,
			},
			mockBehavior: func(args args, post domain.Post) {
				rows := mock.NewRows([]string{"author_id", "text", "created_at", "updated_at"}).AddRow(
					post.AuthorID, post.Text, post.CreatedAt, post.UpdatedAt)

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
			got, err := repos.GetByID(context.Background(), tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("error getting post by id: %s", err.Error())
			}

			// Check for similarity of post.
			if !reflect.DeepEqual(got, tt.want) {
				t.Error("error user are not similar")
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
	type args struct{ id, authorId uuid.UUID }

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
			args:    args{id: uuid.UUID{}, authorId: uuid.UUID{}},
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
			args:    args{post: domain.Post{ID: uuid.UUID{}, AuthorID: uuid.UUID{}, Text: "text"}},
			wantErr: false,
			mockBehavior: func(args args) {
				mock.ExpectExec(fmt.Sprintf(`UPDATE "%s"`, postgres.PostTable)).
					WithArgs(args.post.Text, args.post.ID, args.post.AuthorID).
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
