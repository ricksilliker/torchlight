package main

import (
	"context"
	"github.com/ricksilliker/torchlight"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"path/filepath"
)

var ReferenceDepth bool
var ReferenceNamespaces bool
var ReferenceFullPath bool

var refCmd = &cobra.Command{
	Use:   "get-references",
	Short: "Get file references related to a scene.",
	Run: func(cmd *cobra.Command, args []string) {
		GetReferences()
	},
}

func init() {
	refCmd.Flags().BoolVar(&ReferenceDepth, "depth", true, "Get the depth associated with the references.")
	refCmd.Flags().BoolVar(&ReferenceNamespaces, "namespace", true, "Get the namespace associated with the references.")
	refCmd.Flags().BoolVar(&ReferenceFullPath, "fullpath", false, "Get the full file path to the references.")
	rootCmd.AddCommand(refCmd)
}

func GetReferences() {
	logger := torchlight.GetLogger(Verbose)
	ctx := context.WithValue(context.Background(), "logger", *logger)

	logger.WithFields(logrus.Fields{
		"use_depth": ReferenceDepth,
		"use_namespaces": ReferenceNamespaces,
		"use_fullpath": ReferenceFullPath,
		"scene_path": ScenePath,
	}).Info("Getting references.")

	sceneData, err := torchlight.GetMayaAsciiSceneData(ScenePath)
	if err != nil {
		logger.WithError(err).Error("could not retrieve file references")
		return
	}

	allRefs := torchlight.GetAllReferences(ctx, sceneData, ScenePath)
	for _, r := range allRefs {
		if !ReferenceFullPath {
			xfp := filepath.FromSlash(r.Filepath)
			r.Filepath = filepath.Base(xfp)
		}

		logger.WithFields(logrus.Fields{
			"data": r,
		}).Info("All file references.")
	}
}