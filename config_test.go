package config_test

import (
	"testing"

	"github.com/drylikov/go_config"
)

// Config struct.
type Config struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

// Test saving.
func TestSave(t *testing.T) {
	err := config.Save("/tmp/whatever/some.json", Config{
		Name:  "drylikov",
		Email: "drylikov@apex.sh",
	})
	assert.NoError(t, err)
}

// Test loading valid config.
func TestLoad_valid(t *testing.T) {
	var c Config
	err := config.Load("/tmp/whatever/some.json", &c)
	assert.NoError(t, err)
	assert.Equal(t, "drylikov", c.Name)
}

// Test loading missing config.
func TestLoad_missing(t *testing.T) {
	var c Config
	err := config.Load("/tmp/nope.json", &c)
	assert.NoError(t, err)
}

// Test saving in home.
func TestSaveHome(t *testing.T) {
	err := config.SaveHome("some.json", Config{
		Name:  "drylikov",
		Email: "drylikov@apex.sh",
	})
	assert.NoError(t, err)
}

// Test loading valid config in home.
func TestLoadHome_valid(t *testing.T) {
	var c Config
	err := config.LoadHome("some.json", &c)
	assert.NoError(t, err)
	assert.Equal(t, "drylikov", c.Name)
}