package tfph

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetStringValue(t *testing.T) {
	cases := []struct {
		Schema map[string]*schema.Schema
		State  *terraform.InstanceState

		Diff  *terraform.InstanceDiff
		Key   string
		Value interface{}
		Ok    bool
	}{
		{
			Schema: map[string]*schema.Schema{
				"availability_zone": {
					Type:     schema.TypeString,
					Optional: true,
				},
			},

			State: &terraform.InstanceState{
				Attributes: map[string]string{
					"availability_zone": "us-east-1",
				},
			},
			Diff:  nil,
			Key:   "",
			Value: "",
			Ok:    true,
		},
	}

	for _, tc := range cases {
		d, err := schema.InternalMap(tc.Schema).Data(tc.State, tc.Diff)
		s := GetStringValue(d, "availability_zone")

		assert.Nil(t, err)
		assert.Equal(t, "us-east-1", s)
	}
}

func TestGetIntValue(t *testing.T) {
	cases := []struct {
		Schema map[string]*schema.Schema
		State  *terraform.InstanceState

		Diff  *terraform.InstanceDiff
		Key   string
		Value interface{}
		Ok    bool
	}{
		{
			Schema: map[string]*schema.Schema{
				"number_of_zones": {
					Type:     schema.TypeInt,
					Optional: true,
				},
			},

			State: &terraform.InstanceState{
				Attributes: map[string]string{
					"number_of_zones": "1",
				},
			},
			Diff:  nil,
			Key:   "",
			Value: "",
			Ok:    true,
		},
	}

	for _, tc := range cases {
		d, err := schema.InternalMap(tc.Schema).Data(tc.State, tc.Diff)
		s := GetIntValue(d, "number_of_zones")

		assert.Nil(t, err)
		assert.Equal(t, 1, s)
	}
}

func TestGetListValueAsStringSlice(t *testing.T) {
	cases := []struct {
		Schema map[string]*schema.Schema
		State  *terraform.InstanceState

		Diff  *terraform.InstanceDiff
		Key   string
		Value interface{}
		Ok    bool
	}{
		{
			Schema: map[string]*schema.Schema{
				"list_of_names": {
					Type: schema.TypeList,
					Elem: &schema.Schema{
						Type: schema.TypeString,
					},
					Optional: true,
				},
			},

			State: &terraform.InstanceState{
				Attributes: map[string]string{
					"list_of_names.#": "3",
					"list_of_names.0": "bob",
					"list_of_names.1": "jess",
					"list_of_names.2": "mike",
				},
			},
			Diff:  nil,
			Key:   "",
			Value: "",
			Ok:    true,
		},
	}

	for _, tc := range cases {
		d, err := schema.InternalMap(tc.Schema).Data(tc.State, tc.Diff)
		s := GetListValueAsStringSlice(d, "list_of_names")

		assert.Nil(t, err)
		assert.Equal(t, []string{"bob", "jess", "mike"}, s)
	}
}

func TestGetListValueAsIntSlice(t *testing.T) {
	cases := []struct {
		Schema map[string]*schema.Schema
		State  *terraform.InstanceState

		Diff  *terraform.InstanceDiff
		Key   string
		Value interface{}
		Ok    bool
	}{
		{
			Schema: map[string]*schema.Schema{
				"list_of_ints": {
					Type: schema.TypeList,
					Elem: &schema.Schema{
						Type: schema.TypeInt,
					},
					Optional: true,
				},
			},

			State: &terraform.InstanceState{
				Attributes: map[string]string{
					"list_of_ints.#": "3",
					"list_of_ints.0": "1",
					"list_of_ints.1": "2",
					"list_of_ints.2": "3",
				},
			},
			Diff:  nil,
			Key:   "",
			Value: "",
			Ok:    true,
		},
	}

	for _, tc := range cases {
		d, err := schema.InternalMap(tc.Schema).Data(tc.State, tc.Diff)
		s := GetListValueAsIntSlice(d, "list_of_ints")

		assert.Nil(t, err)
		assert.Equal(t, []int{1, 2, 3}, s)
	}
}

func TestGetSetValueAsStringSlice(t *testing.T) {
	cases := []struct {
		Schema map[string]*schema.Schema
		State  *terraform.InstanceState

		Diff  *terraform.InstanceDiff
		Key   string
		Value interface{}
		Ok    bool
	}{
		{
			Schema: map[string]*schema.Schema{
				"list_of_names": {
					Type: schema.TypeSet,
					Elem: &schema.Schema{
						Type: schema.TypeString,
					},
					Optional: true,
				},
			},

			State: &terraform.InstanceState{
				Attributes: map[string]string{
					"list_of_names.#":    "3",
					"list_of_names.bob":  "bob",
					"list_of_names.jess": "jess",
					"list_of_names.mike": "mike",
				},
			},
			Diff:  nil,
			Key:   "",
			Value: "",
			Ok:    true,
		},
	}

	for _, tc := range cases {
		d, err := schema.InternalMap(tc.Schema).Data(tc.State, tc.Diff)
		s := GetSetValueAsStringSlice(d, "list_of_names")

		assert.Nil(t, err)
		assert.Equal(t, []string{"mike", "bob", "jess"}, s)
	}
}

func TestGetSetValueAsIntSlice(t *testing.T) {
	cases := []struct {
		Schema map[string]*schema.Schema
		State  *terraform.InstanceState

		Diff  *terraform.InstanceDiff
		Key   string
		Value interface{}
		Ok    bool
	}{
		{
			Schema: map[string]*schema.Schema{
				"list_of_ints": {
					Type: schema.TypeSet,
					Elem: &schema.Schema{
						Type: schema.TypeInt,
					},
					Optional: true,
				},
			},

			State: &terraform.InstanceState{
				Attributes: map[string]string{
					"list_of_ints.#": "3",
					"list_of_ints.0": "1",
					"list_of_ints.1": "2",
					"list_of_ints.2": "3",
				},
			},
			Diff:  nil,
			Key:   "",
			Value: "",
			Ok:    true,
		},
	}

	for _, tc := range cases {
		d, err := schema.InternalMap(tc.Schema).Data(tc.State, tc.Diff)
		s := GetSetValueAsIntSlice(d, "list_of_ints")

		assert.Nil(t, err)
		assert.Equal(t, []int{2, 3, 1}, s)
	}
}
