package UserApi

import (
	"github.com/joho/godotenv"
	. "goa.design/goa/v3/dsl"
	"log"
	"os"
)

var _ = JWTSecurity("jwt", func() {
	Scope("api:access", "api:access")
})

// API describes the global properties of the API server.
var _ = API("User", func() {
	Title("User Service")
	Description("Service for User CRUD")
	Server("user", func() {
		Host("host", func() {
			// load .env file
			err := godotenv.Load(".env")
			if err != nil {
				log.Fatalf("Error loading .env file")
			}
			URI(os.Getenv("GOA_HOST_PORT"))
		})
	})
})

// Service describes a service
var _ = Service("user", func() {
	Description("CRUD for users")
	// Method describes a service method (endpoint)
	Method("create", func() {
		// Payload describes the method payload
		// Here the payload is an object that consists of two fields
		Payload(func() {
			Attribute("email", String, "User's email address", func() {
				Pattern(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
				Example("meridamotore@gmail.com")
			})
			Attribute("password", String, "User's password", func() {
				Example("MyP@ssw0rd")
			})
			Description("Create user")
		})
		// Result describes the method result
		// Here the result is Result(User)
		Result(String)
		// HTTP describes the HTTP transport mapping
		HTTP(func() {
			// Requests to the service consist of HTTP GET requests
			// The payload fields are encoded as path parameters
			POST("user/create")
			// Responses use a "200 OK" HTTP status
			// The result is encoded in the response body
			Response(func() {
				Code(StatusOK)
				Headers(func() {
					// "name" sent in the header metadata
					Attribute("Content-Type: application/json")
				})
			})
		})
	})
	Method("read", func() {
		Security("jwt", func() {
			Scope("api:access")
		})
		Payload(func() {
			Token("jwt", String, "JWT token")
		})
		Result(String)
		HTTP(func() {
			GET("user")
			Response(func() {
				Code(StatusOK)
				Headers(func() {
					Attribute("Content-Type: application/json")
				})
			})
		})
	})
	Method("update", func() {
		Security("jwt", func() {
			Scope("api:access")
		})
		Payload(func() {
			Token("jwt", String, "JWT token")
			Attribute("name", String, "User name")
			Attribute("surname", String, "User surname")
			Attribute("address", String, "User address")
			Attribute("phone", String, "phone")
			Description("Create user")
		})
		Result(String)
		HTTP(func() {
			PUT("user/update")
			Response(func() {
				Code(StatusOK)
				Headers(func() {
					Attribute("Content-Type: application/json")
				})
			})
		})
	})
	Method("delete", func() {
		Security("jwt", func() {
			Scope("api:access")
		})
		Payload(func() {
			Token("jwt", String, "JWT token")
			Attribute("id", UInt, "Delete id")
		})
		Result(String)
		HTTP(func() {
			DELETE("user/delete")
			Response(func() {
				Code(StatusOK)
				Headers(func() {
					Attribute("Content-Type: application/json")
				})
			})
		})
	})
	Method("token", func() {
		Payload(func() {
			Attribute("email", String, "User's email address", func() {
				Pattern(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
				Example("meridamotore@gmail.com")
			})
			Attribute("password", String, "User's password", func() {
				Example("MyP@ssw0rd")
			})
			Description("Get token")
		})
		Result(String)
		HTTP(func() {
			POST("auth/token")
			Response(func() {
				Code(StatusOK)
			})
		})
	})
	Files("/openapi3.json", "./gen/http/openapi3.json")
})
var User = ResultType("application/vnd.goa.user", func() {
	Attributes(func() {
		Attribute("id", UInt)
		Attribute("email", String)
		Attribute("password", String)
		Attribute("name", String)
		Attribute("surname", String)
		Attribute("phone", String)
		Attribute("address", String)
	})
})
