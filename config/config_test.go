package config

import (
	"testing"
)

func TestConfig(t *testing.T) {
	var err error

	t.Run("Test Config creation", func(t *testing.T) {
		_, err = NewConfig(WithConfigPath("../.loganalyzer.json"))
		if err != nil {
			t.Error(err)
		}
	})
	t.Run("Init with wrong file path", func(t *testing.T) {
		err = Init(WithConfigPath(".loganalyzer.json"))
		if err == nil {
			t.Error("Error is expected but got nil")
		}
	})
	t.Run("Failing Load config", func(t *testing.T) {
		err = Load()
		if err == nil {
			t.Error("Expected error but got nil")
		}
	})
	t.Run("Init with valid file path", func(t *testing.T) {
		err = Init(WithConfigPath("../.loganalyzer.json"))
		if err != nil {
			t.Error(err)
		}
	})
	t.Run("Load config", func(t *testing.T) {
		err = Load()
		if err != nil {
			t.Error(err)
		}
	})
	t.Run("Set config", func(t *testing.T) {
		Set("foo", "bar")
	})
	t.Run("Get config", func(t *testing.T) {
		v := Get("foo")
		if v == nil {
			t.Error("Expected bar but got nil")
		}
	})
	t.Run("Scan config", func(t *testing.T) {
		var v string
		err = Scan("foo", &v)
		if err != nil {
			t.Error(err)
		}
		if v != "bar" {
			t.Errorf("Expected bar but got %v", v)
		}
	})
	t.Run("Config name", func(t *testing.T) {
		s := String()
		if s != "Viper config" {
			t.Errorf("Expected Viper config but got %v", s)
		}
	})
}
