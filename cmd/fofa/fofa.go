package fofa

import (
	"github.com/EwanSunn/secScan/internal/fofa"
	"github.com/desertbit/grumble"
)

var Fofa = &grumble.Command{
	Name:  "fofa",
	Help:  "search with fofa.info",
	Usage: "fofa --email --key -t",
	Run:   runFofa,
	Flags: func(f *grumble.Flags) {
		f.Bool("s", "show", false, "show fofa syntax")
		f.String("c", "config", "./conf.yml", "fofa config file")
		f.String("t", "target", "", "fofa search target")
		f.Int("z", "size", 20, "fofa search size")
		f.String("o", "output", "./result/fofa.csv", "fofa result file")
	},
}

func runFofa(ctx *grumble.Context) (err error) {
	if ctx.Flags.Bool("show") == false {
		config := ctx.Flags.String("config")
		target := ctx.Flags.String("target")
		size := ctx.Flags.Int("size")
		output := ctx.Flags.String("output")
		fofa.Run(config, target, output, size)
	} else {
		fofa.ShowSyntax()
	}

	return err
}
