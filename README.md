# 🚧 API Routes
## `[POST] /api/login`
### Body

```
{
username: string
password: string
}
```

### Response
- ✅ Status: **200** 

```
auth: true,
error: false,
message: "Logged In",

cookies: {
token: JWT
}
```

- ❌ Status: **401**

```
auth: false,
error: false,
message: "Invalid data",
```

## `[POST] /api/status`
### Body

```
headers: {
    Authorization: JWT
}
```

### Response
- ✅ Status: **202**

```
auth: true,
error: false,
message: "Authenticated"
```

- ❌ Status: **403**

```
auth: false,
error: false,
message: "Not authorized"
```

## `[GET] /api/profile/get` — not ready yet

### Body

```
headers: {
Authorization: JWT
},

username: string
```

### Response
- ✅ Status: **200**

```
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

- ❌ Status:  **401, 404**

    - **401** when not authorized
    - **404** when the user doesn't exist

## `[POST] /api/files/upload`

### Body

```
headers: {
Authorization: JWT
},

file: formFile
```