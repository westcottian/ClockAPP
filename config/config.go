package config

import (
    "github.com/spf13/viper"
    "log"
    "github.com/fsnotify/fsnotify"
    "time"
)

type Reader interface {
    GetAllKeys() []string
    Get(key string) interface{}
    GetBool(key string) bool
    GetString(key string) string
}

var (
        Secstr string
        Minstr string
        Hrstr string
    )


type viperConfigReader struct {
    viper *viper.Viper
}

var ConfReader *viperConfigReader

func (v viperConfigReader) GetAllKeys() []string{
    return v.viper.AllKeys()
}

func (v viperConfigReader) Get(key string) interface{} {
    return v.viper.Get(key)
}

func (v viperConfigReader) GetBool(key string) bool {
    return v.viper.GetBool(key)
}

func (v viperConfigReader) GetString(key string) string {
    return v.viper.GetString(key)
}


func init() {
    v:= viper.New()
    v.SetConfigType("env")
    v.SetConfigName("timer")
    v.AddConfigPath(".")

    err := v.ReadInConfig()

    if err != nil {
        log.Panic("Not able to read configuration", err.Error())
    }

    ConfReader = &viperConfigReader{
        viper: v,
    }
    
    go func() {
        for {
            time.Sleep(time.Second * 5)
            v.WatchConfig()
            v.OnConfigChange(func(e fsnotify.Event) {
		Loadconfiguration()
            })
        }
    }()
}

func Loadconfiguration() {

    Secstr = ConfReader.GetString("second")
    Minstr = ConfReader.GetString("minute")
    Hrstr = ConfReader.GetString("hour")
}
