package models

import (
  "fmt"
  "github.com/robfig/revel"
)

type Post struct {
  Title              string
  Slug               string
  Body               string
  CreatedAt          int
}

func (post *Post) Validate(v *revel.Validation) {
  v.Check(post.Title,
    revel.Required{},
  )

  v.Check(post.Slug,
    revel.Required{},
  )

  v.Check(post.Body,
    revel.Required{},
  )
}

func (post *Post) Save() {

}

func (post *Post) Last(){
  val, found := revel.Config.String("app.name")

  if(!found) { return }

  fmt.Println(val)

  post.Title = "Hi"
}
