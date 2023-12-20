package collections

func Filter[TYPE interface{}](collection []TYPE, predicate func(item TYPE) bool) []TYPE {
	var filtered []TYPE
	for _, item := range collection {
		if predicate(item) {
			filtered = append(filtered, item)
		}
	}
	return filtered
}
