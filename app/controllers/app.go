package controllers

import (
  // "fmt"
  "github.com/robfig/revel"
  "brontesaurus-blog/app/models"
)

type Application struct {
	*revel.Controller
}

func (c Application) Index() revel.Result {
  // p := models.Post{"A Title", "a-title", "a cool body", 0}
  // p.Validate(c.Validation)
  // fmt.Println(p)

  models.Post.Last()

  greeting := "Aloha World"
  return c.Render(greeting)
	// return c.Render()
}
