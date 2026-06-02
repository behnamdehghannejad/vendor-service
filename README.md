# Vendor Service

Inventory and vendor management service built with Go, Gin, and Prometheus.

## Features

- Vendor management
- Product management
- Inventory management
- Inventory reservation
- Inventory transaction tracking
- Health checks
- Prometheus metrics

---

# Base URL

```text
http://localhost:8080
```

---

# Health & Monitoring

## Health Check

```http
GET /ping
```

### Response

```json
{
  "pong": true
}
```

---

## Prometheus Metrics

```http
GET /metrics
```

---

# Inventory APIs

Inventory resources are identified using:

```text
{vendorID}-{productID}
```

Example:

```text
1-100
```

---

## Search Inventories

Returns inventories filtered by vendor and/or product.

```http
GET /api/v1/inventories
```

### Query Parameters

| Name       | Type    | Required |
| ---------- | ------- | -------- |
| vendor_id  | integer | No       |
| product_id | integer | No       |

### Example

```http
GET /api/v1/inventories?vendor_id=1
```

### Response

```json
{
  "items": [
    {
      "vendor_id": 1,
      "product_id": 100,
      "quantity": 500,
      "reserved": 50
    }
  ]
}
```

---

## Get Inventory

Returns a single inventory record.

```http
GET /api/v1/inventories/{vendorID}-{productID}
```

### Example

```http
GET /api/v1/inventories/1-100
```

### Response

```json
{
  "vendor_id": 1,
  "product_id": 100,
  "quantity": 500,
  "reserved": 50
}
```

---

## Upsert Inventory

Creates or updates inventory quantity.

```http
PUT /api/v1/inventories/{vendorID}-{productID}
```

### Request Body

```json
{
  "quantity": 500
}
```

### Response

```http
204 No Content
```

---

## Reserve Inventory

Reserves inventory for a request.

```http
POST /api/v1/inventories/{vendorID}-{productID}/reserve
```

### Request Body

```json
{
  "reserve": 10,
  "request_id": "REQ-1001"
}
```

### Response

```http
204 No Content
```

---

# Vendor APIs

## Create Vendor

```http
POST /api/v1/vendors
```

### Request Body

```json
{
  "code": "AMZ",
  "name": "Amazon",
  "email": "vendor@example.com",
  "phone": "+123456789",
  "address": "Seattle, WA"
}
```

### Response

```http
201 Created
```

---

## Get Vendor

```http
GET /api/v1/vendors/{id}
```

### Response

```json
{
  "id": 1,
  "code": "AMZ",
  "name": "Amazon",
  "email": "vendor@example.com",
  "phone": "+123456789",
  "address": "Seattle, WA",
  "active": true,
  "created_at": "2026-01-01T12:00:00Z",
  "updated_at": "2026-01-01T12:00:00Z"
}
```

---

## Search Vendors

```http
GET /api/v1/vendors
```

### Query Parameters

| Name        | Type   | Description         |
| ----------- | ------ | ------------------- |
| code        | string | Vendor code         |
| search_name | string | Partial vendor name |
| active      | string | active or deactive  |

### Example

```http
GET /api/v1/vendors?search_name=amaz
```

### Response

```json
{
  "items": [
    {
      "id": 1,
      "code": "AMZ",
      "name": "Amazon",
      "email": "vendor@example.com",
      "phone": "+123456789",
      "address": "Seattle, WA",
      "active": true,
      "created_at": "2026-01-01T12:00:00Z",
      "updated_at": "2026-01-01T12:00:00Z"
    }
  ]
}
```

---

## Delete Vendor

```http
DELETE /api/v1/vendors/{id}
```

### Response

```http
204 No Content
```

---

# Product APIs

## Create Product

```http
POST /api/v1/products
```

### Request Body

```json
{
  "name": "Laptop",
  "description": "16GB RAM, 512GB SSD"
}
```

### Response

```http
201 Created
```

---

## Get Product

```http
GET /api/v1/products/{id}
```

### Response

```json
{
  "id": 10,
  "name": "Laptop",
  "description": "16GB RAM, 512GB SSD",
  "active": true,
  "created_at": "2026-01-01T12:00:00Z",
  "updated_at": "2026-01-01T12:00:00Z"
}
```

---

## Search Products

```http
GET /api/v1/products
```

### Query Parameters

| Name        | Type   |
| ----------- | ------ |
| search_name | string |

### Example

```http
GET /api/v1/products?search_name=laptop
```

### Response

```json
{
  "items": [
    {
      "id": 10,
      "name": "Laptop",
      "description": "16GB RAM, 512GB SSD",
      "active": true,
      "created_at": "2026-01-01T12:00:00Z",
      "updated_at": "2026-01-01T12:00:00Z"
    }
  ]
}
```

---

## Update Product

```http
PATCH /api/v1/products/{id}
```

### Request Body

```json
{
  "name": "Laptop Pro",
  "description": "32GB RAM, 1TB SSD"
}
```

### Response

```http
204 No Content
```

---

# Transaction APIs

Transaction records track inventory reservations and status changes.

## Search Histories

```http
GET /api/v1/transactions
```

### Query Parameters

| Name       | Type    | Description                                    |
| ---------- | ------- | ---------------------------------------------- |
| activation | string  | active or deactive                             |
| vendor_id  | integer | Vendor identifier                              |
| product_id | integer | Product identifier                             |
| status     | string  | CREATED, RUNNING, PAID, READY, SENT, DELIVERED |

### Example

```http
GET /api/v1/transactions?status=PAID&vendor_id=1
```

### Response

```json
{
  "items": [
    {
      "quantity": 5,
      "vendor_id": 1,
      "product_id": 100,
      "status": "PAID",
      "active": true,
      "created_at": "2026-01-01T10:00:00Z",
      "updated_at": "2026-01-01T10:30:00Z"
    }
  ]
}
```

---

# Status Values

Transaction status supports the following values:

```text
CREATED
RUNNING
PAID
READY
SENT
DELIVERED
```

---

# HTTP Status Codes

| Code | Description           |
| ---- | --------------------- |
| 200  | Success               |
| 201  | Created               |
| 204  | No Content            |
| 400  | Validation Error      |
| 404  | Resource Not Found    |
| 409  | Conflict              |
| 500  | Internal Server Error |

---

# Example Workflow

1. Create a vendor.
2. Create a product.
3. Add inventory for a vendor-product pair.
4. Query inventory availability.
5. Reserve inventory.
6. Track inventory transaction through status changes.

---

# Observability

Prometheus metrics are exposed at:

```http
GET /metrics
```

Example Prometheus configuration:

```yaml
scrape_configs:
  - job_name: vendor-service
    static_configs:
      - targets:
          - localhost:8080
```
