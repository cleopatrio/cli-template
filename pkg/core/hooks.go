package core

import (
	"clitemplate/pkg/dbsql"
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func (c *CommandState) ConnectDatabase(cmd *cobra.Command, args []string) {
	dbpath := viper.Get("database.path")

	db, err := dbsql.ConnectDatabase(dbsql.DBConnectOptions{
		Adapter:        dbsql.SQLAdapter(cmd.Flag("adapter").Value.String()),
		DSN:            *c.Flags.DatabaseURL,
		Filename:       dbpath.(string),
		VerboseLogging: c.Flags.VerboseLogging,
	})
	if err != nil {
		panic(err)
	}

	c.Database = db
}

func (c *CommandState) BeforeHook(cmd *cobra.Command, args []string) {
	c.SetFormatter(cmd, args)

	if !c.Flags.TimeExecutions {
		return
	}

	c.ExecutionStartTime = time.Now()
}

func (c *CommandState) AfterHook(cmd *cobra.Command, args []string) {
	if !c.Flags.TimeExecutions {
		return
	}

	fmt.Fprintln(
		os.Stderr,
		append([]any{"Elapsed time:", time.Since(c.ExecutionStartTime)}, c.ExecutionExitLog...)...,
	)
}
