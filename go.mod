module github.com/hi019/fiber-boilerplate

go 1.15

require (
	github.com/facebook/ent v0.4.2
	github.com/go-playground/universal-translator v0.17.0 // indirect
	github.com/go-playground/validator v9.31.0+incompatible
	github.com/gofiber/fiber v1.14.3-0.20200908022242-0b8785f8d91d
	github.com/gofiber/session v1.2.5
	github.com/leodido/go-urn v1.2.0 // indirect
	github.com/mattn/go-sqlite3 v1.14.2
	github.com/mitchellh/go-homedir v1.1.0
	github.com/rs/zerolog v1.19.0
	github.com/spf13/cobra v1.0.0
	github.com/spf13/viper v1.4.0
	github.com/stretchr/testify v1.6.1
	golang.org/x/crypto v0.0.0-20200820211705-5c72a883971a
)

replace github.com/gofiber/session => ../session
