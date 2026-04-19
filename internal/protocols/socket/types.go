package socket

type ChartResponse struct {
	Digits    int      `json:"digits"`
	RateInfos []Candle `json:"rateInfos"`
}
