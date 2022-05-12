package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/eCommSim/authentication/database"
	"github.com/gin-gonic/gin"
)

func Signup(c *gin.Context) {
	r := c.Request
	rw := c.Writer
	// extra error handling should be done at server side to prevent malicious attacks
	if _, ok := r.Header["Email"]; !ok {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte("Email Missing\n"))
		return
	}
	if _, ok := r.Header["Username"]; !ok {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte("Username Missing\n"))
		return
	}
	if _, ok := r.Header["Password"]; !ok {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte("Password Missing\n"))
		return
	}
	if _, ok := r.Header["Fullname"]; !ok {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte("Fullname Missing\n"))
		return
	}

	if _, ok := r.Header["Role"]; !ok {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte("Role Missing\n"))
		return
	}

	if err := database.AddUser(r.Header["Email"][0], r.Header["Username"][0], r.Header["Password"][0],
		r.Header["Fullname"][0], r.Header["Role"][0]); err != nil {
		rw.WriteHeader(http.StatusConflict)
		rw.Write([]byte(err.Error() + "\n"))
		return
	}

	rw.WriteHeader(http.StatusOK)
	rw.Write([]byte("User Created\n"))
}

func Getuser(c *gin.Context) {
	r := c.Request
	rw := c.Writer
	// extra error handling should be done at server side to prevent malicious attacks
	if _, ok := r.Header["Email"]; !ok {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte("email missing\n"))
		return
	}

	usr, ok := database.GetUser(r.Header["Email"][0], r.Header["Password"][0])
	if !ok {
		rw.WriteHeader(http.StatusConflict)
		rw.Write([]byte("user not found\n"))
		return
	}

	rw.WriteHeader(http.StatusOK)
	rw.Write([]byte(fmt.Sprintf("User found: %v\n", usr)))
}

func Signin(c *gin.Context) {
	r := c.Request
	rw := c.Writer
	// extra error handling should be done at server side to prevent malicious attacks
	if _, ok := r.Header["Email"]; !ok {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte("email missing\n"))
		return
	}

	usr, ok := database.GetUser(r.Header["Email"][0], r.Header["Password"][0])
	if !ok {
		rw.WriteHeader(http.StatusConflict)
		rw.Write([]byte("user not found\n"))
		return
	}

	validToken, err := database.GenerateJWT(usr.Email, usr.Role)
	if err != nil {
		c.Writer.WriteHeader(http.StatusInternalServerError)
		c.Writer.Write([]byte(err.Error() + "\n"))
	}

	var token database.Token
	token.Email = usr.Email
	token.Role = usr.Role
	token.TokenString = validToken

	c.Writer.WriteHeader(http.StatusOK)
	c.Writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(c.Writer).Encode(token)
}

// func isAuthorized(handler http.HandlerFunc) gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		if c.Request.Header["Token"] == nil {
// 			c.Writer.WriteHeader(http.StatusNotAcceptable)
// 			c.Writer.Write([]byte("not token found\n"))
// 			return
// 		}

// 		key := os.Getenv("SECRET_KEY")

// 		token, err := jwt.Parse(c.Request.Header["Token"][0], func(t *jwt.Token) (interface{}, error) {
// 			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
// 				return nil, fmt.Errorf("error parsing token")
// 			}
// 			return key, nil
// 		})

// 		if err != nil {
// 			c.Writer.WriteHeader(http.StatusNotAcceptable)
// 			c.Writer.Write([]byte("token expired\n"))
// 		}

// 		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
// 			switch claims["role"] {
// 			case "admin":
// 				c.Request.Header.Set("role", "admin")
// 				handler.ServeHTTP(c.Writer, c.Request)
// 			case "user":
// 				c.Request.Header.Set("role", "user")
// 				handler.ServeHTTP(c.Writer, c.Request)
// 			default:
// 				c.Writer.WriteHeader(http.StatusUnauthorized)
// 				c.Writer.Write([]byte("not authorize\n"))
// 			}
// 			// if claims["role"] == "admin" {
// 			// 	c.Request.Header.Set("role", "admin")
// 			// 	handler.ServeHTTP(c.Writer, c.Request)
// 			// 	return
// 			// } else if claims["role"] == "user" {
// 			// 	c.Request.Header.Set("role", "admin")
// 			// 	handler.ServeHTTP(c.Writer, c.Request)
// 			// 	return
// 			// }
// 		}
// 	}
// }
