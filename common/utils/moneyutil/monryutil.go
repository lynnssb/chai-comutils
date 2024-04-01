package moneyutil

import "github.com/shopspring/decimal"

func ConvertFenToYuan(fen int64) float64 {
	yuan := decimal.NewFromInt(fen).Div(decimal.NewFromInt(100)).Round(2)
	result, _ := yuan.Float64()
	return result
}

func ConvertYuanToFen(yuan float64) int64 {
	d := decimal.New(1, 2)
	fen := decimal.NewFromFloat(yuan).Mul(d).IntPart()
	return fen
}
