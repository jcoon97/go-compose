package cmd

import (
	"github.com/jcoon97/go-compose/writer"
	"github.com/spf13/cobra"
)

const envMongoDatabase = "MONGO_INITDB_DATABASE"

var mongoDatabase string

var generateMongoCmd = &cobra.Command{
	Use:   "mongo -d DATABASE",
	Short: "Generates a simple Docker Compose file for a Mongo server",
	Run: func(cmd *cobra.Command, args []string) {
		e := make(map[string]string)
		e[envMongoDatabase] = mongoDatabase

		s := writer.Service{
			Image:       "mongo:latest",
			Environment: e,
			Ports:       []string{"27017-27019:27017-27019"},
			Restart:     "always",
		}

		c := writer.NewConfiguration(s)
		c.WriteFile(outputPath)
	},
}

func init() {
	generateCmd.AddCommand(generateMongoCmd)

	generateMongoCmd.PersistentFlags().StringVarP(&mongoDatabase, "database", "d", "", "Specifies the name of the database to be created")
	generateMongoCmd.MarkPersistentFlagRequired("database")
}
