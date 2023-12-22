package collections

func Map[TYPE interface{}, RETURN interface{}](collection []TYPE, transform func(item TYPE) RETURN) []RETURN {
	var result []RETURN
	for _, item := range collection {
		transformed := transform(item)
		result = append(result, transformed)
	}
	return result
}
