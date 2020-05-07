func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	//     first find the small and large array 
		x := len(nums1)
		y := len(nums2)
		
		var sm []int
		var lg []int
		
		if x <= y {
			sm = nums1
			lg = nums2
		} else {
			sm = nums2
			lg = nums1
		}
		
		
	//     Once that has been established try to find the partition
		
		start := 0
		end := len(sm) - 1
		
	//     create a partition to split both arrays into equal size halves
	//     when the two left parts are compared to the two right sides
	//     as the median is the value that can divide the combined sorted list
	//     into two even parts
		pos_sm :=  int(start+end) / 2
		pos_lg := int((x+y+1) / 2) - pos_sm
		
		max_left_sm := 0
		max_left_lg := 0
		
		min_right_sm := 0
		min_right_lg := 0
		
		for {
			
	//         if you are at the end of the slice then you should set it to 
	//         -+ inf as appropriate
			if pos_sm == 0  {
				max_left_sm = math.MinInt64
			} else {
				max_left_sm = sm[pos_sm - 1]
			}
			
			if pos_lg == 0 {
				max_left_lg = math.MinInt64
			} else {
				max_left_lg = lg[pos_lg - 1]
			}
				
			if pos_sm == len(sm) {
				min_right_sm = math.MaxInt64
			} else {
				min_right_sm = sm[pos_sm]
			}
			
			if pos_lg == len(lg) {
				min_right_lg = math.MaxInt64
			} else {
				min_right_lg = lg[pos_lg]
			}
			
	//         if the condition is met then you can break out
	//         else if the condition is not met move to partition 
	//         in the right direction
			
			if (max_left_sm <= min_right_lg) && (max_left_lg <= min_right_sm) {
				break
			} else if max_left_sm > min_right_lg {
				pos_sm--
				
			} else {
				pos_sm++
			}
			
			pos_lg = int((x+y+1) / 2) - pos_sm
		}
		
		if (x + y) % 2 != 0 {
			return math.Max(float64(max_left_sm), float64(max_left_lg))
		} else {
			mx := math.Max(float64(max_left_sm), float64(max_left_lg))
			mn := math.Min(float64(min_right_sm), float64(min_right_lg))
			return (mx+mn)/2
		}
		
		return 0
		
	}