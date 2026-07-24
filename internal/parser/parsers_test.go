package parser

import (
	"bytes"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/yuin/goldmark"
)

func TestGetMetadataMdFile(t *testing.T) {
	tests := []struct {
		name        string
		content     string
		expected    map[string]any
		expectError bool
	}{
		{
			name: "valid frontmatter with two separators",
			content: `---
title: My Page
author: John Doe
---
# Content
This is the body`,
			expected: map[string]any{
				"title":  "My Page",
				"author": "John Doe",
			},
			expectError: false,
		},
		{
			name: "frontmatter with extra whitespace",
			content: `---
key1 : value1  
key2:   value2   
---
body`,
			expected: map[string]any{
				"key1": "value1",
				"key2": "value2",
			},
			expectError: false,
		},
		{
			name: "single separator (unclosed frontmatter) - reads until EOF",
			content: `---
key: value
another: data`,
			expected: map[string]any{
				"key":     "value",
				"another": "data",
			},
			expectError: false,
		},
		{
			name:        "no separator at all",
			content:     `just plain text without any ---`,
			expected:    map[string]any{},
			expectError: false,
		},
		{
			name: "multiple separators (stops at second)",
			content: `---
a: 1
---
b: 2
---
c: 3`,
			expected: map[string]any{
				"a": "1",
			},
			expectError: false,
		},
		{
			name:        "empty file",
			content:     "",
			expected:    map[string]any{},
			expectError: false,
		},
		{
			name: "only separators",
			content: `---
---`,
			expected:    map[string]any{},
			expectError: false,
		},
		{
			name:        "file does not exist",
			content:     "", // не используется, т.к. файл не создаётся
			expected:    map[string]any{},
			expectError: true,
		},
		{
			name: "line without colon",
			content: `---
invalid line
---
`,
			expected:    map[string]any{},
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var filePath string
			if tt.name == "file does not exist" {
				filePath = "/non/existent/file/xyz"
			} else {

				tmpFile, err := os.CreateTemp(t.TempDir(), "test_*.md")
				if err != nil {
					t.Fatalf("failed to create temp file: %v", err)
				}
				defer os.Remove(tmpFile.Name())

				if _, err := tmpFile.WriteString(tt.content); err != nil {
					t.Fatalf("failed to write to temp file: %v", err)
				}
				tmpFile.Close()

				filePath = tmpFile.Name()
			}

			got, err := GetMetadataMdFile(filePath)

			if tt.expectError {
				if err == nil {
					t.Errorf("expected error, got nil")
				}
				// Не проверяем результат при ошибке (обычно пустой map)
				return
			}
			if err != nil {
				t.Errorf("unexpected error: %v", err)
				return
			}

			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("expected %v, got %v", tt.expected, got)
			}
		})
	}
}

