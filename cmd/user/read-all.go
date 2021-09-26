package user

import (
	"context"
	"log"
	"time"

	"alwaysremember.com/mongodb-cli/pkg/db"
	"alwaysremember.com/mongodb-cli/pkg/model"
	"github.com/spf13/cobra"
)

var batchMode bool

var readAllUsers = &cobra.Command{
	Use:   "read-all",
	Short: "Read all user entries",
	Run: func(cmd *cobra.Command, args []string) {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		client := db.CreateClient(ctx)
		defer client.Disconnect(ctx)

		collection := client.Database(databaseName).Collection(collectionName)

		cursor, err := collection.Find(ctx, model.User{})
		if err != nil {
			log.Fatalf("Error while calling collection.Find: '%s'", err)
		}

		if batchMode {
			log.Println("Reading in batchMode")
			defer cursor.Close(ctx)
			for cursor.Next(ctx) {
				var user model.User
				if err = cursor.Decode(&user); err != nil {
					log.Fatalf("Error while calling cursor.Decode: '%s'", err)
				}

				log.Printf("%+v", user)
			}

		} else {
			log.Println("Reading all at once")

			var users []model.User
			if err = cursor.All(ctx, &users); err != nil {
				log.Fatalf("Error while calling cursor.All: '%s'", err)
			}

			for _, u := range users {
				log.Printf("%+v", u)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(readAllUsers)

	readAllUsers.Flags().BoolVar(&batchMode, "batch", false, "Read users in batch mode")
}
