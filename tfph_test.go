package tfph

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
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

func TestDoesNotContainString(t *testing.T) {
	tester := []string{"abc", "efd", "xyz"}
	assert.Equal(t, true, DoesNotContainString(tester, "ab1c"))
}

func TestDoesNotContainString_Invalid(t *testing.T) {
	tester := []string{"abc", "efd", "xyz"}
	assert.Equal(t, false, DoesNotContainString(tester, "abc"))
}

func TestErrsFromDiags(t *testing.T) {
	expected := fmt.Errorf("Severity: 0 | Summary: The summer of error one, | Detail: The Details of error one\nSeverity: 0 | Summary: The summer of error two, | Detail: The Details of error two\n")
	diags := diag.Diagnostics{
		{
			Severity: diag.Error,
			Summary:  "The summer of error one",
			Detail:   "The Details of error one",
		},
		{
			Severity: diag.Error,
			Summary:  "The summer of error two",
			Detail:   "The Details of error two",
		},
	}

	assert.Equal(t, expected, ErrsFromDiags(diags))
}

func TestErrsFromDiags_NoErrors(t *testing.T) {
	diags := diag.Diagnostics{}

	assert.Equal(t, nil, ErrsFromDiags(diags))
}
