package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	"github.com/yurifrl/logapi/pkg/file"
	"github.com/yurifrl/logapi/pkg/store"
)

func addReadCmd() *cobra.Command {

	cmd := &cobra.Command{
		Use:   "read",
		Short: "Starts and http server",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			// Create Store
			fileStore := store.Create()

			// Create File Reader
			fs := afero.NewOsFs()
			f := file.New(
				fs,
				logrus.New(),
				fileStore,
			)

			err := f.Sync(fileName)
			if err != nil {
				logrus.Error(os.Stderr, err)
				os.Exit(2)
			}

			all, err := fileStore.GetAll()
			if err != nil {
				logrus.Error(os.Stderr, err)
				os.Exit(2)
			}

			out, err := json.Marshal(all)
			if err != nil {
				logrus.Error(os.Stderr, err)
				os.Exit(2)
			}

			// Print to stdout
			fmt.Println(string(out))
		},
	}

	return cmd
}

func init() {
	cobra.OnInitialize(func() {
		v.AutomaticEnv()
	})

	RootCmd.AddCommand(addReadCmd())
}
