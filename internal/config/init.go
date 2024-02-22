package config

import (
	"os"
	"path/filepath"

	"github.com/derhabicht/wendover/internal/logging"
)

func init() {
	cwd, _ := os.Getwd()
	logging.Trace().Str("cwd", cwd).Msg("")

	cfgDir, err := CfgDir()
	if err != nil {
		logging.Error().Err(err).Msgf("could not find configuration directory: %s", cfgDir)
	} else {
		logging.Debug().Str("usrCfgDir", cfgDir).Msgf("configuration directory found: %s", cfgDir)
	}

	SetConfigType("yaml")
	SetConfigName("config")
	AddConfigPath(filepath.Join(cfgDir, "cfg"))
	err = ReadInConfig()
	if err != nil {
		path := filepath.Join(cfgDir, "cfg", "config.yaml")
		logging.Warn().Err(err).Str("path", path).Msg("config file not found, creating a default config")
		err = WriteConfigAs(path)
		if err != nil {
			logging.Error().Err(err).Msg("failed to create default config file")
		}
	}
}
