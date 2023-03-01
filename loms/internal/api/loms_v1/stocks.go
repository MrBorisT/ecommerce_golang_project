package loms_v1

import (
	"context"
	"route256/loms/internal/model"
	desc "route256/loms/pkg/loms_v1"
)

func (i *Implementation) Stocks(ctx context.Context, req *desc.StocksRequest) (*desc.StocksResponse, error) {
	stocks, err := i.lomsService.Stocks(ctx, req.GetSku())
	if err != nil {
		return nil, err
	}

	convertedStocks := convertStocksToDesc(stocks)

	return &desc.StocksResponse{
		Stocks: convertedStocks,
	}, nil
}

func convertStocksToDesc(stocks []model.Stock) []*desc.Stock {
	convertedStocks := make([]*desc.Stock, 0, len(stocks))

	for _, stock := range stocks {
		convertedStocks = append(convertedStocks, &desc.Stock{
			WarehouseID: stock.WarehouseID,
			Count:       stock.Count,
		})
	}

	return convertedStocks
}
