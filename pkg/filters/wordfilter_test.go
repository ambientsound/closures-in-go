package filters_test

import (
	"github.com/ambientsound/closures-in-go/pkg/filters"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBadWordFilter(t *testing.T) {
	assert.NoError(t, filters.BadWordFilter("foo to the bar"))
	assert.NoError(t, filters.BadWordFilter("bar to the baz"))
	assert.Error(t, filters.BadWordFilter("beer to the belly"))
}
