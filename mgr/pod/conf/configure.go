package conf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the protoc-gen-go-gframe package it is being compiled against.
import (
	"github.com/BurntSushi/toml"
)

var (
	rootConf = &Root{}
)

type Root struct {
	// TODO mapping to config file
}

func GetRoot() Root {
	if rootConf == nil {
		return Root{}
	}
	return *rootConf
}

func BindToml(fpath string) error {
	r := Root{}
	_, err := toml.DecodeFile(fpath, &r)
	if err != nil {
		return err
	}
	rootConf = &r
	// TODO watch file modify event
	return nil
}
