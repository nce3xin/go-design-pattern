package factorymethod

// 配置解析器的接口
type IRuleConfigParser interface {
	Parse(data []byte) error
}

// json格式解析器
type JsonRuleConfigParser struct {
}

// override
func (jrcp JsonRuleConfigParser) Parse(data []byte) error {
	//TODO implement me
	panic("implement me")
}

// yaml格式解析器
type yamlRuleConfigParser struct {
}

// override
func (yrcp yamlRuleConfigParser) Parse(data []byte) error {
	//TODO implement me
	panic("implement me")
}

// 创建配置解析器的工厂的接口
type IRuleConfigParserFactory interface {
	CreateParser() IRuleConfigParser
}

// 创建json格式解析器的工厂
type JsonRuleConfigParserFactory struct {
}

// override
func (jrcpf JsonRuleConfigParserFactory) CreateParser() IRuleConfigParser {
	return JsonRuleConfigParser{}
}

// 创建yaml格式解析器的工厂
type YamlRuleConfigParserFactory struct {
}

// override
func (yrcpf YamlRuleConfigParserFactory) CreateParser() IRuleConfigParser {
	return yamlRuleConfigParser{}
}

func NewIRuleConfigParserFactory(configType string) IRuleConfigParserFactory {
	switch configType {
	case "json":
		return JsonRuleConfigParserFactory{}
	case "yaml":
		return YamlRuleConfigParserFactory{}
	}
	return nil
}
