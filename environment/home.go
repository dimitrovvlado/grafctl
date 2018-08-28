package environment

import (
	"os"
	"path/filepath"
	"runtime"
)

// Home describes the root location of the CLI configuration
type Home string

// String returns Home as a string.
// Implements fmt.Stringer.
func (h Home) String() string {
	return os.ExpandEnv(string(h))
}

// Path returns config's home folder with appended relative paths
func (h Home) Path(elem ...string) string {
	p := []string{h.String()}
	p = append(p, elem...)
	return filepath.Join(p...)
}

// ConfigFile returns the path to the config.json file
func (h Home) ConfigFile() string {
	return h.Path("config.json")
}

func homeDir() string {
	if runtime.GOOS == "windows" {

		// First prefer the HOME environmental variable
		if home := os.Getenv("HOME"); len(home) > 0 {
			if _, err := os.Stat(home); err == nil {
				return home
			}
		}
		if homeDrive, homePath := os.Getenv("HOMEDRIVE"), os.Getenv("HOMEPATH"); len(homeDrive) > 0 && len(homePath) > 0 {
			homeDir := homeDrive + homePath
			if _, err := os.Stat(homeDir); err == nil {
				return homeDir
			}
		}
		if userProfile := os.Getenv("USERPROFILE"); len(userProfile) > 0 {
			if _, err := os.Stat(userProfile); err == nil {
				return userProfile
			}
		}
	}
	return os.Getenv("HOME")
}
