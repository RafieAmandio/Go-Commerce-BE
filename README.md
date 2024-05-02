# Go-Commerce 🛒

Go-Commerce is a simple backend solution for commerce operations, designed to empower businesses with a lightweight and efficient system. It provides endpoints for managing products, customers, carts, and orders, along with an Entity-Relationship Diagram (ERD) to guide the database design.

## Entity-Relationship Diagram (ERD) 📊

```plaintext
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
APP_PORT=6000
DB_HOST=ep-lingering-resonance-a1749ljv.ap-southeast-1.aws.neon.tech
DB_NAME=mooi
DB_USERNAME=mooi_owner
DB_PASSWORD=LJADBRjX3Yu4
JWT_SECRET=LJADBRjX3Yu4asdaw312nj1d9wh10sla273A
