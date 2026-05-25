package architecture_test

import (
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestDomainDoesNotImportTransportOrInfrastructure(t *testing.T) {
	t.Parallel()

	forbidden := []string{
		"net/http",
		"database/sql",
		"github.com/gin-gonic/gin",
		"github.com/labstack/echo",
		"github.com/gofiber/fiber/v2",
		"/{{ cookiecutter.module_path }}/internal/transport/",
		"/{{ cookiecutter.module_path }}/internal/repository/",
		"/{{ cookiecutter.module_path }}/internal/platform/",
	}

	assertNoForbiddenImports(t, filepath.Join("..", "domain"), forbidden)
}

func TestServiceDoesNotImportTransport(t *testing.T) {
	t.Parallel()

	forbidden := []string{
		"net/http",
		"github.com/gin-gonic/gin",
		"github.com/labstack/echo",
		"github.com/gofiber/fiber/v2",
		"/{{ cookiecutter.module_path }}/internal/transport/",
	}

	assertNoForbiddenImports(t, filepath.Join("..", "service"), forbidden)
}

func checkFileImports(t *testing.T, path string, forbidden []string) error {
	t.Helper()

	file, err := parser.ParseFile(token.NewFileSet(), path, nil, parser.ImportsOnly)
	if err != nil {
		return err
	}

	for _, imported := range file.Imports {
		importPath := strings.Trim(imported.Path.Value, "\"")
		for _, blocked := range forbidden {
			if importPath == blocked || strings.Contains(importPath, blocked) {
				t.Fatalf("%s imports forbidden package %s", path, importPath)
			}
		}
	}

	return nil
}

func assertNoForbiddenImports(t *testing.T, root string, forbidden []string) {
	t.Helper()

	err := filepath.WalkDir(root, func(path string, d os.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}

		if d.IsDir() || !strings.HasSuffix(path, ".go") || strings.HasSuffix(path, "_test.go") {
			return nil
		}

		return checkFileImports(t, path, forbidden)
	})
	if err != nil {
		t.Fatal(err)
	}
}
