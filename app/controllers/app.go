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
  post := models.Post{}
  post.Last()
  fmt.Println(post.Title)

  return c.Render(post)
}
