package main

import (
	"os"
	"path/filepath"

	"github.com/ag7if/go-files"
	"github.com/pkg/errors"

	"github.com/derhabicht/wendover/build"
	"github.com/derhabicht/wendover/internal/config"
	"github.com/derhabicht/wendover/internal/logging"
)

func CopyAssets(projectPath string, logger logging.Logger) error {
	defaultCfgDir := filepath.Join(projectPath, "config")

	cfgDir, err := config.CfgDir()
	if err != nil {
		return errors.WithStack(err)
	}
	destAssetDir := filepath.Join(cfgDir, "assets")

	logging.Info().Str("file", "default_activity_config.yaml").Msg("copying asset")
	actCfg, err := files.NewFile(filepath.Join(defaultCfgDir, "default_activity_config.yaml"), logger.DefaultLogger())
	if err != nil {
		return errors.WithStack(err)
	}
	_, err = actCfg.Copy(destAssetDir)
	err = build.ClearFileExistsError(err)
	if err != nil {
		logging.Warn().Err(err).Str("file", "default_activity_config.yaml").Msg("failed to copy asset")
	}

	return nil
}

func main() {
	logging.InitLogging("info", true)

	logger := logging.Logger{}

	err := build.CreateConfigDirectories()
	if err != nil {
		logging.Error().Err(err).Msg("failed to create config directories")
		os.Exit(1)
	}

	err = build.CreateCacheDirectories()
	if err != nil {
		logging.Error().Err(err).Msg("failed to create cache directories")
		os.Exit(1)
	}

	err = CopyAssets(os.Args[1], logger)
	if err != nil {
		logging.Error().Err(err).Msg("failed to copy assets")
		os.Exit(1)
	}
}
