package internal

import (
	"context"
	"errors"
	"os"
	"path/filepath"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	krxcode "github.com/sunglim/go-korea-stock-code/code"
	"gopkg.in/yaml.v2"
	"sunglim.github.com/sunglim/systemtrading/internal/metrics"
	"sunglim.github.com/sunglim/systemtrading/internal/options"
	"sunglim.github.com/sunglim/systemtrading/log"
	"sunglim.github.com/sunglim/systemtrading/order"
	"sunglim.github.com/sunglim/systemtrading/order/koreainvestment"
	ki "sunglim.github.com/sunglim/systemtrading/pkg/koreainvestment"
)

func Wrapper(opts *options.Options) {
	log.SetTelegramLoggerByToken(opts.TelegramToken, opts.TelegramChatId)

	config := opts.ConfigFile
	if config == "" {
		return
	}

	RunOrDie := func(ctx context.Context) {
		koreainvestment.Initialize(opts.KoreaInvestmentUrl, opts.KoreaInvestmentAppKey, opts.KoreaInvestmentSecret,
			ki.KoreaInvestmentAccount{
				CANO:         opts.KoreaInvestmentAccount,
				ACNT_PRDT_CD: "01",
			})

		go order.StrategryBuyEveryDay(krxcode.Code기업은행, "12:05")

		go order.StrategryBuyEveryDay(krxcode.Code신한지주, "12:06")

		go order.StrategryBuyEveryDayIfBelowAverage("12:00", opts.BuyEveryDayIfBelowAverageConfig.BelowAverage)

		go order.StrategryBuyEveryDayIfLowerThan("13:00", []order.StrategryOrder{
			{
				Code:     krxcode.Code부국증권,
				Price:    17500,
				Quantity: 2,
			},
			{
				Code:     krxcode.CodeKB금융,
				Price:    48000,
				Quantity: 1,
			},
			{
				Code:     krxcode.Code삼성카드,
				Price:    28400,
				Quantity: 1,
			},
			{
				Code:     krxcode.Code삼성전자,
				Price:    60000,
				Quantity: 5,
			},
			{
				Code:     krxcode.Code하나금융지주,
				Price:    33000,
				Quantity: 1,
			},
			{
				Code:     krxcode.CodeBNK금융지주,
				Price:    6500,
				Quantity: 3,
			},
			{
				Code:     krxcode.Code기업은행,
				Price:    9600,
				Quantity: 2,
			},
			{
				Code:     krxcode.CodeDGB금융지주,
				Price:    7000,
				Quantity: 4,
			},
			{
				Code:     krxcode.Code우리금융지주,
				Price:    11300,
				Quantity: 1,
			},
			{
				Code:     krxcode.Code신한지주,
				Price:    32000,
				Quantity: 10,
			},
			{
				Code:     krxcode.Code케이티앤지,
				Price:    80000,
				Quantity: 1,
			},
			{
				Code:     "102110", // tiger 200
				Price:    29000,
				Quantity: 10,
			},
			{
				Code:     "148020", // kbstar 200
				Price:    29000,
				Quantity: 10,
			},
		})

		sellStrategry := order.NewStrategySellEveryDayIfAverageIsHigherThanAveragePercentage("13:01", []order.StrategryBuyEveryDayIfBelowOrder{{}})
		go sellStrategry.Start()

		store := metrics.MetricStore{}
		store.ListenAndServe(":8080")
	}

	cfgViper := viper.New()
	cfgViper.SetConfigType("yaml")
	cfgViper.SetConfigFile(config)
	err := cfgViper.ReadInConfig()
	if err != nil {
		if errors.Is(err, viper.ConfigFileNotFoundError{}) {
			panic(errors.New("not found config file"))
		}
		panic(err)
	}
	cfgViper.OnConfigChange(func(in fsnotify.Event) {
		panic(errors.New("not implemented yet"))
	})
	cfgViper.WatchConfig()

	configFile, err := os.ReadFile(filepath.Clean(config))
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(configFile, &opts.BuyEveryDayIfBelowAverageConfig)
	if err != nil {
		panic(err)
	}

	RunOrDie(context.Background())
	select {}
}
