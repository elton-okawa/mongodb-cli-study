package cmd

import (
	"context"
	"log"
	"time"

	"alwaysremember.com/mongodb-cli/pkg/db"
	"github.com/spf13/cobra"
	"go.mongodb.org/mongo-driver/bson"
)

var databaseName, collectionName, id string

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a MongoDB document",
	Long:  `Create a MongoDB document`,
	Run: func(cmd *cobra.Command, args []string) {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		client := db.CreateClient(ctx)
		defer client.Disconnect(ctx)

		collection := client.Database(databaseName).Collection(collectionName)

		res, err := collection.InsertOne(ctx, bson.D{
			{Key: "id", Value: id},
		})

		if err != nil {
			log.Fatalf("Error while creating document: %s\n", err)
		}

		log.Printf("Added '%s' to database '%s' and collection '%s'\n", res.InsertedID, databaseName, collectionName)
	},
}

func init() {
	rootCmd.AddCommand(createCmd)

	createCmd.Flags().StringVarP(&databaseName, "database", "d", "", "Database name")
	createCmd.Flags().StringVarP(&collectionName, "collection", "c", "", "Collection name")
	createCmd.Flags().StringVarP(&id, "id", "i", "", "Document id")
}
