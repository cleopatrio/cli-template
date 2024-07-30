package core

import (
	"fmt"
	"os"
	"time"

	"github.com/oleoneto/redic/app"
	"github.com/oleoneto/redic/app/domain/protocols"
	dbsql "github.com/oleoneto/redic/app/pkg/repositories/sql"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func (c *CommandState) ConnectDatabase(cmd *cobra.Command, args []string) {
	dbpath := viper.Get("database.path")

	db, err := dbsql.ConnectDatabase(protocols.DBConnectOptions{
		Adapter:        protocols.SQLAdapter(cmd.Flag("adapter").Value.String()),
		DSN:            *c.Flags.DatabaseURL,
		Filename:       dbpath.(string),
		VerboseLogging: c.Flags.VerboseLogging,
	})
	if err != nil {
		panic(err)
	}

	c.Database = db

	app.New(protocols.DBConnectOptions{
		DB: c.Database,
	})
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
