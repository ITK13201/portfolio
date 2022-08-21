package createsuperuser

import (
	"context"
	"fmt"
	"github.com/ITK13201/portfolio/backend/infrastructure"
	"github.com/ITK13201/portfolio/backend/internal/crypto"
	"github.com/spf13/cobra"
	"golang.org/x/term"
	"syscall"
)

func Run(cmd *cobra.Command, args []string) {
	var username string
	var password string
	var passwordAgain string

	sqlClient := infrastructure.NewSqlClient()
	ctx := context.Background()

	for {
		fmt.Printf("Enter Username: ")
		fmt.Scan(&username)
		for {
			fmt.Printf("Enter Password: ")
			bytePassword, _ := term.ReadPassword(int(syscall.Stdin))
			password = string(bytePassword)
			fmt.Println()
			fmt.Printf("Enter Password (again): ")
			bytePasswordAgain, _ := term.ReadPassword(int(syscall.Stdin))
			passwordAgain = string(bytePasswordAgain)
			if password == passwordAgain {
				break
			}
		}
		hashedPassword := crypto.HashAndSalt(password)
		_, err := sqlClient.User.Create().SetUsername(username).SetHashedPassword(hashedPassword).Save(ctx)
		if err == nil {
			break
		}
	}
	fmt.Printf("Superuser created Successfully.\n")
}
