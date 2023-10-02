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

		scheduler := order.StrategryBuyEveryDayIfLowerThan("21:57", []order.StrategryOrder{
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

	RunOrDie(context.Background())

	select {}
}
