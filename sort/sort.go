package sort

import "sort"

func String[T any](list []T, f func(T) string) []string {
	newList := []string{}
	for _, v := range list {
		newList = append(newList, f(v))
	}
	sort.Strings(newList)
	return newList
}
