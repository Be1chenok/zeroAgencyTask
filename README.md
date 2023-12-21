# Zero Agency Task

[![Go Reference](https://pkg.go.dev/badge/golang.org/x/example.svg)](https://pkg.go.dev/golang.org/x/example)


## Clone the project

```
$ git clone https://github.com/Be1chenok/zeroAgencyTask
$ cd zeroAgencyTask
```

## Launch a project

```
$ make run
```
## If it starts up for the first time
## Execute migrations 

```
$ make migrate-up
$ make migrate-down
```

## Logs
`go.uber.org/zap`

## API server provides the following endpoints:
*  `POST /register` - registration

`Example body`

```
{
  "Email":"user@user.go",
  "Username":"user",
  "Password":"userpass"
}
```

*  `POST /login` - sigIn

`Example body`

```
{
  "Username":"user",
  "Password":"userpass"
}
```

*  `GET /logout` - logout need access token
*  `GET /refresh` - refresh token need refresh token
*  `GET /fullLogout` - full logout neen access token

*  `POST /edit/:id` - edit news

`Example body`

```
{
  "Id": 64,
  "Title": "Lorem ipsum",
  "Content": "Dolor sit amet <b>foo</b>",
  "Categories": [1,2,3]
}
```

*  `GET /list` - get news list (default 10 elements)
```
Pagination:
/list?page=2&size=15
```
