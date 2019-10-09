// Package list implements a to-do list.
package list

// Every item on the list has a description and a priority.
type Item struct {
	Text     string
	Priority int
}

// The to-do list consists of a arbitrary number of items.
type List []Item

// Filter calls a function on each item in the list, and returns
// a new list with all the items for which the function returned true.
func (list List) Filter(callback func(Item) bool) List {

	filteredList := make(List, 0)

	for _, item := range list {
		if callback(item) {
			filteredList = append(filteredList, item)
		}
	}

	return filteredList
}
