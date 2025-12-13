## POST /login

Request

```json
{
  "email": "",
  "password": ""
}
```

Response 200

```json
{
  "message": "success",
  "data": "JWT TOKEN"
}
```

## GET /products

request param

```json
{
  "paginate": true,
  "search": ""
}
```

Response

```json
{
    "message": "success",
    "data": [
        {
            "id": products id,
            "name": "Product name",
            "description": "Product description"
        }
    ]
}
```

## GET /products/:id

Response

```json
{
    "message": "success",
    "data": {
        "id": products id,
        "name": "Product name",
        "description": "Product description"
    }

}
```

## PUT /products/:id/update

Response

```json
{
    "message": "success",
    "data": {
        "id": products id,
        "name": "Product name",
        "description": "Product description"
    }

}
```
