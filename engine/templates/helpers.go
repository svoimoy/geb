package templates

import (
	"fmt"
	"os"
	"strings"

	"github.com/aymerick/raymond"
)

func (template *Template) Render(design interface{}) (string, error) {
	tpl := (*raymond.Template)(template)
	return tpl.Exec(design)
}

func RenderTemplate(template *Template, design interface{}) (string, error) {
	tpl := (*raymond.Template)(template)
	return tpl.Exec(design)
}

func AddHelpers(tpl *raymond.Template) {
	add_template_helpers(tpl)
}

func add_template_helpers(tpl *raymond.Template) {

	tpl.RegisterHelper("concat2", helper_concat2)
	tpl.RegisterHelper("concat3", helper_concat3)
	tpl.RegisterHelper("concat4", helper_concat4)
	tpl.RegisterHelper("concat5", helper_concat5)
	tpl.RegisterHelper("join2", helper_join2)
	tpl.RegisterHelper("join3", helper_join3)
	tpl.RegisterHelper("join4", helper_join4)
	tpl.RegisterHelper("join5", helper_join5)

	tpl.RegisterHelper("lwidth", helper_lwidth)
	tpl.RegisterHelper("rwidth", helper_rwidth)
	tpl.RegisterHelper("printf", helper_printf)
	tpl.RegisterHelper("lower", helper_lower)
	tpl.RegisterHelper("upper", helper_upper)
	tpl.RegisterHelper("title", helper_title)

	tpl.RegisterHelper("contains", helper_contains)
	tpl.RegisterHelper("split", helper_split)
	tpl.RegisterHelper("replace", helper_replace)
	tpl.RegisterHelper("hasprefix", helper_hasprefix)
	tpl.RegisterHelper("hassuffix", helper_hassuffix)
	tpl.RegisterHelper("trimprefix", helper_trimprefix)
	tpl.RegisterHelper("trimsuffix", helper_trimsuffix)
	tpl.RegisterHelper("substr", helper_substr)

	tpl.RegisterHelper("eq", helper_eq)
	tpl.RegisterHelper("ne", helper_ne)

	tpl.RegisterHelper("getenv", helper_getenv)
}

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
func helper_substr(str string, start, end int) string {
	if end == -1 {
		end = len(str)
	}
	return str[start:end]
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
