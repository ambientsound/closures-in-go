package list_test

import (
	"github.com/ambientsound/closures-in-go/pkg/filters"
	"github.com/ambientsound/closures-in-go/pkg/list"
	"github.com/stretchr/testify/assert"
	"testing"
)

// Test data fixture
var myTODOList = list.List{
	list.Item{Text: "write a talk", Priority: 10},
	list.Item{Text: "go to the meetup", Priority: 15},
	list.Item{Text: "speak at the meetup", Priority: 30},
	list.Item{Text: "drink beer", Priority: 5},
	list.Item{Text: "sleep", Priority: 0},
}

// Generate a closure that filters out items with priority less than N.
func priorityGreaterThan(priority int) func(list.Item) bool {
	return func(item list.Item) bool {
		return item.Priority >= priority
	}
}

// Generate a closure that blocks items where these words are found.
func withoutWords(words []string) func(list.Item) bool {
	wordFilter := filters.BadWordFilter(words)

	return func(item list.Item) bool {
		return nil == wordFilter(item.Text)
	}
}

func TestList_Filter(t *testing.T) {
	importantList := myTODOList.Filter(priorityGreaterThan(10))

	expectedList := list.List{
		list.Item{Text: "write a talk", Priority: 10},
		list.Item{Text: "go to the meetup", Priority: 15},
		list.Item{Text: "speak at the meetup", Priority: 30},
	}

	assert.Equal(t, expectedList, importantList)
}

func TestList_Filter_Many(t *testing.T) {
	safeImportantList := myTODOList.
		Filter(priorityGreaterThan(10)).
		Filter(withoutWords([]string{"meetup"}))

	expectedList := list.List{
		list.Item{Text: "write a talk", Priority: 10},
	}

	assert.Equal(t, expectedList, safeImportantList)
}
