// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0

package products

import (
	"database/sql"
)

type Product struct {
	ID          int32
	UserID      int32
	Name        string
	Description string
	Price       int32
	UploadDate  sql.NullTime
}
