package cobrautil

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/mikeschinkel/prefsctl/logargs"
	"github.com/mikeschinkel/prefsctl/stdlibex"
)

const (
	ConfigName    = "prefs"
	ConfigFileExt = "yaml"
)

type ConfigArgs struct {
	Username string
	Password string
	Filepath string
	Options  OptionsMap
}

type Config struct {
	InternalUsername string                    `json:"username"`
	ServerHost       string                    `json:"server_host"`
	ServerPort       int                       `json:"server_port"`
	ExtraInfo        map[string]map[string]any `json:"extra_info"`
	listenAddress    string
	password         string
	filepath         string
	project          string
	dnsProvider      string
	options          OptionsMap
	initialized      bool
}

func ConfigLogArgs(c Config) []any {
	return []any{
		logargs.OptionsLogArg, c.Options(),
		logargs.InitializedLogArg, c.Initialized(),
		logargs.CacheAuthLogArg, c.CacheAuth(),
	}
}

func (c *Config) Options() OptionsMap {
	return c.options
}

func (c *Config) Filepath() string {
	return c.filepath
}

func NewConfig(args *ConfigArgs) *Config {
	if args == nil {
		args = &ConfigArgs{}
	}
	return &Config{
		options:          args.Options,
		InternalUsername: args.Username,
		password:         args.Password,
		filepath:         args.Filepath,
		ExtraInfo:        make(map[string]map[string]any),
	}
}

func (c *Config) ServerUsername() string {
	return c.InternalUsername
}

func (c *Config) ServerPassword() string {
	return c.password
}

func (c *Config) Initialized() bool {
	return c.initialized
}

// CacheAuth returns true as auth should be caches for the CLI
func (c *Config) CacheAuth() bool {
	return true
}

func (c *Config) Initialize(_ Context) (err error) {
	if c.initialized {
		goto end
	}
	c.initialized = true
end:
	return nil
}

func (c *Config) GetExtraInfo(group string) map[string]any {
	values, ok := c.ExtraInfo[group]
	if !ok {
		values = make(map[string]any)
	}
	return values
}

func (c *Config) AddExtraInfo(group string, value fmt.Stringer) {
	_, ok := c.ExtraInfo[group]
	if !ok {
		c.ExtraInfo[group] = make(map[string]any)
	}
	c.ExtraInfo[group][value.String()] = value
}
func (c *Config) DeleteExtraInfo(group string, value fmt.Stringer) {
	values, ok := c.ExtraInfo[group]
	if !ok {
		goto end
	}
	delete(values, value.String())
end:
}

func SaveConfig(ctx Context, cfg *Config, file string) error {
	return stdlibex.MarshalJSONFile(ctx, file, cfg)
}

type OnErrFunc func(string, error)

func SaveConfigWithOnErr(ctx Context, cfg *Config, onErr OnErrFunc) (err error) {
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
