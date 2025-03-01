package common

import "strconv"

const (
	DEFAULT_SIZE = 20
	MAXIMUM_SIZE = 100

	LAST_VIEW_ID_DEFAULT = 0
)

type TechbloghubPaging struct {
	Cursor int
	Size   int
}

func GenerateTechPaging(cursorStr string, sizeStr string) TechbloghubPaging {
	return TechbloghubPaging{
		Cursor: toCursor(cursorStr),
		Size:   toSize(sizeStr),
	}
}

func toCursor(lastViewedIdStr string) int {
	lastViewId, err := strconv.Atoi(lastViewedIdStr)
	if err != nil {
		return LAST_VIEW_ID_DEFAULT
	}
	return lastViewId
}

func toSize(sizeStr string) int {
	size, err := strconv.Atoi(sizeStr)
	if err != nil {
		return DEFAULT_SIZE
	}
	return ceil(size)
}

func ceil(limit int) int {
	if limit >= MAXIMUM_SIZE {
		return MAXIMUM_SIZE
	}
	return limit
}

func (t TechbloghubPaging) HasNextPage(size int) bool {
	return t.Size < size
}
