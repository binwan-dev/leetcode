func middleNode(head *ListNode) *ListNode{
	lens:=getListLen(head)
	idx:=0
	startIdx:=lens/2;

	for{
		if head==nil || head.Next==nil|| idx==startIdx{
			return head
		}
        idx+=1
		head=head.Next
	}
}

func getListLen(head *ListNode) int{
	lens:=0
	for{
		if head == nil{
			return lens
		}

		lens+=1
		head=head.Next
	}
}
