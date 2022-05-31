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

package domain

import "testing"

// Testing validate a post.
func TestPost_Validate(t *testing.T) {
	// Testing args.
	type args struct{ text string }

	// Tests structures.
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "OK",
			args:    args{text: "Hello world!"},
			wantErr: false,
		},
	}

	// Conducting tests in various structures.
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Creating a new post.
			post := Post{Text: tt.args.text}

			// Validate post.
			err := post.Validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("error validation post: %s", err.Error())
			}
		})
	}
}
