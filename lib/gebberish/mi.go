package gebberish

import (
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.ibm.com/hofstadter-io/geb/engine/utils"
)

func Mi(rule string, arg int) (string, error) {
	orig, err := get_current()
	if err != nil {
		return orig, errors.Wrap(err, "in MI(...)")
	}
	MI := orig
	L := len(MI)

	fmt.Println(MI, rule, arg)
	switch rule {

	case "init", "setup":
		dir := "$HOME/.geb/.gebberish"
		dir, err := utils.ResolvePath(dir)
		if err != nil {
			return "", err
		}

		err = os.MkdirAll(dir, 0755)
		if err != nil {
			return "", err
		}

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
		indicies := map[int]int{}
		cnt := 0
		for i := 0; i < len(MI)-2; i++ {
			if MI[i:i+3] == "III" {
				indicies[cnt] = i
				cnt++
			}
		}
		if cnt > 0 {
			pos, ok := indicies[arg]
			if ok {
				MI = MI[:pos] + "U" + MI[pos+3:]
			} else {
				return orig, errors.New("mi-rule-3 does not work for that pos")
			}
		} else {
			return orig, errors.New("mi-rule-3 does not apply")
		}

	// mi-rule-4:    if mi-string contains a 'UU',    you may drop it (remove it)
	case "4", "r4", "rule-4", "rule4":
		indicies := map[int]int{}
		cnt := 0
		for i := 0; i < len(MI)-1; i++ {
			if MI[i:i+2] == "UU" {
				indicies[cnt] = i
				cnt++
			}
		}
		if cnt > 0 {
			pos, ok := indicies[arg]
			if ok {
				MI = MI[:pos] + MI[pos+2:]
			} else {
				return orig, errors.New("mi-rule-4 does not work for that pos")
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
	dir := "$HOME/.geb/.gebberish"
	dir = os.ExpandEnv(dir)

	fn := filepath.Join(dir, "mi.txt")
	err := utils.FileExists(fn)
	if err != nil {
		set_current(MI)
		return MI, nil
	}

	bytes, err := ioutil.ReadFile(fn)
	if err != nil {
		return MI, errors.Wrap(err, "in get current")
	}

	return string(bytes), nil
}

func set_current(MI string) error {
	fmt.Println("setting MI:", MI)

	dir := "$HOME/.geb/.gebberish"
	dir = os.ExpandEnv(dir)

	err := os.MkdirAll(dir, 0755)
	if err != nil {
		return err
	}

	fn := filepath.Join(dir, "mi.txt")
	err = ioutil.WriteFile(fn, []byte(MI), 0644)
	if err != nil {
		return errors.Wrap(err, "in set current")
	}

	return nil
}
