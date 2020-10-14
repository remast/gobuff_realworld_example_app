package grifts

import (
	"gobuff_realworld_example_app/actions"

	"github.com/gobuffalo/buffalo"
)

func init() {
	buffalo.Grifts(actions.App())
}
