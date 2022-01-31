package apic

import (
	"fmt"
	"math"
	"sort"
)

// Supported data types
const (
	DataTypeString  = "string"
	DataTypeArray   = "array"
	DataTypeInteger = "integer"
	DataTypeNumber  = "number"
	DataTypeObject  = "object"
)

// SubscriptionPropertyBuilder - used to build a subscription schmea property
type SubscriptionPropertyBuilder interface {
	SetName(name string) SubscriptionPropertyBuilder
	SetDescription(description string) SubscriptionPropertyBuilder
	SetRequired() SubscriptionPropertyBuilder
	SetReadOnly() SubscriptionPropertyBuilder
	SetHidden() SubscriptionPropertyBuilder
	SetAPICRefField(field string) SubscriptionPropertyBuilder
	Build() (*SubscriptionSchemaPropertyDefinition, error)
}

type SubscriptionStringPropertyBuilder interface {
	SubscriptionPropertyBuilder
	SetEnumValues(values []string) SubscriptionStringPropertyBuilder
	SetSortEnumValues() SubscriptionStringPropertyBuilder
	SetFirstEnumValue(value string) SubscriptionStringPropertyBuilder
	AddEnumValue(value string) SubscriptionStringPropertyBuilder
}

type SubscriptionNumberPropertyBuilder interface {
	SubscriptionPropertyBuilder
	SetMinValue(min float64) SubscriptionNumberPropertyBuilder
	SetMaxValue(max float64) SubscriptionNumberPropertyBuilder
}

type SubscriptionIntegerPropertyBuilder interface {
	SubscriptionPropertyBuilder
	SetMinValue(min int64) SubscriptionIntegerPropertyBuilder
	SetMaxValue(max int64) SubscriptionIntegerPropertyBuilder
}

type SubscriptionArrayPropertyBuilder interface {
	SubscriptionPropertyBuilder
	AddArrayItem(item SubscriptionPropertyBuilder) SubscriptionArrayPropertyBuilder
	SetMinArrayItems(min int) SubscriptionArrayPropertyBuilder
	SetMaxArrayItems(max int) SubscriptionArrayPropertyBuilder
}

type SubscriptionObjectPropertyBuilder interface {
	SubscriptionPropertyBuilder
	AddProperty(property SubscriptionPropertyBuilder) SubscriptionObjectPropertyBuilder
}

type stringSchemaProperty struct {
	SubscriptionPropertyBuilder
}

type objectSchemaProperty struct {
	SubscriptionObjectPropertyBuilder
}

func NewSubscriptionSchemaObjectPropertyBuilder() SubscriptionObjectPropertyBuilder {
	return &objectSchemaProperty{}
}

// schemaProperty - holds all the info needed to create a subscrition schema property
type schemaProperty struct {
	SubscriptionPropertyBuilder
	err            error
	name           string
	description    string
	apicRefField   string
	enums          []string
	minValue       *float64 // We use a pointer to differentiate the "blank value" from a choosen 0 min value
	maxValue       *float64 // We use a pointer to differentiate the "blank value" from a choosen 0 max value
	items          []SubscriptionSchemaPropertyDefinition
	minItems       int
	maxItems       int
	required       bool
	readOnly       bool
	hidden         bool
	dataType       string
	sortEnums      bool
	firstEnumValue string

	properties map[string]SubscriptionSchemaPropertyDefinition
}

// NewSubscriptionSchemaPropertyBuilder - Creates a new subscription schema property builder
func NewSubscriptionSchemaPropertyBuilder() SubscriptionPropertyBuilder {
	return &schemaProperty{
		enums: make([]string, 0),
	}
}

// SetName - sets the name of the property
func (p *schemaProperty) SetName(name string) SubscriptionPropertyBuilder {
	p.name = name
	return p
}

// SetDescription - set the description of the property
func (p *schemaProperty) SetDescription(description string) SubscriptionPropertyBuilder {
	p.description = description
	return p
}

// SetEnumValues - add a list of enum values to the property
func (p *schemaProperty) SetEnumValues(values []string) SubscriptionPropertyBuilder {
	dict := make(map[string]bool, 0)

	// use a temp map to filter out any duplicate values from the input
	for _, value := range values {
		if _, ok := dict[value]; !ok {
			dict[value] = true
			p.enums = append(p.enums, value)
		}
	}

	return p
}

// SetSortEnumValues - indicates to sort the enums
func (p *schemaProperty) SetSortEnumValues() SubscriptionPropertyBuilder {
	p.sortEnums = true
	return p
}

