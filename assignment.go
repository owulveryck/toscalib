package toscalib

import (
	"fmt"
	"reflect"
	"strconv"
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
		for k, v := range m {
			if isFunction(k) {
				p.Function = k
				args := make([]interface{}, 1)
				args[0] = v
				p.Args = args
			}
			if isOperator(k) {
				p.Expression = ConstraintClause{Operator: k, Values: v}
			}
		}
		return nil
	}

	var mm map[string][]string
	if err := unmarshal(&mm); err == nil {
		for k, v := range mm {
			if isFunction(k) {
				p.Function = k
				args := make([]interface{}, len(v))
				for i, a := range v {
					args[i] = a
				}
				p.Args = args
			}
			if isOperator(k) {
				p.Expression = ConstraintClause{Operator: k, Values: v}
			}
		}
		return nil
	}

	// ex. concat function with another function embedded inside
	var mmm map[string][]interface{}
	if err := unmarshal(&mmm); err == nil {
		for k, v := range mmm {
			if isFunction(k) {
				p.Function = k
				p.Args = v
			}
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

// Evaluate gets the value of an Assignment, including the evaluation of expression or function
func (p *Assignment) Evaluate(std *ServiceTemplateDefinition, ctx string) interface{} {
	// TODO(kenjones): Add support for the evaluation of ConstraintClause
	if p.Value != nil {
		return p.Value
	}

	switch p.Function {
	case ConcatFunc:
		var output string
		for _, val := range p.Args {
			switch reflect.TypeOf(val).Kind() {
			case reflect.String:
				output = fmt.Sprintf("%s%s", output, val)
			case reflect.Int:
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

	case GetInputFunc:
		if len(p.Args) == 1 {
			return std.GetInputValue(p.Args[0].(string), false)
		}

	case GetPropFunc:
		if len(p.Args) == 2 {
			if nt := std.findNodeTemplate(p.Args[0].(string), ctx); nt != nil {
				if pa, ok := nt.Properties[p.Args[1].(string)]; ok {
					return pa.Evaluate(std, nt.Name)
				}
			}
		}
		if len(p.Args) >= 3 {
			nt, rnt := getNTByArgs(std, ctx, p.Args)
			if nt == nil {
				break
			}
			if rnt != nil {
				if prop := rnt.findProperty(p.Args[2].(string), p.Args[1].(string)); prop != nil {
					return prop.evaluate(std, rnt.Name, get(3, p.Args))
				}
			}
			if prop := nt.findProperty(p.Args[2].(string), p.Args[1].(string)); prop != nil {
				return prop.evaluate(std, nt.Name, get(3, p.Args))
			}
			if prop := nt.findProperty(p.Args[1].(string), ""); prop != nil {
				return prop.evaluate(std, nt.Name, get(2, p.Args))
			}
		}

	case GetAttrFunc:
		if len(p.Args) == 2 {
			if nt := std.findNodeTemplate(p.Args[0].(string), ctx); nt != nil {
				if pa, ok := nt.Attributes[p.Args[1].(string)]; ok {
					return pa.Evaluate(std, nt.Name)
				}
			}
		}
		if len(p.Args) >= 3 {
			nt, rnt := getNTByArgs(std, ctx, p.Args)
			if nt == nil {
				break
			}
			if rnt != nil {
				if attr := rnt.findAttribute(p.Args[2].(string), p.Args[1].(string)); attr != nil {
					return attr.evaluate(std, rnt.Name, get(3, p.Args))
				}
			}
			if attr := nt.findAttribute(p.Args[2].(string), p.Args[1].(string)); attr != nil {
				return attr.evaluate(std, nt.Name, get(3, p.Args))
			}
			if attr := nt.findAttribute(p.Args[1].(string), ""); attr != nil {
				return attr.evaluate(std, nt.Name, get(2, p.Args))
			}
		}
	}

	return nil
}
