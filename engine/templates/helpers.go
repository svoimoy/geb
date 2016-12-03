package templates

import (
	"fmt"
	"strings"

	"github.com/aymerick/raymond"
)

func AddHelpers(tpl *raymond.Template) {
	add_template_helpers(tpl)
}

func add_template_helpers(tpl *raymond.Template) {

	tpl.RegisterHelper("concat", helper_concat)
	tpl.RegisterHelper("lwidth", helper_lwidth)
	tpl.RegisterHelper("rwidth", helper_rwidth)
	tpl.RegisterHelper("printf", helper_printf)
	tpl.RegisterHelper("lower", helper_lower)
	tpl.RegisterHelper("upper", helper_upper)
	tpl.RegisterHelper("title", helper_title)

	tpl.RegisterHelper("eq", helper_eq)

}

func helper_concat(prefix, value, suffix string, options *raymond.Options) string {
	return prefix + value + suffix
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

func helper_eq(lhs, rhs string) string {
	if lhs == rhs {
		return lhs
	}
	return ""
}
