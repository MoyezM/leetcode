func maxArea(height []int) int {
    area := -1
    start := 0
    end := len(height) -1 
    
    for start != end {
        a := min(height[start], height[end]) * (end - start)
        
        if a > area {
            area = a
        }
        
        if (height[start] < height[end]) {
            start++
        } else {
            end--
        }
    }
    return area
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}