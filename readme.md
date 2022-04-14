## Structure API

/user/login

```bash
Method: POST,
Parameter:
 "email": string,
 "password": string

Sample Response:
{
    "data": {
        "email": "wahyu.mailist@gmail.com",
        "name": "Wahyu Ade Pratama",
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6IndhaHl1Lm1haWxpc3RAZ21haWwuY29tIiwiZXhwIjoxNjUyNDgyNTkyLCJpc3MiOiJ3YWh5dS5tYWlsaXN0QGdtYWlsLmNvbSJ9.4_JEPeKzbnS5ozQBgMLPLroer_7rQmrL8ByUIYxfT6g"
    },
    "message": "Login Success",
    "status": "200"
}
```
/user/register

```bash
Method: POST,
Parameter:
 "email": string,
 "name": string,
 "password": string

Sample Response:
{
    "data": {
        "email": "wahyuadepratama@gmail.com",
        "name": "Wahyu Ade Pratama",
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6IndhaHl1YWRlcHJhdGFtYUBnbWFpbC5jb20iLCJleHAiOjE2NTI0ODI3NjgsImlzcyI6IndhaHl1YWRlcHJhdGFtYUBnbWFpbC5jb20ifQ.sB2s_Qyl9TqenJ4dOatClv7Ldj1Ex2KKbYnHlXqRvNo"
    },
    "message": "Registration Success",
    "status": "200"
}
```
/weather/update

```bash
Method: PUT,
Parameter Body:
 "token": string

Sample Response:
{
    "data": {
        "lat": -6.2953,
        "lon": 106.6383,
        "timezone": "Asia/Jakarta",
        "current": {
            "pressure": 1008,
            "humidity": 87,
            "wind_speed": 2.57,
            "weather": [
                {
                    "id": 721,
                    "main": "Haze",
                    "description": "haze"
                }
            ]
        }
    },
    "message": "Data updated successfully from API",
    "status": "200"
}
```
/weather/all

```bash
Method: GET,

Sample Response:
{
    "data": {
        "id": 8,
        "lat": -6.2953,
        "lon": 106.6383,
        "timezone": "Asia/Jakarta",
        "pressure": 1008,
        "humidity": 87,
        "wind_speed": 2.57,
        "created_at": "2022-04-14T06:02:20+07:00",
        "weather": [
            {
                "id": 721,
                "weather_id": 8,
                "main": "Haze",
                "description": "haze"
            }
        ]
    },
    "message": "Data loaded successfully",
    "status": "200"
}
```
