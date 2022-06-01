package recursion

func CollectOddValues(intList []int) []int {

	var oddList = new([]int)
	collectHelper(intList, oddList, 0)
	return *oddList
}

func collectHelper(intList []int, oddList *[]int, index int) []int {
	if index == len(intList) {
		return *oddList
	}
	if intList[0]%2 != 0 {
		*oddList = append(*oddList, intList[0])
	}
	return collectHelper(intList, oddList, index+1)
}
