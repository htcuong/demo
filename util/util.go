package util

const (
	DefaultOffset = 0
	DefaultLimit  = 10
)

func OffsetFromPage(page int, limit int) (offset int) {
	offset = DefaultOffset

	if page == 0 {
		page = 1
	}

	if limit == 0 {
		limit = DefaultLimit
	}

	return (page * limit) - limit
}
