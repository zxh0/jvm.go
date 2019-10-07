package classfile

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestJava8HW(t *testing.T) {
	bytes, err := ioutil.ReadFile("../test/testdata/java8/HelloWorld.class")
	require.NoError(t, err)

	cf, err := Parse(bytes)
	require.NoError(t, err)
	require.Equal(t, uint16(52), cf.MajorVersion)
	require.Equal(t, uint16(0), cf.MinorVersion)
	require.Equal(t, 34, len(cf.ConstantPool))
	require.Equal(t, uint16(0x21), cf.AccessFlags)
	require.Equal(t, "HelloWorld", cf.GetThisClassName())
	require.Equal(t, "java/lang/Object", cf.GetSuperClassName())
	require.Equal(t, []string{}, cf.GetInterfaceNames())
	require.Equal(t, 0, len(cf.Fields))
	require.Equal(t, 2, len(cf.Methods))
	require.Equal(t, 1, len(cf.AttributeTable))
	require.Equal(t, "HelloWorld.java", cf.GetUTF8(cf.GetSourceFileIndex()))
}

func TestJava13HW(t *testing.T) {
	bytes, err := ioutil.ReadFile("../test/testdata/java13/HelloWorld.class")
	require.NoError(t, err)

	cf, err := Parse(bytes)
	require.NoError(t, err)
	require.Equal(t, uint16(57), cf.MajorVersion)
	require.Equal(t, uint16(0), cf.MinorVersion)
}

func TestModuleInfo(t *testing.T) {
	bytes, err := ioutil.ReadFile("../test/testdata/java13/module-info.class")
	require.NoError(t, err)

	cf, err := Parse(bytes)
	require.NoError(t, err)
	require.Equal(t, uint16(57), cf.MajorVersion)
	require.Equal(t, uint16(0), cf.MinorVersion)
	modAttr, _ := cf.GetModuleAttribute()
	require.Equal(t, uint16(6), modAttr.ModuleNameIndex)
}
