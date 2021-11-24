package configs

// Server config
var Server *server

// MongoDB Config
var MongoDB *mongodb

// Constant config
var Constant *constant

// Auth config
var Auth *auth

// Build config
var Build *build

func init() {
	Server = setupServer()
	MongoDB = setupMongoDB()
	Constant = setupConstant()
	Auth = setupAuth()
	Build = setupBuild()
}
