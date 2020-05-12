func threeSum(nums []int) [][]int {
    
    var ret [][]int
//  sorting allows us easily check for duplicates
//  it was a nighmare otherwise
    sort.Ints(nums)
    
    for i, num := range nums {
//      adjust for duplicate values
        
        if i > 0 && num == nums[i-1] || i > len(nums) - 2 {
			continue
		}
        
        t := - num 
        i, j := i+1, len(nums)-1
        
        for i < j {
            sum := nums[i] + nums[j]
                 
            if sum == t {
                ret = append(ret, []int{num, nums[i], nums[j]})
//              update the indexs
                i++
                j--
                
//              adjust for duplicate values
                for i < j && nums[j] == nums[j+1] {
                    j--
                }
                
                for i < j && nums[i] == nums[i-1] {
                    i++
                }
                
            } else if sum > t {
                j--                
            } else {
                i++
            }
        }
        
    }
    
    return ret    
}
