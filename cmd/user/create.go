package user

import (
	"context"
	"log"
	"time"

	"alwaysremember.com/mongodb-cli/pkg/db"
	"github.com/spf13/cobra"
	"go.mongodb.org/mongo-driver/bson"
)

var name string

var createUser = &cobra.Command{
	Use:   "create",
	Short: "Create an user",
	Run: func(cmd *cobra.Command, args []string) {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		client := db.CreateClient(ctx)
		defer client.Disconnect(ctx)

		collection := client.Database(databaseName).Collection(collectionName)

		res, err := collection.InsertOne(ctx, bson.D{
			{Key: "name", Value: name},
		})

		if err != nil {
			log.Fatalf("Error while creating user: %s\n", err)
		}

		log.Printf("Added '%s' to database '%s' and collection '%s'\n", res.InsertedID, databaseName, collectionName)
	},
}

func init() {
	rootCmd.AddCommand(createUser)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// templateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	createUser.Flags().StringVarP(&name, "name", "n", "", "User name")
}
