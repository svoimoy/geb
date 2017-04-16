package templates

// package privateFiles

import (
	// HOFSTADTER_START import
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"strings"

	"github.com/aymerick/raymond"
	"github.com/clbanning/mxj"
	"github.com/codemodus/kace"
	"github.com/ghodss/yaml"
	"github.com/kr/pretty"
	"github.com/naoina/toml"
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
func addTemplateHelpers(tpl *raymond.Template) {
	// HOFSTADTER_START addTemplateHelpers
	tpl.RegisterHelper("concat2", helper_concat2)
	tpl.RegisterHelper("concat3", helper_concat3)
	tpl.RegisterHelper("concat4", helper_concat4)
	tpl.RegisterHelper("concat5", helper_concat5)
	tpl.RegisterHelper("join2", helper_join2)
	tpl.RegisterHelper("join3", helper_join3)
	tpl.RegisterHelper("join4", helper_join4)
	tpl.RegisterHelper("join5", helper_join5)

	tpl.RegisterHelper("yaml", helper_yaml)
	tpl.RegisterHelper("toml", helper_toml)
	tpl.RegisterHelper("json", helper_json)
	tpl.RegisterHelper("xml", helper_xml)
	tpl.RegisterHelper("indent", helper_indent)
	tpl.RegisterHelper("pprint", helper_pretty)
	tpl.RegisterHelper("pretty", helper_pretty)
	tpl.RegisterHelper("lwidth", helper_lwidth)
	tpl.RegisterHelper("rwidth", helper_rwidth)
	tpl.RegisterHelper("printf", helper_printf)
	tpl.RegisterHelper("lower", helper_lower)
	tpl.RegisterHelper("upper", helper_upper)
	tpl.RegisterHelper("title", helper_title)

	tpl.RegisterHelper("camel", helper_camel)
	tpl.RegisterHelper("camelT", helper_camelT)
	tpl.RegisterHelper("snake", helper_snake)
	tpl.RegisterHelper("snakeU", helper_snakeU)
	tpl.RegisterHelper("kebab", helper_kebab)
	tpl.RegisterHelper("kebabU", helper_kebabU)

	tpl.RegisterHelper("contains", helper_contains)
	tpl.RegisterHelper("split", helper_split)
	tpl.RegisterHelper("replace", helper_replace)
	tpl.RegisterHelper("hasprefix", helper_hasprefix)
	tpl.RegisterHelper("hassuffix", helper_hassuffix)
	tpl.RegisterHelper("trimprefix", helper_trimprefix)
	tpl.RegisterHelper("trimsuffix", helper_trimsuffix)
	tpl.RegisterHelper("trimto", helper_trimto_first)
	tpl.RegisterHelper("trimfrom", helper_trimfrom_first)
	tpl.RegisterHelper("trimto_first", helper_trimto_first)
	tpl.RegisterHelper("trimfrom_first", helper_trimfrom_first)
	tpl.RegisterHelper("trimto_last", helper_trimto_last)
	tpl.RegisterHelper("trimfrom_last", helper_trimfrom_last)
	tpl.RegisterHelper("substr", helper_substr)
	tpl.RegisterHelper("getprefix", helper_getprefix)
	tpl.RegisterHelper("getsuffix", helper_getsuffix)
	tpl.RegisterHelper("getbetween", helper_getbetween)

	tpl.RegisterHelper("builtin", helper_builtin)
	tpl.RegisterHelper("ternary", helper_ternary)

	tpl.RegisterHelper("length", helper_length)
	tpl.RegisterHelper("identity", helper_identity)
	tpl.RegisterHelper("thelist", helper_thelist)
	tpl.RegisterHelper("sublist", helper_sublist)
	tpl.RegisterHelper("rsublist", helper_rsublist)
	tpl.RegisterHelper("reverse", helper_reverse)

	tpl.RegisterHelper("eq", helper_eq)
	tpl.RegisterHelper("ne", helper_ne)

	tpl.RegisterHelper("getenv", helper_getenv)

	// HOFSTADTER_END   addTemplateHelpers
	return
}

// HOFSTADTER_BELOW

func helper_concat2(s1, s2 string) string {
	return s1 + s2
}
func helper_concat3(s1, s2, s3 string) string {
	return s1 + s2 + s3
}
func helper_concat4(s1, s2, s3, s4 string) string {
	return s1 + s2 + s3 + s4
}
func helper_concat5(s1, s2, s3, s4, s5 string) string {
	return s1 + s2 + s3 + s4 + s5
}

func helper_join2(sep, s1, s2 string) string {
	return strings.Join([]string{s1, s2}, sep)
}
func helper_join3(sep, s1, s2, s3 string) string {
	return strings.Join([]string{s1, s2, s3}, sep)
}
func helper_join4(sep, s1, s2, s3, s4 string) string {
	return strings.Join([]string{s1, s2, s3, s4}, sep)
}
func helper_join5(sep, s1, s2, s3, s4, s5 string) string {
	return strings.Join([]string{s1, s2, s3, s4, s5}, sep)
}

func helper_pretty(value interface{}) string {
	return fmt.Sprintf("%# v", pretty.Formatter(value))
}

func helper_yaml(value interface{}) string {
	bytes, err := yaml.Marshal(value)
	if err != nil {
		return err.Error()
	}
	return string(bytes)
}

func helper_toml(value interface{}) string {
	bytes, err := toml.Marshal(value)
	if err != nil {
		return err.Error()
	}
	return string(bytes)
}

func helper_indent(value, indent string) string {
	ret := ""
	lines := strings.Split(value, "\n")
	for _, line := range lines {
		ret += indent + line + "\n"
	}
	return ret
}

func helper_json(value interface{}) string {
	bytes, err := json.MarshalIndent(value, "", "  ")
	if err != nil {
		return err.Error()
	}
	return string(bytes)
}

func helper_xml(value interface{}) string {
	mv := mxj.Map(value.(map[string]interface{}))
	bytes, err := mv.XmlIndent("", "  ")
	if err != nil {
		return err.Error()
	}
	return string(bytes)
}

func helper_lwidth(width string, value string) string {
	fmt_str := "%-" + width + "s"
	return fmt.Sprintf(fmt_str, value)
}

func helper_rwidth(width string, value string) string {
	fmt_str := "%-" + width + "s"
	return fmt.Sprintf(fmt_str, value)
}

func helper_printf(fmt_str string, args ...interface{}) string {
	return fmt.Sprintf(fmt_str, args...)
}

func helper_lower(value string) string {
	return strings.ToLower(value)
}

func helper_upper(value string) string {
	return strings.ToUpper(value)
}

func helper_title(value string) string {
	return strings.ToTitle(value)
}

func helper_camel(value string) string {
	return kace.Camel(value, false)
}

func helper_camelT(value string) string {
	return fmt.Sprint(kace.Camel(value, true))
}

func helper_snake(value string) string {
	return kace.Snake(value)
}

func helper_snakeU(value string) string {
	return kace.SnakeUpper(value)
}

func helper_kebab(value string) string {
	return kace.Kebab(value)
}

func helper_kebabU(value string) string {
	return kace.KebabUpper(value)
}

func helper_contains(str, srch string) string {
	if strings.Contains(str, srch) {
		return "true"
	}
	return ""
}

func helper_split(str, sep string) []string {
	return strings.Split(str, sep)
}

func helper_replace(str, old, new string, cnt int) string {
	return strings.Replace(str, old, new, cnt)
}
func helper_hasprefix(str, pre string) string {
	if strings.HasPrefix(str, pre) {
		return "true"
	}
	return ""
}
func helper_hassuffix(str, suf string) string {
	if strings.HasSuffix(str, suf) {
		return "true"
	}
	return ""
}
func helper_trimprefix(str, pre string) string {
	return strings.TrimPrefix(str, pre)
}
func helper_trimsuffix(str, suf string) string {
	return strings.TrimSuffix(str, suf)
}

func helper_trimto_first(str, pre string, keep bool) string {
	pos := strings.Index(str, pre)
	if pos >= 0 {
		if keep {
			return str[pos:]
		}
		return str[pos+len(pre):]
	}
	return str
}

func helper_trimfrom_first(str, pre string, keep bool) string {
	pos := strings.Index(str, pre)
	if pos >= 0 {
		if keep {
			return str[:pos+len(pre)]
		}
		return str[:pos]
	}
	return str
}

func helper_trimto_last(str, pre string, keep bool) string {
	pos := strings.LastIndex(str, pre)
	if pos >= 0 {
		if keep {
			return str[pos:]
		}
		return str[pos+len(pre):]
	}
	return str
}

func helper_trimfrom_last(str, pre string, keep bool) string {
	pos := strings.LastIndex(str, pre)
	if pos >= 0 {
		if keep {
			return str[:pos+len(pre)]
		}
		return str[:pos]
	}
	return str
}

func helper_substr(str string, start, end int) string {
	if end == -1 {
		end = len(str)
	}
	return str[start:end]
}

func helper_getprefix(str, suf string) string {
	pos := strings.Index(str, suf)
	if pos >= 0 {
		return str[:pos]
	}
	return str
}

func helper_getsuffix(str, suf string) string {
	pos := strings.Index(str, suf)
	if pos >= 0 {
		return str[pos+1:]
	}
	return str
}

func helper_getbetween(str, lhs, rhs string) string {
	lpos := strings.Index(str, lhs)
	rpos := strings.LastIndex(str, rhs)
	if lpos < 0 {
		lpos = 0
	} else {
		lpos += 1
	}
	if rpos < 0 {
		rpos = len(str)
	}
	return str[lpos:rpos]
}

var known_builtins = map[string]struct{}{
	"bool":        struct{}{},
	"byte":        struct{}{},
	"error":       struct{}{},
	"float":       struct{}{},
	"float32":     struct{}{},
	"float64":     struct{}{},
	"complex64":   struct{}{},
	"complex128":  struct{}{},
	"int":         struct{}{},
	"int8":        struct{}{},
	"int16":       struct{}{},
	"int32":       struct{}{},
	"int64":       struct{}{},
	"uint":        struct{}{},
	"uint8":       struct{}{},
	"uint16":      struct{}{},
	"uint32":      struct{}{},
	"uint64":      struct{}{},
	"rune":        struct{}{},
	"string":      struct{}{},
	"object":      struct{}{},
	"interface{}": struct{}{},
}

func helper_builtin(str string) string {
	_, ok := known_builtins[str]
	if ok {
		return "true"
	}
	return ""
}

func helper_ternary(first, second string) string {
	if first != "" {
		return first
	}
	return second
}

func helper_length(list interface{}) interface{} {
	val := reflect.ValueOf(list)
	if val.Kind() == reflect.Slice || val.Kind() == reflect.Array {
		return val.Len()
	}
	return "not an array"
}

func helper_identity(thing interface{}) interface{} {
	return thing
}

func helper_thelist(thing interface{}) interface{} {
	val := reflect.ValueOf(thing)
	if val.Kind() == reflect.Slice || val.Kind() == reflect.Array {
		return "IS an array!"
	}
	return "not an array"
}

func helper_sublist(list interface{}, start, count int) interface{} {
	val := reflect.ValueOf(list)
	if val.Kind() == reflect.Slice || val.Kind() == reflect.Array {
		return val.Slice(start, start+count).Interface()
	}
	return "not an array"
}

func helper_rsublist(list interface{}, start, count int) interface{} {

	val := reflect.ValueOf(list)
	if val.Kind() == reflect.Slice || val.Kind() == reflect.Array {
		L := val.Len()
		last := L - start
		first := L - start - count
		return val.Slice(first, last).Interface()
	}
	return "not an array"
}

func helper_reverse(list interface{}) interface{} {
	val := reflect.ValueOf(list)
	if val.Kind() == reflect.Slice || val.Kind() == reflect.Array {
		L := val.Len()
		rev := make([]interface{}, 0, L)
		for i := 0; i < L; i++ {
			elem := val.Index(L - 1 - i)
			rev = append(rev, elem)
		}
		return rev
	}
	return "not an array"
}

func helper_eq(lhs, rhs string) string {
	if lhs == rhs {
		return lhs
	}
	return ""
}

func helper_ne(lhs, rhs string) string {
	if lhs != rhs {
		return lhs
	}
	return ""
}

func helper_getenv(env_var string) string {
	ret := os.Getenv(env_var)
	if ret == "" {
		return fmt.Sprintf("ENV_VAR: %q is empty", env_var)
	}
	return ret
}
