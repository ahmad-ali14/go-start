package main

import (
	"fmt"
	"math"

	// importing local pckg by using the relaive path starting from 'src/github.com'
	"github.com/ahmad-ali14/go_starter/03_pkgs/str_util"
)

func main() {
	// math core pckg
	fmt.Println(math.Sqrt2)

	// own pckg
	fmt.Println(str_util.Reverse("Ahmad"))
}
