package internal

type ProjectFile struct {
  PackageName          string `json:"packageName"`
  QualifiedPackageName string `json:"qualifiedPackageName"`
  FilePath             string `json:"filePath"`
  ModPath              string `json:"modPath"`
  Imports              []string `json:"imports"`
  ExternalImports      []string `json:"externalImports"`
}
