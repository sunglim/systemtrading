# How to run


## With config file

go run main.go --config-file config.yaml

### Config for StrategryBuyEveryDayIfBelowAverage strategry

```yaml
BuyEveryDayIfBelowAverage:
  ExecutionTime: "12:00"
  CodeAndQuantity:
    # Nongsim holdings
    - code: "072710"
      quantity: 1
    # Woori financial
    - code: "316140"
      quantity: 4
    # DBG financial
    - code: "139130"
      quantity: 4
    # KBstar 200
    - code: "148020"
      quantity: 4
```
