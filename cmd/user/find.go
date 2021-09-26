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
	"go.mongodb.org/mongo-driver/mongo/options"
)

var findUser = &cobra.Command{
	Use:   "find",
	Short: "Find user",
	Run: func(cmd *cobra.Command, args []string) {
		// TODO handle these errors
		id, _ := cmd.Flags().GetString("id")
		name, _ := cmd.Flags().GetString("name")
		age, _ := cmd.Flags().GetInt("age")

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		client := db.CreateClient(ctx)
		defer client.Disconnect(ctx)

		collection := client.Database(databaseName).Collection(collectionName)

		opts := options.Find()
		opts.SetSort(bson.D{{Key: "age", Value: -1}})

		var objId primitive.ObjectID
		if len(id) > 0 {
			var err error
			objId, err = primitive.ObjectIDFromHex(id)
			if err != nil {
				log.Fatalf("Error while converting id to objId: '%s'", err)
			}
		} else {
			objId = primitive.NilObjectID
		}

		cursor, err := collection.Find(ctx, model.User{ID: objId, Name: name, Age: age}, opts)
		if err != nil {
			log.Fatalf("Error while calling collection.Find: '%s'", err)
		}

		var users []model.User
		if err = cursor.All(ctx, &users); err != nil {
			log.Fatalf("Error while calling cursor.All: '%s'", err)
		}

		for _, u := range users {
			log.Printf("%+v", u)
		}
	},
}

func init() {
	rootCmd.AddCommand(findUser)

	findUser.Flags().String("id", "", "User id to find")
	findUser.Flags().String("name", "", "User name to find")
	findUser.Flags().Int("age", 0, "User age to find")
}
