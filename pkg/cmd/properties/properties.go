package properties

import (
	"encoding/json"
	goflag "flag"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/Axway/agent-sdk/pkg/util"
	"github.com/Axway/agent-sdk/pkg/util/errors"
	"github.com/Axway/agent-sdk/pkg/util/log"
	"github.com/spf13/cobra"
	flag "github.com/spf13/pflag"
	"github.com/spf13/viper"
)

// ErrInvalidSecretReference - Error for parsing properties with secret reference
var ErrInvalidSecretReference = errors.Newf(1411, "invalid secret reference - %s, please check the value for %s config")

// QA EnvVars
const qaEnforceDurationLowerLimit = "QA_ENFORCE_DURATION_LOWER_LIMIT"

// SecretPropertyResolver - interface for resolving property values with secret references
type SecretPropertyResolver interface {
	ResolveSecret(secretRef string) (string, error)
}

// Properties - Root Command Properties interface for all configs to use for adding and parsing values
type Properties interface {
	// Methods for adding yaml properties and command flag
	AddStringProperty(name string, defaultVal string, description string)
	AddStringPersistentFlag(name string, defaultVal string, description string)
	AddStringFlag(name string, description string)
	AddDurationProperty(name string, defaultVal time.Duration, description string)
	AddIntProperty(name string, defaultVal int, description string)
	AddBoolProperty(name string, defaultVal bool, description string)
	AddBoolFlag(name, description string)
	AddStringSliceProperty(name string, defaultVal []string, description string)
	AddObjectSliceProperty(name string, objectPropertyNames []string)

	// Methods to get the configured properties
	StringPropertyValue(name string) string
	StringFlagValue(name string) (bool, string)
	DurationPropertyValue(name string) time.Duration
	IntPropertyValue(name string) int
	BoolPropertyValue(name string) bool
	BoolPropertyValueOrTrue(name string) bool // Use this method when the default value, no config given, is true
	BoolFlagValue(name string) bool
	StringSlicePropertyValue(name string) []string
	ObjectSlicePropertyValue(name string) []map[string]interface{}

	// Methods to set a property
	SetStringFlagValue(name string, value string)

	// Log Properties
	MaskValues(name string)
	DebugLogProperties()
	SetAliasKeyPrefix(aliasKeyPrefix string)
}

var aliasKeyPrefix string

type properties struct {
	Properties
	rootCmd                  *cobra.Command
	envIntfArrayPropValues   map[string][]map[string]interface{}
	envIntfArrayPropertyKeys map[string]map[string]bool
	secretResolver           SecretPropertyResolver
	flattenedProperties      map[string]string
}

var expansionRegEx *regexp.Regexp

func init() {
	expansionRegEx = regexp.MustCompile(`\$\{(\w+):(.*)\}|\$\{(\w+)\}`)
}

// NewProperties - Creates a new Properties struct
func NewProperties(rootCmd *cobra.Command) Properties {
	cmdprops := &properties{
		rootCmd:                  rootCmd,
		envIntfArrayPropertyKeys: make(map[string]map[string]bool),
		flattenedProperties:      make(map[string]string),
	}

	return cmdprops
}

// NewPropertiesWithSecretResolver - Creates a new Properties struct with secret resolver for string property/flag
func NewPropertiesWithSecretResolver(rootCmd *cobra.Command, secretResolver SecretPropertyResolver) Properties {
	cmdprops := &properties{
		rootCmd:                  rootCmd,
		envIntfArrayPropertyKeys: make(map[string]map[string]bool),
		flattenedProperties:      make(map[string]string),
		secretResolver:           secretResolver,
	}

	return cmdprops
}

//SetAliasKeyPrefix -
func SetAliasKeyPrefix(keyPrefix string) {
	aliasKeyPrefix = keyPrefix
}

//GetAliasKeyPrefix -
func GetAliasKeyPrefix() string {
	return aliasKeyPrefix
}

func (p *properties) bindOrPanic(key string, flg *flag.Flag) {
	if err := viper.BindPFlag(key, flg); err != nil {
		panic(err)
	}
	if aliasKeyPrefix != "" {
		if err := viper.BindPFlag(aliasKeyPrefix+"."+key, flg); err != nil {
			panic(err)
		}
	}
}

