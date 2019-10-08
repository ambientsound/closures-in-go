package filters_test

import (
	"github.com/ambientsound/closures-in-go/pkg/filters"
	"github.com/stretchr/testify/assert"
	"testing"
)

var badWords = []string{
	"beer",
}

func TestBadWordFilter(t *testing.T) {
	badWordFilter := filters.BadWordFilter(badWords)

	assert.NoError(t, badWordFilter("foo to the bar"))
	assert.NoError(t, badWordFilter("bar to the baz"))
	assert.Error(t, badWordFilter("beer to the belly"))
}
