package utils

func SliceUnique[T comparable](s []T) []T {
	if len(s) <= 0 {
		return s
	}
	inResult := make(map[T]bool)
	rs := make([]T, 0, len(s))
	for _, str := range s {
		if _, ok := inResult[str]; !ok {
			inResult[str] = true
			rs = append(rs, str)
		}
	}
	result := make([]T, len(rs))
	_ = copy(result, rs)
	return result
}

func SliceContains[T comparable](slice []T, target T) bool {
	for _, e := range slice {
		if e == target {
			return true
		}
	}
	return false
}

func SliceToMap[T comparable, M comparable](list []T, f func(T) M) map[M]T {
	result := map[M]T{}
	for idx := range list {
		item := list[idx]
		result[f(item)] = item
	}
	return result
}

func SliceRemoveFunc[E any](s []E, f func(E) bool) []E {
	if len(s) <= 0 {
		return s
	}
	for i, v := range s {
		if f(v) {
			return SliceRemoveFunc(append(s[:i], s[i+1:]...), f)
		}
	}
	return s
}

func SliceFindItemFunc[E any](s []E, f func(E) bool) (rs E, idx int) {
	if len(s) <= 0 {
		return rs, -1
	}
	for i, v := range s {
		if f(v) {
			return v, i
		}
	}
	return rs, -1
}

func SliceAddFirst[E any](s []E, insertValue E) []E {
	if len(s) <= 0 {
		return s
	}
	res := make([]E, len(s)+1)
	copy(res[1:], s)
	res[0] = insertValue
	return res
}

func SliceInsertWithIndex[T any](array []T, value T, index int) []T {
	if len(array) <= 0 {
		return array
	}
	return append(array[:index], append([]T{value}, array[index:]...)...)
}

func SliceRemoveWithInndex[T any](array []T, index int) []T {
	if len(array) <= 0 {
		return array
	}
	return append(array[:index], array[index+1:]...)
}

func SliceMoveElement[T any](array []T, srcIndex int, dstIndex int) []T {
	if len(array) <= 0 {
		return array
	}
	value := array[srcIndex]
	return SliceInsertWithIndex(SliceRemoveWithInndex(array, srcIndex), value, dstIndex)
}
