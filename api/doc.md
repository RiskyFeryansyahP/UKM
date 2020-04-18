# UKM Organization STIKI MALANG
----------------------------------------------------------------------------------------------

### Documentation URIs API endpoint
---------------------------------------------------------------------------------------------

the application we build is use APIs to communicated with mobile front-end, and this session, i will make a documentation of APIs that will be used.

prefix url staging development `https://ukm-backend.herokuapp.com/api/v0/`

#### User service APIs

user service is providing a service related to user, example :

- user can register / sign up
URL APIs register of the user : `https://ukm-backend.herokuapp.com/api/v0/user/register`
- user can login / sign in
URL APIs login of the user : `https://ukm-backend.herokuapp.com/api/v0/user/login`

###### User Register

Below I will describe the fields that will be sent in the **Request Body**

| Field             | Description               | Option    | Type      |
| ----------------- | ------------------------- | --------- | --------- |
| Email             | email of the user         | Required  | string    |
| Password          | password of the user      | Required  | string    |
| First Name        | First Name the user       | Required  | string    |
| Last Name         | Last Name the user        | Required  | string    |
| Jenis Kelamin     | Gender the user           | Required  | string    |
| Year Generation   | Year Generation the user  | Optional  | int       |
| Phone             | Phone Number the user     | Required  | string    |
| Role              | Role user                 | Required  | int       |
| Image Profile     | Image profile of the user | Optional  | string    |

Example sent request for registration user :

`POST https://ukm-backend.herokuapp.com/api/v0/user/register`

`Content-Type : application/json`

**Request Body** :
```json
{
	"email": "171111000@mhs.stiki.ac.id",
	"password": "risky",
	"role_id": 1,
	"userprofile": {
		"firstName": "Risky",
		"lastName": "Feryansyah",
		"jk": "Male",
		"year_generation": "2017",
		"phone": "083834121715",
		"status": true,
		"img_profile": ""
	}
}
```

**Response Success From Server** :

```json
{
    "statuscode": 200,
    "status": true,
    "result": {
        "id": 3,
        "email": "171111000@mhs.stiki.ac.id",
        "password": "risky",
        "created_at": "2020-04-08T12:44:16.041149051Z",
        "updated_at": "2020-04-08T12:44:16.041149051Z",
        "edges": {
            "Profile": null,
            "Role": null
        }
    }
}
```

If the required option is empty or null, the **Response from the server** will be like this below :

```json
{
    "statuscode": 400,
    "message": "ent: validator failed for field \"email\": value is less than the required length"
}
```

###### User Login

Below I will describe the fields that will be sent in the **Request Body**

| Field             | Description               | Option    | Type      |
| ----------------- | ------------------------- | --------- | --------- |
| Email             | email of the user         | Required  | string    |
| Password          | password of the user      | Required  | string    |

Example sent request for Login user :

`POST https://ukm-backend.herokuapp.com/api/v0/user/login`

`Content-Type : application/json`

**Request Body** :

```json
{
	"email": "171111000@mhs.stiki.ac.id",
	"password": "risky"
}
```

**Response Success From Server** :

```json
{
    "statuscode": 200,
    "status": true,
    "result_profile": {
        "id": 1,
        "email": "171111040@mhs.stiki.ac.id",
        "password": "risky",
        "created_at": "2020-04-10T07:10:35.252282Z",
        "updated_at": "2020-04-10T07:10:35.252282Z",
        "edges": {
            "Profile": {
                "id": 1,
                "firstName": "Risky",
                "lastName": "Pribadi",
                "jk": "Male",
                "alamat": "Jln Gadang",
                "tanggal_lahir": "07-10-1998",
                "year_generation": "2017",
                "phone": "083834121715",
                "status": true,
                "created_at": "2020-04-10T07:10:35.252282Z",
                "updated_at": "2020-04-10T07:12:51.428912Z",
                "edges": {
                    "Owner": null,
                    "Ukm": null
                }
            },
            "Role": {
                "id": 1,
                "value": "DEVELOPER",
                "edges": {
                    "Access": null
                }
            }
        }
    }
}
```

If the required option is empty or email password failed, the **Response from the server** will be like this below :

```json
{
    "statuscode": 401,
    "message": "ent: profile not found"
}
```

#### Profile service APIs

Profile service is providing a service related to user, example :

