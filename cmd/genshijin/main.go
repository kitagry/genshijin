package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/kitagry/genshijin"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		genshiGo := genshijin.ToGenshijin(sc.Text())
		fmt.Println(genshiGo)
	}
}
