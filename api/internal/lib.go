package internal

import (
	"encoding/json"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path"
	"strings"

	"golang.org/x/mod/modfile"
	"golang.org/x/mod/module"
)

func GenerateFromPackage(artifactPath string) {
	projName := readProjectName()
	projFiles := readProject()

	projRepr := make(map[string]ProjectFile)

	for fileName, ast := range projFiles {
		imports := make([]string, 0)
		externalImports := make([]string, 0)
		qualifiedPkgName := projName + "/" + strings.TrimSuffix(fileName, ".go")

    // remove last part
    for qualifiedPkgName[len(qualifiedPkgName)-1] != '/' {
      qualifiedPkgName = qualifiedPkgName[:len(qualifiedPkgName)-1]
    }
    if qualifiedPkgName[len(qualifiedPkgName)-1] != '/' {
      panic("qualifiedPkgName should end with /")
    }
    qualifiedPkgName = path.Clean(qualifiedPkgName[0 : len(qualifiedPkgName)-1])

		for _, impo := range ast.Imports {
			importPath := strings.Trim(impo.Path.Value, "\"")
			if strings.HasPrefix(importPath, projName) {
				escaped, err := module.EscapePath(importPath)
				if err != nil {
					panic(err)
				}
				imports = append(imports, escaped)
			} else {
				externalImports = append(externalImports, importPath)
			}
		}

		projRepr[qualifiedPkgName] = ProjectFile{
			PackageName:          ast.Name.Name,
			QualifiedPackageName: qualifiedPkgName,
			FilePath:             fileName,
			Imports:              imports,
			ExternalImports:      externalImports,
		}
	}

	writeArtifact(artifactPath, projRepr)
}

func writeArtifact(artifactPath string, projRepr map[string]ProjectFile) {
	bytes, err := json.Marshal(projRepr)
	if err != nil {
		panic(err)
	}

	err = os.WriteFile(artifactPath, bytes, os.ModePerm)
	if err != nil {
		panic(err)
	}

	println("Artifact written to " + artifactPath)
}

func readProjectName() string {
	if _, err := os.Stat("go.mod"); err != nil {
		panic("go.mod not found")
	}

	data, err := os.ReadFile("go.mod")
	if err != nil {
		panic(err)
	}

	f, err := modfile.Parse("go.mod", data, nil)
	if err != nil {
		panic(err)
	}

	return f.Module.Mod.Path
}

func readProject() map[string]*ast.File {
	dirs := []string{"."}
	asts := make(map[string]*ast.File)

	for len(dirs) > 0 {
		dir := dirs[0]
		dirs = dirs[1:]

		files, err := os.ReadDir(dir)
		if err != nil {
			panic(err)
		}

		for _, f := range files {
			if f.IsDir() {
				newDir := dir + "/" + f.Name()
				dirs = append(dirs, newDir)
				continue
			}

			if !strings.HasSuffix(f.Name(), ".go") {
				continue
			}

			filename := dir + "/" + f.Name()
			parsed, err := parser.ParseFile(token.NewFileSet(), filename, nil, 0)
			if err != nil {
				panic(err)
			}
			asts[filename] = parsed
		}
	}

	return asts
}
