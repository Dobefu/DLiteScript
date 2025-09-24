package evaluator

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/Dobefu/DLiteScript/internal/ast"
	"github.com/Dobefu/DLiteScript/internal/datavalue"
	"github.com/Dobefu/DLiteScript/internal/errorutil"
)

func TestEvaluateImportStatement(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    *ast.ImportStatement
		expected datavalue.Value
	}{
		{
			name: "import statement",
			input: &ast.ImportStatement{
				Path: &ast.StringLiteral{
					Value: "../examples/09_imports/test.dl",
					Range: ast.Range{
						Start: ast.Position{Offset: 0, Line: 0, Column: 0},
						End:   ast.Position{Offset: 5, Line: 0, Column: 0},
					},
				},
				Namespace: "test",
				Alias:     "",
				Range: ast.Range{
					Start: ast.Position{Offset: 0, Line: 0, Column: 0},
					End:   ast.Position{Offset: 5, Line: 0, Column: 0},
				},
			},
			expected: datavalue.Null(),
		},
		{
			name: "import statement with alias",
			input: &ast.ImportStatement{
				Path: &ast.StringLiteral{
					Value: "../examples/09_imports/utils.dl",
					Range: ast.Range{
						Start: ast.Position{Offset: 0, Line: 0, Column: 0},
						End:   ast.Position{Offset: 5, Line: 0, Column: 0},
					},
				},
				Namespace: "test",
				Alias:     "test",
				Range: ast.Range{
					Start: ast.Position{Offset: 0, Line: 0, Column: 0},
					End:   ast.Position{Offset: 5, Line: 0, Column: 0},
				},
			},
			expected: datavalue.Null(),
		},
		{
			name: "import statement in current namespace",
			input: &ast.ImportStatement{
				Path: &ast.StringLiteral{
					Value: "../examples/09_imports/utils.dl",
					Range: ast.Range{
						Start: ast.Position{Offset: 0, Line: 0, Column: 0},
						End:   ast.Position{Offset: 5, Line: 0, Column: 0},
					},
				},
				Namespace: "test",
				Alias:     "_",
				Range: ast.Range{
					Start: ast.Position{Offset: 0, Line: 0, Column: 0},
					End:   ast.Position{Offset: 5, Line: 0, Column: 0},
				},
			},
			expected: datavalue.Null(),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			cwd, err := os.Getwd()

			if err != nil {
				t.Fatalf("error getting CWD: %s", err.Error())
			}

			ev := NewEvaluator(io.Discard)
			ev.SetCurrentFilePath(cwd)
			rawResult, err := ev.evaluateImportStatement(test.input)

			if err != nil {
				t.Fatalf("error evaluating %s: %s", test.input.Expr(), err.Error())
			}

			if rawResult.Value.DataType != test.expected.DataType {
				t.Fatalf(
					"expected \"%s\", got \"%s\"",
					test.expected.DataType.AsString(),
					rawResult.Value.DataType.AsString(),
				)
			}
		})
	}
}

func TestEvaluateImportStatementErr(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    *ast.ImportStatement
		content  string
		expected string
	}{
		{
			name: "import statement with invalid path",
			input: &ast.ImportStatement{
				Path: &ast.StringLiteral{
					Value: "bogus",
					Range: ast.Range{
						Start: ast.Position{Offset: 0, Line: 0, Column: 0},
						End:   ast.Position{Offset: 5, Line: 0, Column: 0},
					},
				},
				Namespace: "test",
				Alias:     "",
				Range: ast.Range{
					Start: ast.Position{Offset: 0, Line: 0, Column: 0},
					End:   ast.Position{Offset: 5, Line: 0, Column: 0},
				},
			},
			content:  "",
			expected: "no such file or directory",
		},
		{
			name: "import statement with invalid UTF-8 character in file",
			input: &ast.ImportStatement{
				Path: &ast.StringLiteral{
					Value: "",
					Range: ast.Range{
						Start: ast.Position{Offset: 0, Line: 0, Column: 0},
						End:   ast.Position{Offset: 5, Line: 0, Column: 0},
					},
				},
				Namespace: "test",
				Alias:     "",
				Range: ast.Range{
					Start: ast.Position{Offset: 0, Line: 0, Column: 0},
					End:   ast.Position{Offset: 5, Line: 0, Column: 0},
				},
			},
			content: "\x80",
			expected: fmt.Sprintf(
				"%s: %s at position 0",
				errorutil.StageTokenize.String(),
				errorutil.ErrorMsgInvalidUTF8Char,
			),
		},
		{
			name: "import statement parse error",
			input: &ast.ImportStatement{
				Path: &ast.StringLiteral{
					Value: "",
					Range: ast.Range{
						Start: ast.Position{Offset: 0, Line: 0, Column: 0},
						End:   ast.Position{Offset: 5, Line: 0, Column: 0},
					},
				},
				Namespace: "test",
				Alias:     "",
				Range: ast.Range{
					Start: ast.Position{Offset: 0, Line: 0, Column: 0},
					End:   ast.Position{Offset: 5, Line: 0, Column: 0},
				},
			},
			content: "func(",
			expected: fmt.Sprintf(
				"%s: %s at position 4",
				errorutil.StageParse.String(),
				fmt.Sprintf(errorutil.ErrorMsgUnexpectedToken, "("),
			),
		},
		{
			name: "import statement evaluate error",
			input: &ast.ImportStatement{
				Path: &ast.StringLiteral{
					Value: "",
					Range: ast.Range{
						Start: ast.Position{Offset: 0, Line: 0, Column: 0},
						End:   ast.Position{Offset: 5, Line: 0, Column: 0},
					},
				},
				Namespace: "test",
				Alias:     "_",
				Range: ast.Range{
					Start: ast.Position{Offset: 0, Line: 0, Column: 0},
					End:   ast.Position{Offset: 5, Line: 0, Column: 0},
				},
			},
			content: "_",
			expected: fmt.Sprintf(
				"%s: %s at position 0",
				errorutil.StageEvaluate.String(),
				fmt.Sprintf(errorutil.ErrorMsgUndefinedIdentifier, "_"),
			),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			currentFilePath := ""

			if test.content != "" {
				tempFile, err := os.CreateTemp("", "")

				if err != nil {
					t.Fatalf("error creating temp file: %s", err.Error())
				}

				defer func() { _ = os.Remove(tempFile.Name()) }()

				_, _ = tempFile.WriteString(test.content)
				_ = tempFile.Close()

				test.input.Path.Value = filepath.Base(tempFile.Name())
				currentFilePath = tempFile.Name()
			}

			ev := NewEvaluator(io.Discard)
			ev.SetCurrentFilePath(currentFilePath)
			_, err := ev.evaluateImportStatement(test.input)

			if err == nil {
				t.Fatalf("expected error, got none")
			}

			if !strings.Contains(err.Error(), test.expected) {
				t.Errorf(
					"expected error to contain \"%s\", got \"%s\"",
					test.expected,
					err.Error(),
				)
			}
		})
	}
}
