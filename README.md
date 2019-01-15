# Sample Dashboard using go lang and vuejs

## Overview

This repository includes both backend and frontend.  Backend is developed using go lang while frontend is via Vue js.

## Note

This is my first time using Go lang since I am really a python/django developer so please pardon me if I haven't met the standards for go programming yet. :)

## Software Requirements

* Go 1.11
* vue-cli 3.3

## Backend setup

1. Go to `backend_go` directory and just run the following command:

```
go run server.go
```

It will automatically install the package dependencies and run the server on `http://localhost:1323` unless you have changed the default port.

## Frontend setup

1. Go to `frontend` directory
2. Install packages by running `npm install` or `npm install --only dev`.  You can also use `yarn` if you prefer using it.
3. Run `npm run dev

**NOTE**

API_URL is set in `.env.dev` file, so if the api url is not the same as the default, please change it also.  Another thing, in the backend (`server.go`), I only allowed CORS from `localhost:8080` so if the front end url changes, please do change the CORS settings in `server.go` file also.

## Credentials

email: test@email.com
password: password
