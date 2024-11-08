package middleware

import (
    "github.com/labstack/echo/v4"
    "github.com/golang-jwt/jwt"
    "net/http"
    "strings"
)

func JWTAuth(next echo.HandlerFunc) echo.HandlerFunc {
    return func(c echo.Context) error {
        authHeader := c.Request().Header.Get("Authorization")
        if authHeader == "" {
            return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Missing Authorization Header"})
        }

        tokenString := strings.Split(authHeader, "Bearer ")[1]
        if tokenString == "" {
            return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Invalid token"})
        }

        token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
            // Validate signing method
            if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
                return nil, echo.NewHTTPError(http.StatusUnauthorized, "Invalid signing method")
            }
            return []byte("your_secret_key"), nil
        })

        if err != nil || !token.Valid {
            return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Invalid or expired token"})
        }

        // Attach user_id to the context
        claims := token.Claims.(*jwt.MapClaims)
        userID := (*claims)["user_id"].(float64)
        c.Set("user_id", int(userID))

        return next(c)
    }
}
