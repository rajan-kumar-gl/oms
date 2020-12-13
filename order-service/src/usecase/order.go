package usecase

import (
	"errors"
	"fmt"
	"time"

	"github.com/eapache/go-resiliency/breaker"
	"github.com/oms/order-service/src/proto"
)

type implimantation struct {
	pl      proto.ProductLocker
	cart    proto.CartService
	product proto.ProductService
	cb      *breaker.Breaker
}

func New(cartRepo proto.CartService, pl proto.ProductLocker, productRepo proto.ProductService) proto.OrderManager {
	return &implimantation{
		pl:      pl,
		cart:    cartRepo,
		product: productRepo,
		cb: breaker.New(
			2,
			1,
			time.Duration(500)*time.Millisecond,
		),
	}
}

func (o *implimantation) OrderProduct(userID int, productID int, howMany int) error {
	productInCart, err := o.cart.IsProductInCart(userID, productID)
	if err != nil {
		return err
	}
	if !productInCart {
		return errors.New("this product should first add to cart")
	}
	quantity, err := o.product.GetQtyOfProduct(productID)
	if err != nil {
		return errors.New("unable to get product qty")
	}

	if quantity < 0 || quantity-howMany <= 0 {
		return fmt.Errorf("%d product not avaliable right now", howMany)
	}

	paymentSuccessful := doPayment()
	if !paymentSuccessful {
		return errors.New("payment failed, Please try again in sometime.")
	}

	resBreaker := o.cb.Run(func() error {
		err = o.updateQuantity(productID, howMany)
		return err
	})

	if resBreaker == breaker.ErrBreakerOpen {
		//TODO :: notify Dev about this
		return errors.New("circuit breaker :: Too many error occured")
	} else if resBreaker != nil {
		return resBreaker
	}

	o.pl.AllocateProductToUser(userID, productID, howMany)

	return nil
}

func (o *implimantation) updateQuantity(productID int, howMany int) error {
	quantity, err := o.product.GetQtyOfProduct(productID)
	if err != nil || quantity < 0 || quantity-howMany <= 0 {
		return errors.New("Oops, this product is out of stock now. You will get your money back soon")
	}

	err = o.product.LockProduct(productID, howMany)
	if err != nil {
		return errors.New("somthing went wrong wil will refund your money soon")
	}

	return nil
}

func doPayment() bool {
	//TODO :: Intregate payment Getway
	return true
}
