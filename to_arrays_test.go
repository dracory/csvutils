package csvutils

import (
	"os"
	"path/filepath"
	"testing"
)

func TestToArrays_Success(t *testing.T) {
	// Create a temporary CSV file
	tmpDir := t.TempDir()
	tmpFile := filepath.Join(tmpDir, "test.csv")

	content := "name,age,city\nAlice,30,NYC\nBob,25,LA\n"
	if err := os.WriteFile(tmpFile, []byte(content), 0644); err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	// Test the function
	result, err := ToArrays(tmpFile)
	if err != nil {
		t.Fatalf("ToArrays failed: %v", err)
	}

	// Verify results
	if len(result) != 3 {
		t.Errorf("Expected 3 rows, got %d", len(result))
	}

	if len(result[0]) != 3 {
		t.Errorf("Expected 3 columns, got %d", len(result[0]))
	}

	if result[0][0] != "name" || result[0][1] != "age" || result[0][2] != "city" {
		t.Errorf("Header mismatch: got %v", result[0])
	}

	if result[1][0] != "Alice" || result[1][1] != "30" || result[1][2] != "NYC" {
		t.Errorf("First row mismatch: got %v", result[1])
	}
}

func TestToArrays_FileNotFound(t *testing.T) {
	_, err := ToArrays("/nonexistent/path/file.csv")
	if err == nil {
		t.Error("Expected error for non-existent file, got nil")
	}
}

func TestToArrays_EmptyFile(t *testing.T) {
	tmpDir := t.TempDir()
	tmpFile := filepath.Join(tmpDir, "empty.csv")

	if err := os.WriteFile(tmpFile, []byte(""), 0644); err != nil {
		t.Fatalf("Failed to create empty file: %v", err)
	}

	result, err := ToArrays(tmpFile)
	if err != nil {
		t.Fatalf("ToArrays failed on empty file: %v", err)
	}

	if len(result) != 0 {
		t.Errorf("Expected empty result, got %d rows", len(result))
	}
}

func TestToArrays_QuotedFields(t *testing.T) {
	tmpDir := t.TempDir()
	tmpFile := filepath.Join(tmpDir, "quoted.csv")

	content := "name,description\n\"Alice Smith\",\"A person, who codes\"\n"
	if err := os.WriteFile(tmpFile, []byte(content), 0644); err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}

	result, err := ToArrays(tmpFile)
	if err != nil {
		t.Fatalf("ToArrays failed: %v", err)
	}

	if len(result) != 2 {
		t.Errorf("Expected 2 rows, got %d", len(result))
	}

	if result[1][1] != "A person, who codes" {
		t.Errorf("Quoted field mismatch: got %v", result[1][1])
	}
}
