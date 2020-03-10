### Install
```
go get -u github.com/gofiber/fiber
go get -u github.com/gofiber/helmet
```
### Example
```go
package main

import (
  "github.com/gofiber/fiber"
  "github.com/gofiber/helmet"
)

func main() {
  app := fiber.New()

  app.Use(helmet.New())

  app.Get("/", func(c *fiber.Ctx) {
    c.Send("Welcome!")
  })

  app.Listen(3000)
}
```
### Test
```curl
curl -I http://localhost:3000
```