// SetFirstEnumValue - Sets a first item for enums. Only needed for sorted enums if you want a specific
// item first in the list
func (p *schemaProperty) SetFirstEnumValue(value string) SubscriptionPropertyBuilder {
	p.firstEnumValue = value
	return p
}

func (p *schemaProperty) enumContains(str string) bool {
	for _, v := range p.enums {
		if v == str {
			return true
		}
	}

	return false
}

// AddEnumValue - add a new value to the enum list
func (p *schemaProperty) AddEnumValue(value string) SubscriptionPropertyBuilder {
	if !p.enumContains(value) {
		p.enums = append(p.enums, value)
	}
	return p
}

// SetRequired - set the property as a required field in the schema
func (p *schemaProperty) SetRequired() SubscriptionPropertyBuilder {
	p.required = true
	return p
}

// SetReadOnly - set the property as a read only property
func (p *schemaProperty) SetReadOnly() SubscriptionPropertyBuilder {
	p.readOnly = true
	return p
}

// SetHidden - set the property as a hidden property
func (p *schemaProperty) SetHidden() SubscriptionPropertyBuilder {
	p.hidden = true
	return p
}

// SetAPICRefField - set the apic reference field for this property
func (p *schemaProperty) SetAPICRefField(field string) SubscriptionPropertyBuilder {
	p.apicRefField = field
	return p
}

// IsString - mark the datatype of the property as a string
func (p *schemaProperty) IsString() SubscriptionPropertyBuilder {
	if p.dataType != "" {
		p.err = fmt.Errorf("The data type cannot be set to %s, it is already set to %s", DataTypeString, p.dataType)
	} else {
		p.dataType = DataTypeString
	}
	return p
}

// IsNumber - mark the datatype of the property as a number
func (p *schemaProperty) IsNumber() SubscriptionPropertyBuilder {
	if p.dataType != "" {
		p.err = fmt.Errorf("The data type cannot be set to %s, it is already set to %s", DataTypeNumber, p.dataType)
	} else {
		p.dataType = DataTypeNumber
	}
	return p
}

// IsInteger - mark the datatype of the property as an integer
func (p *schemaProperty) IsInteger() SubscriptionPropertyBuilder {
	if p.dataType != "" {
		p.err = fmt.Errorf("The data type cannot be set to %s, it is already set to %s", DataTypeInteger, p.dataType)
	} else {
		p.dataType = DataTypeInteger
	}
	return p
}

// SetMinValue - set the minimum allowed value
func (p *schemaProperty) SetMinValue(min float64) SubscriptionPropertyBuilder {
	p.minValue = &min
	return p
}

// SetMaxArrayItems - set the maximum items allowed in the the array
func (p *schemaProperty) SetMaxValue(max float64) SubscriptionPropertyBuilder {
	p.maxValue = &max
	return p
}

// IsArray - mark the datatype of the property as an array
func (p *schemaProperty) IsArray() SubscriptionPropertyBuilder {
	if p.dataType != "" {
		p.err = fmt.Errorf("The data type cannot be set to %s, it is already set to %s", DataTypeArray, p.dataType)
	} else {
		p.dataType = DataTypeArray
	}
	return p
}

// SetArrayItems - store allowed items for the array
func (p *schemaProperty) SetArrayItems(items []SubscriptionPropertyBuilder) SubscriptionPropertyBuilder {
	for _, item := range items {
		p.AddArrayItem(item)
	}
	return p
}

func (p *schemaProperty) AddArrayItem(item SubscriptionPropertyBuilder) SubscriptionPropertyBuilder {
	def, err := item.Build()
	if err == nil {
		p.items = append(p.items, *def)
	} else {
		p.err = err
	}
	return p
}

// SetMinArrayItems - set the minimum items expected in the the array
func (p *schemaProperty) SetMinArrayItems(min int) SubscriptionPropertyBuilder {
	if min < 1 {
		p.err = fmt.Errorf("The min array items must be greater than 0")
	} else {
		p.minItems = min
	}
	return p
}

// SetMaxArrayItems - set the maximum items allowed in the the array
func (p *schemaProperty) SetMaxArrayItems(max int) SubscriptionPropertyBuilder {
	if max < 1 {
		p.err = fmt.Errorf("The max array items must be greater than 0")
	} else {
		p.maxItems = max
	}
	return p
}

// IsObject - mark the datatype of the property as an object
func (p *schemaProperty) IsObject() SubscriptionPropertyBuilder {
	if p.dataType != "" {
		p.err = fmt.Errorf("The data type cannot be set to %s, it is already set to %s", DataTypeObject, p.dataType)
	} else {
		p.dataType = DataTypeObject
	}
	return p
}

