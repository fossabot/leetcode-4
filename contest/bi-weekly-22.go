package leetcode

type Seats [][]int

func maxNumberOfFamilies(n int, reservedSeats [][]int) int {
	// seatTaken := make([][]bool, n)

	m := Seats(reservedSeats)
	sort.Sort(&m)

	// for i := 0; i < n; i++ {
	// 	seatTaken[i] = make([]bool, 10)
	// }

	// for _, seat := range reservedSeats {
	// 	seatTaken[seat[0]-1][seat[1]-1] = true
	// }
	res := 0
	for i := 0; i < n; i++ {
		test2345 := testIndex(seatTaken[i], 2, 3, 4, 5)
		test6789 := testIndex(seatTaken[i], 6, 7, 8, 9)
		if test2345 || test6789 {
			if test2345 {
				res++
			}
			if test6789 {
				res++
			}
		} else {
			test4567 := testIndex(seatTaken[i], 4, 5, 6, 7)
			if test4567 {
				res++
			}
		}
	}
	return res
}

func testIndex(reservedSeats []bool, i, j, k, l int) bool {
	if !reservedSeats[i-1] && !reservedSeats[j-1] && !reservedSeats[k-1] && !reservedSeats[l-1] {
		return true
	}
	return false
}

func (m Seats) Len() int { return len(m) }
func (m Seats) Less(i, j int) bool {
	if m[i][0] < m[j][0] {
		return true
	}
	return m[i][1] < m[j][1]
}

func (m *Seats) Swap(i, j int) { m[i], m[j] = m[j], m[i] }
