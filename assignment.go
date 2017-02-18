package toscalib

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
)

// Assignment supports Value evaluation
type Assignment struct {
	Value      interface{}
	Function   string
	Args       []interface{}
	Expression ConstraintClause
}

// UnmarshalYAML converts YAML text to a type
func (p *Assignment) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var s string
	if err := unmarshal(&s); err == nil {
		p.Value = s
		return nil
	}

	var m map[string]string
	if err := unmarshal(&m); err == nil {
		processed := false
		for k, v := range m {
			if isFunction(k) {
				processed = true
				p.Function = k
				args := make([]interface{}, 1)
				args[0] = v
				p.Args = args
			}
			if isOperator(k) {
				processed = true
				p.Expression = ConstraintClause{Operator: k, Values: v}
			}
		}
		if !processed {
			p.Value = m
		}
		return nil
	}

	var mm map[string][]string
	if err := unmarshal(&mm); err == nil {
		processed := false
		for k, v := range mm {
			if isFunction(k) {
				processed = true
				p.Function = k
				args := make([]interface{}, len(v))
				for i, a := range v {
					args[i] = a
				}
				p.Args = args
			}
			if isOperator(k) {
				processed = true
				p.Expression = ConstraintClause{Operator: k, Values: v}
			}
		}
		if !processed {
			p.Value = mm
		}
		return nil
	}

	// ex. concat function with another function embedded inside
	var mmm map[string][]interface{}
	if err := unmarshal(&mmm); err == nil {
		processed := false
		for k, v := range mmm {
			if isFunction(k) {
				processed = true
				p.Function = k
				p.Args = v
			}
		}
		if !processed {
			p.Value = mmm
		}
		return nil
	}

	// Value is list of values
	var mmmm []interface{}
	if err := unmarshal(&mmmm); err == nil {
		p.Value = mmmm
		return nil
	}

	// Value is map of values
	var mmmmm map[string]interface{}
	if err := unmarshal(&mmmmm); err == nil {
		p.Value = mmmmm
		return nil
	}

	var res interface{}
	_ = unmarshal(&res)
	return fmt.Errorf("Cannot parse Property %v", res)
}

func newAssignmentFunc(val interface{}) *Assignment {
	rval := reflect.ValueOf(val)
	switch rval.Kind() {
	case reflect.Map:
		for k, v := range rval.Interface().(map[interface{}]interface{}) {
			if !isFunction(k.(string)) {
				continue
			}
			// Convert it to a Assignment
			rv := reflect.ValueOf(v)
			switch rv.Kind() {
			case reflect.String:
				return &Assignment{Function: k.(string), Args: []interface{}{rv.Interface()}}
			default:
				return &Assignment{Function: k.(string), Args: rv.Interface().([]interface{})}
			}
		}
	}
	return nil
}

func (p *Assignment) lookupValueArg(arg string) interface{} {
	v := reflect.ValueOf(p.Value)
	switch v.Kind() {
	case reflect.Slice:
		if i, err := strconv.ParseInt(arg, 10, 0); err == nil {
			return v.Index(int(i)).Interface()
		}
	case reflect.Map:
		return v.MapIndex(reflect.ValueOf(arg)).Interface()
	}
	return v.Interface()
}

func (p *Assignment) evaluate(std *ServiceTemplateDefinition, ctx, arg string) interface{} {
	if arg != "" {
		val := p.lookupValueArg(arg)

		// handle the scenario when the property value is another
		// function call.
		if pa := newAssignmentFunc(val); pa != nil {
			return pa.Evaluate(std, ctx)
		}

		if len(p.Args) != 0 {
			tmp := clone(*p)
			pa, _ := tmp.(Assignment)
			pa.Value = val
			return pa.lookupValueArg(p.Args[0].(string))
		}

		return val
	}
	return p.Evaluate(std, ctx)
}

func getNTByArgs(std *ServiceTemplateDefinition, ctx string, args []interface{}) (*NodeTemplate, *NodeTemplate) {
	nt := std.findNodeTemplate(args[0].(string), ctx)
	if nt == nil {
		return nil, nil
	}

	if r := nt.GetRequirement(args[1].(string)); r != nil {
		if rnt := std.GetNodeTemplate(r.Node); rnt != nil {
			return nt, rnt
		}
	}
	return nt, nil
}

func (p *Assignment) evalConcat(std *ServiceTemplateDefinition, ctx string) interface{} {
	var output string
	for _, val := range p.Args {
		switch reflect.TypeOf(val).Kind() {
		case reflect.String, reflect.Int:
			output = fmt.Sprintf("%s%s", output, val)
		case reflect.Map:
			if pa := newAssignmentFunc(val); pa != nil {
				if o := pa.Evaluate(std, ctx); o != nil {
					output = fmt.Sprintf("%s%s", output, o)
				}
			}
		}
	}
	return output
}

