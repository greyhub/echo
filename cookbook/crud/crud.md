# CRUD

## Create User
```zsh
curl -X POST \
  -H 'Content-Type: application/json' \
  -d '{"name":"Joe Smith"}' \
  localhost:1323/users
```

## Get Users
```zsh
curl localhost:1323/users
```

## Get User
```zsh
curl localhost:1323/users/1
```

## Update User
```zsh
curl -X PUT \
  -H 'Content-Type: application/json' \
  -d '{"name":"Joe"}' \
  localhost:1323/users/1
```

## Delete User
```zsh
curl -X DELETE localhost:1323/users/1
```