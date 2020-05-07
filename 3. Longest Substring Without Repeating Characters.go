func lengthOfLongestSubstring(s string) int {
    
    if len(s) <= 1 {
        return len(s)
    }
    
    hash := [128]int{}
    
    max := 0
    cur := 1
    
    front := 0
    end := 1
    
    hash[s[front]]++
    
    for end != len(s) {

        if hash[s[end]] > 0 {
            hash[s[front]]--
            front++           
            cur--
            
        } else {
                    
            hash[s[end]]++
            cur++
            end++
            if cur > max {
                max = cur
            }
        }
    }
       
    return max
}