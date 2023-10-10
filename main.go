package main

import (
	"os"
	"sort"

	"github.com/pubgo/dix"
	"github.com/pubgo/funk/assert"
	"github.com/pubgo/funk/config"
	"github.com/pubgo/lava/core/logging"
	"github.com/pubgo/lava/core/orm"
	"github.com/urfave/cli/v3"
	"gitlab.asants.com/bas/migrate/cmds/migratecmd"
	"gitlab.asants.com/bas/migrate/migrations"

	_ "github.com/pubgo/lava/core/orm/drivers/postgres"
)

type conf struct {
	Log *logging.Config        `yaml:"logger"`
	Db  map[string]*orm.Config `yaml:"db"`
}

func main() {
	var di = dix.New(dix.WithValuesNull())
	di.Provide(config.Load[conf])
	di.Provide(orm.NewClients)
	di.Provide(logging.New)
	di.Provide(migrations.New)
	di.Provide(migratecmd.New)
	di.Inject(func(cmd []*cli.Command) {
		var app = &cli.App{
			Name:                   "migration",
			Suggest:                true,
			UseShortOptionHandling: true,
			Commands:               cmd,
		}
		sort.Sort(cli.FlagsByName(app.Flags))
		sort.Sort(cli.CommandsByName(app.Commands))
		assert.Must(app.Run(os.Args))
	})
}
