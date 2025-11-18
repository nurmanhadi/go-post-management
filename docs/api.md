## API Endpoints    

### Create Post

**Endpoint:** `POST /api/posts/`

**Request Body:**

```json
{
  "user_id": 2,
  "description": "test222"
}
```

**Response Body:**

```json
{
  "data": "OK",
  "path": "/api/posts/"
}
```

**Response Code**

201, 404, 400

---

### Update Post

**Endpoint:** `PUT /api/posts/{id}`

**Request Body:**

```json
{
  "description": "test2"
}
```

**Response Body:**

```json
{
  "data": "OK",
  "path": "/api/posts/"
}
```

**Response Code**

200, 404, 400

---

### Get By Id

**Endpoint:** `GET /api/posts/{id}`

**Response Body:**

```json
{
  "data": {
    "id": 1,
    "description": "test",
    "total_like": 3,
    "total_comment": 1,
    "user": {
      "id": 1,
      "username": "test10",
      "name": {
        "first_name": null,
        "last_name": null
      },
      "avatar_url": null
    },
    "comments": [
      {
        "id": 4,
        "description": "test",
        "user": {
          "id": 2,
          "username": "test11",
          "name": {
            "first_name": "testing",
            "last_name": "testing juga"
          },
          "avatar_url": null
        },
        "created_at": "2025-11-18T13:41:33.11749+07:00",
        "updated_at": "2025-11-18T13:41:33.11749+07:00"
      }
    ],
    "created_at": "2025-11-17T14:00:53.287151+07:00",
    "updated_at": "2025-11-17T14:30:59.733125+07:00"
  },
  "path": "/api/posts/1"
}
```

**Response Code**

200, 404

---

### Delete Post

**Endpoint:** `DELETE /api/posts/{id}`

**Response Body:**

```json
{
  "data": "OK",
  "path": "/api/posts/{id}"
}
```

**Response Code**

200, 404

---

### Like Post

**Endpoint:** `POST /api/posts/likes`

**Request Body:**

```json
{
  "post_id": 1,
  "user_id": 1
}
```

**Response Body:**

```json
{
  "data": "OK",
  "path": "/api/posts/likes"
}
```

**Response Code**

201, 404, 400, 409

---

### Unlike Post

**Endpoint:** `DELETE /api/posts/likes`

**Request Body:**

```json
{
  "post_id": 1,
  "user_id": 1
}
```

**Response Body:**

```json
{
  "data": "OK",
  "path": "/api/posts/likes"
}
```

**Response Code**

200, 404, 400

---

### Comment Post

**Endpoint:** `POST /api/posts/comments`

**Request Body:**

```json
{
  "post_id": 1,
  "user_id": 1,
  "description": "test5"
}
```

**Response Body:**

```json
{
  "data": "OK",
  "path": "/api/posts/comments"
}
```

**Response Code**

201, 404, 400

---

### Delete Post

**Endpoint:** `DELETE /api/posts/comments`

**Request Body:**

```json
{
  "post_id": 1,
  "user_id": 1
}
```

**Response Body:**

```json
{
  "data": "OK",
  "path": "/api/posts/comments"
}
```

**Response Code**

200, 404, 400