func (p *properties) AddObjectSliceProperty(envPrefix string, intfPropertyNames []string) {
	envPrefix = strings.ReplaceAll(envPrefix, ".", "_")
	envPrefix = strings.ToUpper(envPrefix)
	if !strings.HasSuffix(envPrefix, "_") {
		envPrefix += "_"
	}
	iPropNames := make(map[string]bool)
	for _, propName := range intfPropertyNames {
		iPropNames[propName] = true
	}

	p.envIntfArrayPropertyKeys[envPrefix] = iPropNames
}

func (p *properties) AddStringProperty(name string, defaultVal string, description string) {
	if p.rootCmd != nil {
		flagName := p.nameToFlagName(name)
		p.rootCmd.Flags().String(flagName, defaultVal, description)
		p.bindOrPanic(name, p.rootCmd.Flags().Lookup(flagName))
		p.rootCmd.Flags().MarkHidden(flagName)
	}
}

func (p *properties) AddStringPersistentFlag(flagName string, defaultVal string, description string) {
	if p.rootCmd != nil {
		flg := goflag.CommandLine.Lookup(flagName)
		if flg == nil {
			goflag.CommandLine.String(flagName, "", description)
			flg = goflag.CommandLine.Lookup(flagName)
		}

		p.rootCmd.PersistentFlags().AddGoFlag(flg)
	}
}

func (p *properties) AddStringFlag(flagName string, description string) {
	if p.rootCmd != nil {
		p.rootCmd.Flags().String(flagName, "", description)
	}
}

func (p *properties) SetStringFlagValue(flagName string, value string) {
	if p.rootCmd != nil {
		p.rootCmd.Flags().Set(flagName, value)
	}
}

func (p *properties) AddStringSliceProperty(name string, defaultVal []string, description string) {
	if p.rootCmd != nil {
		flagName := p.nameToFlagName(name)
		p.rootCmd.Flags().StringSlice(flagName, defaultVal, description)
		p.bindOrPanic(name, p.rootCmd.Flags().Lookup(flagName))
		p.rootCmd.Flags().MarkHidden(flagName)
	}
}

func (p *properties) AddDurationProperty(name string, defaultVal time.Duration, description string) {
	if p.rootCmd != nil {
		flagName := p.nameToFlagName(name)
		p.rootCmd.Flags().Duration(flagName, defaultVal, description)
		p.bindOrPanic(name, p.rootCmd.Flags().Lookup(flagName))
		p.rootCmd.Flags().MarkHidden(flagName)
	}
}

func (p *properties) AddIntProperty(name string, defaultVal int, description string) {
	if p.rootCmd != nil {
		flagName := p.nameToFlagName(name)
		p.rootCmd.Flags().Int(flagName, defaultVal, description)
		p.bindOrPanic(name, p.rootCmd.Flags().Lookup(flagName))
		p.rootCmd.Flags().MarkHidden(flagName)
	}
}

func (p *properties) AddBoolProperty(name string, defaultVal bool, description string) {
	if p.rootCmd != nil {
		flagName := p.nameToFlagName(name)
		p.rootCmd.Flags().Bool(flagName, defaultVal, description)
		p.bindOrPanic(name, p.rootCmd.Flags().Lookup(flagName))
		p.rootCmd.Flags().MarkHidden(flagName)
	}
}

func (p *properties) AddBoolFlag(flagName string, description string) {
	if p.rootCmd != nil {
		p.rootCmd.Flags().Bool(flagName, false, description)
	}
}

func (p *properties) StringSlicePropertyValue(name string) []string {
	val := viper.Get(name)

	// special check to differentiate between yaml and commandline parsing. For commandline, must
	// turn it into an array ourselves
	switch val := val.(type) {
	case string:
		p.addPropertyToFlatMap(name, val)
		return p.convertStringToSlice(fmt.Sprintf("%v", viper.Get(name)))
	default:
		return viper.GetStringSlice(name)
	}
}

func (p *properties) convertStringToSlice(value string) []string {
	slc := strings.Split(value, ",")
	for i := range slc {
		slc[i] = strings.TrimSpace(slc[i])
	}
	return slc
}

