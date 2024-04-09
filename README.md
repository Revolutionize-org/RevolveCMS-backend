<h3>Backend</h3>

---

> Start container

docker compose up --build

> Backend

The service run on port 3000 on the endpoint /graphql

> Run migration

docker build -t migrator ./migrator

source .env

docker run --network host migrator -path=/migrations/ -database "$POSTGRES_URL" up

> How to add a user

Everything is done in migrator/query/index.js

You can add your own user, an admin is created by default.

run : `node migrator/query/index.js`

**ALL CREDIENTIALS ARE IN .env**
