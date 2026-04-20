package main

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"mongoConnector/models"
)

func TestLoadConfig_ValidFile(t *testing.T) {
	tempDir := t.TempDir()
	configPath := filepath.Join(tempDir, "config.json")

	want := models.Config{
		TestIt:       true,
		URI:          "http://example.com",
		ScreenWidth:  320,
		ScreenHeight: 240,
		TileSize:     16,
	}

	data, err := json.Marshal(want)
	if err != nil {
		t.Fatalf("json.Marshal failed: %v", err)
	}

	if err := os.WriteFile(configPath, data, 0o644); err != nil {
		t.Fatalf("failed to write config file: %v", err)
	}

	got, err := loadConfig(configPath)
	if err != nil {
		t.Fatalf("loadConfig returned error: %v", err)
	}

	if got != want {
		t.Fatalf("expected %+v, got %+v", want, got)
	}
}

func TestLoadConfig_InvalidJSON(t *testing.T) {
	tempDir := t.TempDir()
	configPath := filepath.Join(tempDir, "config.json")

	if err := os.WriteFile(configPath, []byte("{invalid json}"), 0o644); err != nil {
		t.Fatalf("failed to write config file: %v", err)
	}

	if _, err := loadConfig(configPath); err == nil {
		t.Fatal("expected an error for invalid JSON, got nil")
	}
}
