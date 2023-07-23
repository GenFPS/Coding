package main

func validMountainArray(arr []int) bool {
	if len(arr) <= 2 {
		return false
	}

	i := 0
	for i+1 < len(arr) && arr[i] < arr[i+1] {
		i++
	}
	/*Вершина не может быть первым или послединим элементом
	(т.е массив строго возрастает до вершины, а после вершини - строго убывает.)
	*/
	if i == 0 || i == len(arr)-1 {
		return false
	}
	for i+1 < len(arr) && arr[i] > arr[i+1] {
		i++
	}
	return i == len(arr)-1
}
