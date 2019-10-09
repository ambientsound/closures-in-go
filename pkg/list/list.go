// Package list implements a TODO list.
// Every item on the list has a description and a priority.
package list

type Item struct {
	Text     string
	Priority int
}

type List []Item

// Filter calls a function on each item in the list.
// Returns a new list where the items having the
func (list List) Filter(callback func(Item) bool) List {
	filteredList := make(List, 0)

	for _, item := range list {
		if callback(item) {
			filteredList = append(filteredList, item)
		}
	}

	return filteredList
}
