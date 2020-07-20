package autoload

/*
	You can just read the .env file on import just by doing

		import _ "github.com/wallester/godotenv/autoload"

	And bob's your mother's brother
*/

import "github.com/wallester/godotenv"

func init() {
	godotenv.Load()
}
