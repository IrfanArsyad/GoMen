package main

import (
	"fmt"
	"os"
	"os/exec"

	"gomen/internal/generator"
)

const version = "1.0.0"

func main() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(0)
	}

	command := os.Args[1]

	switch command {
	case "serve":
		runGoCommand()

	case "migrate":
		runGoCommand("-migrate")

	case "seed":
		runGoCommand("-seed")

	case "make:controller":
		if len(os.Args) < 3 {
			fmt.Println("Error: Controller name is required")
			fmt.Println("Usage: gomen make:controller <ControllerName>")
			os.Exit(1)
		}
		generator.MakeController(os.Args[2])

	case "make:model":
		if len(os.Args) < 3 {
			fmt.Println("Error: Model name is required")
			fmt.Println("Usage: gomen make:model <ModelName>")
			os.Exit(1)
		}
		generator.MakeModel(os.Args[2])

	case "make:migration":
		if len(os.Args) < 3 {
			fmt.Println("Error: Migration name is required")
			fmt.Println("Usage: gomen make:migration <migration_name>")
			os.Exit(1)
		}
		generator.MakeMigration(os.Args[2])

	case "make:service":
		if len(os.Args) < 3 {
			fmt.Println("Error: Service name is required")
			fmt.Println("Usage: gomen make:service <ServiceName>")
			os.Exit(1)
		}
		generator.MakeService(os.Args[2])

	case "make:request":
		if len(os.Args) < 3 {
			fmt.Println("Error: Request name is required")
			fmt.Println("Usage: gomen make:request <RequestName>")
			os.Exit(1)
		}
		generator.MakeRequest(os.Args[2])

	case "make:middleware":
		if len(os.Args) < 3 {
			fmt.Println("Error: Middleware name is required")
			fmt.Println("Usage: gomen make:middleware <MiddlewareName>")
			os.Exit(1)
		}
		generator.MakeMiddleware(os.Args[2])

	case "make:seeder":
		if len(os.Args) < 3 {
			fmt.Println("Error: Seeder name is required")
			fmt.Println("Usage: gomen make:seeder <SeederName>")
			os.Exit(1)
		}
		generator.MakeSeeder(os.Args[2])

	case "make:resource":
		if len(os.Args) < 3 {
			fmt.Println("Error: Resource name is required")
			fmt.Println("Usage: gomen make:resource <ResourceName>")
			os.Exit(1)
		}
		generator.MakeResource(os.Args[2])

	case "list":
		printCommands()

	case "version", "-v", "--version":
		fmt.Printf("GoMen CLI v%s\n", version)

	case "help", "-h", "--help":
		printUsage()

	default:
		fmt.Printf("Unknown command: %s\n", command)
		fmt.Println("Run 'gomen list' to see available commands")
		os.Exit(1)
	}
}

func runGoCommand(args ...string) {
	cmdArgs := append([]string{"run", "main.go"}, args...)
	cmd := exec.Command("go", cmdArgs...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	if err := cmd.Run(); err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			os.Exit(exitErr.ExitCode())
		}
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}

func printUsage() {
	fmt.Println(`
   ____       __  __
  / ___| ___ |  \/  | ___ _ __
 | |  _ / _ \| |\/| |/ _ \ '_ \
 | |_| | (_) | |  | |  __/ | | |
  \____|\___/|_|  |_|\___|_| |_|

  GoMen CLI - Code Generator & Application Manager

Usage:
  gomen <command> [arguments]

Application Commands:
  serve                     Start the application server
  migrate                   Run database migrations
  seed                      Run database seeders

Generator Commands:
  make:controller <Name>    Create a new controller
  make:model <Name>         Create a new model
  make:migration <name>     Create a new migration file
  make:service <Name>       Create a new service
  make:request <Name>       Create a new request validation
  make:middleware <Name>    Create a new middleware
  make:seeder <Name>        Create a new seeder
  make:resource <Name>      Create model, controller, service, and request

Other Commands:
  list                      Show all available commands
  version                   Show CLI version
  help                      Show this help message

Examples:
  gomen serve
  gomen migrate
  gomen seed
  gomen make:controller Product
  gomen make:model Product
  gomen make:migration create_products_table
  gomen make:resource Product
`)
}

func printCommands() {
	fmt.Println(`
Application Commands:
  serve              Start the application server
  migrate            Run database migrations
  seed               Run database seeders

Generator Commands:
  make:controller    Create a new controller
  make:model         Create a new model
  make:migration     Create a new migration file
  make:service       Create a new service
  make:request       Create a new request validation
  make:middleware    Create a new middleware
  make:seeder        Create a new seeder
  make:resource      Create model, controller, service, and request (full resource)
`)
}
