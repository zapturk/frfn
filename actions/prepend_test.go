package actions_test

import (
	"bytes"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/zapturk/frfn/actions"
)

func TestPrependFileName(t *testing.T) {
	// Helper function to create a temporary directory and files for testing
	createTestDir := func(t *testing.T, files []string) string {
		t.Helper()
		dir, err := os.MkdirTemp("", "testdir")
		if err != nil {
			t.Fatalf("Failed to create temp dir: %v", err)
		}
		for _, file := range files {
			if err := os.WriteFile(filepath.Join(dir, file), []byte("dummy content"), 0o666); err != nil {
				t.Fatalf("Failed to create dummy file: %v", err)
			}
		}
		return dir
	}

	// Helper function to mock stdin
	mockStdin := func(t *testing.T, input string) func() {
		t.Helper()
		oldStdin := os.Stdin
		r, w, err := os.Pipe()
		if err != nil {
			t.Fatalf("Failed to create pipe: %v", err)
		}
		os.Stdin = r
		go func() {
			defer w.Close()
			w.WriteString(input + "\n")
		}()
		return func() {
			os.Stdin = oldStdin
		}
	}

	t.Run("force prepend", func(t *testing.T) {
		// Test setup
		initialFiles := []string{"foo.txt", "test.txt"}
		expectedFiles := []string{"bar_foo.txt", "bar_test.txt"}
		testDir := createTestDir(t, initialFiles)
		defer os.RemoveAll(testDir)

		originalDir, err := os.Getwd()
		if err != nil {
			t.Fatalf("Failed to get current directory: %v", err)
		}
		if err := os.Chdir(testDir); err != nil {
			t.Fatalf("Failed to change directory: %v", err)
		}
		defer os.Chdir(originalDir)

		oldStdout := os.Stdout
		r, w, _ := os.Pipe()
		os.Stdout = w

		// Action
		err = actions.PrependFileName("bar_", true)

		// Assertions
		w.Close()
		var buf bytes.Buffer
		buf.ReadFrom(r)
		os.Stdout = oldStdout
		output := buf.String()

		if err != nil {
			t.Errorf("PrependFileName() error = %v, wantErr %v", err, false)
		}

		files, err := ioutil.ReadDir(".")
		if err != nil {
			t.Fatalf("Failed to read dir: %v", err)
		}

		if len(files) != len(expectedFiles) {
			t.Errorf("Expected %d files, but got %d", len(expectedFiles), len(files))
		}

		foundFiles := make(map[string]bool)
		for _, f := range files {
			foundFiles[f.Name()] = true
		}

		for _, expectedFile := range expectedFiles {
			if !foundFiles[expectedFile] {
				t.Errorf("Expected file %s not found", expectedFile)
			}
		}

		expectedMsg := "foo.txt was changed to bar_foo.txt"
		if !strings.Contains(output, expectedMsg) {
			t.Errorf("Expected output to contain %q, but got %q", expectedMsg, output)
		}
	})

	t.Run("confirm with y", func(t *testing.T) {
		// Test setup
		initialFiles := []string{"foo.txt", "test.txt"}
		expectedFiles := []string{"bar_foo.txt", "bar_test.txt"}
		testDir := createTestDir(t, initialFiles)
		defer os.RemoveAll(testDir)

		originalDir, err := os.Getwd()
		if err != nil {
			t.Fatalf("Failed to get current directory: %v", err)
		}
		if err := os.Chdir(testDir); err != nil {
			t.Fatalf("Failed to change directory: %v", err)
		}
		defer os.Chdir(originalDir)

		restoreStdin := mockStdin(t, "y")
		defer restoreStdin()

		oldStdout := os.Stdout
		r, w, _ := os.Pipe()
		os.Stdout = w

		// Action
		err = actions.PrependFileName("bar_", false)

		// Assertions
		w.Close()
		var buf bytes.Buffer
		buf.ReadFrom(r)
		os.Stdout = oldStdout
		output := buf.String()

		if err != nil {
			t.Errorf("PrependFileName() error = %v, wantErr %v", err, false)
		}

		files, err := ioutil.ReadDir(".")
		if err != nil {
			t.Fatalf("Failed to read dir: %v", err)
		}

		if len(files) != len(expectedFiles) {
			t.Errorf("Expected %d files, but got %d", len(expectedFiles), len(files))
		}

		foundFiles := make(map[string]bool)
		for _, f := range files {
			foundFiles[f.Name()] = true
		}

		for _, expectedFile := range expectedFiles {
			if !foundFiles[expectedFile] {
				t.Errorf("Expected file %s not found", expectedFile)
			}
		}

		expectedMsg := "foo.txt will change to bar_foo.txt"
		if !strings.Contains(output, expectedMsg) {
			t.Errorf("Expected output to contain %q, but got %q", expectedMsg, output)
		}
	})

	t.Run("confirm with yes", func(t *testing.T) {
		// Test setup
		initialFiles := []string{"foo.txt", "test.txt"}
		expectedFiles := []string{"bar_foo.txt", "bar_test.txt"}
		testDir := createTestDir(t, initialFiles)
		defer os.RemoveAll(testDir)

		originalDir, err := os.Getwd()
		if err != nil {
			t.Fatalf("Failed to get current directory: %v", err)
		}
		if err := os.Chdir(testDir); err != nil {
			t.Fatalf("Failed to change directory: %v", err)
		}
		defer os.Chdir(originalDir)

		restoreStdin := mockStdin(t, "yes")
		defer restoreStdin()

		oldStdout := os.Stdout
		r, w, _ := os.Pipe()
		os.Stdout = w

		// Action
		err = actions.PrependFileName("bar_", false)

		// Assertions
		w.Close()
		var buf bytes.Buffer
		buf.ReadFrom(r)
		os.Stdout = oldStdout
		output := buf.String()

		if err != nil {
			t.Errorf("PrependFileName() error = %v, wantErr %v", err, false)
		}

		files, err := ioutil.ReadDir(".")
		if err != nil {
			t.Fatalf("Failed to read dir: %v", err)
		}

		if len(files) != len(expectedFiles) {
			t.Errorf("Expected %d files, but got %d", len(expectedFiles), len(files))
		}

		foundFiles := make(map[string]bool)
		for _, f := range files {
			foundFiles[f.Name()] = true
		}

		for _, expectedFile := range expectedFiles {
			if !foundFiles[expectedFile] {
				t.Errorf("Expected file %s not found", expectedFile)
			}
		}

		expectedMsg := "foo.txt will change to bar_foo.txt"
		if !strings.Contains(output, expectedMsg) {
			t.Errorf("Expected output to contain %q, but got %q", expectedMsg, output)
		}
	})

	t.Run("confirm with enter", func(t *testing.T) {
		// Test setup
		initialFiles := []string{"foo.txt", "test.txt"}
		expectedFiles := []string{"bar_foo.txt", "bar_test.txt"}
		testDir := createTestDir(t, initialFiles)
		defer os.RemoveAll(testDir)

		originalDir, err := os.Getwd()
		if err != nil {
			t.Fatalf("Failed to get current directory: %v", err)
		}
		if err := os.Chdir(testDir); err != nil {
			t.Fatalf("Failed to change directory: %v", err)
		}
		defer os.Chdir(originalDir)

		restoreStdin := mockStdin(t, "")
		defer restoreStdin()

		oldStdout := os.Stdout
		r, w, _ := os.Pipe()
		os.Stdout = w

		// Action
		err = actions.PrependFileName("bar_", false)

		// Assertions
		w.Close()
		var buf bytes.Buffer
		buf.ReadFrom(r)
		os.Stdout = oldStdout
		output := buf.String()

		if err != nil {
			t.Errorf("PrependFileName() error = %v, wantErr %v", err, false)
		}

		files, err := ioutil.ReadDir(".")
		if err != nil {
			t.Fatalf("Failed to read dir: %v", err)
		}

		if len(files) != len(expectedFiles) {
			t.Errorf("Expected %d files, but got %d", len(expectedFiles), len(files))
		}

		foundFiles := make(map[string]bool)
		for _, f := range files {
			foundFiles[f.Name()] = true
		}

		for _, expectedFile := range expectedFiles {
			if !foundFiles[expectedFile] {
				t.Errorf("Expected file %s not found", expectedFile)
			}
		}

		expectedMsg := "foo.txt will change to bar_foo.txt"
		if !strings.Contains(output, expectedMsg) {
			t.Errorf("Expected output to contain %q, but got %q", expectedMsg, output)
		}
	})

	t.Run("deny with n", func(t *testing.T) {
		// Test setup
		initialFiles := []string{"foo.txt", "test.txt"}
		expectedFiles := []string{"foo.txt", "test.txt"}
		testDir := createTestDir(t, initialFiles)
		defer os.RemoveAll(testDir)

		originalDir, err := os.Getwd()
		if err != nil {
			t.Fatalf("Failed to get current directory: %v", err)
		}
		if err := os.Chdir(testDir); err != nil {
			t.Fatalf("Failed to change directory: %v", err)
		}
		defer os.Chdir(originalDir)

		restoreStdin := mockStdin(t, "n")
		defer restoreStdin()

		oldStdout := os.Stdout
		r, w, _ := os.Pipe()
		os.Stdout = w

		// Action
		err = actions.PrependFileName("bar_", false)

		// Assertions
		w.Close()
		var buf bytes.Buffer
		buf.ReadFrom(r)
		os.Stdout = oldStdout
		output := buf.String()

		if err != nil {
			t.Errorf("PrependFileName() error = %v, wantErr %v", err, false)
		}

		files, err := ioutil.ReadDir(".")
		if err != nil {
			t.Fatalf("Failed to read dir: %v", err)
		}

		if len(files) != len(expectedFiles) {
			t.Errorf("Expected %d files, but got %d", len(expectedFiles), len(files))
		}

		foundFiles := make(map[string]bool)
		for _, f := range files {
			foundFiles[f.Name()] = true
		}

		for _, expectedFile := range expectedFiles {
			if !foundFiles[expectedFile] {
				t.Errorf("Expected file %s not found", expectedFile)
			}
		}

		expectedMsg := "No change was made"
		if !strings.Contains(output, expectedMsg) {
			t.Errorf("Expected output to contain %q, but got %q", expectedMsg, output)
		}
	})
}
