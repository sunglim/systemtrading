package options

import (
	"fmt"
	"os"

	"github.com/prometheus/common/version"
	"github.com/spf13/cobra"
)

// A list of configuratable parameters.
type Options struct {
	Help bool `yaml:"help"`

	TelegramChatId int64  `yaml:"telegram_chat_id"`
	TelegramToken  string `yaml:"telegram_token"`

	KoreaInvestmentUrl     string `yaml:"koreainvestment_url"`
	KoreaInvestmentAppKey  string `yaml:"koreainvestment_appkey"`
	KoreaInvestmentSecret  string `yaml:"koreainvestment_appsecret"`
	KoreaInvestmentAccount string `yaml:"koreainvestment_account"`

	// filled out when the config file has relevant configs.
	BuyEveryDayIfBelowAverageConfig BuyEveryDayIfBelowAverage
	BuyEveryDayIfLowerThanConfig    BuyEveryDayIfLowerThan

	ConfigFile string

	cmd *cobra.Command
}

func NewOptions() *Options {
	return &Options{}
}

func (o *Options) GetConfigFile() string {
	return o.ConfigFile
}

func (o *Options) AddFlags(cmd *cobra.Command) {
	o.cmd = cmd

	versionCommand := &cobra.Command{
		Use:   "version",
		Short: "Print version information.",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("%s\n", version.Print("systemtrading"))
		},
	}
	cmd.AddCommand(versionCommand)

	// is called when parsing failed.
	o.cmd.Flags().Usage = func() {
		_, _ = fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		o.cmd.Flags().PrintDefaults()
	}

	o.cmd.Flags().Int64Var(&o.TelegramChatId, "telegram-chat-id", 0, "telegram chat ID")
	o.cmd.Flags().StringVar(&o.TelegramToken, "telegram-token", "", "telegram token")

	o.cmd.Flags().StringVar(&o.ConfigFile, "config-file", "", "A path to config file")

	o.addRequiredFlag(&o.KoreaInvestmentUrl, "koreainvestment-url", "The endpoint of KoreaInvesment API")
	o.addRequiredFlag(&o.KoreaInvestmentAppKey, "koreainvestment-appkey", "Your KoreaInvesment Appkey to call APIs")
	o.addRequiredFlag(&o.KoreaInvestmentSecret, "koreainvestment-appsecret", "Your KoreaInvestment AppSecret to call APIs")
	o.addRequiredFlag(&o.KoreaInvestmentAccount, "koreainvestment-account", "Your KoreaInvestment Account including '-' e.g. 123456-01")
}

func (o *Options) addRequiredFlag(p *string, name, usage string) {
	// The default value is empty which is meaningless. Because this flag is mandatory.
	o.cmd.Flags().StringVar(p, name, "", usage)
	o.cmd.MarkFlagRequired(name)
}

func (o *Options) Parse() error {
	err := o.cmd.Execute()
	return err
}

var InitCommands = &cobra.Command{
	Use:   "systemtrading",
	Short: "Add arguments for system trading",
	Args:  cobra.NoArgs,
}
