package v2ray

import (
	"io"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	mobasset "golang.org/x/mobile/asset"

	v2core "github.com/v2fly/v2ray-core/v5"
	v2filesystem "github.com/v2fly/v2ray-core/v5/common/platform/filesystem"
	v2serial "github.com/v2fly/v2ray-core/v5/infra/conf/serial"
	_ "github.com/v2fly/v2ray-core/v5/main/distro/all"

	v2applog "github.com/v2fly/v2ray-core/v5/app/log"
	v2commlog "github.com/v2fly/v2ray-core/v5/common/log"
)

const (
	v2Asset = "v2ray.location.asset"
)

var (
	server v2core.Server
)

func start(cfgStr string) {
	config, err := v2serial.LoadJSONConfig(strings.NewReader(cfgStr))
	if err != nil {
		log.Fatalf("v2ray load config err:%s", err.Error())
		return
	}

	instance, err := v2core.New(config)
	if err != nil {
		log.Fatalf("create v2ray core err:%s", err.Error())
		return
	}

	if err := instance.Start(); err != nil {
		log.Fatalf("start v2ray core err:%s", err.Error())
		return
	}
	server = instance
	runtime.GC()
	return
}

// Delegate Funcation
func TestConfig(ConfigureFileContent string) error {
	_, err := v2serial.LoadJSONConfig(strings.NewReader(ConfigureFileContent))
	return err
}

func Start(assetPath string, cfgStr string) {
	if len(assetPath) > 0 {
		os.Setenv(v2Asset, assetPath)
	}
	v2filesystem.NewFileReader = func(path string) (io.ReadCloser, error) {
		if _, err := os.Stat(path); os.IsNotExist(err) {
			_, file := filepath.Split(path)
			return mobasset.Open(file)
		}
		return os.Open(path)
	}

	v2applog.RegisterHandlerCreator(v2applog.LogType_Console,
		func(lt v2applog.LogType,
			options v2applog.HandlerCreatorOptions) (v2commlog.Handler, error) {
			return v2commlog.NewLogger(createStdoutLogWriter()), nil
		})

	Stop()
	start(cfgStr)
}

func Stop() {
	if server != nil {
		server.Close()
		server = nil
	}
}
