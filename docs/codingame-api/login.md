# Login

Used to collect a Token for a registered User.

**URL** : `/CodingamerRemoteService/loginSiteV2`

**Method** : `POST`

**Auth required** : NO

**Data constraints** :
When sending login data use an JSON array, as an example below. Third element is an boolean value and when set to false the rememberMe cookie is not returned. If value is not set then it defaults to false.

```json
[ "[Email address]", "[Password]", true ]
```

## Success Response

**Code** : `200 OK`

**Content example** : Returnes a huge data structure, you can get an example of how it looks like by going [here](../../src/login_response.go)