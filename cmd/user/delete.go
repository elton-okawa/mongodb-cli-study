package user

import (
	"context"
	"log"
	"time"

	"alwaysremember.com/mongodb-cli/pkg/db"
	"alwaysremember.com/mongodb-cli/pkg/model"
	"github.com/spf13/cobra"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var deleteUser = &cobra.Command{
	Use:   "delete",
	Short: "Delete user",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		id := args[0]

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		client := db.CreateClient(ctx)
		defer client.Disconnect(ctx)

		collection := client.Database(databaseName).Collection(collectionName)

		objId, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			log.Fatalf("Error while converting id to objId: '%s'", err)
		}

		if res, err := collection.DeleteOne(ctx, model.User{ID: objId}); err == nil {
			log.Printf("Successfully delete '%v' documents", res.DeletedCount)
		} else {
			log.Fatalf("Error while calling collection.DeleteOne: '%s'", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(deleteUser)
}
