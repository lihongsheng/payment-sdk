package tools

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateRandomUnix(t *testing.T) {
	ran := GenerateRandomUnix(4)
	t.Log(ran)
	assert.Equal(t, 4, len(ran))
	ran = GenerateRandomUnix(4)
	t.Log(ran)
	assert.Equal(t, 4, len(ran))
}

func TestGenerateRandomDigits(t *testing.T) {
	ran := GenerateRandomDigits(4)
	t.Log(ran)
	assert.Equal(t, 4, len(ran))
	ran = GenerateRandomDigits(4)
	t.Log(ran)
	assert.Equal(t, 4, len(ran))
}

func TestGenerateRandom(t *testing.T) {
	ran := NewGenerate(4)
	r := ran.Generate()
	t.Log(r)
	assert.Equal(t, 26, len(r))
	r = ran.Generate()
	t.Log(r)
	assert.Equal(t, 26, len(r))
	r = ran.Generate()
	t.Log(r)
	assert.Equal(t, 26, len(r))
}
