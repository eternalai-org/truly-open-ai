package utils

// MapClone clone a map
func MapClone(m map[string]interface{}) map[string]interface{} {
	cp := make(map[string]interface{})
	for k, v := range m {
		vm, ok := v.(map[string]interface{})
		if ok {
			cp[k] = MapClone(vm)
		} else {
			cp[k] = v
		}
	}

	return cp
}
