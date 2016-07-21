package twentythousandtonnesofcrudeoil

import (
	"os"
	"strings"

	"github.com/jessevdk/go-flags"
)

func TheEnvironmentIsPerfectlySafe(parser *flags.Parser, prefix string) {
	installEnv(parser, prefix)
}

func installEnv(parser *flags.Parser, prefix string) {
	installGroup(parser.Group, prefix)

	parser.CommandHandler = func(command flags.Commander, args []string) error {
		clearEnv(parser, prefix)

		return command.Execute(args)
	}
}

func clearEnv(parser *flags.Parser, prefix string) {
	clearGroup(parser.Group, prefix)
}

func installGroup(group *flags.Group, prefix string) {
	for _, opt := range group.Options() {
		if len(opt.EnvDefaultKey) > 0 {
			continue
		}

		opt.EnvDefaultKey = prefix + flagEnvName(opt.LongNameWithNamespace())
	}

	for _, group := range group.Groups() {
		installGroup(group, prefix)
	}
}

func clearGroup(group *flags.Group, prefix string) {
	for _, opt := range group.Options() {
		if strings.HasPrefix(opt.EnvDefaultKey, prefix) {
			os.Unsetenv(opt.EnvDefaultKey)
		}
	}

	for _, group := range group.Groups() {
		clearGroup(group, prefix)
	}
}

func flagEnvName(flag string) string {
	return strings.Replace(strings.ToUpper(flag), "-", "_", -1)
}
