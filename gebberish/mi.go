package gebberish

import (
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
	"path/filepath"
	"strings"

	"github.ibm.com/hofstadter-io/geb/engine/utils"
)

func Mi(rule string, arg int) (string, error) {
	orig, err := get_current()
	if err != nil {
		return orig, errors.Wrap(err, "in MI(...)")
	}
	MI := orig
	L := len(MI)
	switch rule {
	case "c", "curr", "current", "s", "stat", "status", "get":
		return MI, nil

	case "reset", "give-up", "giveup", "start-over", "startover", ":[":
		MI = "MI"

	// mi-rule-1:    if mi-string ends in 'I',        you may add a 'U'
	case "1", "r1", "rule-1", "rule1":
		if string(MI[L-1]) == "I" {
			MI += "U"
		} else {
			return orig, errors.New("mi-rule-1 does not apply")
		}

	// mi-rule-2:    suppose mi-string = 'Mx',          then you may make it 'Mxx'
	case "2", "r2", "rule-2", "rule2":
		MI = "M" + MI[1:] + MI[1:]

	// mi-rule-3:    if mi-string contains an 'III',  you may replace it with 'U'
	case "3", "r3", "rule-3", "rule3":
		cnt := 0
		for i, _ := range MI {
			if MI[i:i+3] == "III" {
				cnt++
			}
		}
		if cnt > 0 {
			if arg >= cnt {
				return orig, errors.New("arg grater than count of " + fmt.Sprint(cnt))
			}
			if arg > -1 {
				idx := strings.LastIndex(MI, "III")
				MI = MI[:idx] + "U" + MI[idx+3:]

			} else {
				cnt := 0
				for idx, _ := range MI {
					if MI[idx:idx+3] == "III" {
						if cnt == arg {
							MI = MI[:idx] + "U" + MI[idx+3:]
						}
						cnt++
					}
				}

			}
		} else {
			return orig, errors.New("mi-rule-3 does not apply")
		}

	// mi-rule-4:    if mi-string contains a 'UU',    you may drop it (remove it)
	case "4", "r4", "rule-4", "rule4":
		cnt := 0
		for i, _ := range MI {
			if MI[i:i+2] == "UU" {
				cnt++
			}
		}
		if cnt > 0 {
			if arg >= cnt {
				return orig, errors.New("arg grater than count of " + fmt.Sprint(cnt))
			}
			if arg > -1 {
				idx := strings.LastIndex(MI, "UU")
				MI = MI[:idx] + MI[idx+2:]

			} else {
				cnt := 0
				for idx, _ := range MI {
					if MI[idx:idx+2] == "UU" {
						if cnt == arg {
							MI = MI[:idx] + MI[idx+2:]
						}
						cnt++
					}
				}

			}
		} else {
			return orig, errors.New("mi-rule-4 does not apply")
		}

	default:
		return orig, errors.New("unknown mi input")
	}

	err = set_current(MI)
	if err != nil {
		return orig, errors.Wrap(err, "in MI(...)")
	}

	return MI, nil
}

func get_current() (string, error) {
	MI := "MI"
	dir := "$HOME/.hofstadter/.gebberish"
	dir, err := utils.ResolvePath(dir)
	if err != nil {
		return MI, nil
	}
	fn := filepath.Join(dir, "mi.txt")

	err = utils.FileExists(fn)
	if err != nil {
		return MI, nil
	}
	bytes, err := ioutil.ReadFile(fn)
	if err != nil {
		return MI, errors.Wrap(err, "in get current")
	}

	return string(bytes), nil
}

func set_current(MI string) error {
	dir := "$HOME/.hofstadter/.gebberish"
	dir, err := utils.ResolvePath(dir)
	if err != nil {
		return nil
	}
	fn := filepath.Join(dir, "mi.txt")

	err = ioutil.WriteFile(fn, []byte(MI), 0644)
	if err != nil {
		return errors.Wrap(err, "in set current")
	}

	return nil
}
