package main

import (
	"encoding/base64"
	"strconv"
	"unicode"

	log "github.com/Masterminds/log-go"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/helmet/v2"
	"github.com/gofrs/uuid"
)

// Create a Version 4 UUID, panicking on error.
// Use this form to initialize package-level variables.
var u1 = uuid.Must(uuid.NewV4())

// Constants
const GatewayVersion = "GATEWAY SERVICE API v0.0.1"
const GatewayDEBUG = true

//https://stackoverflow.com/questions/22593259/check-if-string-is-int
func isInt(s string) bool {
	for _, c := range s {
		if !unicode.IsDigit(c) {
			return false
		}
	}
	return true
}

//https://play.golang.org/p/RnEBFCJ9h0
func IsBase64(s string) bool {
	_, err := base64.StdEncoding.DecodeString(s)
	return err == nil
}

//https://gist.github.com/is73/de4f38e1d8da157fe33e
func BytesToString(data []byte) string {
	return string(data[:])
}

func StringToInt(data string) int {
	n, _ := strconv.Atoi(data)
	return n
}

// https://stackoverflow.com/questions/56140620/how-to-get-original-file-size-from-base64-encode-string
func calcOrigBinaryLength(datas string) int {

	l := len(datas)

	// count how many trailing '=' there are (if any)
	eq := 0
	if l >= 2 {
		if datas[l-1] == '=' {
			eq++
		}
		if datas[l-2] == '=' {
			eq++
		}

		l -= eq
	}

	// basically:
	//
	// eq == 0 :    bits-wasted = 0
	// eq == 1 :    bits-wasted = 2
	// eq == 2 :    bits-wasted = 4

	// each base64 character = 6 bits

	// so orig length ==  (l*6 - eq*2) / 8

	return (l*3 - eq) / 4
}

//MAIN
func main() {
	// Use an external setup function in order
	// to configure the app in tests as well
	app := Setup()

	// Start server with https/ssl enabled on http://localhost:443
	log.Fatal(app.Listen(":3001"))
	//Use and see: https://github.com/gofiber/fiber/issues/1077
	//log.Fatal(app.ListenTLS(":3000", SSLCertPath, SSLKeyPath))
}

// Setup Setup a fiber app with all of its routes
func Setup() *fiber.App {
	// Fiber instance
	app := fiber.New()

	if GatewayDEBUG {
		// Default middleware config
		app.Use(logger.New())
		// Log config
		app.Use(logger.New(logger.Config{
			Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
		}))
	}

	// Enable Helmet (Anti XSS)
	app.Use(helmet.New())

	// Routes
	app.Get("/", hello)
	app.Get("/version", version)

	// Return the configured app
	return app
}

//Generic response
func hello(c *fiber.Ctx) error {
	// Create a Version 4 UUID.
	u2, err := uuid.NewV4()
	if err != nil {
		log.Errorf("Failed to generate UUID: %v", err)
	}
	log.Infof("Generated Version 4 UUID %v", u2)
	return c.SendString("Hello, World 👋! " + u2.String())
}

//Return Gateway Version
func version(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"success": true,
		"message": GatewayVersion,
	})
}
