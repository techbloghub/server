package schemasupport

import (
	"database/sql/driver"

	"github.com/lib/pq"
)

type PostingTags []string

func (t *PostingTags) Scan(value interface{}) error {
	if value == nil {
		*t = []string{}
		return nil
	}

	var arr []string
	if err := pq.Array(&arr).Scan(value); err != nil {
		return err
	}

	*t = arr
	return nil
}

func (t PostingTags) Value() (driver.Value, error) {
	return pq.Array(t).Value()
}

func (t *PostingTags) ToStringSlice() []string {
	tags := make([]string, len(*t))
	for i, tag := range *t {
		tags[i] = tag
	}
	return tags
}
