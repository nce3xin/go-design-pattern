package example2

// 统一多个类的接口设计

// ASensitiveWordsFilter A过滤系统
type ASensitiveWordsFilter struct {
}

// text是原始文本，函数输出用***替换敏感词之后的文本
func (a *ASensitiveWordsFilter) filterSexyWords(text string) string {
	return ""
}

func (a *ASensitiveWordsFilter) filterPoliticalWords(text string) string {
	return ""
}

// BSensitiveWordsFilter B过滤系统
type BSensitiveWordsFilter struct {
}

func (b *BSensitiveWordsFilter) filter(text string) {

}

// CSensitiveWordsFilter C过滤系统
type CSensitiveWordsFilter struct {
}

// C过滤系统提供的接口
func (c *CSensitiveWordsFilter) filter(text, mask string) {

}

type ISensitiveWordsFilter interface {
	filter(text string) string
}

type ASensitiveWordsFilterAdaptor struct {
	aFilter *ASensitiveWordsFilter
}

func (A ASensitiveWordsFilterAdaptor) filter(text string) string {
	maskedText := A.aFilter.filterSexyWords(text)
	maskedText = A.aFilter.filterPoliticalWords(maskedText)
	return maskedText
}

// ... 省略BSensitiveWordsFilterAdaptor、CSensitiveWordsFilterAdaptor

type RiskManagement struct {
	filters []ISensitiveWordsFilter
}

func (r *RiskManagement) addSensitiveWordsFilter(f ISensitiveWordsFilter) {
	r.addSensitiveWordsFilter(f)
}

func (r *RiskManagement) filterSensitiveWords(text string) string {
	maskedText := text
	for _, filter := range r.filters {
		maskedText = filter.filter(text)
	}
	return maskedText
}
