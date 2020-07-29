// lib package provides base functions for skp

package lib

import (
	. "github.com/batkiz/skp/utils"
	"github.com/tidwall/gjson"
	. "io/ioutil"
	"os"
	"path/filepath"
)

// ReadManifest
func ReadManifest(path string) Manifest {
	json, err := ReadFile(path)

	HandleErr(err)

	app, err := UnmarshalManifest(json)
	HandleErr(err)

	return app
}

func loadConfig(file string) string {
	config, err := ReadFile(file)
	HandleErr(err)

	return string(config)
}

func getConfig(name string) string {
	return gjson.Get(scoopConfig, name).String()
}



// core envs

func getScoopDir() string {
	if os.Getenv("SCOOP") != "" {
		return os.Getenv("SCOOP")
	} else if getConfig("rootPath") != "" {
		return getConfig("rootPath")
	} else {
		return filepath.Join(os.Getenv("USERPROFILE"), "SCOOP")
	}
}

func getGlobalDir() string {
	if os.Getenv("SCOOP_GLOBAL") != "" {
		return os.Getenv("SCOOP_GLOBAL")
	} else if getConfig("globalPath") != "" {
		return getConfig("globalPath")
	} else {
		return filepath.Join(os.Getenv("ProgramData"), "SCOOP")
	}
}

func getCacheDir() string {
	if os.Getenv("SCOOP_CACHE") != "" {
		return os.Getenv("SCOOP_CACHE")
	} else if getConfig("cachePath") != "" {
		return getConfig("cachePath")
	} else {
		return filepath.Join(ScoopDir, "cache")
	}
}

func getConfigHome() string {
	if os.Getenv("XDG_CONFIG_HOME") != "" {
		return os.Getenv("XDG_CONFIG_HOME")
	} else {
		return filepath.Join(os.Getenv("USERPROFILE"), ".config")
	}
}

var ScoopDir = getScoopDir()
var GlobalDir = getGlobalDir()
var CacheDir = getCacheDir()
var ConfigFile = filepath.Join(getConfigHome(), "scoop", "config.json")
var scoopConfig = loadConfig(ConfigFile)


