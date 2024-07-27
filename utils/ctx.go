package utils

import (
	"slices"
)

type Ctx struct {
	Ignore []string
}

func NewCtx() *Ctx {
	return &Ctx{
		Ignore: []string{},
	}
}

func (ctx *Ctx) IsTypeIgnored(ignoredType string) bool {
	return slices.Contains(ctx.Ignore, ignoredType)
}

func (ctx *Ctx) AppendIgnored(ignoredType string) {
	if !slices.Contains(ctx.Ignore, ignoredType) {
		ctx.Ignore = append(ctx.Ignore, ignoredType)
	}
}
