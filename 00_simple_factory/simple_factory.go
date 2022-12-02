package simplefactory

import (
	"k8s.io/klog/v2"
)

// 接口
type IRuleConfigParser interface {
	Parse(data []byte) error
}

// 实现类：yaml格式
type YamlRuleConfigParser struct {
}

func (yrcp YamlRuleConfigParser) Parse(data []byte) error {
	klog.Infof("YamlRuleConfigParser")
	return nil
}

// 实现类：json格式
type JsonRuleConfigParser struct {
}

func (jrcp JsonRuleConfigParser) Parse(data []byte) error {
	klog.Infof("JsonRuleConfigParser")
	return nil
}

func NewRuleConfigParser(configType string) IRuleConfigParser {
	// 简单工厂，这里的switch case语句是无法消除的
	switch configType {
	case "json":
		return JsonRuleConfigParser{}
	case "yaml":
		return YamlRuleConfigParser{}
	}
	return nil
}
