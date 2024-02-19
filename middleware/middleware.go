// middleware/auth.go

package middleware

import (
   "backend-project/util"
   "github.com/gofiber/fiber/v2"
)

func AuthMiddleware(c *fiber.Ctx) error {
   cookie := c.Cookies("jwt")
   if cookie == "" {
      return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
         "message": "Unauthorized",
      })
   }

   userID, err := util.ParseJwt(cookie)
   if err != nil {
      return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
         "message": "Invalid token",
      })
   }

   c.Locals("userID", userID)
   return c.Next()
}