func (p *properties) parseStringValueForKey(key string) string {
	s := strings.TrimSpace(viper.GetString(key))
	if strings.Index(s, "$") == 0 {
		matches := expansionRegEx.FindAllSubmatch([]byte(s), -1)
		if len(matches) > 0 {
			expSlice := matches[0]
			if len(expSlice) > 2 {
				envVar := string(expSlice[1])
				defaultVal := ""
				if envVar == "" {
					if len(expSlice) >= 4 {
						envVar = strings.Trim(string(expSlice[3]), "\"")
					}
				} else {
					if len(expSlice) >= 3 {
						defaultVal = strings.Trim(string(expSlice[2]), "\"")
					}
				}

				if envVar != "" {
					s = os.Getenv(envVar)
					if s == "" && defaultVal != "" {
						s = defaultVal
					}
				}
			}
		}
	}
	return s
}

func (p *properties) parseStringValue(key string) string {
	var s string
	if aliasKeyPrefix != "" {
		s = p.parseStringValueForKey(aliasKeyPrefix + "." + key)
	}
	// If no alias or no value parsed for alias key
	if s == "" {
		s = p.parseStringValueForKey(key)
	}
	return s
}

func (p *properties) resolveSecretReference(cfgName, cfgValue string) string {
	if p.secretResolver != nil {
		secretValue, err := p.secretResolver.ResolveSecret(cfgValue)
		if err != nil {
			// only log the error and return empty string,
			// validation on config triggers the agent to return error to root command
			log.Error(ErrInvalidSecretReference.FormatError(err.Error(), cfgName))
			cfgValue = ""
		}
		if secretValue != "" {
			cfgValue = secretValue
		}
	}
	return cfgValue
}

func (p *properties) StringPropertyValue(name string) string {
	s := p.parseStringValue(name)
	s = p.resolveSecretReference(name, s)
	p.addPropertyToFlatMap(name, s)
	return s
}

func (p *properties) StringFlagValue(name string) (bool, string) {
	flag := p.rootCmd.Flag(name)
	if flag == nil || flag.Value.String() == "" {
		return false, ""
	}
	fv := flag.Value.String()
	fv = p.resolveSecretReference(name, fv)
	p.addPropertyToFlatMap(name, fv)
	return true, fv
}

func (p *properties) DurationPropertyValue(name string) time.Duration {
	lowerLimit, _ := time.ParseDuration("30s")

	s := p.parseStringValue(name)
	d, _ := time.ParseDuration(s)

	// Get config value and check if duration is less than 30s, check if allow lower limits
	if isDurationLowerLimitEnforced() && d < lowerLimit {
		flagName := p.nameToFlagName(name)
		flag := p.rootCmd.Flag(flagName)

		// since config value is < 30s, get duration default value
		d, _ = time.ParseDuration(flag.DefValue)

		if d >= lowerLimit {
			// if defaultValue is > 30s, then just set the lower limit
			d = lowerLimit
			log.Warnf("config %s has been set to the lower limit value of %s. Please update this value greater than the lower limit if necessary", name, d)
		}
	}

	p.addPropertyToFlatMap(name, s)
	return d
}
func (p *properties) IntPropertyValue(name string) int {
	s := p.parseStringValue(name)
	i, _ := strconv.Atoi(s)

	p.addPropertyToFlatMap(name, s)
	return i
}

func (p *properties) BoolPropertyValue(name string) bool {
	return p.boolPropertyValue(name, false)
}

func (p *properties) BoolPropertyValueOrTrue(name string) bool {
	return p.boolPropertyValue(name, true)
}

func (p *properties) boolPropertyValue(name string, defVal bool) bool {
	s := p.parseStringValue(name)
	if s == "" {
		return defVal
	}
	b, _ := strconv.ParseBool(s)

	p.addPropertyToFlatMap(name, s)
	return b
}

func (p *properties) BoolFlagValue(name string) bool {
	flag := p.rootCmd.Flag(name)
	if flag == nil {
		return false
	}
	if flag.Value.String() == "true" {
		return true
	}
	return false
}

func (p *properties) nameToFlagName(name string) (flagName string) {
	parts := strings.Split(name, ".")
	flagName = parts[0]
	for _, part := range parts[1:] {
		flagName += strings.Title(part)
	}
	return
}

