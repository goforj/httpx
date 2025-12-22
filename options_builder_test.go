package httpx

import (
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestOptionFunctionsMatchBuilderMethods(t *testing.T) {
	root, err := findRepoRoot()
	if err != nil {
		t.Fatalf("repo root: %v", err)
	}

	fset := token.NewFileSet()
	pkgs, err := parser.ParseDir(fset, root, func(info os.FileInfo) bool {
		name := info.Name()
		return strings.HasPrefix(name, "options_") && strings.HasSuffix(name, ".go")
	}, 0)
	if err != nil {
		t.Fatalf("parse options: %v", err)
	}

	var pkg *ast.Package
	for _, p := range pkgs {
		if p.Name == "httpx" {
			pkg = p
			break
		}
	}
	if pkg == nil {
		t.Fatal("package httpx not found")
	}

	methods := map[string]struct{}{}
	functions := map[string]struct{}{}

	for _, file := range pkg.Files {
		for _, decl := range file.Decls {
			fn, ok := decl.(*ast.FuncDecl)
			if !ok || fn.Name == nil {
				continue
			}
			if fn.Recv != nil && len(fn.Recv.List) > 0 {
				if receiverType(fn.Recv.List[0].Type) == "OptionBuilder" {
					methods[fn.Name.Name] = struct{}{}
				}
				continue
			}
			if !fn.Name.IsExported() {
				continue
			}
			if isOptionBuilderReturn(fn.Type.Results) {
				functions[fn.Name.Name] = struct{}{}
			}
		}
	}

	var missing []string
	for name := range methods {
		if _, ok := functions[name]; !ok {
			missing = append(missing, name)
		}
	}
	for name := range functions {
		if _, ok := methods[name]; !ok {
			missing = append(missing, name)
		}
	}

	if len(missing) > 0 {
		t.Fatalf("option mismatch between functions and builder methods: %v", missing)
	}
}

func receiverType(expr ast.Expr) string {
	switch v := expr.(type) {
	case *ast.Ident:
		return v.Name
	case *ast.StarExpr:
		return receiverType(v.X)
	case *ast.IndexExpr:
		return receiverType(v.X)
	case *ast.IndexListExpr:
		return receiverType(v.X)
	default:
		return ""
	}
}

func isOptionBuilderReturn(results *ast.FieldList) bool {
	if results == nil || len(results.List) == 0 {
		return false
	}
	last := results.List[len(results.List)-1]
	if named, ok := last.Type.(*ast.Ident); ok && named.Name == "OptionBuilder" {
		return true
	}
	if sel, ok := last.Type.(*ast.SelectorExpr); ok {
		return sel.Sel.Name == "OptionBuilder"
	}
	return false
}

func findRepoRoot() (string, error) {
	wd, err := os.Getwd()
	if err != nil {
		return "", err
	}
	for {
		if _, err := os.Stat(filepath.Join(wd, "go.mod")); err == nil {
			return wd, nil
		}
		parent := filepath.Dir(wd)
		if parent == wd {
			return "", os.ErrNotExist
		}
		wd = parent
	}
}
