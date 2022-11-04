package match

import "testing"

func TestRabinKarp(t *testing.T) {
	src := "owen欧文雪儿"
	for i := 0; i < 26; i++ {
		v := rune('a' + i)
		target := string([]rune{v})
		t.Logf("finding [%v] in src [%v], index is %v", src, target, RabinKarp(src, target))
	}

	target := "欧文雪"
	t.Logf("finding [%v] in src [%v], index is %v", src, target, RabinKarp(src, target))
	target = "ow"
	t.Logf("finding [%v] in src [%v], index is %v", src, target, RabinKarp(src, target))
}