func TestGetHTMLContent(t *testing.T) {
	tests := []struct {
		name          string
		setupDir      func(t *testing.T) string
		expectedCount int
		expectError   bool
		content       string
		expectContent string
		isFiles       bool
	}{
		{
			name: "success load 1 file",
			content: `---
title: Test
---
# Heading
Paragraph`,
			expectContent: `# Heading
Paragraph`,
			setupDir: func(t *testing.T) string {
				tmpDir := t.TempDir()
				if err := os.WriteFile(filepath.Join(tmpDir, "file.txt"), []byte("text"), 0644); err != nil {
					t.Fatal(err)
				}
				return tmpDir
			},
			expectError:   false,
			expectedCount: 1,
		},
		{
			name:          "load 1 file without metadata",
			content:       `Just plain text`,
			expectContent: `Just plain text`,
			setupDir: func(t *testing.T) string {
				tmpDir := t.TempDir()
				if err := os.WriteFile(filepath.Join(tmpDir, "file.txt"), []byte("text"), 0644); err != nil {
					t.Fatal(err)
				}
				return tmpDir
			},
			expectedCount: 1,
			expectError:   false,
		},
		{
			name: "load 2 files",
			content: `---
title: Test
---
# Heading
Paragraph`,
			expectContent: `# Heading
Paragraph`,
			setupDir: func(t *testing.T) string {
				tmpDir := t.TempDir()
				if err := os.WriteFile(filepath.Join(tmpDir, "file.txt"), []byte("text"), 0644); err != nil {
					t.Fatal(err)
				}
				return tmpDir
			},
			expectedCount: 2,
			isFiles:       true,
			expectError:   false,
		},
		{
			name: "directory with only non-md files",
			setupDir: func(t *testing.T) string {
				tmpDir := t.TempDir()
				if err := os.WriteFile(filepath.Join(tmpDir, "file.txt"), []byte("text"), 0644); err != nil {
					t.Fatal(err)
				}
				return tmpDir
			},
			expectedCount: 0,
			expectError:   false,
			content:       "",
			expectContent: "",
		},
		{
			name: "non-existent directory",
			setupDir: func(t *testing.T) string {
				return filepath.Join(t.TempDir(), "non-existent")
			},
			expectedCount: 0,
			expectError:   true,
			content:       "",
			expectContent: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dir := tt.setupDir(t)
			if tt.isFiles && len(tt.content) >= 1 {
				if err := os.WriteFile(filepath.Join(dir, "file1.md"), []byte(tt.content), 0644); err != nil {
					t.Fatal(err)
				}
				if err := os.WriteFile(filepath.Join(dir, "file2.md"), []byte(tt.content), 0644); err != nil {
					t.Fatal(err)
				}
			} else {
				if len(tt.content) >= 1 {
					if err := os.WriteFile(filepath.Join(dir, "file1.md"), []byte(tt.content), 0644); err != nil {
						t.Fatal(err)
					}
				}
			}

			result, err := GetHTMLContent(dir)

			if tt.expectError {
				if err == nil {
					t.Errorf("expected error, got nil")
				}
				return
			}
			if err != nil {
				t.Errorf("unexpected error: %v", err)
				return
			}

			if len(result) != tt.expectedCount {
				t.Errorf("expected %d files, got %d", tt.expectedCount, len(result))
			}
			for _, res := range result {
				var trueData bytes.Buffer

				if err := goldmark.Convert([]byte(tt.expectContent), &trueData); err != nil {
					t.Errorf("err convert data")
				}

				if !bytes.Equal(res.Data.Bytes(), trueData.Bytes()) {
					t.Errorf("err is not equal %s and %s", res.Data.String(), trueData.String())
				}
			}
		})
	}
}

func TestGetMarkDownFiles(t *testing.T) {
	tests := []struct {
		name        string
		setupDir    func(t *testing.T) string
		expected    []string // имена файлов (только базовые имена)
		expectError bool
	}{
		{
			name: "find .md files",
			setupDir: func(t *testing.T) string {
				tmp := t.TempDir()
				os.WriteFile(filepath.Join(tmp, "a.md"), []byte(""), 0644)
				os.WriteFile(filepath.Join(tmp, "b.md"), []byte(""), 0644)
				os.WriteFile(filepath.Join(tmp, "c.txt"), []byte(""), 0644)
				os.Mkdir(filepath.Join(tmp, "sub"), 0755)
				os.WriteFile(filepath.Join(tmp, "sub", "d.md"), []byte(""), 0644)
				return tmp
			},
			expected:    []string{"a.md", "b.md", "d.md"},
			expectError: false,
		},
		{
			name: "no .md files",
			setupDir: func(t *testing.T) string {
				tmp := t.TempDir()
				os.WriteFile(filepath.Join(tmp, "a.txt"), []byte(""), 0644)
				return tmp
			},
			expected:    []string{},
			expectError: false,
		},
		{
			name: "non-existent directory",
			setupDir: func(t *testing.T) string {
				return filepath.Join(t.TempDir(), "missing")
			},
			expected:    nil,
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dir := tt.setupDir(t)

			files, err := GetMarkDownFiles(dir)

			if tt.expectError {
				if err == nil {
					t.Errorf("expected error, got nil")
				}
				return
			}
			if err != nil {
				t.Errorf("unexpected error: %v", err)
				return
			}

			baseNames := make([]string, len(files))
			for i, f := range files {
				baseNames[i] = filepath.Base(f)
			}

			expectedMap := make(map[string]bool)
			for _, e := range tt.expected {
				expectedMap[e] = true
			}
			for _, name := range baseNames {
				if !expectedMap[name] {
					t.Errorf("unexpected file %s", name)
				}
				delete(expectedMap, name)
			}
			if len(expectedMap) > 0 {
				t.Errorf("missing files: %v", expectedMap)
			}
		})
	}
}

