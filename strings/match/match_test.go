package match

import (
	"testing"
)

//	func TestR(t *testing.T) {
//		// [1.23 - 4.78)
//		for i := 0; i < 100; i++ {
//			fmt.Println(rand.Float64()*(4.78-1.23) + 1.23)
//		}
//	}
func testStringMatch(t *testing.T, fn func(string, string) int) {
	src := "owen欧文雪儿❤"

	for i := 0; i < len([]rune(src)); i++ {
		r := []rune(src)[i]
		target := string([]rune{r})
		t.Logf("finding [%v] in src [%v], index is %v", src, target, fn(src, target))
	}

	target := "欧文雪"
	t.Logf("finding [%v] in src [%v], index is %v", src, target, fn(src, target))

	target = "ow"
	t.Logf("finding [%v] in src [%v], index is %v", src, target, fn(src, target))
}

func TestMatch(t *testing.T) {
	testStringMatch(t, BruteForce)
	testStringMatch(t, RabinKarp)
	testStringMatch(t, KMP)
}
