package internal

import (
	"container/list"
	"context"
	"errors"
	"os"
	"path/filepath"

	"github.com/fsnotify/fsnotify"
	"github.com/go-co-op/gocron"
	"github.com/spf13/viper"
	krxcode "github.com/sunglim/go-korea-stock-code/code"
	"gopkg.in/yaml.v2"
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

	koreainvestment.Initialize(opts.KoreaInvestmentUrl, opts.KoreaInvestmentAppKey, opts.KoreaInvestmentSecret,
		ki.KoreaInvestmentAccount{
			CANO:         opts.KoreaInvestmentAccount,
			ACNT_PRDT_CD: "01",
		})

	gocrons := list.New()
	RunOrDie := func(ctx context.Context) {
		logger := log.Default()
		logger.Println("Start cron jobs")

		gocrons.PushBack(order.StrategryBuyEveryDay(krxcode.Code기업은행, "22:01"))

		gocrons.PushBack(order.StrategryBuyEveryDay(krxcode.Code신한지주, "12:06"))

		buyConfig := opts.BuyEveryDayIfBelowAverageConfig.BuyEveryDayIfBelowAverage
		gocrons.PushBack(order.StrategryBuyEveryDayIfBelowAverage(buyConfig.ExecutionTime, buyConfig.CodeAndQuantity))

		buyLowerThanConfig := opts.BuyEveryDayIfLowerThanConfig.BuyEveryDayIfLowerThan
		scheduler := order.StrategryBuyEveryDayIfLowerThan(buyLowerThanConfig.ExecutionTime, buyLowerThanConfig.CodeAndQuantityAndPrice)

		gocrons.PushBack(scheduler)

		sellStrategry := order.NewStrategySellEveryDayIfAverageIsHigherThanAveragePercentage("13:01", []order.StrategryBuyEveryDayIfBelowOrder{{}})

		gocrons.PushBack(sellStrategry.Start())
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
		for e := gocrons.Front(); e != nil; e = e.Next() {
			val := e.Value.(*gocron.Scheduler)
			val.Stop()
		}

		gocrons.Init()

		RunOrDie(context.Background())
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

	err = yaml.Unmarshal(configFile, &opts.BuyEveryDayIfLowerThanConfig)
	if err != nil {
		panic(err)
	}

	RunOrDie(context.Background())

	select {}
}
