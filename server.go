package main
import (
  "github.com/labstack/echo"
)
func main() {
  e := echo.New() // 1
  e.Logger.Fatal(e.Start(":8080")) // 2
}
