package add

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	git "gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
	// HOFSTADTER_START import
	// HOFSTADTER_END   import
)

// HOFSTADTER_START const
// HOFSTADTER_END   const

// HOFSTADTER_START var
// HOFSTADTER_END   var

// HOFSTADTER_START init
// HOFSTADTER_END   init

/*
Where's your docs doc?!
*/
func AddGitRepo(opts Options) (err error) {
	// HOFSTADTER_START AddGitRepo

	// Check the location
	if opts.Location == "vendor" {
		opts.Location = filepath.Join("vendor", opts.Url[strings.LastIndex(opts.Url, "/")+1:])
	}

	fmt.Println("Writing to:", opts.Location)

	// Get the repository
	repo, err := git.PlainClone(opts.Location, false, &git.CloneOptions{
		URL: opts.Url,
	})
	if err != nil {
		return
	}

	// Get the working tree
	tree, err := repo.Worktree()
	if err != nil {
		return
	}

	checkoutOpts := &git.CheckoutOptions{
		Force: false,
	}
	// Checkout the right thing
	if opts.Commit != "" {
		checkoutOpts.Hash = plumbing.NewHash(opts.Commit)
	} else if opts.Tag != "" {
		checkoutOpts.Branch = plumbing.ReferenceName(opts.Tag)
	} else {
		checkoutOpts.Branch = plumbing.ReferenceName("refs/heads/" + opts.Branch)
	}

	err = tree.Checkout(checkoutOpts)
	if err != nil {
		return
	}

	subs, err := tree.Submodules()
	if err != nil {
		return
	}

	err = subs.Init()
	if err != nil {
		return
	}

	err = os.RemoveAll(filepath.Join(opts.Location, ".git"))
	if err != nil {
		return
	}

	// HOFSTADTER_END   AddGitRepo
	return
}
