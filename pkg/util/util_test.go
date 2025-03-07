package util

import (
	"io/ioutil"
	"net/url"
	"os"
	"reflect"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestComputeHash(t *testing.T) {
	val, err := ComputeHash("this is a test")
	assert.Nil(t, err)
	val1, err := ComputeHash("this is a test1")
	assert.Nil(t, err)
	assert.NotEqual(t, val, val1)
	val, err = ComputeHash("this is a test1")
	assert.Nil(t, err)
	assert.Equal(t, val, val1)
}

func TestStringSliceContains(t *testing.T) {
	strSlice := []string{"foo", "bar"}
	assert.True(t, StringSliceContains(strSlice, "foo"))
	assert.False(t, StringSliceContains(strSlice, "foobar"))
}

func TestRemoveDuplicateValuesFromStringSlice(t *testing.T) {
	strSlice := []string{"foo", "bar", "foo", "bar", "foobar"}
	newSlice := RemoveDuplicateValuesFromStringSlice(strSlice)
	assert.Equal(t, 3, len(newSlice))
	assert.True(t, reflect.DeepEqual(newSlice, []string{"foo", "bar", "foobar"}))
}

func TestMaskValue(t *testing.T) {
	value := MaskValue("12345")
	assert.Equal(t, "*****", value)
}

func TestGetURLHostName(t *testing.T) {
	host := GetURLHostName("http://axway.com/abcd")
	assert.Equal(t, host, "axway.com")

	host = GetURLHostName("axway")
	assert.Equal(t, "", host)
}

func TestGetProxyURL(t *testing.T) {
	url := &url.URL{
		Scheme: "http",
		Host:   "axway.com",
		Path:   "abcd",
	}

	proxyurl := GetProxyURL(url)
	// assert.Nil(t, err)
	assert.NotNil(t, proxyurl)

	u, err := proxyurl(nil)
	assert.Nil(t, err)
	assert.NotNil(t, u)
	assert.Equal(t, url, u)

	url.Host = ""
	proxyurl = GetProxyURL(url)
	u, err = proxyurl(nil)
	assert.Nil(t, err)
	assert.Nil(t, u)

	proxyurl = GetProxyURL(nil)
	u, err = proxyurl(nil)
	assert.Nil(t, err)
	assert.Nil(t, u)
}

func TestLoadEnvFromFile(t *testing.T) {
	err := LoadEnvFromFile("foobar")
	assert.NotNil(t, err)

	err = LoadEnvFromFile("./testdata/env_vars.txt")
	assert.Nil(t, err)

	assert.Equal(t, "https://bbunny.dev.test.net", os.Getenv("CENTRAL_URL"))
	i, _ := strconv.ParseInt(os.Getenv("CENTRAL_INTVAL1"), 10, 0)
	assert.Equal(t, int64(15), i)
	b, _ := strconv.ParseBool(os.Getenv("CENTRAL_SSL_INSECURESKIPVERIFY"))
	assert.True(t, b)

	// These keys in the env_vars.txt all have values followed by a TAB char
	// this test is to verify that they get parsed correctly
	b, _ = strconv.ParseBool(os.Getenv("CENTRAL_USAGEREPORTING_OFFLINE"))
	assert.True(t, b)
	i, _ = strconv.ParseInt(os.Getenv("CENTRAL_INTVAL2"), 10, 0)
	assert.Equal(t, int64(20), i)
	assert.Equal(t, "https://test.net", os.Getenv("CENTRAL_AUTH_URL"))
}

func TestMergeMapStringInterface(t *testing.T) {
	m1 := map[string]interface{}{
		"foo": "foo1",
		"baz": "baz1",
		"aaa": "aaa1",
	}
	m2 := map[string]interface{}{
		"foo":  "foo2",
		"baz":  "baz2",
		"quux": "quux2",
		"asdf": "asdf2",
	}

	result := MergeMapStringInterface(m1, m2)
	assert.Equal(t, m1["aaa"], result["aaa"])
	assert.Equal(t, m2["foo"], result["foo"])
	assert.Equal(t, m2["baz"], result["baz"])
	assert.Equal(t, m2["quux"], result["quux"])
	assert.Equal(t, m2["asdf"], result["asdf"])

	m3 := map[string]interface{}{
		"foo":  "foo3",
		"zxcv": "zxcv3",
	}

	resul2t := MergeMapStringInterface(m1, m2, m3)
	assert.Equal(t, m1["aaa"], resul2t["aaa"])
	assert.Equal(t, m2["baz"], resul2t["baz"])
	assert.Equal(t, m2["quux"], resul2t["quux"])
	assert.Equal(t, m2["asdf"], resul2t["asdf"])
	assert.Equal(t, m3["foo"], resul2t["foo"])
	assert.Equal(t, m3["zxcv"], resul2t["zxcv"])

	result3 := MergeMapStringInterface(nil)
	assert.NotNil(t, result3)

	result4 := MergeMapStringInterface(m1, nil)
	assert.NotNil(t, result4)
}

func TestMergeMapStringString(t *testing.T) {
	m1 := map[string]string{
		"foo": "foo1",
		"baz": "baz1",
		"aaa": "aaa1",
	}
	m2 := map[string]string{
		"foo":  "foo2",
		"baz":  "baz2",
		"quux": "quux2",
		"asdf": "asdf2",
	}

	result := MergeMapStringString(m1, m2)
	assert.Equal(t, m1["aaa"], result["aaa"])
	assert.Equal(t, m2["foo"], result["foo"])
	assert.Equal(t, m2["baz"], result["baz"])
	assert.Equal(t, m2["quux"], result["quux"])
	assert.Equal(t, m2["asdf"], result["asdf"])

	m3 := map[string]string{
		"foo":  "foo3",
		"zxcv": "zxcv3",
	}

	result2 := MergeMapStringString(m1, m2, m3)
	assert.Equal(t, m1["aaa"], result2["aaa"])
	assert.Equal(t, m2["baz"], result2["baz"])
	assert.Equal(t, m2["quux"], result2["quux"])
	assert.Equal(t, m2["asdf"], result2["asdf"])
	assert.Equal(t, m3["foo"], result2["foo"])
	assert.Equal(t, m3["zxcv"], result2["zxcv"])

	result3 := MergeMapStringString(nil)
	assert.NotNil(t, result3)

	result4 := MergeMapStringString(m1, nil)
	assert.NotNil(t, result4)
}

func TestMapStringInterfaceToStringString(t *testing.T) {
	m1 := map[string]interface{}{
		"foo":  "foo1",
		"baz":  false,
		"aaa":  1,
		"test": `{"a":"a","b":["1","2","3"]}`,
		"nil":  nil,
	}
	result := MapStringInterfaceToStringString(m1)

	assert.Equal(t, "foo1", result["foo"])
	assert.Equal(t, "false", result["baz"])
	assert.Equal(t, "1", result["aaa"])
	assert.Equal(t, `{"a":"a","b":["1","2","3"]}`, result["test"])
	assert.Equal(t, "", result["nil"])
}

func TestParsePort(t *testing.T) {
	p := ParsePort(nil)
	assert.Equal(t, 0, p)

	u, _ := url.Parse("http://test:222")
	p = ParsePort(u)
	assert.Equal(t, 222, p)

	u, _ = url.Parse("http://test")
	p = ParsePort(u)
	assert.Equal(t, 80, p)

	u, _ = url.Parse("noscheme://test")
	p = ParsePort(u)
	assert.Equal(t, 0, p)
}

func TestParseAddr(t *testing.T) {
	addr := ParseAddr(nil)
	assert.Equal(t, "", addr)

	u, _ := url.Parse("http://test:222")
	addr = ParseAddr(u)
	assert.Equal(t, "test:222", addr)

	u, _ = url.Parse("http://test")
	addr = ParseAddr(u)
	assert.Equal(t, "test:80", addr)
}

func TestComputeKIDFromDER(t *testing.T) {
	key, err := ioutil.ReadFile("testdata/public_key")
	if err != nil {
		t.Errorf("unable to read public_key")
	}
	res, err := ComputeKIDFromDER(key)
	if err != nil {
		t.Errorf("unable to compute kid")
	}
	expected := "1wzYoslzjo-ROTN1CUWPQYtTUqrqiaDO96fAAmb7JvA"
	if res != expected {
		t.Fail()
	}

	// der file format
	key, err = ioutil.ReadFile("testdata/public_key.der")
	if err != nil {
		t.Errorf("unable to read public_key.der")
	}
	res, err = ComputeKIDFromDER(key)
	if err != nil {
		t.Errorf("unable to compute kid")
	}
	expected = "iXcfstYFMANhYzgPwMWJxIQdfLQBqWjdiwyl7e4xv6Q"
	if res != expected {
		t.Fail()
	}
}

func TestReadPrivateKey(t *testing.T) {
	cases := []struct {
		description  string
		privKeyFile  string
		passwordFile string
	}{
		{
			description: "no password",
			privKeyFile: "testdata/private_key.pem",
		},
		{
			description:  "with empty password file",
			privKeyFile:  "testdata/private_key.pem",
			passwordFile: "testdata/password_empty",
		},
		{
			description:  "with password",
			privKeyFile:  "testdata/private_key_with_pwd.pem",
			passwordFile: "testdata/password",
		},
	}

	for _, testCase := range cases {
		if _, err := ReadPrivateKeyFile(testCase.privKeyFile, testCase.passwordFile); err != nil {
			t.Errorf("testcase: %s: failed to read rsa key %s", testCase.description, err)
		}
	}
}

func TestReadPublicKeyFile(t *testing.T) {
	cases := []struct {
		description   string
		publicKeyFile string
	}{
		{
			description:   "with public key",
			publicKeyFile: "testdata/public_key",
		},
	}
	for _, testCase := range cases {
		if _, err := ReadPublicKeyBytes(testCase.publicKeyFile); err != nil {
			t.Errorf("testcase: %s: failed to read public key %s", testCase.description, err)
		}
	}
}

func TestGetStringFromMapInterface(t *testing.T) {
	cases := []struct {
		data        map[string]interface{}
		key         string
		expectedVal string
	}{
		{
			data:        map[string]interface{}{"key": "valid"},
			key:         "key",
			expectedVal: "valid",
		},
		{
			data:        map[string]interface{}{"key": 10},
			key:         "invalidKey",
			expectedVal: "",
		},
		{
			data:        map[string]interface{}{"key": 10},
			expectedVal: "",
		},
	}
	for _, testCase := range cases {
		ret := GetStringFromMapInterface(testCase.key, testCase.data)
		assert.Equal(t, testCase.expectedVal, ret)
	}
}

func TestGetStringArrayFromMapInterface(t *testing.T) {
	cases := []struct {
		data        map[string]interface{}
		key         string
		expectedVal []string
	}{
		{
			data:        map[string]interface{}{"key": []string{"val1", "val2"}},
			key:         "key",
			expectedVal: []string{"val1", "val2"},
		},
		{
			data:        map[string]interface{}{"key": []interface{}{"val1", "val2"}},
			key:         "key",
			expectedVal: []string{"val1", "val2"},
		},
		{
			data:        map[string]interface{}{"key": []string{"val1", "val2"}},
			key:         "invalidKey",
			expectedVal: []string{},
		},
		{
			data:        map[string]interface{}{"key": []interface{}{10, "val1"}},
			key:         "key",
			expectedVal: []string{"val1"},
		},
		{
			data:        map[string]interface{}{"key": []interface{}{10, 10}},
			key:         "key",
			expectedVal: []string{},
		},
		{
			data:        map[string]interface{}{"key": []int{10}},
			expectedVal: []string{},
		},
	}
	for _, testCase := range cases {
		ret := GetStringArrayFromMapInterface(testCase.key, testCase.data)
		assert.Equal(t, testCase.expectedVal, ret)
	}
}

func TestConvertToDomainNameCompliant(t *testing.T) {
	name := ConvertToDomainNameCompliant("Abc.Def")
	assert.Equal(t, "abc.def", name)
	name = ConvertToDomainNameCompliant(".Abc.Def")
	assert.Equal(t, "abc.def", name)
	name = ConvertToDomainNameCompliant(".Abc...De/f")
	assert.Equal(t, "abc--.de-f", name)
	name = ConvertToDomainNameCompliant("Abc.D-ef")
	assert.Equal(t, "abc.d-ef", name)
	name = ConvertToDomainNameCompliant("Abc.Def=")
	assert.Equal(t, "abc.def", name)
	name = ConvertToDomainNameCompliant("A..bc.Def")
	assert.Equal(t, "a--bc.def", name)
}
