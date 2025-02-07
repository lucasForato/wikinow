package types

import (
	"wikinow/internal/parser"

	"github.com/stretchr/testify/suite"
)

type CtxSuite struct {
	suite.Suite
	ctx *parser.Ctx
}
