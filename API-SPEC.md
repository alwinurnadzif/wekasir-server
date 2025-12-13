## AUTH

### POST /login

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

## PRODUCTS

### GET /products

request param

```json
{
  "paginate": false
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
            "description": "Product description",
            "qty": "product qty"
        }
    ]
}
```

request param

```json
{
  "paginate": true,
  "search": "",
  "page": 0,
  "size": 10
}
```

Response

```json
{
    "message": "success",
    "data": {
        "first": "boolean",
        "last": "boolean",
        "max_page": "maximal page",
        "page": "current page",
        "size": "items length",
        "total": "total products",
        "total_pages": "total pages",
        "items": [
                {
                    "id": products id,
                    "name": "Product name",
                    "description": "Product description",
                    "qty": "product qty"
                }
            ]
        }
}
```

### POST /products

Request

```json
{
  "name": "",
  "description": "",
  "qty": "number"
}
```

Response

```json
{
  "message": "success",
  "data": {
    "id": "product id",
    "name": "product name",
    "description": "product description",
    "qty": "product qty"
  }
}
```

### GET /products/:id

Response

```json
{
    "message": "success",
    "data": {
        "id": products id,
        "name": "Product name",
        "description": "Product description"
        "qty": "Product description"
    }

}
```

### PUT /products/:id/update

Request

```json
{
  "name": "",
  "description": "",
  "qty": "number"
}
```

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

### DELETE /products/:id/delete

Response

```json
{
  "message": "success"
}
```
