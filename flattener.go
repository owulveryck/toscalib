package toscalib

import "github.com/kenjones-cisco/mergo"

type flatTypes struct {
	Capabilities  map[string]CapabilityType
	Interfaces    map[string]InterfaceType
	Relationships map[string]RelationshipType
	Nodes         map[string]NodeType
	Groups        map[string]GroupType
	Policies      map[string]PolicyType
}

// TODO(kenjones): Switch to reflect if possible in the future (if simple)

func flattenCapType(name string, s ServiceTemplateDefinition) CapabilityType {
	if ct, ok := s.CapabilityTypes[name]; ok {
		if ct.DerivedFrom != "" {
			parent := flattenCapType(ct.DerivedFrom, s)
			// mergo does not handle merging Slices so the rt items
			// will wipe away, capture the values here.
			sources := parent.ValidSources

			_ = mergo.MergeWithOverwrite(&parent, ct)

			// now copy them back in using append, if the child type had
			// any previously, otherwise it will duplicate the parents.
			if len(ct.ValidSources) > 0 {
				parent.ValidSources = append(parent.ValidSources, sources...)
			}
			return parent
		}
		return ct
	}
	return CapabilityType{}
}

func flattenIntfType(name string, s ServiceTemplateDefinition) InterfaceType {
	if it, ok := s.InterfaceTypes[name]; ok {
		if it.DerivedFrom != "" {
			parent := flattenIntfType(it.DerivedFrom, s)
			_ = mergo.MergeWithOverwrite(&parent, it)
			return parent
		}
		return it
	}
	return InterfaceType{}
}

func flattenRelType(name string, s ServiceTemplateDefinition) RelationshipType {
	if rt, ok := s.RelationshipTypes[name]; ok {
		if rt.DerivedFrom != "" {
			parent := flattenRelType(rt.DerivedFrom, s)
			// mergo does not handle merging Slices so the rt items
			// will wipe away, capture the values here.
			targets := parent.ValidTarget

			_ = mergo.MergeWithOverwrite(&parent, rt)

			// now copy them back in using append, if the child type had
			// any previously, otherwise it will duplicate the parents.
			if len(rt.ValidTarget) > 0 {
				parent.ValidTarget = append(parent.ValidTarget, targets...)
			}
			return parent
		}
		return rt
	}
	return RelationshipType{}
}

func flattenNodeType(name string, s ServiceTemplateDefinition) NodeType {
	if nt, ok := s.NodeTypes[name]; ok {
		if nt.DerivedFrom != "" {
			parent := flattenNodeType(nt.DerivedFrom, s)
			// mergo does not handle merging Slices so the nt items
			// will wipe away, capture the values here.
			reqs := parent.Requirements

			_ = mergo.MergeWithOverwrite(&parent, nt)

			// now copy them back in using append, if the child node type had
			// any previously, otherwise it will duplicate the parents.
			if len(nt.Requirements) > 0 {
				parent.Requirements = append(parent.Requirements, reqs...)
			}
			return parent
		}
		return nt
	}
	return NodeType{}
}

func flattenGroupType(name string, s ServiceTemplateDefinition) GroupType {
	if gt, ok := s.GroupTypes[name]; ok {
		if gt.DerivedFrom != "" {
			parent := flattenGroupType(gt.DerivedFrom, s)
			// mergo does not handle merging Slices so the nt items
			// will wipe away, capture the values here.
			reqs := parent.Requirements
			members := parent.Members

			_ = mergo.MergeWithOverwrite(&parent, gt)

			// now copy them back in using append, if the child type had
			// any previously, otherwise it will duplicate the parents.
			if len(gt.Requirements) > 0 {
				parent.Requirements = append(parent.Requirements, reqs...)
			}
			if len(gt.Members) > 0 {
				parent.Members = append(parent.Members, members...)
			}
			return parent
		}
		return gt
	}
	return GroupType{}
}

func flattenPolicyType(name string, s ServiceTemplateDefinition) PolicyType {
	if pt, ok := s.PolicyTypes[name]; ok {
		if pt.DerivedFrom != "" {
			parent := flattenPolicyType(pt.DerivedFrom, s)
			// mergo does not handle merging Slices so the items
			// will wipe away, capture the values here.
			targets := parent.Targets

			_ = mergo.MergeWithOverwrite(&parent, pt)

			// now copy them back in using append, if the child type had
			// any previously, otherwise it will duplicate the parents.
			if len(pt.Targets) > 0 {
				parent.Targets = append(parent.Targets, targets...)
			}
			return parent
		}
		return pt
	}
	return PolicyType{}
}

func flattenHierarchy(s ServiceTemplateDefinition) flatTypes {
	var flats flatTypes

	flats.Capabilities = make(map[string]CapabilityType)
	for name := range s.CapabilityTypes {
		flats.Capabilities[name] = flattenCapType(name, s)
	}

	flats.Interfaces = make(map[string]InterfaceType)
	for name := range s.InterfaceTypes {
		flats.Interfaces[name] = flattenIntfType(name, s)
	}

	flats.Relationships = make(map[string]RelationshipType)
	for name := range s.RelationshipTypes {
		flats.Relationships[name] = flattenRelType(name, s)
	}

	flats.Nodes = make(map[string]NodeType)
	for name := range s.NodeTypes {
		flats.Nodes[name] = flattenNodeType(name, s)
	}

	for k, v := range flats.Nodes {
		for name, capDef := range v.Capabilities {
			capDef.extendFrom(flats.Capabilities[capDef.Type])
			v.Capabilities[name] = capDef
		}
		for name, iDef := range v.Interfaces {
			// during merge the Type is not always properly inherited, so set it
			// from the parent.
			if iDef.Type == "" {
				iDef.Type = flats.Nodes[v.DerivedFrom].Interfaces[name].Type
			}
			iDef.extendFrom(flats.Interfaces[iDef.Type])
			v.Interfaces[name] = iDef
		}
		flats.Nodes[k] = v
	}

	flats.Groups = make(map[string]GroupType)
	for name := range s.GroupTypes {
		flats.Groups[name] = flattenGroupType(name, s)
	}

	flats.Policies = make(map[string]PolicyType)
	for name := range s.PolicyTypes {
		flats.Policies[name] = flattenPolicyType(name, s)
	}

	return flats
}
