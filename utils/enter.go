package utils

func FindMissingElements[T comparable](A, B []T) []T {
	// 将 B 转换为 map 以提高查找效率
	bMap := make(map[T]struct{})
	for _, item := range B {
		bMap[item] = struct{}{}
	}

	// 遍历 A，收集 B 中不存在的元素
	var missing []T
	for _, item := range A {
		if _, exists := bMap[item]; !exists {
			missing = append(missing, item)
		}
	}
	return missing
}
