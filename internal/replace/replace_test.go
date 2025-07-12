package replace_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/winebarrel/terraform-provider-multireplace/internal/replace"
)

func TestReplace_OK(t *testing.T) {
	assert := assert.New(t)
	require := require.New(t)
	out, err := replace.Replace("London Bridge Is Falling Down, Falling down, falling down", "Falling", "Winding")
	require.NoError(err)
	assert.Equal("London Bridge Is Winding Down, Winding down, falling down", out)
}

func TestReplace_OK_Regexp(t *testing.T) {
	assert := assert.New(t)
	require := require.New(t)
	out, err := replace.Replace("London Bridge Is Falling Down, Falling down, falling down", "/(?i)(falling|down)/", "💥")
	require.NoError(err)
	assert.Equal("London Bridge Is 💥 💥, 💥 💥, 💥 💥", out)
}

func TestReplace_Err_Regexp(t *testing.T) {
	assert := assert.New(t)
	_, err := replace.Replace("blah blah blah", "/)/", "N/A")
	assert.ErrorContains(err, "error parsing regexp")
}

func TestReplace_OK_NonRegexp(t *testing.T) {
	assert := assert.New(t)
	require := require.New(t)
	out, err := replace.Replace("blah/blah/blah", "/", " ")
	require.NoError(err)
	assert.Equal("blah blah blah", out)
}

func TestMultiReplace_OK(t *testing.T) {
	assert := assert.New(t)
	require := require.New(t)
	out, err := replace.MultiReplace(
		"London Bridge Is Falling Down, Falling down, falling down",
		map[string]string{"Falling": "Winding", "falling": "jumping"},
	)
	require.NoError(err)
	assert.Equal("London Bridge Is Winding Down, Winding down, jumping down", out)
}

func TestMultiReplace_OK_Regexp(t *testing.T) {
	assert := assert.New(t)
	require := require.New(t)
	out, err := replace.MultiReplace(
		"London Bridge Is Falling Down, Falling down, falling down",
		map[string]string{"/(?i)falling/": "raising", "/(?i)down/": "up"},
	)
	require.NoError(err)
	assert.Equal("London Bridge Is raising up, raising up, raising up", out)
}
