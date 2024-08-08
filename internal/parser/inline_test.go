package parser

import (
	"html/template"
	"testing"
	"wikinow/internal/utils"

	"github.com/stretchr/testify/suite"
)

// Setting Up ------------------------------------------------------------------

func (s *InlineSuite) SetupTest() {
	s.ctx = CreateCtx()
}

type InlineSuite struct {
	suite.Suite
	ctx *Ctx
}

func TestInlineSuite(t *testing.T) {
	suite.Run(t, new(InlineSuite))
}

// Bold ------------------------------------------------------------------------

func (s *InlineSuite) Test_should_parse_bold() {
	str := "**bold**"
	result := ParseInline(str, s.ctx, nil)
	s.Equal("<strong>bold</strong>", utils.RemoveClass(result))
}

func (s *InlineSuite) Test_should_parse_bold_twice() {
	str := "**1** **2**"
	result := ParseInline(str, s.ctx, nil)
	s.Equal("<strong>1</strong> <strong>2</strong>", utils.RemoveClass(result))
}

// Italic ----------------------------------------------------------------------

func (s *InlineSuite) Test_should_parse_italic() {
	str := "*italic*"
	result := ParseInline(str, s.ctx, nil)
	s.Equal("<i>italic</i>", utils.RemoveClass(result))
}

func (s *InlineSuite) Test_should_parse_italic_twice() {
	str := "*italic* *italic*"
	result := ParseInline(str, s.ctx, nil)
	s.Equal("<i>italic</i> <i>italic</i>", utils.RemoveClass(result))
}

// Bold Italic -----------------------------------------------------------------

func (s *InlineSuite) Test_should_parse_bold_italic() {
	str := "***bold italic***"
	result := ParseInline(str, s.ctx, nil)
	// This gets fixed when it is rendered, dont worry :)
	s.Equal("<strong><i>bold italic</strong></i>", utils.RemoveClass(result))
}

func (s *InlineSuite) Test_should_parse_bold_italic_twice() {
	str := "***bold italic*** ***bold italic***"
	result := ParseInline(str, s.ctx, nil)
	s.Equal("<strong><i>bold italic</strong></i> <strong><i>bold italic</strong></i>", utils.RemoveClass(result))
}

func (s *InlineSuite) Test_should_parse_bold_italic_underline() {
	str := "___bold italic underline___"
	result := ParseInline(str, s.ctx, nil)
	s.Equal("<strong><i>bold italic underline</strong></i>", utils.RemoveClass(result))
}

func (s *InlineSuite) Test_should_parse_bold_italic_underline_twice() {
	str := "___bold italic underline___ ___bold italic underline___"
	result := ParseInline(str, s.ctx, nil)
	s.Equal("<strong><i>bold italic underline</strong></i> <strong><i>bold italic underline</strong></i>", utils.RemoveClass(result))
}

// Image -----------------------------------------------------------------------

func (s *InlineSuite) Test_should_parse_image() {
	str := "![alt](src)"
	result := ParseInline(str, s.ctx, nil)
	s.Equal("<img src=\"src\" alt=\"alt\" />", utils.RemoveClass(result))
}

func (s *InlineSuite) Test_should_parse_image_by_reference() {
	lines := []string{
		"![1]: src",
	}
	LoadCtx(s.ctx, &lines)

	str := "![alt][1]"
	result := ParseInline(str, s.ctx, nil)

	s.Equal("<img src=\"src\" alt=\"alt\" />", utils.RemoveClass(result))
}

// Link ------------------------------------------------------------------------

func (s *InlineSuite) Test_should_parse_link() {
	str := "[text](href)"
	result := ParseInline(str, s.ctx, nil)
	s.Equal("<a href=\"href\" target=\"_blank\">text</a>", utils.RemoveClass(result))
}

func (s *InlineSuite) Test_should_parse_link_by_reference() {
	lines := []string{
		"[1]: href",
	}
	LoadCtx(s.ctx, &lines)
	str := "[text][1]"
	result := ParseInline(str, s.ctx, nil)
	s.Equal("<a href=\"href\" target=\"_blank\">text</a>", utils.RemoveClass(result))
}

// Variable --------------------------------------------------------------------

func (s *InlineSuite) Test_should_parse_variable() {
	lines := []string{
		"---",
		"test: test",
		"---",
	}
	LoadCtx(s.ctx, &lines)

	str := "this is a $var(test)"
	result := ParseInline(str, s.ctx, nil)

	s.Equal("this is a test", utils.RemoveClass(result))
}

// Code ------------------------------------------------------------------------

func (s *InlineSuite) Test_should_parse_inline_code() {
	str := "`code`"
	result := ParseInline(str, s.ctx, nil)
	s.Equal("<code>code</code>", utils.RemoveClass(result))
}

func (s *InlineSuite) Test_should_parse_inline_code_twice() {
	str := "`code` `code`"
	result := ParseInline(str, s.ctx, nil)
	s.Equal("<code>code</code> <code>code</code>", utils.RemoveClass(result))
}

func (s *InlineSuite) Test_should_parse_code_block() {
	lines := []string{
		"---",
		"title: test",
		"file: testdata/testfile.md",
		"---",
	}
	if err := LoadCtx(s.ctx, &lines); err != nil {
		s.Fail(err.Error())
		return
	}

	str := "$code(file, 1, 2)"

	result := parseCodeBlock(str, s.ctx, new(MockFileReader))

	s.Equal("<pre><code>2</code></pre>", utils.RemoveClass(template.HTML(result)))
}

// Link to another file --------------------------------------------------------

func (s *InlineSuite) Test_should_parse_link_to_another_file() {
  lines := []string{
    "---",
    "title: test",
    "file: testdata/testfile.md",
    "---",
  }
  if err := LoadCtx(s.ctx, &lines); err != nil {
    s.Fail(err.Error())
    return
  }
  str := "$link(text, testdata/testfile.md)"
  result := ParseInline(str, s.ctx)
  s.Equal("<a href=\"testdata/testfile.md\">text</a>", utils.RemoveClass(result))
}

// Mocks -----------------------------------------------------------------------

type MockFileReader struct{}

func (r MockFileReader) ReadFile(path string) ([]byte, error) {
	str := "1\n2\n3\n4\n5"
	return []byte(str), nil
}
