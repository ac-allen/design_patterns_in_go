package main

import "fmt"

type Bird struct {
  Age int
}

func (b *Bird) Fly() {
  if b.Age >= 10 {
    fmt.Println("Flying!")
  }
}

type Lizard struct {
  Age int
}

func (l *Lizard) Crawl() {
  if l.Age < 10 {
    fmt.Println("Crawling!")
  }
}

type Dragon struct {
  Bird
  Lizard
}

func (d *Dragon) Age() int {
	return d.Bird.Age
  }
  
  func (d *Dragon) SetAge(age int) {
	d.Bird.Age = age
	d.Lizard.Age = age
  }

func main() {
  d := Dragon{}
  d.SetAge(10)
  d.Bird.Age = 20
  fmt.Println(d.Bird.Age)
  fmt.Println(d.Lizard.Age)
  d.Fly()
  d.Crawl()
}