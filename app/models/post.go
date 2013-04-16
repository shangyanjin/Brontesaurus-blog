package models

import (
  "fmt"
  "encoding/json"
  "github.com/garyburd/redigo/redis"
  "github.com/robfig/revel"
)

var c redis.Conn

type Post struct {
  Title              string
  Slug               string
  Body               string
  CreatedAt          int
}

func log_fatal(err error) {
  if err != nil {
    fmt.Println(err)
  }
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

  c, err := redis.Dial("tcp", ":6379")
  log_fatal(err)

  post_key, err := redis.String(c.Do("lindex", val+":posts", 0))
  log_fatal(err)

  data, err := redis.Bytes(c.Do("get", post_key))

  er := json.Unmarshal(data, &post)
  log_fatal(er)
}
