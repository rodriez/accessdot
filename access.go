package accessdot

import (
	"strconv"
)

type Access struct {
	Type  string
	Key   string
	Index int
}

func NewAccess(char, buff string) Access {
	switch char {
	case "]":
		i, _ := strconv.Atoi(buff)
		return NewIndexAccess(i)
	default:
		return NewKeyAccess(buff)
	}
}

func NewKeyAccess(key string) Access {
	return Access{
		Type: "key",
		Key:  key,
	}
}

func NewIndexAccess(i int) Access {
	return Access{
		Type:  "index",
		Index: i,
	}
}
