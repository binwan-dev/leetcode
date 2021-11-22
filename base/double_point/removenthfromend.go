func removeNthFromEnd(head *ListNode, n int) *ListNode {
	s,f:=&ListNode{Val:0,Next:head},&ListNode{Val:0,Next:head}

	for i:=0;i<n;i++{
		f=f.Next
	}

	for{
		if f.Next==nil{
            if s.Next==head{
                return head.Next
            }
			s.Next=s.Next.Next
            return head
		}
		s=s.Next
		f=f.Next
	}
	return head
}
