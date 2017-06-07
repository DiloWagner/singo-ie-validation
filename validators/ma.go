package validators

// MA struct - Maranhão
// Implements the Validator interface
type MA struct {
}

// IsValid func
func (v MA) IsValid(insc string) bool {

	rule := NewRule()

	if !rule.IsCorrectSize(insc, 9) {
		return false
	}

	if !rule.IsStartWith(insc, "12") {
		return false
	}

	return rule.ValidateDefaultRule(insc, 8, 11)
}