// String array containing any sensitive data that needs to be masked with "*" (asterisks)
// Add any sensitive data here using flattened key format
var maskValues = make([]string, 0)

func (p *properties) addPropertyToFlatMap(key, value string) {
	for _, maskValue := range maskValues {
		match, _ := regexp.MatchString("\\b"+strings.TrimSpace(maskValue)+"\\b", key)
		if match {
			value = util.MaskValue(value)
		}
	}
	p.flattenedProperties[key] = value
}

func (p *properties) MaskValues(maskedKeys string) {
	maskValues = strings.Split(maskedKeys, ",")
}

func (p *properties) DebugLogProperties() {
	data, _ := json.MarshalIndent(p.flattenedProperties, "", " ")
	log.Debugf("%s\n", data)
}

// enforce the lower limit by default
func isDurationLowerLimitEnforced() bool {
	if val := os.Getenv(qaEnforceDurationLowerLimit); val != "" {
		ret, err := strconv.ParseBool(val)
		if err != nil {
			log.Errorf("Invalid value %s for env variable %s", val, qaEnforceDurationLowerLimit)
			return true
		}
		return ret
	}
	return true
}

// ObjectSlicePropertyValue
func (p *properties) ObjectSlicePropertyValue(name string) []map[string]interface{} {
	p.readEnv()
	name = strings.ReplaceAll(name, ".", "_")
	name = strings.ToUpper(name)
	if !strings.HasSuffix(name, "_") {
		name += "_"
	}
	return p.envIntfArrayPropValues[name]
}

func (p *properties) readEnv() {
	if p.envIntfArrayPropValues != nil {
		return
	}

	p.envIntfArrayPropValues = make(map[string][]map[string]interface{})
	envVarsMap := p.parseEnvPropertiesFlatMap()

	for prefix, eVals := range envVarsMap {
		propValues, ok := p.envIntfArrayPropValues[prefix]
		if !ok {
			propValues = make([]map[string]interface{}, 0)
		}

		for _, val := range eVals {
			v := p.convertFlatMap(val)
			propValues = append(propValues, v)
		}
		p.envIntfArrayPropValues[prefix] = propValues
	}
}

func (p *properties) convertFlatMap(flatMap map[string]string) map[string]interface{} {
	m := make(map[string]interface{})
	for key, val := range flatMap {
		ok := strings.Contains(key, ".")
		if !ok {
			m[key] = val
		} else {
			k := strings.Split(key, ".")
			p.setChildMapProperty(m, k, val)
		}
	}
	return m
}

func (p *properties) setChildMapProperty(parentMap map[string]interface{}, childKeys []string, val string) {
	cm, ok := parentMap[childKeys[0]]
	if !ok {
		cm = make(map[string]interface{})
	}

	childMap, ok := cm.(map[string]interface{})
	if ok {
		if len(childKeys) > 2 {
			p.setChildMapProperty(childMap, childKeys[1:], val)
		} else {
			childMap[childKeys[1]] = val
		}
		parentMap[childKeys[0]] = cm
	}

}

func (p *properties) parseEnvPropertiesFlatMap() map[string]map[string]map[string]string {
	envVarsMap := make(map[string]map[string]map[string]string)
	for _, element := range os.Environ() {
		variable := strings.SplitN(element, "=", 2)
		name := variable[0]
		val := variable[1]
		for prefix, iPropNames := range p.envIntfArrayPropertyKeys {
			if strings.HasPrefix(name, prefix) {
				n := strings.ReplaceAll(name, prefix, "")
				elements := strings.Split(name, "_")
				lastSuffix := elements[len(elements)-1]
				_, ok := envVarsMap[prefix]

				if !ok {
					envVarsMap[prefix] = make(map[string]map[string]string)
				}

				m, ok := envVarsMap[prefix][lastSuffix]
				if !ok {
					m = make(map[string]string)
				}
				for pName := range iPropNames {
					propName := strings.ReplaceAll(pName, ".", "_")
					propName = strings.ToUpper(propName)
					if strings.HasPrefix(n, propName) {
						m[pName] = val
						envVarsMap[prefix][lastSuffix] = m
					}
				}

			}
		}
	}
	return envVarsMap
}
