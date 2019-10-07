package list_test

import (
	"github.com/ambientsound/closures-in-go/pkg/filters"
	"github.com/ambientsound/closures-in-go/pkg/list"
	"github.com/stretchr/testify/assert"
	"testing"
)

// Test data fixture
var myList = list.List{
	list.Item{Text: "write a talk", Priority: 10},
	list.Item{Text: "go to the meetup", Priority: 15},
	list.Item{Text: "speak at the meetup", Priority: 30},
	list.Item{Text: "drink beer", Priority: 5},
	list.Item{Text: "sleep", Priority: 0},
}

func TestList_Filter(t *testing.T) {
	callback := func(item list.Item) bool {
		return item.Priority >= 10
	}

	filteredList := myList.Filter(callback)

	expectedList := list.List{
		list.Item{Text: "write a talk", Priority: 10},
		list.Item{Text: "go to the meetup", Priority: 15},
		list.Item{Text: "speak at the meetup", Priority: 30},
	}

	assert.Equal(t, expectedList, filteredList)
}

func TestList_Filter_Badwords(t *testing.T) {
	callback := func(item list.Item) bool {
		return nil == filters.BadWordFilter(item.Text)
	}

	filteredList := myList.Filter(callback)

	assert.NotContains(t, filteredList, list.Item{Text: "drink beer", Priority: 5})
}
