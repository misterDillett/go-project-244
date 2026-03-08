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

    expected := `{
    common: {
      + follow: false
        setting1: Value 1
      - setting2: 200
      - setting3: true
      + setting3: null
      + setting4: blah blah
      + setting5: {
            key5: value5
        }
        setting6: {
            doge: {
              - wow:
              + wow: so much
            }
            key: value
          + ops: vops
        }
    }
    group1: {
      - baz: bas
      + baz: bars
        foo: bar
      - nest: {
            key: value
        }
      + nest: str
    }
  - group2: {
        abc: 12345
        deep: {
            id: 45
        }
    }
  + group3: {
        deep: {
            id: {
                number: 45
            }
        }
        fee: 100500
    }
}`
    assert.Equal(t, expected, result)
}

func TestGenDiff_NestedJSON_Plain(t *testing.T) {
    result, err := GenDiff(
        "testdata/fixture/file1_nested.json",
        "testdata/fixture/file2_nested.json",
        "plain",
    )
    require.NoError(t, err)

    expected := `Property 'common.follow' was added with value: false
Property 'common.setting2' was removed
Property 'common.setting3' was updated. From true to null
Property 'common.setting4' was added with value: 'blah blah'
Property 'common.setting5' was added with value: [complex value]
Property 'common.setting6.doge.wow' was updated. From '' to 'so much'
Property 'common.setting6.ops' was added with value: 'vops'
Property 'group1.baz' was updated. From 'bas' to 'bars'
Property 'group1.nest' was updated. From [complex value] to 'str'
Property 'group2' was removed
Property 'group3' was added with value: [complex value]`
    assert.Equal(t, expected, result)
}

func TestGenDiff_NestedJSON_JSON(t *testing.T) {
    result, err := GenDiff(
        "testdata/fixture/file1_nested.json",
        "testdata/fixture/file2_nested.json",
        "json",
    )
    require.NoError(t, err)

    assert.Contains(t, result, "common")
    assert.Contains(t, result, "group1")
    assert.Contains(t, result, "group2")
    assert.Contains(t, result, "group3")
    assert.Contains(t, result, "follow")
    assert.Contains(t, result, "setting6")
}

func TestGenDiff_NestedYAML_Stylish(t *testing.T) {
    result, err := GenDiff(
        "testdata/fixture/file1_nested.yaml",
        "testdata/fixture/file2_nested.yaml",
        "stylish",
    )
    require.NoError(t, err)

    assert.Contains(t, result, "common")
    assert.Contains(t, result, "group1")
    assert.Contains(t, result, "+ follow: false")
    assert.Contains(t, result, "- setting2: 200")
    assert.Contains(t, result, "doge")
}

func TestGenDiff_NestedYAML_Plain(t *testing.T) {
    result, err := GenDiff(
        "testdata/fixture/file1_nested.yaml",
        "testdata/fixture/file2_nested.yaml",
        "plain",
    )
    require.NoError(t, err)

    assert.Contains(t, result, "Property 'common.follow' was added with value: false")
    assert.Contains(t, result, "Property 'common.setting2' was removed")
    assert.Contains(t, result, "Property 'common.setting6.doge.wow' was updated")
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