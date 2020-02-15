package main

import (
	"database/sql"
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	flt1      = 87.3
	flt2      = 1002.00003
	flt3      = 0.0006
	uniString = "café"
)

func TestStringSliceToCommaDelimitedFilled(t *testing.T) {
	assert := assert.New(t)

	slice := make([]string, 4)
	slice[0] = "beavis"
	slice[1] = "butt-head"
	slice[2] = "ren"
	slice[3] = "stimpy"

	result := StringSliceToCommaDelimited(slice)

	assert.NotNil(result)
	assert.Equal(result, "beavis,butt-head,ren,stimpy", "The strings should match")
}

func TestStringSliceToCommaDelimitedEmpty(t *testing.T) {
	assert := assert.New(t)

	slice := make([]string, 0)

	result := StringSliceToCommaDelimited(slice)

	assert.NotNil(result)
	assert.Equal(result, "")
}

func TestStringSliceToCommaDelimitedSingle(t *testing.T) {
	assert := assert.New(t)

	slice := make([]string, 1)
	slice[0] = "z"

	result := StringSliceToCommaDelimited(slice)

	assert.NotNil(result)
	assert.Equal(result, "z")
}

func TestConvertUnicodeStringToString(t *testing.T) {
	assert := assert.New(t)

	result := ConvertUnicodeStringToString("caf\xc3\xa9")

	assert.NotNil(result)
	assert.Equal(uniString, result)

}

func TestNullFloatToFloat64(t *testing.T) {
	assert := assert.New(t)

	nullFloat1 := NullFloatToFloat64(flt1)
	assert.NotNil(nullFloat1)
	assert.IsType(sql.NullFloat64{}, nullFloat1, "The result should be sql.NullFloat64")
	assert.Equal(nullFloat1.Float64, flt1, "The values should be equal")
	assert.Equal(nullFloat1.Valid, true, "The value should be true")

	nullFloat2 := NullFloatToFloat64(flt2)
	assert.NotNil(nullFloat2)
	assert.IsType(sql.NullFloat64{}, nullFloat2, "The result should be sql.NullFloat64")
	assert.Equal(nullFloat2.Float64, flt2, "The values should be equal")
	assert.Equal(nullFloat2.Valid, true, "The value should be true")

	nullFloat3 := NullFloatToFloat64(flt3)
	assert.NotNil(nullFloat3)
	assert.IsType(sql.NullFloat64{}, nullFloat3, "The result should be sql.NullFloat64")
	assert.Equal(nullFloat3.Float64, flt3, "The values should be equal")
	assert.Equal(nullFloat3.Valid, true, "The value should be true")

	newNullFloat := NullFloatToFloat64(0.0)
	assert.IsType(sql.NullFloat64{}, newNullFloat, "The result should be sql.NullFloat64")
	assert.Equal(newNullFloat.Valid, false, "The value should be false")
}

func TestFloat64ToString(t *testing.T) {
	assert := assert.New(t)

	// Need to shave off the trailing zeros
	// "87.300000" != "87.3"
	re := regexp.MustCompile(`^(.*?)\.(.*?)[0]*$`)

	result1 := Float64ToString(flt1)
	match1 := re.ReplaceAllString(result1, `$1.$2`)
	assert.NotNil(result1)
	assert.Equal(match1, "87.3")

	result2 := Float64ToString(flt2)
	match2 := re.ReplaceAllString(result2, `$1.$2`)
	assert.NotNil(result2)
	assert.Equal(match2, "1002.00003")

	result3 := Float64ToString(flt3)
	match3 := re.ReplaceAllString(result3, `$1.$2`)
	assert.NotNil(result3)
	assert.Equal(match3, "0.0006")
}
