package csvutils

import (
	"os"
	"path/filepath"
	"testing"
)

func TestToMaps_Success(t *testing.T) {
	// Create a temporary CSV file
	tmpDir := t.TempDir()
	tmpFile := filepath.Join(tmpDir, "test.csv")

	content := "name,age,city\nAlice,30,NYC\nBob,25,LA\n"
	if err := writeTestFile(tmpFile, content); err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	// Test the function
	result, err := ToMaps(tmpFile, nil)
	if err != nil {
		t.Fatalf("ToMaps failed: %v", err)
	}

	// Verify results
	if len(result) != 2 {
		t.Errorf("Expected 2 data rows, got %d", len(result))
	}

	// Check first row
	if result[0]["name"] != "Alice" {
		t.Errorf("Expected name 'Alice', got '%s'", result[0]["name"])
	}
	if result[0]["age"] != "30" {
		t.Errorf("Expected age '30', got '%s'", result[0]["age"])
	}
	if result[0]["city"] != "NYC" {
		t.Errorf("Expected city 'NYC', got '%s'", result[0]["city"])
	}

	// Check second row
	if result[1]["name"] != "Bob" {
		t.Errorf("Expected name 'Bob', got '%s'", result[1]["name"])
	}
}

func TestToMaps_WithReplacements(t *testing.T) {
	tmpDir := t.TempDir()
	tmpFile := filepath.Join(tmpDir, "test.csv")

	content := "user name,user age,city\nAlice,30,NYC\n"
	if err := writeTestFile(tmpFile, content); err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	replacements := map[string]string{
		"user name": "username",
		"user age":  "age",
	}

	result, err := ToMaps(tmpFile, replacements)
	if err != nil {
		t.Fatalf("ToMaps failed: %v", err)
	}

	// Check that replacements were applied
	if _, ok := result[0]["username"]; !ok {
		t.Errorf("Expected 'username' key after replacement, got keys: %v", result[0])
	}

	if _, ok := result[0]["user name"]; ok {
		t.Errorf("Old key 'user name' should not exist after replacement")
	}

	if result[0]["username"] != "Alice" {
		t.Errorf("Expected username 'Alice', got '%s'", result[0]["username"])
	}
}

func TestToMaps_HeaderTrimming(t *testing.T) {
	tmpDir := t.TempDir()
	tmpFile := filepath.Join(tmpDir, "test.csv")

	content := " name , age , city \nAlice,30,NYC\n"
	if err := writeTestFile(tmpFile, content); err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	result, err := ToMaps(tmpFile, nil)
	if err != nil {
		t.Fatalf("ToMaps failed: %v", err)
	}

	// Check that headers were trimmed
	if _, ok := result[0]["name"]; !ok {
		t.Errorf("Expected trimmed 'name' key, got keys: %v", result[0])
	}
}

func TestToMaps_OnlyHeader(t *testing.T) {
	tmpDir := t.TempDir()
	tmpFile := filepath.Join(tmpDir, "test.csv")

	content := "name,age,city\n"
	if err := writeTestFile(tmpFile, content); err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	result, err := ToMaps(tmpFile, nil)
	if err != nil {
		t.Fatalf("ToMaps failed: %v", err)
	}

	if len(result) != 0 {
		t.Errorf("Expected empty result for header-only file, got %d rows", len(result))
	}
}

func TestToMaps_EmptyFile(t *testing.T) {
	tmpDir := t.TempDir()
	tmpFile := filepath.Join(tmpDir, "empty.csv")

	if err := writeTestFile(tmpFile, ""); err != nil {
		t.Fatalf("Failed to create empty file: %v", err)
	}

	result, err := ToMaps(tmpFile, nil)
	if err != nil {
		t.Fatalf("ToMaps failed on empty file: %v", err)
	}

	if len(result) != 0 {
		t.Errorf("Expected empty result, got %d rows", len(result))
	}
}

func TestToMaps_FileNotFound(t *testing.T) {
	_, err := ToMaps("/nonexistent/path/file.csv", nil)
	if err == nil {
		t.Error("Expected error for non-existent file, got nil")
	}
}

func TestToMaps_MismatchedColumns(t *testing.T) {
	tmpDir := t.TempDir()
	tmpFile := filepath.Join(tmpDir, "test.csv")

	content := "name,age,city\nAlice,30\nBob,25,LA,Extra\n"
	if err := writeTestFile(tmpFile, content); err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	_, err := ToMaps(tmpFile, nil)
	if err == nil {
		t.Error("Expected error for mismatched columns, got nil")
	}
}

// Helper function to write test files
func writeTestFile(path string, content string) error {
	return os.WriteFile(path, []byte(content), 0644)
}
