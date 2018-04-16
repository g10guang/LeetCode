package sort_test

import (
	"testing"
	"github.com/g10guang/leetcode/sort"
	"strings"
)

func TestFrequencySort(t *testing.T) {
	suits := []struct {
		input  string
		expect string
	}{
		{input: "tree", expect: "eetr"},
		{input: "2a554442f544asfasssffffasss", expect: "sssssssffffff44444aaaa55522"},
		{input: "his s he a ha he  ha ae", expect: "        hhhhhaaaaeeessi"},
		{input: "Mymommaalwayssaid,\"Lifewaslikeaboxofchocolates.Youneverknowwhatyou'regonnaget.", expect: "aaaaaaaaaoooooooooeeeeeeennnnsssswwwwiiillltttyyymmmrruukk..hhggffcc\"dbvYMLx,'"},
	}
	for _, s := range suits {
		if r := sort.FrequencySort(s.input); strings.Compare(s.expect, r) != 0 {
			t.Errorf("input=%v\n expect=%v\n return=%v\n", s.input, s.expect, r)
		}
	}
}
