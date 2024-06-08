package search

// LinearSearch implements linear search - O(N)
// Data can be unordered
// https://frontendmasters.com/courses/algorithms/linear-search-kata-setup/
func LinearSearch(a []int, x int) bool {
	for i := 0; i < len(a); i++ {
		if a[i] == x {
			return true
		}
	}

	return false
}
