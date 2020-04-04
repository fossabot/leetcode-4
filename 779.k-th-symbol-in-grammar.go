/*
 * @lc app=leetcode id=779 lang=golang
 *
 * [779] K-th Symbol in Grammar
 *
 * https://leetcode.com/problems/k-th-symbol-in-grammar/description/
 *
 * algorithms
 * Medium (36.97%)
 * Likes:    346
 * Dislikes: 121
 * Total Accepted:    25.7K
 * Total Submissions: 69.6K
 * Testcase Example:  '1\n1'
 *
 * On the first row, we write a 0. Now in every subsequent row, we look at the
 * previous row and replace each occurrence of 0 with 01, and each occurrence
 * of 1 with 10.
 *
 * Given row N and index K, return the K-th indexed symbol in row N. (The
 * values of K are 1-indexed.) (1 indexed).
 *
 *
 * Examples:
 * Input: N = 1, K = 1
 * Output: 0
 *
 * Input: N = 2, K = 1
 * Output: 0
 *
 * Input: N = 2, K = 2
 * Output: 1
 *
 * Input: N = 4, K = 5
 * Output: 1
 *
 * Explanation:
 * row 1: 0
 * row 2: 01
 * row 3: 0110
 * row 4: 01101001
 *
 *
 * Note:
 *
 *
 * N will be an integer in the range [1, 30].
 * K will be an integer in the range [1, 2^(N-1)].
 *
 *
 */

package leetcode

// @lc code=start
func kthGrammar(N int, K int) int {
	// bArr is the byte array representing Nth row
	bArr := nthRow(N)

	// bI is the index in byte array
	bI := (K - 1) / 8

	b := bArr[bI]

	// bB is the bit index in byte
	bB := (K - 1) % 8

	return int((b >> bB) & 0x01)
}

func nthRow(N int) []byte {
	if N == 1 {
		return []byte{0}
	} else if N == 2 {
		// binary 0b01000000
		return []byte{0x40}
	} else if N == 3 {
		// binary 0b01100000
		return []byte{0x60}
	} else if N == 4 {
		// binary 0b01101001
		return []byte{0x69}
	}
	bArr := nthRow(N - 1)
	lenArr := len(bArr)
	newBArr := make([]byte, lenArr*2)

	for i, b := range bArr {
		newBArr[2*i] = nextRow(b)
		b = (b & 0x0F) << 4
		newBArr[2*i+1] = nextRow(b)
	}
	return newBArr
}

func nextRow(b byte) byte {
	var res byte
	for i := 0; i < 4; i++ {
		if b&(0x80>>i) == 0 {
			res |= (0x40 >> (2 * i))
		} else {
			res |= (0x80 >> (2 * i))
		}
	}
	return res
}

// @lc code=end
