package common

import "strconv"

const (
	DEFAULT_SIZE = 20
	MAXIMUM_SIZE = 100

	CURSOR_DEFAULT = 0
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

func toCursor(cursorStr string) int {
	cursor, err := strconv.Atoi(cursorStr)
	if err != nil {
		return CURSOR_DEFAULT
	}
	return cursor
}

func toSize(sizeStr string) int {
	size, err := strconv.Atoi(sizeStr)
	if err != nil {
		return DEFAULT_SIZE
	}
	if size >= MAXIMUM_SIZE {
		return MAXIMUM_SIZE
	}
	return size
}

func (t TechbloghubPaging) HasNextPage(size int) bool {
	return t.Size < size
}
