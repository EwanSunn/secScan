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
		f.String("e", "email", "", "fofa account email")
		f.String("k", "key", "", "fofa account key")
		f.String("t", "target", "", "fofa search target")
		f.Int("z", "size", 20, "fofa search size")
	},
}

func runFofa(ctx *grumble.Context) (err error) {
	if ctx.Flags.Bool("show") == false {
		email := ctx.Flags.String("email")
		key := ctx.Flags.String("key")
		target := ctx.Flags.String("target")
		size := ctx.Flags.Int("size")
		fofa.Run(email, key, target, size)
	} else {
		fofa.ShowSyntax()
	}

	return err
}
