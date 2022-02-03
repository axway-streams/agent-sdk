package apic

import (
	"fmt"
	"sort"
)

// Supported data types
const (
	DataTypeString  = "string"
	DataTypeNumber  = "number"
	DataTypeInteger = "integer"
	DataTypeArray   = "array"
	DataTypeObject  = "object"
)

type SubscriptionPropertyBuilder interface {
	Build() (*SubscriptionSchemaPropertyDefinition, error)
}

// schemaProperty - holds all the info needed to create a subscrition schema property
type propertyBuilder struct {
	err          error
	name         string
	description  string
	apicRefField string
	required     bool
	readOnly     bool
	hidden       bool
	dataType     string
	SubscriptionPropertyBuilder
}

// NewSubscriptionSchemaPropertyBuilder - Creates a new subscription schema property builder
func NewSubscriptionSchemaPropertyBuilder() *propertyBuilder {
	return &propertyBuilder{}
}

// SetName - sets the name of the property
func (p *propertyBuilder) SetName(name string) *propertyBuilder {
	p.name = name
	return p
}

// SetDescription - set the description of the property
func (p *propertyBuilder) SetDescription(description string) *propertyBuilder {
	p.description = description
	return p
}

// SetRequired - set the property as a required field in the schema
func (p *propertyBuilder) SetRequired() *propertyBuilder {
	p.required = true
	return p
}

// SetReadOnly - set the property as a read only property
func (p *propertyBuilder) SetReadOnly() *propertyBuilder {
	p.readOnly = true
	return p
}

// SetHidden - set the property as a hidden property
func (p *propertyBuilder) SetHidden() *propertyBuilder {
	p.hidden = true
	return p
}

// SetAPICRefField - set the apic reference field for this property
func (p *propertyBuilder) SetAPICRefField(field string) *propertyBuilder {
	p.apicRefField = field
	return p
}

// Build - create a string SubscriptionSchemaPropertyDefinition for use in the subscription schema builder
func (p *propertyBuilder) Build() (*SubscriptionSchemaPropertyDefinition, error) {
	if p.err != nil {
		return nil, p.err
	}
	if p.name == "" {
		return nil, fmt.Errorf("Cannot add a subscription schema property without a name")
	}

	if p.dataType == "" {
		return nil, fmt.Errorf("Subscription schema property named %s must have a data type", p.name)
	}

	prop := &SubscriptionSchemaPropertyDefinition{
		Name:        p.name,
		Title:       p.name,
		Type:        p.dataType,
		Description: p.description,
		APICRef:     p.apicRefField,
		ReadOnly:    p.readOnly,
		Required:    p.required,
	}

	if p.hidden {
		prop.Format = "hidden"
	}

	return prop, nil
}

/**
  string property datatype
*/

type stringPropertyBuilder struct {
	schemaProperty *propertyBuilder
	sortEnums      bool
	firstEnumValue string
	enums          []string
	SubscriptionPropertyBuilder
}

func (p *propertyBuilder) IsString() *stringPropertyBuilder {
	p.dataType = DataTypeString
	return &stringPropertyBuilder{
		schemaProperty: p,
	}
}

