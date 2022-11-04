package match

import (
	"testing"
)

func TestBruteForce(t *testing.T) {
	src := "owen欧文雪儿"
	for i := 0; i < 26; i++ {
		v := rune('a' + i)
		target := string([]rune{v})
		t.Logf("finding [%v] in src [%v], index is %v", src, target, BruteForce(src, target))
	}

	target := "欧文雪"
	t.Logf("finding [%v] in src [%v], index is %v", src, target, BruteForce(src, target))
}
