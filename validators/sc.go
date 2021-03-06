package validators

// SC struct - Santa Catarina
// Implements the Validator interface
type SC struct {
}

// IsValid func
func (v SC) IsValid(insc string) bool {

	rule := NewRule()

	if !rule.IsCorrectSize(insc, 9) {
		return false
	}

	return rule.ValidateDefaultRule(insc, 8, 11)
}
