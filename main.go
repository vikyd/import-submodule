package main

import (
	"fmt"

	"github.com/vikyd/submodule01"
	"github.com/vikyd/submodule01/module01"
	module01v2 "github.com/vikyd/submodule01/module01/v2"
	submodule01v2 "github.com/vikyd/submodule01/v2"
)

func main() {
	fmt.Println("vim-go")
	submodule01.F01()
	submodule01v2.F01()
	module01.F02()
	module01v2.F02()
}
