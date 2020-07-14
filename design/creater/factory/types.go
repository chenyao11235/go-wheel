package factory

type ruleConfig struct {
}


//Parser 用于解析配置文件
type Parser interface {
	parse() *ruleConfig
}

//JSONRuleConfigParser 解析json文件
type JSONRuleConfigParser struct {
}

func (p *JSONRuleConfigParser) parse() *ruleConfig {
	return nil
}

//XMLRuleConfigParser 解析xml文件
type XMLRuleConfigParser struct {
}

func (p *XMLRuleConfigParser) parse() *ruleConfig {
	return nil
}

//YamlRuleConfigParser 解析yaml文件
type YamlRuleConfigParser struct {
}

func (p *YamlRuleConfigParser) parse() *ruleConfig {
	return nil
}
