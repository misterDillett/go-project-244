package code

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGenDiff_FlatJSON(t *testing.T) {
	result, err := GenDiff(
		"testdata/fixture/file1.json",
		"testdata/fixture/file2.json",
		"stylish",
	)
	require.NoError(t, err)

	expected := `{
  - follow: false
    host: hexlet.io
  - proxy: 123.234.53.22
  - timeout: 50
  + timeout: 20
  + verbose: true
}`
	assert.Equal(t, expected, result)
}

func TestGenDiff_FlatYAML(t *testing.T) {
	result, err := GenDiff(
		"testdata/fixture/file1.yaml",
		"testdata/fixture/file2.yaml",
		"stylish",
	)
	require.NoError(t, err)

	expected := `{
  - follow: false
    host: hexlet.io
  - proxy: 123.234.53.22
  - timeout: 50
  + timeout: 20
  + verbose: true
}`
	assert.Equal(t, expected, result)
}

func TestGenDiff_NestedJSON_Stylish(t *testing.T) {
	result, err := GenDiff(
		"testdata/fixture/file1_nested.json",
		"testdata/fixture/file2_nested.json",
		"stylish",
	)
	require.NoError(t, err)

	assert.Contains(t, result, "common")
	assert.Contains(t, result, "group1")
	assert.Contains(t, result, "+ follow: false")
	assert.Contains(t, result, "- setting2: 200")
	assert.Contains(t, result, "setting1: Value 1")
}

func TestGenDiff_NestedJSON_Plain(t *testing.T) {
	result, err := GenDiff(
		"testdata/fixture/file1_nested.json",
		"testdata/fixture/file2_nested.json",
		"plain",
	)
	require.NoError(t, err)

	assert.Contains(t, result, "Property 'common.follow' was added with value: false")
	assert.Contains(t, result, "Property 'common.setting2' was removed")
	assert.Contains(t, result, "Property 'common.setting3' was updated. From true to null")
}

func TestGenDiff_NestedJSON_JSON(t *testing.T) {
	result, err := GenDiff(
		"testdata/fixture/file1_nested.json",
		"testdata/fixture/file2_nested.json",
		"json",
	)
	require.NoError(t, err)

	assert.Contains(t, result, "{")
	assert.Contains(t, result, "}")
	assert.Contains(t, result, "common")
	assert.Contains(t, result, "group1")
}

func TestGenDiff_NestedYAML(t *testing.T) {
	result, err := GenDiff(
		"testdata/fixture/file1_nested.yaml",
		"testdata/fixture/file2_nested.yaml",
		"stylish",
	)
	require.NoError(t, err)

	assert.Contains(t, result, "common")
	assert.Contains(t, result, "group1")
	assert.Contains(t, result, "+ follow: false")
}

func TestGenDiff_InvalidFormat(t *testing.T) {
	_, err := GenDiff(
		"testdata/fixture/file1.json",
		"testdata/fixture/file2.json",
		"invalid",
	)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "unknown format")
}

func TestGenDiff_NonexistentFile(t *testing.T) {
	_, err := GenDiff(
		"nonexistent.json",
		"testdata/fixture/file2.json",
		"stylish",
	)
	assert.Error(t, err)
}