func TestInit(t *testing.T) {
	tests := []struct {
		name         string
		setupDir     func(t *testing.T) string
		expectNonNil bool
	}{
		{
			name: "valid templates directory",
			setupDir: func(t *testing.T) string {
				dir := t.TempDir()
				createTestTemplate(t, dir, "base.html", "{{define \"base\"}}Base{{end}}")
				createTestTemplate(t, dir, "page.html", "{{define \"page\"}}Page{{end}}")
				return dir
			},
			expectNonNil: true,
		},
		{
			name: "empty directory (no .html files)",
			setupDir: func(t *testing.T) string {
				return t.TempDir()
			},
			expectNonNil: false,
		},
		{
			name: "non-existent directory",
			setupDir: func(t *testing.T) string {
				return filepath.Join(t.TempDir(), "non-existent")
			},
			expectNonNil: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dir := tt.setupDir(t)
			g := Init(dir)

			if g == nil {
				t.Errorf("Init returned nil")
			}
			if (g.Templates != nil) != tt.expectNonNil {
				t.Errorf("expected Templates non-nil = %v, got %v", tt.expectNonNil, g.Templates != nil)
			}
		})
	}
}

func TestLoadTemplates(t *testing.T) {
	tests := []struct {
		name      string
		setupDir  func(t *testing.T) string
		expectNil bool
		expectLen int // количество определённых шаблонов
	}{
		{
			name: "multiple html templates",
			setupDir: func(t *testing.T) string {
				dir := t.TempDir()
				createTestTemplate(t, dir, "a.html", "")
				createTestTemplate(t, dir, "b.html", "")
				return dir
			},
			expectNil: false,
			expectLen: 2,
		},
		{
			name: "no .html files",
			setupDir: func(t *testing.T) string {
				dir := t.TempDir()
				createTestTemplate(t, dir, "ignore.txt", "text")
				return dir
			},
			expectNil: true,
			expectLen: 0,
		},
		{
			name: "empty directory",
			setupDir: func(t *testing.T) string {
				return t.TempDir()
			},
			expectNil: true,
			expectLen: 0,
		},
		{
			name: "non-existent directory",
			setupDir: func(t *testing.T) string {
				return filepath.Join(t.TempDir(), "missing")
			},
			expectNil: true,
			expectLen: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dir := tt.setupDir(t)
			tmpl := loadTemplates(dir)

			if (tmpl == nil) != tt.expectNil {
				t.Errorf("expected non-nil = %v, got %v", tt.expectNil, tmpl == nil)
			}
			if tmpl != nil {
				defs := tmpl.Templates()
				if len(defs) != tt.expectLen {
					t.Errorf("expected %d defined templates, got %d", tt.expectLen, len(defs))
				}
			}
		})
	}
}

func TestGetTemplate(t *testing.T) {
	dir := t.TempDir()
	createTestTemplate(t, dir, "base.html", "")
	createTestTemplate(t, dir, "page.html", "")

	g := Init(dir)
	if g == nil || g.Templates == nil {
		t.Fatalf("Init failed: g or Templates is nil")
	}

	tests := []struct {
		name         string
		templateName string
		expectFound  bool
	}{
		{
			name:         "existing template",
			templateName: "base.html",
			expectFound:  true,
		},
		{
			name:         "another existing template",
			templateName: "page.html",
			expectFound:  true,
		},
		{
			name:         "non-existing template",
			templateName: "nonexistent.html",
			expectFound:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tmpl := g.GetTemplate(tt.templateName)
			if !(tmpl == nil) != tt.expectFound {
				t.Errorf("expected found = %v, got %v", tt.expectFound, tmpl != nil)
			}
			if tmpl != nil {
				if tmpl.Name() != tt.templateName {
					t.Errorf("expected template name %q, got %q", tt.templateName, tmpl.Name())
				}
			}
		})
	}

	t.Run("nil Templates", func(t *testing.T) {
		gNil := &GlobalTemplate{Templates: nil}
		tmpl := gNil.GetTemplate("any")
		if tmpl != nil {
			t.Errorf("expected nil, got %v", tmpl)
		}
	})
}

func createTestTemplate(t *testing.T, dir, name, content string) {
	path := filepath.Join(dir, name)
	if err := os.WriteFile(path, []byte(content), 0644); err != nil {
		t.Fatal(err)
	}
}
