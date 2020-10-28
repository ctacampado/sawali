# sawali
Sawali is an e-commerce backend platform where you can mix and match features depending on your needs.

## Sawali services
* User Management
* Authorizer
* Inventory Management
* Order Management
* Delivery Management
* Supplier Management
* Product Catalog
* Cart
* Payment
* Sales
* Loyalty Pogram Management
* Analytics

## User Management
This service is for onboarding users. Includes registration, login, and user account management.

## Authorizer
This service is for authorizing users to be able to access other services.

#### Basic flow:
1. User logs in through a client
2. `User management` authenticates the user
3. Upon successful authentication, `user management` calls `authorizer` for a `new token request`
4. Upon receiving the `new token request`, the `Authorizer` generates a new token for the authenticated user
   and returns a `new token response` to `user management`
5. `User management` returns a `login success response` together with the token token to the client
6. The token is then used by the client for every service request made by the user

## Inventory Management

## Order Management

## Delivery Management

## Supplier Management

## Product Catalog

## Cart

## Payment

## Sales

## Loyalty Pogram Management

## Analytics