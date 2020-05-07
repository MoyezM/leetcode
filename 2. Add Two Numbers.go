func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    var ret = new(ListNode)
    t := ret
    
    var carry int
    
    for l1 != nil || l2 != nil {
        
        var x int
        var y int
        
        
        if l1 != nil {
            x = l1.Val
            l1 = l1.Next
        } 
        
        if l2 != nil {
            y = l2.Val
            l2 = l2.Next
        }
        
        sum := carry + x + y
        
        carry = (sum) / 10
        
        t.Next = new(ListNode)
        t.Next.Val = (sum) % 10
        t = t.Next
        
        if carry > 0 {
            t.Next = new(ListNode)
            t.Next.Val = carry
        }
        
    }
    
    return ret.Next
    
}