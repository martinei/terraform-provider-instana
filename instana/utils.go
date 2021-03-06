package instana

import (
	"fmt"

	"github.com/gessnerfl/terraform-provider-instana/instana/restapi"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/rs/xid"
)

//RandomID generates a random ID for a resource
func RandomID() string {
	xid := xid.New()
	return xid.String()
}

//ReadStringArrayParameterFromResource reads a string array parameter from a resource
func ReadStringArrayParameterFromResource(d *schema.ResourceData, key string) []string {
	if attr, ok := d.GetOk(key); ok {
		var array []string
		items := attr.([]interface{})
		for _, x := range items {
			item := x.(string)
			array = append(array, item)
		}
		return array
	}
	return nil
}

//ReadStringSetParameterFromResource reads a string set parameter from a resource and returns it as a slice of strings
func ReadStringSetParameterFromResource(d *schema.ResourceData, key string) []string {
	if attr, ok := d.GetOk(key); ok {
		var array []string
		set := attr.(*schema.Set)
		for _, x := range set.List() {
			item := x.(string)
			array = append(array, item)
		}
		return array
	}
	return nil
}

//ConvertStringToInterfaceSlice converts a string slice to interface slice for e.g. persistene in TF state
func ConvertStringToInterfaceSlice(input []string) []interface{} {
	if len(input) > 0 {
		result := make([]interface{}, len(input))
		for i, v := range input {
			result[i] = v
		}
		return result
	}
	return []interface{}{}
}

//ConvertSeverityFromInstanaAPIToTerraformRepresentation converts the integer representation of the Instana API to the string representation of the Terraform provider
func ConvertSeverityFromInstanaAPIToTerraformRepresentation(severity int) (string, error) {
	if severity == restapi.SeverityWarning.GetAPIRepresentation() {
		return restapi.SeverityWarning.GetTerraformRepresentation(), nil
	} else if severity == restapi.SeverityCritical.GetAPIRepresentation() {
		return restapi.SeverityCritical.GetTerraformRepresentation(), nil
	} else {
		return "INVALID", fmt.Errorf("%d is not a valid severity", severity)
	}
}

//ConvertSeverityFromTerraformToInstanaAPIRepresentation converts the string representation of the Terraform to the int representation of the Instana API provider
func ConvertSeverityFromTerraformToInstanaAPIRepresentation(severity string) (int, error) {
	if severity == restapi.SeverityWarning.GetTerraformRepresentation() {
		return restapi.SeverityWarning.GetAPIRepresentation(), nil
	} else if severity == restapi.SeverityCritical.GetTerraformRepresentation() {
		return restapi.SeverityCritical.GetAPIRepresentation(), nil
	} else {
		return -1, fmt.Errorf("%s is not a valid severity", severity)
	}
}

//GetIntPointerFromResourceData gets a int value from the resource data and either returns a pointer to the value or nil if the value is not defined
func GetIntPointerFromResourceData(d *schema.ResourceData, key string) *int {
	val, ok := d.GetOk(key)
	if ok {
		intValue := val.(int)
		return &intValue
	}
	return nil
}

//GetFloat64PointerFromResourceData gets a float64 value from the resource data and either returns a pointer to the value or nil if the value is not defined
func GetFloat64PointerFromResourceData(d *schema.ResourceData, key string) *float64 {
	val, ok := d.GetOk(key)
	if ok {
		floatValue := val.(float64)
		return &floatValue
	}
	return nil
}

//GetStringPointerFromResourceData gets a string value from the resource data and either returns a pointer to the value or nil if the value is not defined
func GetStringPointerFromResourceData(d *schema.ResourceData, key string) *string {
	val, ok := d.GetOk(key)
	if ok {
		stringValue := val.(string)
		return &stringValue
	}
	return nil
}
