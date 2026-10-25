# StockFlow

StockFlow is a mini ERP module web

## TechStack
Frontend: 
- React 
- TypeScript 

Backend:
- Go 
- PostgreSQL 

Deployment:
- Docker
- nginx

## API endpoint
"http://api:8080/api/products" --> only reachable by Docker container exec

1. Get all products
```
docker compose exec api wget -qO- http://localhost:8080/api/products
```
Results:
```
{"data":[{"id":1,"sku":"SPT-BLUE-F40","name":"Flatshoes","price":"220000","stock":27,"isActive":true,"createdAt":"2026-10-23T03:20:52.485933Z","lastMovementAt":null},{"id":2,"sku":"SPT-PINK-F39","name":"Flatshoes","price":"215000","stock":18,"isActive":true,"createdAt":"2026-10-23T03:20:52.485933Z","lastMovementAt":null}, ...]}
```

2. Get all products with certain query

```
docker compose exec api wget -qO- http://localhost:8080/api/products?search=Flatshoes
```
Results:
```
{"data":[{"id":1,"sku":"SPT-BLUE-F40","name":"Flatshoes","price":"220000","stock":27,"isActive":true,"createdAt":"2026-10-23T03:20:52.485933Z","lastMovementAt":null},{"id":2,"sku":"SPT-PINK-F39","name":"Flatshoes","price":"215000","stock":18,"isActive":true,"createdAt":"2026-10-23T03:20:52.485933Z","lastMovementAt":null},{"id":3,"sku":"SPT-BLUE-F39","name":"Flatshoes","price":"215000","stock":12,"isActive":true,"createdAt":"2026-10-23T03:20:52.485933Z","lastMovementAt":null},{"id":4,"sku":"SPT-GREN-F40","name":"Flatshoes","price":"199000","stock":76,"isActive":true,"createdAt":"2026-10-23T03:20:52.485933Z","lastMovementAt":null}]}
```

## FE endpoint
"http://localhost/" --> Products List Page

## How to Run
```
docker compose up --build
```
Build the image and containers and run containers run away. Access front-end at http://localhost/
To stop in the same terminal, use Ctrl+C

```
docker compose down -v
```
Stopping the app in a different terminal

## Design Choices
1. Maintain last_movement_at in the row of products table: use Transaction when updating stock in order to not use additional JOIN query every time we want to access product's data. The JOIN query could be problematic later if the data size becomes big and full row scan is done to only get the most recent data (in the case index has not been implemented yet).
2. Use CORS middleware to communicate between backend and frontend.

## What Would I Do Next
1. Use transactions when updating stocks and creating orders. To update stock and create orders, tables are being joined so the ACID (Atomicity, Consistency, Integrity, Durability) must be maintained. Transaction provider function would be defined in repository level. I would also add Adapters (consisting of repos that have involvement with any transaction) in the entity package so it can be both accessed by repo level and service level (service needs Adapter for NewXXXService function). To make sure no deadlock happens when there are concurrent transactions, I'd use FOR UPDATE with consistent ordering in query.
2. If I had more time, I'd use TDD (test driven development) approach for this project. This is to make sure that all possible test cases are covered and when implementing the code, those edge cases will be taken into consideration. I haven't implemented any testing code yet.
3. Create useDebounce function in frontend for query. The current code is only using frontend side query.
4. Use error middleware in beckend to enable errors with field using gin binding or custom go validator if needed for the order cncellation part.
5. To insert sales_order_items, I'd use query builder in order to commit it as one query, not multiple.

I used about unfocused 5.5 hours to do this project. I am still learning to improve my familiarity with syntaxes and the speed of my typing.