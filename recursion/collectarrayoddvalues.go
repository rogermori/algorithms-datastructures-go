package recursion

func CollectOddValuesWithHelperFunction(intList []int) []int {

	var oddList = new([]int)
	collectHelper(intList, oddList, 0)
	return *oddList
}

func collectHelper(intList []int, oddList *[]int, index int) []int {
	if index == len(intList) {
		return *oddList
	}
	if intList[index]%2 != 0 {
		*oddList = append(*oddList, intList[index])
	}
	return collectHelper(intList, oddList, index+1)
}

func CollectOddValuesPure(intList []int) []int {
	var newList []int
	if len(intList) == 0 {
		return newList
	}
	if intList[0]%2 != 0 {
		newList = append(newList, intList[0])
	}
	return append(newList, CollectOddValuesPure(intList[1:])...)
}
