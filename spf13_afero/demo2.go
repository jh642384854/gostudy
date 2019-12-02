package main

import (
	"fmt"
	"github.com/spf13/afero"
)

func main() {
 	afeFs := afero.NewOsFs()
 	fmt.Println(afero.DirExists(afeFs,"upload33"))

}
