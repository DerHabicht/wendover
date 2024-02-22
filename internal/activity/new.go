package activity

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/ag7if/go-files"
	"github.com/pkg/errors"

	"github.com/derhabicht/wendover/internal/config"
	"github.com/derhabicht/wendover/internal/logging"
)

func openEditor(configFile files.File) error {
	editor := os.Getenv("EDITOR")
	if editor == "" {
		return errors.New("failed to find editor; please set the $EDITOR environment variable")
	}

	cmd := exec.Command(editor, configFile.FullPath())
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Start()
	if err != nil {
		return errors.WithStack(err)
	}

	err = cmd.Wait()
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func copyDefaultConfig(name string, logger logging.Logger) (files.File, error) {
	cfgDir, err := config.CfgDir()
	if err != nil {
		return files.File{}, errors.WithStack(err)
	}

	defaultCfg, err := files.NewFile(filepath.Join(cfgDir, "assets", "default_activity_config.yaml"), logger.DefaultLogger())
	if err != nil {
		return files.File{}, errors.WithStack(err)
	}

	newCfg, err := defaultCfg.CopyAndRename(filepath.Join(cfgDir, "cfg", "activities"), fmt.Sprintf("%s.yaml", name))
	if err != nil {
		msg := err.Error()
		if strings.Contains(msg, "file exists") {
			return files.File{}, errors.Errorf("An activity called %s already exists, please pick another name.", name)
		}

		return files.File{}, errors.WithStack(err)
	}

	return newCfg, nil
}

func CreateNewActivity(name string, logger logging.Logger) error {
	logger.Info().Msg("Creating new activity config...")
	actCfg, err := copyDefaultConfig(name, logger)
	if err != nil {
		return errors.WithStack(err)
	}

	err = openEditor(actCfg)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}
