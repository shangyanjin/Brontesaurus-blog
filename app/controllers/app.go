package controllers

import (
  "fmt"
  "github.com/robfig/revel"
  "brontesaurus-blog/app/models"
)

type Application struct {
	*revel.Controller
}

func (c Application) Index() revel.Result {
  // a := models.Post{"A Title", "a-title", "a cool body", 0}
  // a.Validate(c.Validation)

  p := models.Post{}
  p.Last()
  fmt.Println(p)

  greeting := "Aloha World"
  return c.Render(greeting)
	// return c.Render()
}
