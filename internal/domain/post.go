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

import (
	"time"

	"github.com/gofrs/uuid"
)

// Post structure.
type Post struct {
	ID        uuid.UUID
	AuthorID  uuid.UUID
	Text      string
	CreatedAt time.Time
	UpdatedAt *time.Time
}

// Validate post.
func (p Post) Validate() error {
	// Check post text length.
	if len(p.Text) > 500 {
		return &Error{Code: CodeInvalidArgument, Message: "Text is too long"}
	}

	return nil
}
