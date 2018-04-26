package apollo

import (
	"github.com/ServiceComb/go-archaius"
	"github.com/ServiceComb/go-archaius/core"
	"github.com/ServiceComb/go-archaius/sources/apollo"
	"github.com/ServiceComb/go-chassis/bootstrap"
	"github.com/ServiceComb/go-chassis/core/archaius"
	"github.com/ServiceComb/go-chassis/core/lager"
	"github.com/zouyx/agollo"
)

func init() {
	bootstrap.InstallPlugin("apollo",
		bootstrap.BootstrapFunc(InitApolloConfigSource))
}

func InitApolloConfigSource() error {
	namespace := "application"
	cluster := "DEV"
	appID := "go-chassis"
	ip := "localhost:8080"

	ac := &agollo.AppConfig{
		AppId:         appID,
		Cluster:       cluster,
		NamespaceName: namespace,
		Ip:            ip,
	}
	apolloSource := apollo.NewApolloConfigSource(ac)
	if err := archaius.DefaultConf.ConfigFactory.AddSource(apolloSource); err != nil {
		lager.Logger.Error("failed to do add source operation!!", err)
		return err
	}

	listener := &ApolloEventListener{
		Name:    "apollo",
		Factory: archaius.DefaultConf.ConfigFactory,
	}
	archaius.DefaultConf.ConfigFactory.RegisterListener(listener, "a*")

	return nil
}

type ApolloEventListener struct {
	Name    string
	Factory goarchaius.ConfigurationFactory
}

//Event is a method
func (e ApolloEventListener) Event(event *core.Event) {
	value := e.Factory.GetConfigurationByKey(event.Key)
	lager.Logger.Infof("config value %s | %s", event.Key, value)
}
