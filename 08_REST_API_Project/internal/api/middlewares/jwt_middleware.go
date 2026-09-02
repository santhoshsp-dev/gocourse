package middlewares

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"restapi/pkg/utils"

	"github.com/golang-jwt/jwt/v5"
)

// --------------- Start: 098 ----------------
// we need to use ContextKey because ctx showing warning like:- should not use built in type string as key for value.
// type ContextKey string // in video: 98, we moved this code from here to pkg/utils/authorize_user.go
// --------------- End: 098 ----------------

func JWTMiddleware(next http.Handler) http.Handler {
	fmt.Println("---------------------- JWT Middleware --------------------")
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("++++++++++++++++ Inside JWT Middleware")

		token, err := r.Cookie("Bearer")
		if err != nil {
			http.Error(w, "Authorization Header Missing", http.StatusUnauthorized)
			return
		}

		jwtSecret := os.Getenv("JWT_SECRET")

		parsedToken, err := jwt.Parse(token.Value, func(token *jwt.Token) (interface{}, error) {
			// Don't forget to validate the alg is what you expect:
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
			return []byte(jwtSecret), nil
			// }, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))
		})
		if err != nil {
			if errors.Is(err, jwt.ErrTokenExpired) {
				http.Error(w, "Token Expired", http.StatusUnauthorized)
				return
			} else if errors.Is(err, jwt.ErrTokenMalformed) {
				http.Error(w, "Token Malformed", http.StatusUnauthorized)
				return
			}
			utils.ErrorHandler(err, "")
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		if parsedToken.Valid {
			log.Println("Valid JWT")
		} else {
			http.Error(w, "Invalid Login Token", http.StatusUnauthorized)
			log.Println("Invalid JWT:", token.Value)
			return
		}

		claims, ok := parsedToken.Claims.(jwt.MapClaims)
		if !ok {
			http.Error(w, "Invalid Login Token", http.StatusUnauthorized)
			log.Println("Invalid Login Token:", token.Value)
			return
		}

		// Now let's use context to carry the claim information across different middlewares across different functions. Now what is context. Well in the context of our use context is associated with the request. The request object has a context method through which you can access the context, and context can contain key value based information. We have already used context when we were learning. Go in our earlier sections and we are going to use context in the same way. We are going to save key value pairs. However, this time we are going to access the request context.

		// --------------- Start: 098 ----------------
		ctx := context.WithValue(r.Context(), utils.ContextKey("role"), claims["role"])
		ctx = context.WithValue(ctx, utils.ContextKey("exppiresAt"), claims["exp"])
		ctx = context.WithValue(ctx, utils.ContextKey("username"), claims["user"])
		ctx = context.WithValue(ctx, utils.ContextKey("userId"), claims["uid"])
		// --------------- End: 098 ----------------

		fmt.Println(ctx)

		next.ServeHTTP(w, r.WithContext(ctx))
		fmt.Println("Sent Response from JWT Middleware")
	})

}
