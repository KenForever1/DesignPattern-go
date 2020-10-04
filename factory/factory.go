package factory

type IRuleConfigParser interface {
	Parse(data []byte)
}

type JsonRuleConfigParser struct{}

func (J JsonRuleConfigParser) Parse(data []byte) {
	panic("implement me")
}

type YamlRuleConfigParser struct{}

func (Y YamlRuleConfigParser) Parse(data []byte) {
	panic("implement me")
}

func NewIRuleConfigParser(t string) IRuleConfigParser {
	switch t {
	case "json":
		return JsonRuleConfigParser{}
	case "yaml":
		return YamlRuleConfigParser{}
	}
	return nil
}
