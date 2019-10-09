package list

// Filtering function where the caller can decide what gets filtered
// based on a callback function which is called once for each item in the list.
// If the callback returns true, the item is returned in a new list.
func (list List) Filter(callback func(Item) bool) List {

	filteredList := make(List, 0)

	for _, item := range list {
		if callback(item) {
			filteredList = append(filteredList, item)
		}
	}

	return filteredList
}
