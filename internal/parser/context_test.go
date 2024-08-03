package parser

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type LoadSuite struct {
	suite.Suite
	ctx *Ctx
}

func (s *LoadSuite) SetupTest() {
	s.ctx = CreateCtx()
}

func (s *LoadSuite) Test_should_load_all_values_to_store() {
	lines := []string{
		"---",
		"key1: value1",
		"key2: value2",
		"key3: value3",
		"---",
	}

	LoadCtx(s.ctx, &lines)

	key1, _ := ReadCtx(s.ctx, "key1")
	key2, _ := ReadCtx(s.ctx, "key2")
	key3, _ := ReadCtx(s.ctx, "key3")

	s.Equal("value1", key1)
	s.Equal("value2", key2)
	s.Equal("value3", key3)
}

func (s *LoadSuite) Test_should_not_read_what_is_not_within_delimeter() {
	lines := []string{
		"---",
		"key1: value1",
		"---",
		"key3: value3",
	}

	LoadCtx(s.ctx, &lines)

	key1, _ := ReadCtx(s.ctx, "key1")
	key3, _ := ReadCtx(s.ctx, "key3")

	s.Equal("value1", key1)
	s.Equal("", key3)
}

func (s *LoadSuite) Test_should_store_link_definitions() {
	lines := []string{
		"[1]: http://url/b.jpg",
		"[link2]: http://url/a.jpg",
		"![2]: http://url/c.jpg",
	}

	LoadCtx(s.ctx, &lines)

	key1, _ := ReadCtx(s.ctx, "1")
	key2, _ := ReadCtx(s.ctx, "link2")
	key3, _ := ReadCtx(s.ctx, "2")

	s.Equal("http://url/b.jpg", key1)
	s.Equal("http://url/a.jpg", key2)
	s.Equal("http://url/c.jpg", key3)
}

func (s *LoadSuite) Test_should_fail_if_keys_are_repeated() {
	lines := []string{
		"---",
		"test: hello",
		"---",
		"[test]: http://url/b.jpg",
	}

	err := LoadCtx(s.ctx, &lines)
	if err == nil {
		s.Fail("Expected error, got nil")
		return
	}

	s.Equal("Duplicate key: test", err.Error())
}

func (s *LoadSuite) Test_should_fail_if_values_are_invalid() {
	lines := []string{
		"---",
		"test: [hello](world)", // this is a link
		"---",
		"[test]: http://url/b.jpg",
	}

	err := LoadCtx(s.ctx, &lines)
	if err == nil {
		s.Fail("Expected error, got nil")
		return
	}

	s.Equal("value can only contain text: [hello](world)", err.Error())
}

func (s *LoadSuite) Test_should_fail_if_title_is_not_set() {
  lines := []string{
    "---",
    "---",
  }
  err := LoadCtx(s.ctx, &lines)
  if err == nil {
    s.Fail("Expected error, got nil")
    return
  }
  s.Equal("title must be set", err.Error())
}

func TestLoadSuite(t *testing.T) {
	suite.Run(t, new(LoadSuite))
}
