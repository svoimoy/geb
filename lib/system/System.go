package system

import (
	// HOFSTADTER_START import
	"fmt"

	"github.com/sergi/go-diff/diffmatchpatch"
	"gopkg.in/src-d/go-git.v4"
	// HOFSTADTER_END   import
)

// HOFSTADTER_START const
// HOFSTADTER_END   const

// HOFSTADTER_START var
// HOFSTADTER_END   var

// HOFSTADTER_START init
// HOFSTADTER_END   init

// HOFSTADTER_BELOW

const (
	text1 = "Lorem ipsum dolor."
	text2 = "Lorem dolor sit amet."
)

func keepImportsAndDeps() {
	r, err := git.PlainClone("directory", false, &git.CloneOptions{
		URL: "https://github.com/hofstadter-io/dsl-library",
	})

	fmt.Println("go-git", r, err)

	dmp := diffmatchpatch.New()

	diffs := dmp.DiffMain(text1, text2, false)

	fmt.Println(dmp.DiffPrettyText(diffs))

}
