## JWT implementation examples

https://pkg.go.dev/github.com/golang-jwt/jwt#example-New-Hmac


## Log Aggregation

https://blog.logrocket.com/kubernetes-log-aggregation/

## Installing and Using PostgreSQL on Ubuntu

https://www.digitalocean.com/community/tutorials/how-to-install-postgresql-on-ubuntu-20-04-quickstart

## Deploying PostgreSQL as a K8s Service

(Alternative for clustering: https://devopscube.com/deploy-postgresql-statefulset/)

https://www.cloudytuts.com/guides/kubernetes/how-to-deploy-postgress-kubernetes/

```bash
kubectl create secret generic postgres \
--from-literal=POSTGRES_USER="auth" \
--from-literal=POSTGRES_PASSWORD="mypass" \
--from-literal=REPMGR_PASSWORD
```

## Setting up Authentication DB

https://medium.com/coding-blocks/creating-user-database-and-adding-access-on-postgresql-8bfcd2f4a91e

```bash
sudo -u postgres psql
```

```sql
create database authenticate;
create user auth with encrypted password 'mypass';
grant all privileges on database authenticate to auth;
```

```bash
sudo adduser auth
sudo -u auth psql -d authenticate
```

```postgres
\conninfo
```

## pgbounce

https://hub.docker.com/r/edoburu/pgbouncer/
https://github.com/edoburu/docker-pgbouncer/tree/master/examples/kubernetes

### Connecting through pgbounce

```bash
sudo -u auth psql -h localhost -p 31058 -d authenticate
```

### Connecting inside the instance

```bash
psql -U auth -d authenticate
```

### Generate password md5

```bash
printf "auth mypass" | openssl md5 -binary | xxd -p
```

### Decoding the password

```bash
kubectl get secret postgres-secrets -o jsonpath="{.data.POSTGRES_PASSWORD}" | base64 --decode
```

## Get pgbouncer service port

```bash
kubectl get svc pgbouncer -o jsonpath="{.spec.ports[0].nodePort}"
```

Alternatively,

```bash
kubectl get svc pgbouncer -o go-template='{{range.spec.ports}}{{if .nodePort}}{{.nodePort}}{{"\n"}}{{end}}{{end}}'
```

## Creating user

```bash
curl -H "Email: john.smith@gmail.com" -H "Fullname: John Smith" -H "Password: mypass" -H "Role: admin" -H "Username: johns" localhost:8080/signup
```

## Get user

```bash
curl -H "Email: john.smith@gmail.com" -H "Password: mypass" localhost:8080/getuser
```
