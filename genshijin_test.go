package genshijin

import (
	"testing"
)

func TestToGenshijin(t *testing.T) {
	testSet := []struct {
		input  string
		output string
	}{
		{"俺はお前をテストする", "オレ オマエ テスト スル"},
		{"すもももももももものうち", "スモモ モモ モモ ウチ"},
	}

	for _, test := range testSet {
		genshiGo := ToGenshijin(test.input)
		if test.output != genshiGo {
			t.Fatalf("オレ '%s' ホシイ ケド オマエ '%s' カエシタ", test.output, genshiGo)
		}
	}
}
