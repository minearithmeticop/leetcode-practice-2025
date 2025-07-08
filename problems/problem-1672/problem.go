package problem1672

// maximumWealth รับ accounts เป็นเมทริกซ์ของจำนวนเงินแต่ละลูกค้าในแต่ละธนาคาร
// คืนค่าความมั่งคั่งสูงสุดของลูกค้าทั้งหมด
func MaximumWealth(accounts [][]int) int {
	maxWealth := 0
	for _, customer := range accounts {
		sum := 0
		for _, money := range customer {
			sum += money
		}
		if sum > maxWealth {
			maxWealth = sum
		}
	}
	return maxWealth
}
