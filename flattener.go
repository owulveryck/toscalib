package toscalib

import "github.com/kenjones-cisco/mergo"

func (s *ServiceTemplateDefinition) flattenCapType(name string) CapabilityType {
	if ct, ok := s.CapabilityTypes[name]; ok {
		if ct.DerivedFrom != "" {
			parent := s.flattenCapType(ct.DerivedFrom)
			_ = mergo.MergeWithOverwrite(&parent, ct)
			return parent
		}
		return ct
	}
	return CapabilityType{}
}

func (s *ServiceTemplateDefinition) flattenNodeType(name string) NodeType {
	if nt, ok := s.NodeTypes[name]; ok {
		if nt.DerivedFrom != "" {
			parent := s.flattenNodeType(nt.DerivedFrom)
			// mergo does not handle merging Slices so the nt items
			// will wipe away, capture the values here.
			reqs := parent.Requirements
			arts := parent.Artifacts

			_ = mergo.MergeWithOverwrite(&parent, nt)

			// now copy them back in using append, if the child node type had
			// any previously, otherwise it will duplicate the parents.
			if len(nt.Requirements) > 0 {
				parent.Requirements = append(parent.Requirements, reqs...)
			}
			if len(nt.Artifacts) > 0 {
				parent.Artifacts = append(parent.Artifacts, arts...)
			}
			return parent
		}
		return nt
	}
	return NodeType{}
}