func (p *schemaProperty) AddProperty(property SubscriptionPropertyBuilder) SubscriptionPropertyBuilder {
	def, err := property.Build()
	if err == nil {
		if p.properties == nil {
			p.properties = make(map[string]SubscriptionSchemaPropertyDefinition, 0)
		}
		p.properties[def.Name] = *def
	} else {
		p.err = err
	}
	return p
}

// Build - create the SubscriptionSchemaPropertyDefinition for use in the subscription schema builder
func (p *schemaProperty) Build() (*SubscriptionSchemaPropertyDefinition, error) {
	if p.err != nil {
		return nil, p.err
	}
	if p.name == "" {
		return nil, fmt.Errorf("Cannot add a subscription schema property without a name")
	}

	if p.dataType == "" {
		return nil, fmt.Errorf("Subscription schema property named %s must have a data type", p.name)
	}

	var anyOfItems *AnyOfSubscriptionSchemaPropertyDefinitions
	if p.items != nil {
		if p.dataType != DataTypeArray {
			return nil, fmt.Errorf("Array items can only be set for schema property with the data type %s", DataTypeArray)
		}

		anyOfItems = &AnyOfSubscriptionSchemaPropertyDefinitions{p.items}
	}

	var requiredProperties []string
	if p.properties != nil {
		if p.dataType != DataTypeObject {
			return nil, fmt.Errorf("Properties can only be set for schema property with the data type %s", DataTypeObject)
		}

		for _, property := range p.properties {
			if property.Required {
				requiredProperties = append(requiredProperties, property.Name)
			}
		}
	}

	if p.minItems > p.maxItems {
		return nil, fmt.Errorf("Max array items (%d) must be greater than min array items (%d)", p.maxItems, p.minItems)
	}

	if p.minItems > 0 && p.dataType != DataTypeArray {
		return nil, fmt.Errorf("Min array items (%d) can only be set for schema property with the data type %s", p.minItems, DataTypeArray)
	}

	if p.maxItems > 0 && p.dataType != DataTypeArray {
		return nil, fmt.Errorf("Max array items (%d) can only be set for schema property with the data type %s", p.maxItems, DataTypeArray)
	}

	if p.minValue != nil && p.dataType != DataTypeInteger && p.dataType != DataTypeNumber {
		return nil, fmt.Errorf("Min value (%f) can only be set for schema property with the data type %s or %s", *p.minValue, DataTypeNumber, DataTypeInteger)
	}

	if p.minValue != nil && p.dataType == DataTypeInteger && math.Mod(*p.minValue, 1.0) != 0 {
		return nil, fmt.Errorf("Min value (%f) can be set only with integer for the data type %s", *p.minValue, DataTypeInteger)
	}

	if p.maxValue != nil && p.dataType != DataTypeInteger && p.dataType != DataTypeNumber {
		return nil, fmt.Errorf("Max value (%f) can only be set for schema property with the data type %s or %s", DataTypeNumber, *p.maxValue, DataTypeInteger)
	}

	if p.maxValue != nil && p.dataType == DataTypeInteger && math.Mod(*p.maxValue, 1.0) != 0 {
		return nil, fmt.Errorf("Max value (%f) can only be set with integer for the data type %s", *p.maxValue, DataTypeInteger)
	}

	if p.minValue != nil && p.maxValue != nil && *p.minValue > *p.maxValue {
		return nil, fmt.Errorf("Max value (%f) must be greater than min value (%f)", *p.maxValue, *p.minValue)
	}

	// sort if specified to do so
	if p.sortEnums {
		sort.Strings(p.enums)
	}

	// append item to start if specified
	if p.firstEnumValue != "" {
		p.enums = append([]string{p.firstEnumValue}, p.enums...)
	}

	prop := &SubscriptionSchemaPropertyDefinition{
		Name:               p.name,
		Title:              p.name,
		Type:               p.dataType,
		Description:        p.description,
		APICRef:            p.apicRefField,
		ReadOnly:           p.readOnly,
		Required:           p.required,
		Enum:               p.enums,
		SortEnums:          p.sortEnums,
		FirstEnumItem:      p.firstEnumValue,
		Items:              anyOfItems,
		MinItems:           p.minItems,
		MaxItems:           p.maxItems,
		Minimum:            p.minValue,
		Maximum:            p.maxValue,
		Properties:         p.properties,
		RequiredProperties: requiredProperties,
	}

	if p.hidden {
		prop.Format = "hidden"
	}

	return prop, nil
}
