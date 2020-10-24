package cmd

import (
	"fmt"
	"github.com/jcoon97/go-compose/writer"
	"github.com/spf13/cobra"
)

const (
	envMySQLDatabase     = "MYSQL_DATABASE"
	envMySQLRootPassword = "MYSQL_ROOT_PASSWORD"
)

var (
	mysqlDatabase     string
	mysqlRootPassword string
)

var generateMySQLCmd = &cobra.Command{
	Use:   "mysql -d DATABASE -P ROOT_PASSWORD [-o OUTPUT_PATH]",
	Short: "Generates a simple Docker Compose file for a MySQL server",
	Run: func(cmd *cobra.Command, args []string) {
		e := make(map[string]string)
		e[envMySQLDatabase] = mysqlDatabase
		e[envMySQLRootPassword] = mysqlRootPassword

		s := writer.Service{
			Image:       "mysql:latest",
			Command:     "--default-authentication-plugin=mysql_native_password",
			Environment: e,
			Ports:       []string{"3306:3306/tcp"},
			Restart:     "always",
		}

		c := writer.NewConfiguration(s)
		cnt, path := c.WriteFile(outputPath)
		fmt.Printf("Successfully wrote %d byte(s) to the following location: %s\n", cnt, path)
	},
}

func init() {
	generateCmd.AddCommand(generateMySQLCmd)

	generateMySQLCmd.PersistentFlags().StringVarP(&mysqlDatabase, "database", "d", "", "Specifies the name of the database to be created")
	generateMySQLCmd.PersistentFlags().StringVarP(&mysqlRootPassword, "root-password", "P", "", "Specifies the root user password for the database")

	generateMySQLCmd.MarkPersistentFlagRequired("database")
	generateMySQLCmd.MarkPersistentFlagRequired("root-password")
}
