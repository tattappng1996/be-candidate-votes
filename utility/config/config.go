package config

import (
	"log"
	"os"
	"path/filepath"
	"reflect"
	"strings"

	"github.com/spf13/viper"
)

// Loan config from external file
func LoadFileConfig[T any]() *T {
	path := os.Getenv("CONFIG_PATH")
	if len(path) == 0 {
		log.Fatal("Can't get CONFIG_PATH from ENV")
	}
	dir := filepath.Dir(path)
	filebase := filepath.Base(path)
	// file name without extension
	filename := strings.TrimSuffix(filebase, filepath.Ext(filebase))

	viper.SetConfigName(filename)
	viper.AddConfigPath(dir)
	viper.AutomaticEnv()
	viper.SetConfigType("yaml")
	// convert _ to dot in env variable
	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_", ".", "_"))

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("fatal error config file: %+v", err)
	}

	var cfg T
	bindEnvs("env", cfg)
	if err := viper.Unmarshal(&cfg); err != nil {
		log.Fatalf("unable to decode config into struct, %+v", err)
	}

	return &cfg
}

// bindEnvs function will bind ymal file to struc model
func bindEnvs(tag string, iface interface{}, parts ...string) {
	ifv := reflect.ValueOf(iface)
	ift := reflect.TypeOf(iface)
	for i := 0; i < ift.NumField(); i++ {
		v := ifv.Field(i)
		t := ift.Field(i)
		tv, ok := t.Tag.Lookup(tag)
		if !ok {
			continue
		}
		switch v.Kind() {
		case reflect.Struct:
			bindEnvs(tag, v.Interface(), append(parts, tv)...)
		default:
			viper.BindEnv(strings.Join(append(parts, tv), "."))
		}
	}
}
