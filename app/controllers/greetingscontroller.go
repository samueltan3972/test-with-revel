package controllers

import (
	"github.com/revel/revel"
)

type GreetingsController struct {
	*revel.Controller
}

func (c GreetingsController) Greetings() revel.Result {
	return c.RenderText("Hello World")
}
