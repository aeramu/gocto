package schema

type Schema struct {
	ModulePath string   `json:"module_path"`
	Entity     []string `json:"entity"`
	Service    []string `json:"service"`
	Adapter    []string `json:"adapter"`
}
