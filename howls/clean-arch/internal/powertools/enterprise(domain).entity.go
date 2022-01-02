// Package PowerTools
// Layer: Enterprise Business
// Entities encapsulate Enterprise business rules.
// They are the least likely to change
// when something external changes.
package PowerTools

type Code string
type PowerTool struct {
	Code Code
	Type string
	Rate int
}

func (t PowerTool) Validate() error {
	// TODO
	return nil
}

func NewPowerTool(code Code, toolType string, rate int) *PowerTool {
	return &PowerTool{
		Code: code,
		Type: toolType,
		Rate: rate,
	}
}
