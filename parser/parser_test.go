package parser

import (
	"sort"
  "testing"

	"github.com/stretchr/testify/suite"
)

type ParserTestSuite struct {
	suite.Suite
	data []Node
}

func (suite *ParserTestSuite) SetupTest() {
	suite.data = []Node{
		&Container{Type: "D", Start: 7, End: 8},
		&Container{Type: "E", Start: 8, End: 9},
		&Container{Type: "B", Start: 3, End: 4},
		&Leaf{Type: "G", Start: 2, End: 3},
		&Container{Type: "A", Start: 1, End: 10},
		&Container{Type: "C", Start: 3, End: 5},
		&Container{Type: "F", Start: 5, End: 9},
	}
}

func (suite *ParserTestSuite) Test_sorting() {
	sort.Sort(ByDiff(suite.data))
	suite.Equal("D", suite.data[0].GetType())
	suite.Equal("E", suite.data[1].GetType())
	suite.Equal("B", suite.data[2].GetType())
	suite.Equal("G", suite.data[3].GetType())
	suite.Equal("C", suite.data[4].GetType())
	suite.Equal("F", suite.data[5].GetType())
	suite.Equal("A", suite.data[6].GetType())
}

// Testing GroupNodes function

func (suite *ParserTestSuite) Test_A_should_be_parent_of_them_all() {
	grouped := GroupNodes(&suite.data)

	suite.Equal(1, len(*grouped))
	a := (*grouped)[0]
	suite.Equal("A", a.GetType())
}

func (suite *ParserTestSuite) Test_A_should_have_three_children() {
	grouped := GroupNodes(&suite.data)
	a := (*grouped)[0]
	aChildren := *a.GetChildren()
	suite.Equal(3, len(aChildren))
}

func (suite *ParserTestSuite) Test_G_C_F_should_be_A_children() {
	grouped := GroupNodes(&suite.data)
	aChildren := *(*grouped)[0].GetChildren()
	suite.Equal("G", aChildren[0].GetType())
	suite.Equal("C", aChildren[1].GetType())
	suite.Equal("F", aChildren[2].GetType())
}

func (suite *ParserTestSuite) Test_C_should_have_one_children() {
	grouped := GroupNodes(&suite.data)
	c := (*(*grouped)[0].GetChildren())[1]
	cChildren := *c.GetChildren()
	suite.Equal(1, len(cChildren))
}

func (suite *ParserTestSuite) Test_C_should_have_B_as_child() {
	grouped := GroupNodes(&suite.data)
	c := (*(*grouped)[0].GetChildren())[1]
	cChildren := *c.GetChildren()
	suite.Equal("B", cChildren[0].GetType())
}

func (suite *ParserTestSuite) Test_F_should_have_two_children() {
	grouped := GroupNodes(&suite.data)
	f := (*(*grouped)[0].GetChildren())[2]
	fChildren := *f.GetChildren()
	suite.Equal(2, len(fChildren))
}

func (suite *ParserTestSuite) Test_F_should_have_D_and_E_as_children() {
	grouped := GroupNodes(&suite.data)
	f := (*(*grouped)[0].GetChildren())[2]
	fChildren := *f.GetChildren()
	suite.Equal("D", fChildren[0].GetType())
	suite.Equal("E", fChildren[1].GetType())
}

func (suite *ParserTestSuite) Test_leaves_should_not_have_children() {
	nodes := []Node{
		&Leaf{Type: "G", Start: 2, End: 3},
		&Leaf{Type: "G", Start: 1, End: 10},
		&Leaf{Type: "G", Start: 2, End: 4},
	}

	grouped := GroupNodes(&nodes)
	suite.Equal(3, len(*grouped))
}

func (suite *ParserTestSuite) Test_should_have_5_levels_of_hierarchy() {
	nodes := []Node{
		&Container{Type: "A", Start: 1, End: 11},
		&Container{Type: "B", Start: 2, End: 10},
		&Container{Type: "C", Start: 3, End: 9},
		&Container{Type: "D", Start: 4, End: 8},
		&Container{Type: "E", Start: 5, End: 7},
	}

	grouped := GroupNodes(&nodes)

	a := (*grouped)[0]
	suite.Equal("A", a.GetType())
	b := (*a.GetChildren())[0]
	suite.Equal("B", b.GetType())
	c := (*b.GetChildren())[0]
	suite.Equal("C", c.GetType())
	d := (*c.GetChildren())[0]
	suite.Equal("D", d.GetType())
	e := (*d.GetChildren())[0]
	suite.Equal("E", e.GetType())
}

func TestParser(t *testing.T) {
	suite.Run(t, new(ParserTestSuite))
}
