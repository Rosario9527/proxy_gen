package v2ray

import (
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"

	mobasset "golang.org/x/mobile/asset"

	v2core "github.com/v2fly/v2ray-core/v4"
	v2filesystem "github.com/v2fly/v2ray-core/v4/common/platform/filesystem"
	v2serial "github.com/v2fly/v2ray-core/v4/infra/conf/serial"
	_ "github.com/v2fly/v2ray-core/v4/main/distro/all"

	v2applog "github.com/v2fly/v2ray-core/v4/app/log"
	v2commlog "github.com/v2fly/v2ray-core/v4/common/log"
)

const (
	v2Asset = "v2ray.location.asset"
)

var pointInstance = &v2RayPoint{}

type v2RayPoint struct {
	v2rayOP sync.Mutex

	v2coreInstance *v2core.Instance
}

func (v *v2RayPoint) zero() {
	v.v2coreInstance = nil
}

func (v *v2RayPoint) start(cfgStr string) {
	config, err := v2serial.LoadJSONConfig(strings.NewReader(cfgStr))
	if err != nil {
		log.Fatalf("v2ray load config err:%s", err.Error())
		return
	}
	log.Println("XX v2ray start config", config)

	instance, err := v2core.New(config)
	if err != nil {
		v.v2coreInstance = nil
		log.Fatalf("create v2ray core err:%s", err.Error())
		return
	}

	log.Println("start v2ray core")
	if err := instance.Start(); err != nil {
		log.Fatalf("start v2ray core err:%s", err.Error())
		return
	}
	v.v2coreInstance = instance
	return
}

func (v *v2RayPoint) stop() {
	defer v.zero()
	if v.v2coreInstance != nil {
		err := v.v2coreInstance.Close()
		log.Printf("v2ray stop err:%v\n", err)
	}
	return
}

//Delegate Funcation
func TestConfig(ConfigureFileContent string) error {
	_, err := v2serial.LoadJSONConfig(strings.NewReader(ConfigureFileContent))
	return err
}

func Start(assetPath string, cfgStr string) {
	pointInstance.v2rayOP.Lock()
	defer pointInstance.v2rayOP.Unlock()
	if pointInstance.v2coreInstance != nil {
		log.Println("v2ray is running, to re-run, Stop first")
		return
	}

	if len(assetPath) > 0 {
		os.Setenv(v2Asset, assetPath)
	}
	//Now we handle read, fallback to gomobile asset (apk assets)
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

	pointInstance.start(cfgStr)
}

func Stop() {
	pointInstance.v2rayOP.Lock()
	defer pointInstance.v2rayOP.Unlock()
	pointInstance.stop()
}
