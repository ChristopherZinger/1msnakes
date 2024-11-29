package arrays

func Reverse[T any](items []*T) {
	for i, j := 0, len(items)-1; i < j; {
		items[i], items[j] = items[j], items[i]
		i++
		j--
	}
}

func Prepend[T any](arr []*T, item *T) []*T {
	Reverse(arr)
	r := append(arr, item)
	Reverse(r)
	return r
}
