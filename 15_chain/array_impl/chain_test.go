package arrayimpl

import "testing"

func TestSensitiveWordFilterChain_Filter(t *testing.T) {
	chain := &SensitiveWordFilterChain{}
	chain.AddFilter(&SexyWordFilter{})
	chain.AddFilter(&PoliticalWordFilter{})
	chain.AddFilter(&AdsWordFilter{})
	legal := chain.Filter("content")
	if !legal {
		// 非法，不发表
	} else {
		// 合法，发表
	}
}
