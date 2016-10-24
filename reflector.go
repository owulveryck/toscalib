package toscalib

func reflectAssignmentProps(src map[string]PropertyAssignment, dest map[string]AttributeAssignment) *map[string]AttributeAssignment {
	for name, def := range src {
		if len(dest) == 0 {
			dest = make(map[string]AttributeAssignment)
		}

		_, ok := dest[name]
		if !ok {
			a := new(AttributeAssignment)
			a.Value = def.Value
			a.Function = def.Function
			a.Args = def.Args
			a.Expression = def.Expression
			dest[name] = *a // dereference pointer
		}
	}
	return &dest
}

func reflectDefinitionProps(src map[string]PropertyDefinition, dest map[string]AttributeDefinition) *map[string]AttributeDefinition {
	for name, def := range src {
		if len(dest) == 0 {
			dest = make(map[string]AttributeDefinition)
		}

		_, ok := dest[name]
		if !ok {
			a := new(AttributeDefinition)
			a.Type = def.Type
			a.Description = def.Description
			a.Default = def.Default
			a.Status = def.Status
			a.EntrySchema = def.EntrySchema
			dest[name] = *a // dereference pointer
		}
	}
	return &dest
}
