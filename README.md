# Go-Commerce 🛒

Go-Commerce is a simple backend solution for commerce operations, designed to empower businesses with a lightweight and efficient system. It provides endpoints for managing products, customers, carts, and orders, along with an Entity-Relationship Diagram (ERD) to guide the database design.

## Entity-Relationship Diagram (ERD) 📊

<img src="https://github.com/RafieAmandio/Go-Commerce-BE/assets/88525718/32a9c44a-e60c-444a-9030-1bd00ad04e72" alt="mermaid-diagram-2024-05-02-194659" style="width:80%;">

| Entity       | Attributes                              |
|--------------|-----------------------------------------|
| Product      | product_id, name, category, price, quantity |
| Customer     | customer_id, username, email, password (hashed), created_at, updated_at |
| Cart         | cart_id, customer_id, created_at, updated_at |
| CartItem     | cart_item_id, cart_id, product_id, quantity, created_at, updated_at |
| Order        | order_id, customer_id, total_amount, created_at, updated_at |
| OrderItem    | order_item_id, order_id, product_id, quantity, created_at, updated_at |

## Endpoints 🚀

| Category           | Endpoint                                           | Description                                |
|--------------------|----------------------------------------------------|--------------------------------------------|
| Product            | GET /products                                      | Get all products.                          |
|                    | GET /products/{category}                           | Get products by category.                  |
|                    | POST /products                                     | Add a new product (admin only).            |
|                    | PUT /products/{id}                                 | Update a product (admin only).             |
|                    | DELETE /products/{id}                              | Delete a product (admin only).             |
| Customer           | POST /register                                     | Register a new customer.                   |
|                    | POST /login                                        | Login a customer.                          |
| Cart               | GET /carts/{customer_id}                          | Get customer's cart.                       |
|                    | POST /carts/{customer_id}/items/{product_id}      | Add a product to the cart.                 |
|                    | DELETE /carts/{customer_id}/items/{cart_item_id}  | Remove a product from the cart.            |
| Order              | POST /orders/{customer_id}                         | Create a new order.                        |
|                    | GET /orders/{customer_id}                          | Get customer's orders.                     |


## Getting Started 🚀

To get started with Go-Commerce, follow these steps:

1. Clone this repository.
2. Set up your database according to the provided ERD.
3. Configure your environment variables:

```plaintext
APP_PORT=<Port number your application will listen on>
DB_HOST=<Host address of your database server>
DB_NAME=<Name of your database>
DB_USERNAME=<Username used to connect to the database>
DB_PASSWORD=<Password used to authenticate with the database>
JWT_SECRET=<Secret key used for JSON Web Token (JWT) encryption>

```
