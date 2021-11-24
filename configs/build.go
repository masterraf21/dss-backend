package configs

import "os"

type build struct {
	Version string
	Commit  string
}

func setupBuild() *build {
	v := &build{
		Version: os.Getenv("VERSION"),
		Commit:  os.Getenv("Commit"),
	}

	return v
}
