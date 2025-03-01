package schemasupport

import (
	"database/sql/driver"

	"github.com/lib/pq"
)

type PostingTags []string

func (s PostingTags) Value() (driver.Value, error) {
	return pq.Array(s).Value()
}

func (s *PostingTags) Scan(value interface{}) error {
	return pq.Array(s).Scan(value)
}

func (s PostingTags) ToStringArray() []string {
	return s
}
