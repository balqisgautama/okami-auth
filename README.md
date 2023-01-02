
# OAuth (Okami Auth)

OAuth adalah sistem untuk authentikasi user <br/>
OAuth memiliki fitur membuat akun, aktivasi akun dan generate user token

## Membuat Akun / Registrasi
### HTTP Request (POST)
> BASE_URL/okami/auth/registration
### Request Parameters (application/json)
| Parameter | Description                                                                                                                                   |
|-----------|-----------------------------------------------------------------------------------------------------------------------------------------------|
| email     | `string` Name of user                                                                                                                         |
| password  | `string` Account password <br/> Length should be within 8 to 20 characters<br/>At least 1 character uppercase, lowercase, numeric and special |
```json
{
    "email":"test@gmail.com",
    "password": "@Testing1"
}
```
### Response Parameters (application/json)
### - Success
```json
{
  "timestamp": 1672063452,
  "status": {
    "success": true,
    "code": "",
    "message": "Please check your mailbox",
    "detail": null
  },
  "Data": null
}
```
### - Failed
```json
{
  "timestamp": 1672063143,
  "status": {
    "success": false,
    "code": "OAUTH-370001-VALIDATION",
    "message": "Key: 'RegistrationForm.Email' Error:Field validation for 'Email' failed on the 'required' tag\nKey: 'RegistrationForm.Password' Error:Field validation for 'Password' failed on the 'required' tag",
    "detail": [
      "SignRequest.go",
      "ValidateRegistration"
    ]
  },
  "Data": null
}
```

## Aktivasi Akun
- ### Send aktivasi link ke email user
### HTTP Request (GET)
> BASE_URL/okami/auth/active/JWT_TOKEN
### Response Parameters (text/html)

- ### Resend aktivasi link ke email user
### HTTP Request (GET)
> BASE_URL/okami/auth/active/resend/JWT_TOKEN
### Response Parameters (text/html)

## Generate Token / Sign In
- ### Step 1
### HTTP Request (POST)
> BASE_URL/okami/auth/sign/in/step1
### Request Parameters (application/json)
| Parameter | Description              |
|-----------|--------------------------|
| uuid      | `string` uuid v4 rfc4122 |
### Response Parameters (application/json)
- ### Success
Header Secret-Token `JWT`
```json
{
  "timestamp": 1672296996,
  "status": {
    "success": true,
    "code": "",
    "message": "",
    "detail": null
  },
  "Data": null
}
```
- ### Failed
```json
{
  "timestamp": 1672202881,
  "status": {
    "success": false,
    "code": "OAUTH-370001-VALIDATION",
    "message": "Key: 'Step1.UUID' Error:Field validation for 'UUID' failed on the 'uuid4_rfc4122' tag",
    "detail": [
      "SignRequest.go",
      "ValidateStep1"
    ]
  },
  "Data": null
}
```

- ### Step 2
### HTTP Request (POST)
> BASE_URL/okami/auth/sign/in/step2
### Request Parameters (application/json)
| Parameter | Description                                                                                                                                   |
|-----------|-----------------------------------------------------------------------------------------------------------------------------------------------|
| email     | `string` Name of user                                                                                                                         |
| password  | `string` Account password <br/> Length should be within 8 to 20 characters<br/>At least 1 character uppercase, lowercase, numeric and special |
### Response Parameters (application/json)
- ### Success
Header Secret-Token `JWT`
```json
{
  "timestamp": 1672296996,
  "status": {
    "success": true,
    "code": "",
    "message": "",
    "detail": null
  },
  "Data": null
}
```

- ### Step 3
### HTTP Request (GET)
> BASE_URL/okami/auth/sign/in/step3/JWT_TOKEN
### Response Parameters (application/json)
- ### Success
```json
{
  "timestamp": 1672306439,
  "status": {
    "success": true,
    "code": "",
    "message": "",
    "detail": null
  },
  "Data": {
    "user_id": 27,
    "email": "test@gmail.com",
    "client_id": "e074c717-de06-4ad0-a50f-94de7b19de0a",
    "status": 2,
    "sysadmin": 0,
    "user_token": "eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJjaWQiOiJlMDc0YzcxNy1kZTA2LTRhZDAtYTUwZi05NGRlN2IxOWRlMGEiLCJleHAiOjE2NzIzMzUyMzksImlhdCI6MTY3MjMwNjQzOSwiaXNzIjoic2VydmVyIiwic3ViIjoiMjcifQ.MK_oKkKUkoffRh2DHD0HaM6MEBOdo7ZGIyD4VajZPgTKQrk68vNlEai8Ah_cFRXS3WKakFAmgyGxtLLTJO66Pg"
  }
}
```