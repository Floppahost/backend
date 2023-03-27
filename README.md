# This repository is now archived, meaning we do not provide support for this anymore.

# 🚧 API Routes

## `[POST] /login`

### Request

```javascript
{
  username: string;
  password: string;
}
```

### Response

- ✅ Status: **200**

```javascript
auth: true,
error: false,
message: "Logged In",

cookies: {
token: JWT
}
```

- ❌ Status: **401**

```javascript
auth: false,
error: false,
message: "Invalid data",
```

## `[POST] /status`

### Request

```javascript
headers: {
  Authorization: JWT;
}
```

### Response

- ✅ Status: **202**

```javascript
auth: true,
error: false,
message: "Authenticated"
```

- ❌ Status: **403**

```javascript
auth: false,
error: false,
message: "Not authorized"
```

## `[GET] /profile/get/<username>` — not ready yet

### Request

```javascript
headers: {
Authorization: JWT
},
```

### Response

- ✅ Status: **200**

```javascript
auth: true,
error: false,
message: "User fetched",

data: {
username: string,
uid: string,
avatar: string,
bio: string,
uploads: string // upload counter—the user can toggle this in settings
}
```

- ❌ Status: **401, 404**

  - **401** when not authorized
  - **404** when the user doesn't exist

## `[POST] /files/upload`

### Request

```javascript
headers: {
Authorization: JWT
},

file: formFile
```

- ✅ Status: **200**

- ❌ Status: **401, 404**

  - **401** when not authorized
  - **404** when the user doesn't exist

## `[POST] /admin/wave`

### Request

```javascript
headers: {
Authorization: JWT
},
```

- ✅ Status: **200**

- ❌ Status: **401, 501**

  - **501** when the invite system is disabled
  - **401** when the user is unauthorized

## `[POST] /auth/register`

### Request

```javascript
body: {
    username: string,
    email: string,
    password: string,
    invite: string
}
```

_If the invite system is disabled, we just ignore the invite_

- ✅ Status: **200**

- ❌ Status: **400**

## `[POST] /admin/invite/user`

### Request

```javascript
headers: {
Authorization: JWT
},

username: string

```

- ✅ Status: **200**

- ❌ Status: **401, 501**

  - **501** when the invite system is disabled
  - **401** when the user is unauthorized
  - **404** when the user doesn't exist

## `[POST] /admin/blacklist`

### Request

```javascript
headers: {
Authorization: JWT
},

username: string
reason: string

```

- ✅ Status: **200**

- ❌ Status: **401, 404**

  - **401** when the user is unauthorized
  - **404** when the user doesn't exist

## `[POST] /admin/unblacklist`

### Request

```javascript
headers: {
Authorization: JWT
},

username: string

```

- ✅ Status: **200**

- ❌ Status: **401, 404**

  - **401** when the user is unauthorized
  - **404** when the user doesn't exist
