package segito

import (
	"database/sql"
	"fmt"
	"strings"
	"unicode/utf8"
)

func StringSliceToCommaDelimited(slice []string) string {
	len := len(slice)
	if len > 0 {
		if len == 1 {
			return slice[0]
		}
		return strings.Join(slice, ",")
	}

	return ""
}

func ConvertUnicodeStringToString(str string) string {
	var out = ""

	for len(str) > 0 {
		r, size := utf8.DecodeRuneInString(str)
		out = fmt.Sprintf("%s%c", out, r)
		str = str[size:]

	}
	return out
}

func NullFloatToFloat64(f float64) sql.NullFloat64 {
	result := sql.NullFloat64{}

	if f != 0 {
		result.Float64 = f
		result.Valid = true
	}

	return result
}

// float64 to string
func Float64ToString(f float64) string {
	return fmt.Sprintf("%f", f)
}
