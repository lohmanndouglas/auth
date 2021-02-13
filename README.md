# auth
next-auth social authentication with golang custom backend

## Solution

This exemple is splited in two parts: 
a) frontend code is a nextjs app with next-auth authentication and,
b) backend is a golang REST API built using golang echo pkg. 

### Frontend

### Backend

#### Database

In this repo the database is an abstration mock of a real database. For while we just save in memory users. 
To use a real database you need to write the connection and implement the methods that are in "backend/repository" directory

## Acknowledgments

I am very gratefull that others contributed sharing their knowledge and help me to build this repo. 

A special thanks to: 
 - Arunoda Susiripala whose proposed the authentication flow with a custom backend and shared a complete youtube video tutorial. 
   - [Click here to check the video](https://rysea.com.br/).
   - [Tutorial document](https://arunoda.me/blog/add-auth-support-to-a-next-js-app-with-a-custom-backend)
 - ejirocodes thats create a clear example of next-auth and share his project in github.
   - [link for next-auth exemple by ejirocodes](https://github.com/ejirocodes/Nextjs_Authentication)


