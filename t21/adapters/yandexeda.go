package adapters

import (
	"encoding/json"

	"github.com/heckphy/wb-l1/t21/s1/yandexeda"
)

type YandexEdaAPIAdapter struct {
	// враппим имеющийся функционал и реализуем требуемый интерфейс
	// в других языках вместо композиции может использоваться наследование
	y *yandexeda.YandexEdaAPI
}

func NewYandexEdaAPIAdapter(baseURL string) *YandexEdaAPIAdapter {
	return &YandexEdaAPIAdapter{
		y: yandexeda.NewAPI(baseURL),
	}
}

// здесь как раз и происходит подгонка
func (a *YandexEdaAPIAdapter) FetchProducts() json.RawMessage {
	return a.y.GetFoods()
}
