package sort

// BubbleSort implements bubble sort - O(N^2)
// Repeatedly steps through the input list, swapping their values if needed until no swaps have
// to be performed during a pass, meaning that the list has become fully sorted.
func BubbleSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				tmp := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = tmp
			}
		}
	}

	return arr
}