- Check detail profile
URL APIs register of the user : `https://ukm-backend.herokuapp.com/api/v0/profile/:email`
- Edit Profile
URL APIs login of the user : `https://ukm-backend.herokuapp.com/api/v0/profile/:email`


##### Check Detail Profile

sometimes, we need to check detail profile again to make sure that the data is up to date, because when we update data, the cache / session / shared preference will must update too.

To check detail profile we need to give `URL Parameter`, example we need to check detail profile of which has **Email 171111040@mhs.stiki.ac.id** :

`GET` `https://ukm-backend.herokuapp.com/api/v0/profile/171111040@mhs.stiki.ac.id`

`Content-Type : application/json`

**Response Success From Server** :

```json
{
    "statuscode": 200,
    "status": true,
    "result": {
        "id": 2,
        "firstName": "Risky",
        "lastName": "Feryansyah",
        "jk": "Male",
        "year_generation": "2017",
        "phone": "083834121715",
        "status": true,
        "created_at": "2020-04-08T12:44:16.041149Z",
        "updated_at": "2020-04-08T12:44:16.041149Z",
        "edges": {
            "Owner": null,
            "Ukm": null
        }
    }
}
```

If you not give a `URL` parameter the Result will be **NOT FOUND**, And if you give wrong email the result will be :

```json
{
    "statuscode": 200,
    "message": "ent: profile not found"
}
```

##### Update Profile

Below I will describe the fields that will be sent in the **Request Body**

| Field             | Description               | Option    | Type      |
| ----------------- | ------------------------- | --------- | --------- |
| First Name        | First Name the user       | Required  | string    |
| Last Name         | Last Name the user        | Required  | string    |
| Year Generation   | Year Generation the user  | Optional  | int       |
| Phone             | Phone Number the user     | Required  | string    |
| Image Profile     | Image profile of the user | Optional  | string    |

Example sent request for Update Profile :

`PUT https://ukm-backend.herokuapp.com/api/v0/profile/:email`

`Content-Type : application/json`

<span style="color:red"> **Note*** </span>

> **Required to give `URL Parameter ':email'` when update data**

**Response Success From Server** :

```json
{
    "statuscode": 200,
    "status": true,
    "result": {
        "id": 2,
        "firstName": "Risky",
        "lastName": "Pribadi",
        "jk": "Male",
        "year_generation": "2017",
        "phone": "083834121715",
        "status": true,
        "created_at": "2020-04-08T12:44:16.041149Z",
        "updated_at": "2020-04-09T12:31:57.887931Z",
        "edges": {
            "Owner": null,
            "Ukm": null
        }
    }
}
```

If you not give a `URL` parameter the Result will be **NOT FOUND**, And if you give wrong email the result will be :

```json
{
    "statuscode": 200,
    "message": "ent: profile not found"
}
```

And if you give valid request body or empty required field, the result will be :

```json
{
    "statuscode": 400,
    "message": "last name can't be empty"
}
```

#### UKM service APIs

UKM service is providing a service related to profile, example :

- Register profile to ukm
URL APIs register profile to ukm : `https://ukm-backend.herokuapp.com/api/v0/ukm/register/profile/:profileID`

##### Register profile to UKM

This service make a user registration to ukm

Below I will describe the fields that will be sent in the **Request Body**

| Field             | Description                   | Option    | Type      |
| ----------------- | ----------------------------- | --------- | --------- |
| Name UKM          | Name Of UKM you want regist   | Required  | string    |

we need to passing a url parameter `:profileID`

`POST https://ukm-backend.herokuapp.com/api/v0/ukm/register/profile/1`

`Content-Type: application/json`

**Request Body** :
```json
{
	"name": "SceN"
}
```

**Response Success From Server** :

```json
{
    "statuscode": 200,
    "status": true,
    "result": {
        "id": 1,
        "firstName": "Risky Feryansyah",
        "lastName": "Pribadi",
        "jk": "Male",
        "year_generation": "2017",
        "phone": "083834121715",
        "status": true,
        "created_at": "2020-04-12T23:56:34.249481Z",
        "updated_at": "2020-04-12T23:56:34.249481Z",
        "edges": {
            "Owner": null,
            "Ukm": [
                {
                    "id": 7,
                    "name": "Kompeni",
                    "created_at": "2020-04-14T01:18:38.998236Z",
                    "updated_at": "2020-04-15T04:10:02.561128Z",
                    "edges": {
                        "Profiles": null
                    }
                }
            ]
        }
    }
}
```