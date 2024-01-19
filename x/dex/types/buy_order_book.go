package types

func NewBuyOrderBook(AmountDenom string, PriceDenom string) BuyOrderBook {
	book := NewOrderBook()
	return BuyOrderBook{
		AmountDenom: AmountDenom,
		PriceDenom:  PriceDenom,
		Book:        &book,
	}
}

func (b *BuyOrderBook) FillSellOrder(order Order) (remainingSellOrder Order, liquidated []Order, gain int32, filled bool) {
	var liquidatedList []Order
	totalGain := int32(0)
	remainingSellOrder = order

	for {
		var match bool
		var liquidation Order
		remainingSellOrder, liquidation, gain, match, filled = b.LiquidateFromSellOrder(remainingSellOrder)
		if !match {
			break
		}

		totalGain += gain
		liquidatedList = append(liquidatedList, liquidation)

		if filled {
			break
		}
	}
	return remainingSellOrder, liquidatedList, totalGain, filled
}

func (b *BuyOrderBook) LiquidateFromSellOrder(order Order) (remainingSellOrder Order, liquidatedBuyOrder Order, gain int32, match bool, filled bool) {
	remainingSellOrder = order

	orderCount := len(b.Book.Orders)
	if orderCount == 0 {
		return order, liquidatedBuyOrder, gain, false, false
	}

	highestBid := b.Book.Orders[orderCount-1]
	if order.Price > highestBid.Price {
		return order, liquidatedBuyOrder, gain, false, false
	}

	liquidatedBuyOrder = *highestBid

	if highestBid.Amount >= order.Amount {
		remainingSellOrder.Amount = 0
		liquidatedBuyOrder.Amount = order.Amount
		gain = order.Amount * highestBid.Price

		highestBid.Amount -= order.Amount
		if highestBid.Amount == 0 {
			b.Book.Orders = b.Book.Orders[:orderCount-1]
		} else {
			b.Book.Orders[orderCount-1] = highestBid
		}
		return remainingSellOrder, liquidatedBuyOrder, gain, true, true
	}

	gain = highestBid.Amount * highestBid.Price
	b.Book.Orders = b.Book.Orders[:orderCount-1]
	remainingSellOrder.Amount -= highestBid.Amount
	return remainingSellOrder, liquidatedBuyOrder, gain, true, false
}
