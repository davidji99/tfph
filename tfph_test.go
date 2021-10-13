package tfph

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseCompositeID_DefaultSeparator(t *testing.T) {
	expected := []string{"id1", "id2", "id3"}
	result, err := ParseCompositeID("id1:id2:id3", 3)

	assert.Nil(t, err)
	assert.Equal(t, expected, result)
}

func TestParseCompositeID_CustomSeparator(t *testing.T) {
	expected := []string{"id1", "id2", "id3"}
	result, err := ParseCompositeID("id1|id2|id3", 3, "|")

	assert.Nil(t, err)
	assert.Equal(t, expected, result)
}

func TestParseCompositeID_MultipleCustomSeparator(t *testing.T) {
	expected := []string{"id1", "id2", "id3"}
	result, err := ParseCompositeID("id1;id2;id3", 3, ";", "|")

	assert.Nil(t, err)
	assert.Equal(t, expected, result)
}

func TestParseCompositeID_Invalid(t *testing.T) {
	result, err := ParseCompositeID("id1|id2|id3", 2, "|")

	assert.Equal(t, []string(nil), result)
	assert.Equal(t, "error: composite ID requires 2 parts separated by a [|] (x|y)", err.Error())
}

func TestContainsString(t *testing.T) {
	tester := []string{"abc", "efd", "xyz"}
	assert.Equal(t, true, ContainsString(tester, "abc"))
}

func TestContainsString_Invalid(t *testing.T) {
	tester := []string{"abc", "efd", "xyz"}
	assert.Equal(t, false, ContainsString(tester, "a2bc"))
}
