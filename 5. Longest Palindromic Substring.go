// DP solution for the problem
// Note : manchester exists but its not obvious

func longestPalindrome(s string) string {
    
    res := ""
    
    table := make([][]bool, len(s)+1)
    
    for i := range table {
        table[i] = make([]bool, len(s)+1)
    }
    
    for i := 1; i < len(s)+1; i++ {
        for j := 0; j < len(s) - i + 1; j++ {
            
            if i == 1 {
                table[j][j+i] = true
            } else if i == 2 && s[j] == s[j+i-1]{
                table[j][j+i] = true
            } else if table[j+1][j+i-1] && s[j] == s[j+i-1] {
                table[j][j+i] = true
            }
            
            if i > len(res) && table[j][j+i] {
                res = s[j:j+i]
            }
        }
    }
    
    return res
}