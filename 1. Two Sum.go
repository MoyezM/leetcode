func twoSum(nums []int, target int) []int {
    hash := make(map[int]int)
    
    for i, num := range nums {
        if val, ok := hash[num]; ok && i != val {
            return  []int{i, val}
        }
        
        hash[target - num] = i
        
    }
    
    return []int{-1,-1}
}