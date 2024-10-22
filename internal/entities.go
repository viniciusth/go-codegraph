package internal

type ProjectFile struct {
    PackageName     string
    FilePath        string
    ModPath         string
    Imports         []string
    ExternalImports []string
}
