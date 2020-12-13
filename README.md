
# Oms
This Service contain three microservice, cart-service, order-service and product service. You need to run all of them sepratelly . 
### Clone Repo inside your workspace 
```
git clone https://github.com/rajan-kumar-gl/oms
```

### Cart Service
This Service can be used for managing cart operations, By this command microservice server will up on port `7902`
```sh
build cart-service/app.go & ./app
```

### Product Service
This Service can be used for managing Product operations By this command microservice server will up on port `7903`
```sh
build product-service/app.go & ./app
```
### Order Service
This Service can be used for placing user orders By this command microservice server will up on port `7904`
```sh
build product-service/app.go & ./app
```

## Useacase
```
An eCommerce company is having a sale, where some products are having lightning deals with an extra discount for 30 mins. When a product is having a lightning deal, it is listed in the landing page and a huge amount of traffic is bombarded on the product’s detail page. One such product X is having stock of 10 and 1,00,000 users have landed on its product details page out of which at time zero 10,000 users have added product in their cart and proceed to checkout. Since available quality is only 10, there will be a race condition among 10,000 users to get 10 quantities. You have to build a system for

  

Order Management, where in  

i. a quantity is allocated to users who completes the payment

ii. no more than 10 quantity can be allocated

iii. focus is to be given on solving the race condition that will arise when 10,000 users race to get only 10 quantities

  

You need to create three micro-services of

i. available quantity

ii. adding product on user’s cart

iii. allocating product to user once payment is completed

You can assume all other micro-services which might be required to complete building this system like Payment Gateway etc.

  

You have to use GoLang for implementing this system. Remember that focus is to be given on

i. handling race condition

ii. handling 10,000 concurrent users at one point of time

iii. interaction among multiple micro-services

iii. implementing API gateway will gain additional points
```

