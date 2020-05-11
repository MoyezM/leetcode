func isMatch(s string, p string) bool {
    
    table := make([][]bool, len(p)+1)
    
    for i := range table {
        table[i] = make([]bool, len(s)+1)
    }
    
    table[0][0] = true
    if len(s) > 0 && len(p) > 0 {
        table[1][1] = s[0] == p[0] || p[0] == '.'
    }
    
//  compensate for the case of zero for *
//  this is essentially dropping the truth value from above
    for i := 0; i <= len(p)-2; i++ {
        table[i+2][0] = table[i][0] && p[i+1] == '*'
    }
    
    
    
    for i := 1; i < len(p) +1; i++ {
        for j := 1; j  < len(s) +1; j++{
            if p[i-1] != '*' {
//                 same charecter, '.' wildcard
                if table[i-1][j-1] && (s[j-1] == p[i-1] || p[i-1] == '.') {
                    table[i][j] = true
                } 
            } else {
//              the case for zero of the preceding element, one or more, match all
                if table[i-2][j] || (table[i][j-1] && (s[j-1]==p[i-2] || p[i-2] == '.')) {
                    table[i][j] = true
                }
            }
        }
    }

    return table[len(p)][len(s)]
}