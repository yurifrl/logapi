package main

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	"github.com/yurifrl/logapi/pkg/conf"
	"github.com/yurifrl/logapi/pkg/file"
	"github.com/yurifrl/logapi/pkg/file/fileserver"
	"github.com/yurifrl/logapi/pkg/server"
	"github.com/yurifrl/logapi/pkg/store"
)

func addServerCmd(s *server.Server) *cobra.Command {
	return &cobra.Command{
		Use:   "server",
		Short: "Starts and http server",
		Long:  ``,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if fileName == "" {
				return fmt.Errorf("File parameter not set")
			}
			return nil
		},
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

			// @TODO: Get this from config
			// @TODO: Sync on file change
			err := f.Sync(fileName)
			if err != nil {
				logrus.Error(os.Stderr, err)
				os.Exit(2)
			}

			// Setup fileserver
			err = fileserver.Setup(logrus.New(), s.Router(), fileStore, f)
			if err != nil {
				logrus.Error(os.Stderr, err)
				os.Exit(2)
			}

			// Start server
			err = s.ListenAndServe()
			if err != nil {
				logrus.Panic(os.Stderr, err)
				os.Exit(2)
			}
			<-conf.Stop.Chan() // Wait until StopChan
			conf.Stop.Wait()   // Wait until everyone cleans up
		},
	}
}

func init() {
	s, err := server.New(logrus.New())
	if err != nil {
		logrus.Panic(os.Stderr, err)
		panic(-1)
	}

	cobra.OnInitialize(func() {
		initEnvs(s)
		v.AutomaticEnv()
	})

	RootCmd.AddCommand(addServerCmd(s))
}

func initEnvs(s *server.Server) {

	// Take the envs and load it on the struct
	err := v.Unmarshal(s)
	if err != nil {
		logrus.Panic(os.Stderr, err)
		panic(-1)
	}
}
