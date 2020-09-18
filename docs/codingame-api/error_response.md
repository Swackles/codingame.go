# Error responses

This is description for codingame error responses.

<br><br>

## Table of contents

- [Error responses](#error-responses)
  - [Table of contents](#table-of-contents)
  - [Responses](#responses)
    - [User not logged in](#user-not-logged-in)
    - [Service not found](#service-not-found)
    - [Internal server error](#internal-server-error)
    - [Email is incorrect](#email-is-incorrect)
    - [Password is incorrect](#password-is-incorrect)
    - [You're not the owner](#youre-not-the-owner)

<br><br>

## Responses

---

### User not logged in

**Code** : `403 Forbidden`

**Content example** :

```json
{
    "code": "UserRequired",
    "message": "Only a logged user is authorized to perform this operation"
}
```

---

### Service not found

**Code** : `422 Unprocessable Entity`

**Content example** :

```json
{
    "id": -3,
    "message": "Service not found: lastactivities.getlastactivities(1)"
}
```

---

### Internal server error

**Code**: `500 Internal server error`

**Content example** :

```json
{
    "id": -1,
    "message": "internal error"
}
```

---

### Email is incorrect

**Code** : `200 OK`

**Content example** :

```json
{
    "error": {
        "id": 334,
        "message": "Malformed email"
    }
}
```

---

### Password is incorrect

**Code** : `200 OK`

**Content example** :

```json
{
    "error": {
        "id": 396,
        "message": "The password you entered is incorrect."
    }
}
```

---

### You're not the owner

If what you're are trying to request, dosen't belong to your account.

**Code** : `422 Unprocessable Entity`

**Content example** :

```json
{
    "id": 511,
    "message": "Only the owner can access puzzles progress"
}
```
