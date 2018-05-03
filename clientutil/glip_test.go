package clientutil

import (
	"testing"
)

var fuzzyAtMentionTests = []struct {
	inputName string
	inputText string
	expected  bool
}{
	{`glipbot`, ` @glipbot hello world!`, true},
	{`glipbot`, ` hello world! @glipbot `, true},
	{`glipbot`, `hello world!`, false},
	{`glipbot`, `@glipdog hello world`, false},
	{`  GLIPBOT  `, ` @GlipBot hello world!`, true},
	{`  GLIPBOT  `, ` @glipbot hello world!`, true}}

func TestIsAtMentionedFuzzy(t *testing.T) {
	for _, tt := range fuzzyAtMentionTests {
		match := IsAtMentionedFuzzy(tt.inputName, tt.inputText)

		if match != tt.expected {
			t.Errorf("clientutil.IsAtMentionedFuzzy error: name[%v] text[%v] result[%v] expected [%v]", tt.inputName, tt.inputText, match, tt.expected)
		}
	}
}
