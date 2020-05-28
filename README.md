# Helmet

![Release](https://img.shields.io/github/release/gofiber/helmet.svg)
[![Discord](https://img.shields.io/badge/discord-join%20channel-7289DA)](https://gofiber.io/discord)
![Test](https://github.com/gofiber/helmet/workflows/Test/badge.svg)
![Security](https://github.com/gofiber/helmet/workflows/Security/badge.svg)
![Linter](https://github.com/gofiber/helmet/workflows/Linter/badge.svg)

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
