package main

import (
  "fmt"
)

type User struct {
  FullName string
}

func NewUser(fullName string) *User {
  return &User{FullName: fullName}
}

func main() {
  john := NewUser("John Doe")
  jane := NewUser("Jane Doe")
  alsoJane := NewUser("Jane Smith")
  fmt.Println(john.FullName)
  fmt.Println(jane.FullName)
  fmt.Println("Memory taken by users:",
    len([]byte(john.FullName)) +
      len([]byte(alsoJane.FullName)) +
      len([]byte(jane.FullName)))
}