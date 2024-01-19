package types

func NewSellOrderBook(AmountDenom string, PriceDenom string) SellOrderBook {
	book := NewOrderBook()
	return SellOrderBook{
		AmountDenom: AmountDenom,
		PriceDenom:  PriceDenom,
		Book:        &book,
	}
}

func (s *SellOrderBook) AppendOrder(creator string, amount int32, price int32) (int32, error) {
	return s.Book.appendOrder(creator, amount, price, Decreasing)
}

func (s *SellOrderBook) FillBuyOrder(order Order) (remainingBuyOrder Order, liquidated []Order, purchase int32, filled bool) {
	var liquidatedList []Order
	totalPurchase := int32(0)
	remainingBuyOrder = order

	for {
		var match bool
		var liquidation Order
		remainingBuyOrder, liquidation, purchase, match, filled = s.LiquidateFromBuyOrder(remainingBuyOrder)
		if !match {
			break
		}

		totalPurchase += purchase
		liquidatedList = append(liquidatedList, liquidation)
		if filled {
			break
		}
	}

	return remainingBuyOrder, liquidatedList, totalPurchase, filled
}

func (s *SellOrderBook) LiquidateFromBuyOrder(order Order) (remainingBuyOrder Order, liquidatedSellOrder Order, purchase int32, match bool, filled bool) {
	remainingBuyOrder = order

	orderCount := len(s.Book.Orders)
	if orderCount == 0 {
		return order, liquidatedSellOrder, purchase, false, false
	}

	lowestAsk := s.Book.Orders[orderCount-1]
	if order.Price < lowestAsk.Price {
		return order, liquidatedSellOrder, purchase, false, false
	}

	liquidatedSellOrder = *lowestAsk

	if lowestAsk.Amount > order.Amount {
		remainingBuyOrder.Amount = 0
		liquidatedSellOrder.Amount = order.Amount
		purchase = order.Amount

		lowestAsk.Amount -= order.Amount
		if lowestAsk.Amount == 0 {
			s.Book.Orders = s.Book.Orders[:orderCount-1]
		} else {
			s.Book.Orders[orderCount-1] = lowestAsk
		}
		return remainingBuyOrder, liquidatedSellOrder, purchase, true, true
	}

	purchase = lowestAsk.Amount
	s.Book.Orders = s.Book.Orders[:orderCount-1]
	remainingBuyOrder.Amount -= lowestAsk.Amount
	return remainingBuyOrder, liquidatedSellOrder, purchase, true, false
}
