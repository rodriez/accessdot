package accessdot_test

import (
	"testing"

	"github.com/rodriez/accessdot"
	"github.com/stretchr/testify/assert"
)

func TestParse_1(t *testing.T) {
	str := "call"

	reader := accessdot.NewExpressionReader(".[]", "EOF")
	list := accessdot.NewParser(reader).Parse(str)

	assert.NotEmpty(t, list)
	assert.Equal(t, "key", list[0].Type)
	assert.Equal(t, "call", list[0].Key)
}

func TestParse_2(t *testing.T) {
	str := "purchase.item.price"

	reader := accessdot.NewExpressionReader(".[]", "EOF")
	list := accessdot.NewParser(reader).Parse(str)

	assert.NotEmpty(t, list)
	assert.Equal(t, "key", list[0].Type)
	assert.Equal(t, "purchase", list[0].Key)
	assert.Equal(t, "key", list[1].Type)
	assert.Equal(t, "item", list[1].Key)
	assert.Equal(t, "key", list[2].Type)
	assert.Equal(t, "price", list[2].Key)
}

func TestParse_3(t *testing.T) {
	str := "cart.items[2].id"

	reader := accessdot.NewExpressionReader(".[]", "EOF")
	list := accessdot.NewParser(reader).Parse(str)

	assert.NotEmpty(t, list)
	assert.Equal(t, "key", list[0].Type)
	assert.Equal(t, "cart", list[0].Key)
	assert.Equal(t, "key", list[1].Type)
	assert.Equal(t, "items", list[1].Key)
	assert.Equal(t, "index", list[2].Type)
	assert.Equal(t, 2, list[2].Index)
	assert.Equal(t, "key", list[3].Type)
	assert.Equal(t, "id", list[3].Key)
}
