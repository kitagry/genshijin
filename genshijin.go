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
		if len(features) > 0 && features[0] != "助詞" {
			res += features[len(features)-1] + " "
		}
	}
	res = strings.TrimSpace(res)
	return
}
