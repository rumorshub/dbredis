package dbredis

import (
	"github.com/roadrunner-server/endure/v2/dep"
	"github.com/roadrunner-server/errors"
)

const PluginName = "db.redis"

type Plugin struct {
	maker RedisMaker
}

func (p *Plugin) Init(cfg Configurer) error {
	const op = errors.Op("db.redis_plugin_init")

	if !cfg.Has(PluginName) {
		return errors.E(op, errors.Disabled)
	}

	return nil
}

func (p *Plugin) Provides() []*dep.Out {
	return []*dep.Out{
		dep.Bind((*Maker)(nil), p.RedisMaker),
	}
}

func (p *Plugin) RedisMaker() RedisMaker {
	return p.maker
}

func (p *Plugin) Name() string {
	return PluginName
}
