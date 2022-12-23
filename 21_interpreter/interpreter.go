package interpreter

import (
	"fmt"
	"strconv"
	"strings"
)

type IExpression interface {
	Interpret(stats map[string]float64) bool
}

type GreaterExpression struct {
	key          string
	thresholdVal float64
}

func (g *GreaterExpression) Interpret(stats map[string]float64) bool {
	v, ok := stats[g.key]
	if !ok {
		return false
	}
	return v > g.thresholdVal
}

type LessExpression struct {
	key          string
	thresholdVal float64
}

func (l *LessExpression) Interpret(stats map[string]float64) bool {
	v, ok := stats[l.key]
	if !ok {
		return false
	}
	return v < l.thresholdVal
}

func NewGreaterExpression(exp string) (*GreaterExpression, error) {
	exp = strings.TrimSpace(exp)
	elements := strings.Split(exp, " ")
	if len(elements) != 3 || elements[1] != ">" {
		return nil, fmt.Errorf("expression %s is invalid", exp)
	}
	key := elements[0]
	thresholdVal, _ := strconv.ParseFloat(elements[2], 64)
	return &GreaterExpression{
		key:          key,
		thresholdVal: thresholdVal,
	}, nil
}

func NewLessExpression(exp string) (*LessExpression, error) {
	exp = strings.TrimSpace(exp)
	elements := strings.Split(exp, " ")
	if len(elements) != 3 || elements[1] != "<" {
		return nil, fmt.Errorf("expression %s is invalid", exp)
	}
	key := elements[0]
	thresholdVal, _ := strconv.ParseFloat(elements[2], 64)
	return &LessExpression{
		key:          key,
		thresholdVal: thresholdVal,
	}, nil
}

type AndExpression struct {
	expressions []IExpression
}

func (a *AndExpression) Interpret(stats map[string]float64) bool {
	for _, exp := range a.expressions {
		if !exp.Interpret(stats) {
			return false
		}
	}
	return true
}

func NewAndExpression(exp string) (*AndExpression, error) {
	exps := strings.Split(exp, "&&")
	newExpressions := make([]IExpression, len(exps))
	for i, e := range exps {
		e = strings.TrimSpace(e)
		if strings.Contains(e, ">") {
			ge, err := NewGreaterExpression(e)
			if err != nil {
				return nil, err
			}
			newExpressions[i] = ge
		} else if strings.Contains(e, "<") {
			le, err := NewLessExpression(e)
			if err != nil {
				return nil, err
			}
			newExpressions[i] = le
		} else {
			return nil, fmt.Errorf("expression %s is invalid", exp)
		}
	}
	return &AndExpression{
		expressions: newExpressions,
	}, nil
}

// AlertRuleInterpreter 告警规则
type AlertRuleInterpreter struct {
	expression IExpression
}

func NewAlertRuleInterpreter(rule string) (*AlertRuleInterpreter, error) {
	exp, err := NewAndExpression(rule)
	if err != nil {
		return nil, err
	}
	return &AlertRuleInterpreter{
		expression: exp,
	}, nil
}

func (a *AlertRuleInterpreter) Interpret(stats map[string]float64) bool {
	return a.expression.Interpret(stats)
}
