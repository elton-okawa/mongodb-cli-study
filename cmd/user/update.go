package user

import (
	"context"
	"log"
	"time"

	"mongodb-cli/pkg/db"
	"mongodb-cli/pkg/model"

	"github.com/spf13/cobra"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var updateUser = &cobra.Command{
	Use:   "update <id>",
	Short: "Update user data",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		id := args[0]

		// TODO handle these errors
		name, _ := cmd.Flags().GetString("name")
		age, _ := cmd.Flags().GetInt("age")

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		client := db.CreateClient(ctx)
		defer client.Disconnect(ctx)

		collection := client.Database(databaseName).Collection(collectionName)

		objId, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			log.Fatalf("Error while converting id to objId: '%s'", err)
		}

		if result, err := collection.UpdateByID(ctx, objId, bson.D{{Key: "$set", Value: model.User{Name: name, Age: age}}}); err == nil {
			log.Printf("Updated '%s' successfully, count: %d", id, result.ModifiedCount)
		} else {
			log.Fatalf("Error while calling collection.Find: '%s'", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(updateUser)

	updateUser.Flags().String("name", "", "User name to update")
	updateUser.Flags().Int("age", 0, "User age to update")
}
