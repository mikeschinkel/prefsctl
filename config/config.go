package config

type Config interface {
	Dir() string
	Filepath() string
	OtherFilepath(filename string) string
}
