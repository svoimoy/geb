package utils
// package publicFiles

import (
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
func Merge(original interface{},update interface{}) (merged interface{},err error) {
	// HOFSTADTER_START Merge
	if original == nil {
		return update, nil
	}

	if update == nil {
		return original, nil
	}

	// call the recursive merge
	return merge(original, update)
	// HOFSTADTER_END   Merge
	return
}



// HOFSTADTER_BELOW
