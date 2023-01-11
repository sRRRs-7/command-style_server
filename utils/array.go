package utils

func StarContains(arr []int64, userId int64) []int64 {
	for i := range arr {
		if arr[i] == userId {
			arr[i] = arr[len(arr)-1]
			return arr[:len(arr)-1]
		}
	}
	arr = append(arr, userId)
	return arr
}

func CreateMap(key, value string, list map[string]string) map[string]string {
	list[key] = value
	return list
}
