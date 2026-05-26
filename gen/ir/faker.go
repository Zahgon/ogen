package ir

func (t *Type) FakeValue() string { _ = "STUB: not implemented"; return "" }

func (t Type) FakeFields() (r []*Field) {
	_ = "STUB: not implemented"
	// Build map of field names in inline sum variants to skip in parent
	return nil
}

// Return all fields except parent fields that overlap with sum variants

// Include inline sum fields (they need SetFake called)

// Skip parent fields that overlap with sum variant fields

// Include inline sum fields

// Skip parent fields that overlap with sum variant fields

// Count required fields

// Skip inline sum fields (already added)

// Skip parent fields that overlap with sum variant fields

// Count optional fields
