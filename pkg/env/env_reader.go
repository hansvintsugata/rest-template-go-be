package env

import (
	"errors"
	"reflect"

	"github.com/go-playground/validator/v10"
	"github.com/mcuadros/go-defaults"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
)

var (
	ErrInvalidDestination = errors.New("pointer to struct is expected")
	v                     *viper.Viper
)

func init() {
	v = viper.New()
	v.AddConfigPath(".")
	v.SetConfigFile(".env")
	v.SetConfigType("env")

	v.AutomaticEnv()
}

func MustProcess(dest interface{}) {
	err := Process(dest)
	if err != nil {
		panic(err)
	}
}

func Process(dest interface{}) (err error) {
	s := reflect.ValueOf(dest)

	if s.Kind() != reflect.Ptr {
		return ErrInvalidDestination
	}

	s = s.Elem()
	if s.Kind() != reflect.Struct {
		return ErrInvalidDestination
	}

	if err = BindEnv(dest, v); err != nil {
		return err
	}

	if err = v.ReadInConfig(); err != nil {
		return err
	}

	defaults.SetDefaults(dest)

	if err = v.Unmarshal(dest, func(config *mapstructure.DecoderConfig) {
		config.TagName = "env"
	}); err != nil {
		return err
	}

	validate := validator.New()
	err = validate.Struct(dest)

	return err
}

func BindEnv(dest interface{}, v *viper.Viper) error {
	t := reflect.TypeOf(dest).Elem()

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		tagValue := field.Tag.Get("env")
		err := v.BindEnv(tagValue)
		if err != nil {
			return err
		}
	}
	return nil
}
