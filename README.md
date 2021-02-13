# auth
next-auth social authentication with golang custom backend

## Solution

This exemple is splited in two parts: 
a) frontend code is a nextjs app with next-auth authentication and,
b) backend is a golang REST API built using golang echo pkg.

TODO: describe auth flow here

## Build with docker

### Using docker compose to build frontend and backend

```sh
    $ docker-compose up
```

Do not forget to create a "frontend/.env.local" from the template file.

### Build frontend and backend separatly

```sh
    $ docker build -t frontend-docker .
    $ docker run -p 3000:3000 --env-file frontend/.env.local frontend-docker
```

```sh
    $ docker build -t backend-docker .
    $ docker run -p 4046:4046 backend-docker
```

### Frontend

TODO: describe here

### Backend

TODO: describe here

#### Database

In this repo the database is an abstraction mock of a real database. For while we just save data in memory. 
To use a real database you need to write the connection and implement the methods that are in "backend/repository" directory

## Acknowledgments

I am very gratefull that others contributed sharing their knowledge and help me to build this repo. 

A special thanks to: 
 - Arunoda Susiripala whose proposed the authentication flow with a custom backend and shared a complete youtube video tutorial. 
   - [Click here to check the video](https://rysea.com.br/).
   - [Tutorial document](https://arunoda.me/blog/add-auth-support-to-a-next-js-app-with-a-custom-backend)
 - ejirocodes thats create a clear example of next-auth and share his project in github.
   - [link for next-auth exemple by ejirocodes](https://github.com/ejirocodes/Nextjs_Authentication)


