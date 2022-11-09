package realworld

var (
	prefixSum []int
)

func PrefixSum(arr []int) []int {
	prefixSum = make([]int, len(arr))

	prefixSum[0] = arr[0]
	for i := 1; i < len(arr); i++ {
		prefixSum[i] = prefixSum[i-1] + arr[i]
	}

	return prefixSum
}
