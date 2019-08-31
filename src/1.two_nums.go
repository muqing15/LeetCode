package src

func twoSum(nums []int, target int) []int {
	var mapArr = make(map[int]int)
	var numArr []int
	for i , v := range nums {
		for k:=i+1;k<len(nums);k++{
			if v+nums[k] == target{
				flag := 0
				for mi , mv := range mapArr{
					if k == mi || k == mv {
						flag = 1
					}
				}
				if flag == 0{
					mapArr[i] = k
				}
			}
		}
	}
	for i , v := range mapArr{
		numArr = append(numArr, i, v)
	}
	return numArr
}