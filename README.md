# address-book

## Task: Simple address-book. 
- Implement address book as a standalone service (single app, in-memory storage). 
- The required fields:
  - username, 
  - address, 
  - phone.
- Endpoints: gRPC/HTTP (from protobuf)

User should be able to add the new users, 
find the existing users by any field (including wildcards), 
update user information and delete users from storage.

## Task: Add DB and deploy to minikube

1. DB (postgres, basic SQL syntax)

2. Docker (docker, docker-compose)

3. Minikube

Improve the previous service, adding external DB as a storage. Use GORM library. 
In-memory storage should be changed to database, you may try to create the tiny deployment,  
or at least run the DB in docker on your own PC with the service, started as a simple 
application on your PC as well

Run DB and service via docker-compose

Optional:

Deploy DB and service to minikube 