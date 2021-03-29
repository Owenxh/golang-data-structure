// link: https://leetcode-cn.com/problems/reverse-bits
package main

const (
	m1 = 0x55555555 // 01010101010101010101010101010101
	m2 = 0x33333333 // 00110011001100110011001100110011
	m3 = 0x0f0f0f0f // 00001111000011110000111100001111
	m4 = 0x00ff00ff // 00000000111111110000000011111111
)

// return bits.Reverse32(n)
func reverseBits(n uint32) uint32 {
	n = (n>>1)&m1 | (n&m1)<<1
	n = (n>>2)&m2 | (n&m2)<<2
	n = (n>>4)&m3 | (n&m3)<<4
	n = (n>>8)&m4 | (n&m4)<<8
	return n>>16 | n<<16
}

func main() {

}
