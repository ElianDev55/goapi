package env



var namesDirs = []string{
	"internal",
	"pkg",
	"pkg/db",
	"tools",
	"tools/seed",
	"tools/config",
}

var nameFiles = []string{
	"main.go",
	"README.md",
	".env",
	".env.example",
}


func GenerateNamesDirs() []string {
	return namesDirs
}

func GenerateNameFiles() []string {
	return nameFiles
}