func (p *Assignment) evalToken(std *ServiceTemplateDefinition, ctx string) interface{} {
	var output string
	var value string
	var token string
	for idx, val := range p.Args {
		switch reflect.TypeOf(val).Kind() {
		case reflect.String:
			// the first and second arg must result be of type string or resolve
			// to be a string.
			if idx == 0 {
				value = val.(string)
			} else if idx == 1 {
				token = val.(string)
			}
		case reflect.Map:
			// the first input could actually be a lookup for the value
			if idx == 0 {
				if pa := newAssignmentFunc(val); pa != nil {
					if o := pa.Evaluate(std, ctx); o != nil {
						value = fmt.Sprintf("%s", o)
					}
				}
			}
		case reflect.Int:
			// the 3rd arg must be an int
			if idx == 2 && value != "" && token != "" {
				output = strings.Split(value, token)[val.(int)]
			}
		}
	}
	if output == "" {
		return nil
	}
	return output
}

func (p *Assignment) evalArtifact(std *ServiceTemplateDefinition, ctx string) interface{} {
	nt, _ := getNTByArgs(std, ctx, p.Args)
	if nt == nil {
		return nil
	}

	if at, ok := nt.Artifacts[p.Args[1].(string)]; ok {
		// set default location to the 'temp|tmp' directory to handle 'LOCAL_FILE' being specified
		// or no location or deploy_path is specified.
		location := os.TempDir()
		if loc := get(2, p.Args); loc != "" && loc != LocalFile {
			location = loc
		} else if at.DeployPath != "" {
			location = at.DeployPath
		}

		destFile, err := copyFile(at.File, location)
		if err != nil {
			return nil
		}
		return destFile
	}
	return nil
}

func (p *Assignment) evalProperty(std *ServiceTemplateDefinition, ctx string) interface{} {
	nt, rnt := getNTByArgs(std, ctx, p.Args)
	if nt == nil {
		return nil
	}

	if len(p.Args) == 2 {
		if prop := nt.findProperty(p.Args[1].(string), ""); prop != nil {
			return prop.evaluate(std, nt.Name, "")
		}
	}
	if len(p.Args) >= 3 {
		if rnt != nil {
			if prop := rnt.findProperty(p.Args[2].(string), p.Args[1].(string)); prop != nil {
				prop.Args = remainder(3, p.Args)
				return prop.evaluate(std, rnt.Name, get(3, p.Args))
			}
		}
		if prop := nt.findProperty(p.Args[2].(string), p.Args[1].(string)); prop != nil {
			prop.Args = remainder(3, p.Args)
			return prop.evaluate(std, nt.Name, get(3, p.Args))
		}
		if prop := nt.findProperty(p.Args[1].(string), ""); prop != nil {
			prop.Args = remainder(2, p.Args)
			return prop.evaluate(std, nt.Name, get(2, p.Args))
		}
	}
	return nil
}

func (p *Assignment) evalAttribute(std *ServiceTemplateDefinition, ctx string) interface{} {
	nt, rnt := getNTByArgs(std, ctx, p.Args)
	if nt == nil {
		return nil
	}

	if len(p.Args) == 2 {
		if attr := nt.findAttribute(p.Args[1].(string), ""); attr != nil {
			return attr.evaluate(std, nt.Name, "")
		}
	}
	if len(p.Args) >= 3 {
		if rnt != nil {
			if attr := rnt.findAttribute(p.Args[2].(string), p.Args[1].(string)); attr != nil {
				attr.Args = remainder(3, p.Args)
				return attr.evaluate(std, rnt.Name, get(3, p.Args))
			}
		}
		if attr := nt.findAttribute(p.Args[2].(string), p.Args[1].(string)); attr != nil {
			attr.Args = remainder(3, p.Args)
			return attr.evaluate(std, nt.Name, get(3, p.Args))
		}
		if attr := nt.findAttribute(p.Args[1].(string), ""); attr != nil {
			attr.Args = remainder(2, p.Args)
			return attr.evaluate(std, nt.Name, get(2, p.Args))
		}
	}
	return nil
}

// Evaluate gets the value of an Assignment, including the evaluation of expression or function
func (p *Assignment) Evaluate(std *ServiceTemplateDefinition, ctx string) interface{} {
	// TODO(kenjones): Add support for the evaluation of ConstraintClause
	if p.Value != nil {
		return p.Value
	}

	switch p.Function {
	case ConcatFunc:
		return p.evalConcat(std, ctx)

	case TokenFunc:
		// there are 3 required args
		if len(p.Args) == 3 {
			return p.evalToken(std, ctx)
		}

	case GetArtifactFunc:
		if len(p.Args) > 1 {
			return p.evalArtifact(std, ctx)
		}

	case GetInputFunc:
		if len(p.Args) == 1 {
			return std.GetInputValue(p.Args[0].(string), false)
		}

	case GetPropFunc:
		return p.evalProperty(std, ctx)

	case GetAttrFunc:
		return p.evalAttribute(std, ctx)
	}

	return nil
}
