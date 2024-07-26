package parser

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type BoldTestSuite struct {
	suite.Suite
	data []Node
}

func (suite *BoldTestSuite) Test_should_create_Bold() {
	raw := "**bold**"
	content := "bold"
	bold := NewBold(raw, content, 0, len(raw))

	suite.Equal("Bold", bold.GetType())
}

func (suite *BoldTestSuite) Test_should_create_a_Bold() {
	raw := "**There is an *italic* inside another bold**"
	content := "There is an *italic* inside another bold"
	bold := NewBold(raw, content, 0, len(raw))

	suite.Equal("Bold", bold.GetType())
}

func (suite *BoldTestSuite) Test_should_have_italic_as_child() {
	raw := "**There is an *italic* inside another bold**"
	content := "There is an *italic* inside another bold"
	bold := NewBold(raw, content, 0, len(raw))

	child := (*bold.GetChildren())[0]

	suite.Equal(1, len(*bold.GetChildren()))
	suite.Equal("Italic", child.GetType())
}

func (suite *BoldTestSuite) Test_should_have_bold_with_nil_children() {
  raw := "****"
  content := ""
  bold := NewBold(raw, content, 0, len(raw))
  suite.Nil(*bold.GetChildren())
}

func (suite *BoldTestSuite) Test_should_return_two_bold() {
  in := "**this** is not bold and **this** is bold"
  res := ParseBold(in)
  suite.Equal(2, len(*res))
  suite.Equal("**this**", (*res)[0].AsContainer().Raw)
  suite.Equal("**this**", (*res)[1].AsContainer().Raw)
}

func TestBold( t *testing.T) {
  suite.Run(t, new(BoldTestSuite))
} 
