package match

import (
	"fmt"
	"strings"
	"testing"
	"time"

	"io.vava.datastructure/util"
)

//	func TestRandomFloat(t *testing.T) {
//		// [1.23 - 4.78)
//		for i := 0; i < 100; i++ {
//			fmt.Println(rand.Float64()*(4.78-1.23) + 1.23)
//		}
//	}

func TestMatch(t *testing.T) {
	src := "owen欧文雪儿❤"
	targes := []string{
		"o", "w", "e", "n", "文", "❤", "欧文雪", "Owen",
	}
	fns := []func(string, string) int{BruteForce, RabinKarp, KMP, KMP2, KMP3, strings.Index}

	for _, target := range targes {
		for _, fn := range fns {
			index := fn(src, target)
			if index >= 0 {
				fmt.Printf("found [%v] at index [%v] use [%10v]\n", target, index, util.ResolveFuncName(fn))
			} else {
				fmt.Printf("can't find [%v]\n", target)
			}
		}
	}
}

func testMatchPerformance(src, target *string, fn func(string, string) int) {
	start := time.Now()
	index := fn(*src, *target)
	fmt.Printf("found target string at index [%v] use [%10v] cost [%v]\n", index, util.ResolveFuncName(fn), time.Now().Sub(start))
}

func TestMatchPerformance1(t *testing.T) {
	src := util.GetPrideAndPrejudiceAsString()
	target := "china"
	testMatchPerformance(&src, &target, BruteForce)
	testMatchPerformance(&src, &target, RabinKarp)
	testMatchPerformance(&src, &target, KMP)
	testMatchPerformance(&src, &target, KMP3)

	// 处理中文返回的不是字符在字符串中的下标
	testMatchPerformance(&src, &target, KMP2)

	// 处理中文返回的不是字符在字符串中的下标
	testMatchPerformance(&src, &target, strings.Index)
}

func TestMatchPerformance2(t *testing.T) {
	n, m := 1_000_000, 10_000
	var sBuilder strings.Builder
	for i := 0; i < n; i++ {
		sBuilder.WriteRune('a')
	}

	var tBuilder strings.Builder
	for i := 0; i < m-1; i++ {
		tBuilder.WriteRune('a')
	}
	tBuilder.WriteRune('b')

	src := sBuilder.String()
	target := tBuilder.String()

	testMatchPerformance(&src, &target, BruteForce)
	testMatchPerformance(&src, &target, RabinKarp)
	testMatchPerformance(&src, &target, KMP)
	testMatchPerformance(&src, &target, KMP2)
	testMatchPerformance(&src, &target, KMP3)
	testMatchPerformance(&src, &target, strings.Index)
}
