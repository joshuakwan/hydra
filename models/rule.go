package models

const (
	ruleRegistryName = "/rules/"
)

// Then defines a successful evaluation action
type Then struct {
	Run        string
	Parameters map[string]string
}

// Expression defines an expression
type Expression struct {
	If   string
	Then *Then
}

// Rule defines a rule
type Rule struct {
	Module      string
	Name        string
	Description string
	Enabled     bool
	Expression  *Expression
}

// func CreateRule(rule Rule) error {
// 	client, destroyFunc, err := storage.CreateClient()
// 	if err != nil {
// 		return err
// 	}
// 	return err
// }
