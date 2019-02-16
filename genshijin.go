package genshijin

import (
	"strings"

	"github.com/ikawaha/kagome/tokenizer"
)

var (
	t tokenizer.Tokenizer = tokenizer.New()
)

// ToGenshijin s ニューリョク ゲンシ ジンゴ カエス
func ToGenshijin(s string) (res string) {
	tokens := t.Tokenize(s)

	for _, token := range tokens {
		features := token.Features()

		if len(features) == 0 || features[0] == "助詞" {
			continue
		}

		if features[len(features)-1] == "*" && features[0] == "名詞" {
			res += string([]rune(s)[token.Start:token.End]) + " "
		}

		if features[len(features)-1] != "*" {
			res += features[len(features)-1] + " "
		}
	}

	res = strings.TrimSpace(res)
	return
}
