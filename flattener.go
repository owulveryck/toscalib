package toscalib

import "github.com/kenjones-cisco/mergo"

type flatTypes struct {
	ArtifactTypes map[string]ArtifactType
	Capabilities  map[string]CapabilityType
	Interfaces    map[string]InterfaceType
	Relationships map[string]RelationshipType
	Nodes         map[string]NodeType
	Groups        map[string]GroupType
	Policies      map[string]PolicyType
}

func flattenArtType(name string, s ServiceTemplateDefinition) ArtifactType {
	if at, ok := s.ArtifactTypes[name]; ok {
		if at.DerivedFrom != "" {
			parent := flattenArtType(at.DerivedFrom, s)

			// clone the parent first before applying any changes
			tmp := clone(parent)
			atm, _ := tmp.(ArtifactType)

			// mergo does not handle merging Slices so the items
			// will wipe away, capture the values here.
			exts := atm.FileExt

			_ = mergo.MergeWithOverwrite(&atm, at)

			// now copy them back in using append, if the child type had
			// any previously, otherwise it will duplicate the parents.
			if len(at.FileExt) > 0 {
				atm.FileExt = append(atm.FileExt, exts...)
			}
			return atm
		}
		return at
	}
	return ArtifactType{}
}

func flattenCapType(name string, s ServiceTemplateDefinition) CapabilityType {
	if ct, ok := s.CapabilityTypes[name]; ok {
		if ct.DerivedFrom != "" {
			parent := flattenCapType(ct.DerivedFrom, s)

			// clone the parent first before applying any changes
			tmp := clone(parent)
			ctm, _ := tmp.(CapabilityType)

			// mergo does not handle merging Slices so the rt items
			// will wipe away, capture the values here.
			sources := ctm.ValidSources

			_ = mergo.MergeWithOverwrite(&ctm, ct)

			// now copy them back in using append, if the child type had
			// any previously, otherwise it will duplicate the parents.
			if len(ct.ValidSources) > 0 {
				ctm.ValidSources = append(ctm.ValidSources, sources...)
			}
			return ctm
		}
		return ct
	}
	return CapabilityType{}
}

func flattenIntfType(name string, s ServiceTemplateDefinition) InterfaceType {
	if it, ok := s.InterfaceTypes[name]; ok {
		if it.DerivedFrom != "" {
			parent := flattenIntfType(it.DerivedFrom, s)

			// clone the parent first before applying any changes
			tmp := clone(parent)
			itm, _ := tmp.(InterfaceType)

			_ = mergo.MergeWithOverwrite(&itm, it)
			return itm
		}
		return it
	}
	return InterfaceType{}
}

func flattenRelType(name string, s ServiceTemplateDefinition) RelationshipType {
	if rt, ok := s.RelationshipTypes[name]; ok {
		if rt.DerivedFrom != "" {
			parent := flattenRelType(rt.DerivedFrom, s)

			// clone the parent first before applying any changes
			tmp := clone(parent)
			rtm, _ := tmp.(RelationshipType)

			// mergo does not handle merging Slices so the rt items
			// will wipe away, capture the values here.
			targets := rtm.ValidTarget

			_ = mergo.MergeWithOverwrite(&rtm, rt)

			// now copy them back in using append, if the child type had
			// any previously, otherwise it will duplicate the parents.
			if len(rt.ValidTarget) > 0 {
				rtm.ValidTarget = append(rtm.ValidTarget, targets...)
			}
			return rtm
		}
		return rt
	}
	return RelationshipType{}
}

func flattenNodeType(name string, s ServiceTemplateDefinition) NodeType {
	if nt, ok := s.NodeTypes[name]; ok {
		if nt.DerivedFrom != "" {
			parent := flattenNodeType(nt.DerivedFrom, s)

			// clone the parent first before applying any changes
			tmp := clone(parent)
			ntm, _ := tmp.(NodeType)

			// mergo does not handle merging Slices so the nt items
			// will wipe away, capture the values here.
			reqs := ntm.Requirements

			_ = mergo.MergeWithOverwrite(&ntm, nt)

			// now copy them back in using append, if the child node type had
			// any previously, otherwise it will duplicate the parents.
			if len(nt.Requirements) > 0 {
				ntm.Requirements = append(ntm.Requirements, reqs...)
			}
			return ntm
		}
		return nt
	}
	return NodeType{}
}

func flattenGroupType(name string, s ServiceTemplateDefinition) GroupType {
	if gt, ok := s.GroupTypes[name]; ok {
		if gt.DerivedFrom != "" {
			parent := flattenGroupType(gt.DerivedFrom, s)

			// clone the parent first before applying any changes
			tmp := clone(parent)
			gtm, _ := tmp.(GroupType)

			// mergo does not handle merging Slices so the nt items
			// will wipe away, capture the values here.
			reqs := gtm.Requirements
			members := gtm.Members

			_ = mergo.MergeWithOverwrite(&gtm, gt)

			// now copy them back in using append, if the child type had
			// any previously, otherwise it will duplicate the parents.
			if len(gt.Requirements) > 0 {
				gtm.Requirements = append(gtm.Requirements, reqs...)
			}
			if len(gt.Members) > 0 {
				gtm.Members = append(gtm.Members, members...)
			}
			return gtm
		}
		return gt
	}
	return GroupType{}
}

func flattenPolicyType(name string, s ServiceTemplateDefinition) PolicyType {
	if pt, ok := s.PolicyTypes[name]; ok {
		if pt.DerivedFrom != "" {
			parent := flattenPolicyType(pt.DerivedFrom, s)

			// clone the parent first before applying any changes
			tmp := clone(parent)
			ptm, _ := tmp.(PolicyType)

			// mergo does not handle merging Slices so the items
			// will wipe away, capture the values here.
			targets := ptm.Targets

			_ = mergo.MergeWithOverwrite(&ptm, pt)

			// now copy them back in using append, if the child type had
			// any previously, otherwise it will duplicate the parents.
			if len(pt.Targets) > 0 {
				ptm.Targets = append(ptm.Targets, targets...)
			}
			return ptm
		}
		return pt
	}
	return PolicyType{}
}

func flattenHierarchy(s ServiceTemplateDefinition) flatTypes {
	var flats flatTypes

	flats.ArtifactTypes = make(map[string]ArtifactType)
	for name := range s.ArtifactTypes {
		flats.ArtifactTypes[name] = flattenArtType(name, s)
	}

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

	for k, v := range flats.Relationships {
		for name, iDef := range v.Interfaces {
			// during merge the Type is not always properly inherited, so set it
			// from the parent.
			if iDef.Type == "" {
				iDef.Type = flats.Relationships[v.DerivedFrom].Interfaces[name].Type
			}
			iDef.extendFrom(flats.Interfaces[iDef.Type])
			v.Interfaces[name] = iDef
		}
		flats.Relationships[k] = v
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
