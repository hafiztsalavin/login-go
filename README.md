# JWTs, access tokens & refresh tokens  

This is just a simple API in Go, with basic authentication using JWT tokens and user management. All written in Go with a PostgreSQL databases. Looks like simililar with simple-jwt-go but i added some role: user, admin and super admin with different permission.

### Installation

1. Add `.env` file and copy `env.sample` to `.env`. and adjust the environment variable to your own environment.
```
cp env.sample .env
```
2. Install Go library that was listed in go.mod. 
```
go mod tidy
```
3. Make sure postgres is already based on your `.env` file
4. Create empty database to migrate database if you haven't. 
```
make migrate
```
5. Make a seed data. 
```
make seed
```
6. Run api.
```
make start
```