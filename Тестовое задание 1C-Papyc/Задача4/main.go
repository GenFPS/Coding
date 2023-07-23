package main

import "fmt"

/*
Задача 4.

Есть массив целых чисел, элементы массива отсортированы по возрастанию.
Есть целое число k.
Необходимо найти и вернуть любую пару чисел, сумма которых равняется числу k.
Если такой пары нет, то необходимо вернуть пустой массив.
*/

// Решается с помощью двух указателей

func findPair(arr []int, target int) []int {
	pairArr := make([]int, 2)

	i, j := 0, len(arr)-1

	for i != j {
		if arr[i]+arr[j] > target {
			j--
		} else if arr[i]+arr[j] < target {
			i++
		} else if arr[i]+arr[j] == target {
			pairArr[0], pairArr[1] = arr[i], arr[j]
			return pairArr
		} else {
			return pairArr
		}
	}
	return pairArr
}

func main() {
	arr := []int{2, 3, 4}
	k := 6

	fmt.Println(findPair(arr, k))
}
