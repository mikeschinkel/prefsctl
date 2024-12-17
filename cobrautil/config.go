package cobrautil

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/mikeschinkel/prefsctl/stdlibex"
)

var (
	ConfigName    = "prefs"
	ConfigFileExt = "yaml"
)

const (
	ConfigFilepathLogArg    = "config_filepath"
	ConfigOptionsLogArg     = "config_options"
	ConfigInitializedLogArg = "config_initialized"
)

type ConfigArgs struct {
	Filepath string
	Options  OptionsMap
}

type Config interface {
	Initialize(Context) error
	Filepath() string
	Options() OptionsMap
}

var _ Config = (*config)(nil)

type config struct {
	filepath    string
	options     OptionsMap
	initialized bool
}

func ConfigLogArgs(c config) []any {
	return []any{
		ConfigFilepathLogArg, c.Filepath(),
		ConfigOptionsLogArg, c.Options(),
		ConfigInitializedLogArg, c.Initialized(),
	}
}

func (c *config) Options() OptionsMap {
	return c.options
}

func (c *config) Filepath() string {
	return c.filepath
}

func NewConfig(args *ConfigArgs) Config {
	if args.Options == nil {
		args.Options = make(OptionsMap)
	}
	return &config{
		filepath: args.Filepath,
		options:  args.Options,
	}
}

func (c *config) Initialized() bool {
	return c.initialized
}

func (c *config) Initialize(_ Context) (err error) {
	if c.initialized {
		goto end
	}
	c.initialized = true
end:
	return nil
}

func SaveConfig(ctx Context, cfg Config, file string) error {
	return stdlibex.MarshalJSONFile(ctx, file, cfg)
}

type OnErrFunc func(string, error)

func SaveConfigWithOnErr(ctx Context, cfg Config, onErr OnErrFunc) (err error) {
	fp := cfg.Filepath()
	err = SaveConfig(ctx, cfg, fp)
	if err != nil {
		onErr(fp, err)
	}
	return err
}

func configFilename() string {
	return fmt.Sprintf("%s.%s", ConfigName, ConfigFileExt)
}

func ConfigFilepath() (file string, _ error) {
	dir, err := ConfigDir()
	if err != nil {
		goto end
	}
	file = filepath.Join(dir, configFilename())
end:
	return file, err
}

func ConfigDir() (dir string, err error) {
	dir = os.Getenv("XDG_CONFIG_HOME")
	if dir != "" {
		goto end
	}
	dir = os.Getenv("HOME")
	if dir == "" {
		err = ErrNoHomeDirVar
		goto end
	}
	dir = filepath.Join(dir, ".config", ConfigName)
end:
	return dir, err
}
