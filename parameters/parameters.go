package parameters

import (
	"net/url"
	"strings"
)

const (
	ParameterTypeQuery ParameterType = iota
	ParameterTypePath
	ParameterTypeBody
)

type ParameterType int

type Parameter struct {
	Name  string
	Value any
	Type  ParameterType
}

func (p *Parameter) SetValue(value any) {
	p.Value = value
}

func NewQueryParameter(name string, value string) Parameter {
	return Parameter{
		Name:  name,
		Value: value,
		Type:  ParameterTypeQuery,
	}
}
func NewPathParameter(name string, value string) Parameter {
	return Parameter{
		Name:  name,
		Value: value,
		Type:  ParameterTypePath,
	}
}

func NewBodyParameter(name string, value any) Parameter {
	return Parameter{
		Name:  name,
		Value: value,
		Type:  ParameterTypeBody,
	}
}

func ApplyURL(u url.URL, parameters ...Parameter) url.URL {
	vals := make(url.Values)

	for _, p := range parameters {
		if val, ok := p.Value.(string); ok {
			switch p.Type {
			case ParameterTypeQuery:
				vals.Set(p.Name, val)
			case ParameterTypePath:
				if val == "" {
					u.Path = strings.Replace(u.Path, "/:"+p.Name, "", 1)
				}
				u.Path = strings.Replace(u.Path, ":"+p.Name, val, 1)
			case ParameterTypeBody:
				continue
			default:
				continue
			}
		}
	}

	return u
}

func CreateBody(parameters ...Parameter) map[string]any {
	body := make(map[string]any, len(parameters))

	for _, parameter := range parameters {
		if parameter.Type == ParameterTypeBody {
			body[parameter.Name] = parameter.Value
		}
	}
	return body
}