// SetEnumValues - add a list of enum values to the property
func (p *stringPropertyBuilder) SetEnumValues(values []string) *stringPropertyBuilder {
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
func (p *stringPropertyBuilder) SetSortEnumValues() *stringPropertyBuilder {
	p.sortEnums = true
	return p
}

// SetFirstEnumValue - Sets a first item for enums. Only needed for sorted enums if you want a specific
// item first in the list
func (p *stringPropertyBuilder) SetFirstEnumValue(value string) *stringPropertyBuilder {
	p.firstEnumValue = value
	return p
}

func (p *stringPropertyBuilder) enumContains(str string) bool {
	for _, v := range p.enums {
		if v == str {
			return true
		}
	}
	return false
}

// AddEnumValue - add a new value to the enum list
func (p *stringPropertyBuilder) AddEnumValue(value string) *stringPropertyBuilder {
	if !p.enumContains(value) {
		p.enums = append(p.enums, value)
	}
	return p
}

// Build - create a string SubscriptionSchemaPropertyDefinition for use in the subscription schema builder
func (p *stringPropertyBuilder) Build() (def *SubscriptionSchemaPropertyDefinition, err error) {

	def, err = p.schemaProperty.Build()
	if err != nil {
		return
	}

	// sort if specified to do so
	if p.sortEnums {
		sort.Strings(p.enums)
	}

	// append item to start if specified
	if p.firstEnumValue != "" {
		p.enums = append([]string{p.firstEnumValue}, p.enums...)
	}
	def.Enum = p.enums
	def.SortEnums = p.sortEnums
	def.FirstEnumItem = p.firstEnumValue
	return def, err
}

/**
  number property datatype builder
*/

type numberPropertyBuilder struct {
	schemaProperty *propertyBuilder
	minValue       *float64 // We use a pointer to differentiate the "blank value" from a choosen 0 min value
	maxValue       *float64 // We use a pointer to differentiate the "blank value" from a choosen 0 max value
	SubscriptionPropertyBuilder
}

func (p *propertyBuilder) IsNumber() *numberPropertyBuilder {
	p.dataType = DataTypeNumber
	return &numberPropertyBuilder{
		schemaProperty: p,
	}
}

// SetMinValue - set the minimum allowed value
func (p *numberPropertyBuilder) SetMinValue(min float64) *numberPropertyBuilder {
	p.minValue = &min
	return p
}

// SetMaxValue - set the maximum allowed value
func (p *numberPropertyBuilder) SetMaxValue(max float64) *numberPropertyBuilder {
	p.maxValue = &max
	return p
}

// Build - create a number SubscriptionSchemaPropertyDefinition for use in the subscription schema builder
func (p *numberPropertyBuilder) Build() (def *SubscriptionSchemaPropertyDefinition, err error) {
	def, err = p.schemaProperty.Build()
	if err != nil {
		return
	}

	if p.minValue != nil && p.maxValue != nil && *p.minValue > *p.maxValue {
		return nil, fmt.Errorf("Max value (%f) must be greater than min value (%f)", *p.maxValue, *p.minValue)
	}

	def.Minimum = p.minValue
	def.Maximum = p.maxValue
	return def, err
}

/**
  integer property datatype builder
*/

type integerPropertyBuilder struct {
	numberPropertyBuilder
}

func (p *propertyBuilder) IsInteger() *integerPropertyBuilder {
	p.dataType = DataTypeInteger
	return &integerPropertyBuilder{
		numberPropertyBuilder{
			schemaProperty: p,
		},
	}
}

// SetMinValue - set the minimum allowed value
func (p *integerPropertyBuilder) SetMinValue(min int64) *integerPropertyBuilder {
	minimum := float64(min)
	p.minValue = &minimum
	return p
}

// SetMaxValue - set the maximum allowed value
func (p *integerPropertyBuilder) SetMaxValue(max int64) *integerPropertyBuilder {
	maximum := float64(max)
	p.maxValue = &maximum
	return p
}

/**
  array property datatype builder
*/

type arrayPropertyBuilder struct {
	schemaProperty *propertyBuilder
	items          []SubscriptionSchemaPropertyDefinition
	minItems       *uint
	maxItems       *uint
	SubscriptionPropertyBuilder
}

func (p *propertyBuilder) IsArray() *arrayPropertyBuilder {
	p.dataType = DataTypeArray
	return &arrayPropertyBuilder{
		schemaProperty: p,
	}
}

func (p *arrayPropertyBuilder) AddArrayItem(item SubscriptionPropertyBuilder) *arrayPropertyBuilder {
	def, err := item.Build()
	if err == nil {
		p.items = append(p.items, *def)
	} else {
		p.schemaProperty.err = err
	}
	return p
}

// SetMinArrayItems - set the minimum items expected in the the array
func (p *arrayPropertyBuilder) SetMinArrayItems(min uint) *arrayPropertyBuilder {
	p.minItems = &min
	return p
}

// SetMaxArrayItems - set the maximum items allowed in the the array
func (p *arrayPropertyBuilder) SetMaxArrayItems(max uint) *arrayPropertyBuilder {
	if max < 1 {
		p.schemaProperty.err = fmt.Errorf("The max array items must be greater than 0")
	} else {
		p.maxItems = &max
	}
	return p
}

func (p *arrayPropertyBuilder) Build() (def *SubscriptionSchemaPropertyDefinition, err error) {
	def, err = p.schemaProperty.Build()
	if err != nil {
		return
	}

	var anyOfItems *AnyOfSubscriptionSchemaPropertyDefinitions
	if p.items != nil {
		anyOfItems = &AnyOfSubscriptionSchemaPropertyDefinitions{p.items}
	}

	if p.minItems != nil && p.maxItems != nil && *p.minItems > *p.maxItems {
		return nil, fmt.Errorf("Max array items (%d) must be greater than min array items (%d)", *p.maxItems, *p.minItems)
	}

	def.Items = anyOfItems
	def.MinItems = p.minItems
	def.MaxItems = p.maxItems
	return def, err
}

/**
  object property datatype builder
*/

type objectPropertyBuilder struct {
	schemaProperty *propertyBuilder
	properties     map[string]SubscriptionSchemaPropertyDefinition
	SubscriptionPropertyBuilder
}

func (p *propertyBuilder) IsObject() *objectPropertyBuilder {
	p.dataType = DataTypeObject
	return &objectPropertyBuilder{
		schemaProperty: p,
	}
}

func (p *objectPropertyBuilder) AddProperty(property SubscriptionPropertyBuilder) *objectPropertyBuilder {
	def, err := property.Build()
	if err == nil {
		if p.properties == nil {
			p.properties = make(map[string]SubscriptionSchemaPropertyDefinition, 0)
		}
		p.properties[def.Name] = *def
	} else {
		p.schemaProperty.err = err
	}
	return p
}

// Build - create the SubscriptionSchemaPropertyDefinition for use in the subscription schema builder
func (p *objectPropertyBuilder) Build() (def *SubscriptionSchemaPropertyDefinition, err error) {
	def, err = p.schemaProperty.Build()
	if err != nil {
		return
	}

	var requiredProperties []string
	if p.properties != nil {
		for _, property := range p.properties {
			if property.Required {
				requiredProperties = append(requiredProperties, property.Name)
			}
		}
	}

	def.Properties = p.properties
	def.RequiredProperties = requiredProperties
	return def, err
}
