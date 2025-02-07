package component_test

import (
	"context"
	"io"
	"testing"
	"wikinow/component"
	"wikinow/internal/handler"
	"wikinow/internal/parser"

	"github.com/PuerkitoBio/goquery"
	"github.com/stretchr/testify/suite"
)

// setup -----------------------------------------------------------------------

func (s *CtxSuite) SetupTest() {
	s.ctx = parser.CreateCtx()
}

type CtxSuite struct {
	suite.Suite
	ctx *parser.Ctx
}

func TestLoadSuite(t *testing.T) {
	suite.Run(t, new(CtxSuite))
}

// tests -----------------------------------------------------------------------

func (s *CtxSuite) Test_should_render_blockquote() {
	r, w := io.Pipe()

	lines := []string{
		"> This is a blockquote",
	}

	astTree, _ := handler.TestAst(lines)

	go func() {
		_ = component.Parser(astTree, &lines, s.ctx).Render(context.Background(), w)
		_ = w.Close()
	}()

	doc, err := goquery.NewDocumentFromReader(r)
	s.Require().NoError(err, "failed to read template")

  html, _ := doc.Find(`[data-testid="blockquote"]`).Html()

	s.Assert().Equal(html, "This is a blockquote<br/> ", "Expects blockquote text to be `This is a blockquote`")
}

func (s *CtxSuite) Test_should_render_multiple_line_blockquote() {
	r, w := io.Pipe()

	lines := []string{
		"> This",
		"> is",
		"> a",
		"> long",
		"> blockquote",
	}

	astTree, _ := handler.TestAst(lines)

	go func() {
		_ = component.Parser(astTree, &lines, s.ctx).Render(context.Background(), w)
		_ = w.Close()
	}()

	doc, err := goquery.NewDocumentFromReader(r)
	s.Require().NoError(err, "failed to read template")

  html, _ := doc.Find(`[data-testid="blockquote"]`).Html()

	s.Assert().Equal(html, "This<br/> is<br/> a<br/> long<br/> blockquote<br/> ", "Expect multiline blockquote text")
}

// func (s *CtxSuite) Test_should_load_nested_blockquotes() {
// 	r, w := io.Pipe()
//
// 	lines := []string{
// 		"> First line",
// 		"> > Nested line",
// 		"> Last line",
// 	}
//
// 	astTree, _ := handler.TestAst(lines)
//
// 	go func() {
// 		_ = component.Parser(astTree, &lines, s.ctx).Render(context.Background(), w)
// 		_ = w.Close()
// 	}()
//
// 	doc, err := goquery.NewDocumentFromReader(r)
// 	s.Require().NoError(err, "failed to read template")
//
//   html, _ := doc.Find(`[data-testid="blockquote"]`).Html()
//
// 	s.Assert().Equal(html, "This<br/> is<br/> a<br/> long<br/> blockquote<br/> ", "Expect multiline blockquote text")
// }

func (s *CtxSuite) Test_should_skip_a_line_on_missing_information() {
	r, w := io.Pipe()

	lines := []string{
		"> First line",
		">",
		"> Last line",
	}

	astTree, _ := handler.TestAst(lines)

	go func() {
		_ = component.Parser(astTree, &lines, s.ctx).Render(context.Background(), w)
		_ = w.Close()
	}()

	doc, err := goquery.NewDocumentFromReader(r)
	s.Require().NoError(err, "failed to read template")

  html, _ := doc.Find(`[data-testid="blockquote"]`).Html()

	s.Assert().Equal(html, "First line<br/> <br/>Last line<br/> ", "Expect multiline blockquote text")
}
