# Simple Bank

Created following <https://www.udemy.com/course/backend-master-class-golang-postgresql-kubernetes/>

## Generating a token

```bash
openssl rand -hex 64 | head -c 32
```

## AWS

configuring cli:

aws access key id/secret from gh ci user

### Secrets

Remember that currently secrets are pulled from aws during the image creation and baked into the image. So if you change a secret, a new image needs to be created.
This is not ideal, I am assuming later in the course we will change it.

#### AWS cli

To get secrets

```bash
aws secretsmanager get-secret-value --secret-id simple_bank --query SecretString
```

and converting it to env file format using jq

```bash
## convert the json inso an array of objects representing key/value pairs
aws secretsmanager get-secret-value --secret-id simple_bank --query SecretString --output text | jq 'to_entries'

## Map the array from ^ into an array of just the values of the "key" properting
aws secretsmanager get-secret-value --secret-id simple_bank --query SecretString --output text | jq 'to_entries|map(.key)'

## to get the values
aws secretsmanager get-secret-value --secret-id simple_bank --query SecretString --output text | jq 'to_entries|map(.value)'

## Use string interpolation to create the format we want
## basically the \ in a string is equivalent to ${} in js
aws secretsmanager get-secret-value --secret-id simple_bank --query SecretString --output text | jq 'to_entries|map("\(.key)=\(.value)")'

## Next write each item of the array out using the array iterator
aws secretsmanager get-secret-value --secret-id simple_bank --query SecretString --output text | jq 'to_entries|map("\(.key)=\(.value)")|.[]'

## next remove the quotes using raw flag
aws secretsmanager get-secret-value --secret-id simple_bank --query SecretString --output text | jq -r 'to_entries|map("\(.key)=\(.value)")|.[]'

## Next send to the .env file

aws secretsmanager get-secret-value --secret-id simple_bank --query SecretString --output text | jq -r 'to_entries|map("\(.key)=\(.value)")|.[]' > .env
```

this required adding a policy for secrets manager to the user

### ECR

Action used <https://github.com/marketplace/actions/amazon-ecr-login-action-for-github-actions>
To create a role for GH actions CD, I used <https://docs.github.com/en/actions/deployment/security-hardening-your-deployments/configuring-openid-connect-in-amazon-web-services>
to help me.

#### Testing an image from ECR

First need to authenticate

```bash
aws ecr get-login-password | docker login --username AWS --password-stdin 105745650186.dkr.ecr.us-east-2.amazonaws.com
```

copy the URI from the image <https://us-east-2.console.aws.amazon.com/ecr/repositories/private/105745650186/simplebank?region=us-east-2>

```bash
docker run -p 8080:8080 <image_uri>
```

## Building the docker image (local)

```bash
docker build -t simplebank:latest .
```

## Running the docker image (local)

```bash
docker run  --rm --network simple-bank_default --name simplebank -p 8080:8080 simplebank:latest
# optional, add -e GIN_MODE=release
```

## Networking docker containers

(when running outside of compose) containers cannot talk to each other using hostname on the default bridge by default.
so they need to be connected to a network together.

```bash
docker network create simplebank-network
docker network connect simplebank-network <container>
```

## Development environment setup

requires go, docker

create a .env file and add values for the following variables

```bash
POSTGRES_HOST=
POSTGRES_PASSWORD=
POSTGRES_USER=
POSTGRES_DB=
POSTGRES_PORT=
POSTGRES_SSL_MODE=
```

the docker compose file will load the .env file, so your postgres instance will be set up with the values from that file. You can then use those same values to connect.

## Postgres

### DB Diagram

Using this tool to create the DB diagram <https://dbdiagram.io/home>

specifically <https://dbdiagram.io/d/628847cdf040f104c16ba34e>

### Migrations

Using <https://github.com/golang-migrate/migrate/tree/master/cmd/migrate>

used "Go Toolchain" to install, all other methods failed for me in WSL2 lol.

```bash
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```

for make, used homebrew

```bash
brew install golang-migrate
```

There is a cool pattern used at work where the tools are installed in a dir in the repo dir, and used in there. Like how npm resolves tools in node modules via package json basically.

```bash
migrate create -ext sql -dir db/migrations -seq init_schema
```

note -seq adds a sequence number to the schema.

to run migrations

- manually

```bash
source .env
migrate -path db/migrations -database "postgresql://${POSTGRES_USER}:${POSTGRES_PASSWORD}@${POSTGRES_HOST}:${POSTGRES_PORT}/${POSTGRES_DB}?sslmode=disable" -verbose up
```

- with make <https://simplernerd.com/make-pass-env-nested/>

```bash
make migratedb
```

Note: might be nice to run the migrations IN docker, need a container that has go installed though etc.

in CI, would not source from .env file likely?

migration log is stored in `schema_migrations` table in the DB. If a problem occurs with a migration, I was able to resolve by removing the row for the migration and trying again after fixing the sql.
using the force flag did not work as expected 😢

### SQLC

<https://github.com/kyleconroy/sqlc>

installation

```bash
go install github.com/kyleconroy/sqlc/cmd/sqlc@latest
```

### SQL Driver

<https://github.com/lib/pq>

### Deadlocks

Intuitive, but it was nice to walk through an example where he had all the queries for a transaction written out
and then using the logs to track what order the queries were running with 2 concurrent txs, opened two terminals, and two db connections, started a tx in each and executed the queries in order for each one.

<https://wiki.postgresql.org/wiki/Lock_Monitoring>

when selecting the account for update, this also locks the transfer because of the foreign key constraint

One way to fix this is to remove the fk constraints...
This is not the best solution though if you need this contraint
Instead we can tell PG that we are only selecting for NO KEY UPDATE

### DB Transactions Isolation levels

#### Read Phenomena

Dirty read - tx reads data written by other uncommitted tx.

non repeatable read - tx reads same row twice and sees different value because it has been modified by another tx

phantom read - tx executes a query multiple times and gets a different set of rows each time due to another tx

serialization anomaly - result of a group of txs impossible to achieve if we run them sequentially in any order without overlapping

#### 4 standard ANSI isolation levels

##### Read Uncommitted

can see data written by other uncommitted txs (dirty read, yes phantom read/non repeatable)

in PG, this behaves the same as Read Committed due to the architecture.

##### Read Committed

can only see data written by committed txns (no dirty read, yes phantom read/non repeatable)

##### Repeatable Read

same read query always returns same result (no dirty read, no phantom read, no non repeatable)

in mysql, two txs can update the same row though, and both updates will be applied.

PG throws an error if two txs try to update the same data.

##### Serializable

Can achieve the same result if txs are executes in some order rather than concurrently. (result will always be the same).

in mysql, the select will be used as a select for share, and a lock will happen on a row that is being updated. The lock has a timeout which will eventually time out. SO tx retry is needed!
also possible for a deadlock to occur if one tx selects a row that is already locked by another txn.

PG throws an error and hints for you to retry if you insert into the same table in multiple txns with the same id.
I think using a UUID from the application could be a good way to help handle this? maybe not.

## CI/CD

Github actions are defined in the top level .github/workflows folder, named by project

this is also an interesting structure I found for a monorepo <https://github.com/zladovan/monorepo/tree/master/.github/workflows>

Look into this method for using an env file <https://stackoverflow.com/questions/67964110/how-to-access-secrets-when-using-flutter-web-with-github-actions/67998780#67998780>

the only problem is that I can't pass it to the postgres service which is super annoying

## Viper

Can read from a remote system like consul or etcd
Has ability to live watch writing to the config file

## Mockgen

note I had to add an empty import for mockgen in the main.go file.

give the path to the package
the name of the interface
output file (so it won't write to stdout)

```sh
mockgen -package mockdb -destination db/mock/store.go github.com/okeefem2/simple_bank/db/sqlc Store
```

## JWT vs Paseto

problems with jwt?
"foot guns"

devs have the option of picking weak algorithms because there are so many options.
forgery is easy if you do not implement jwts properly (example the issue where setting header algorithm to none bypasses signature checking in some libraries)

also setting the header to a symmetric alg when you know (via the public key) that the server is using an asymmetric alg.
then sign with the public key.
the server will use the public key, but the header will tell the server to use the symmetric alg instead (I feel like this is an implementation fault) and authentication will pass.
implementation needs to verify the alg header to match what the server uses.

in short, if you use jwts you probably shouldn't roll your own, but use a well knowm implementation (auth0 😉)

### PASETO

devs do not need to choose alg, just the version of paseto to use.
only 2 most recent versions are accepted.

paseto encrypts and authenticates all data in the token with a secret key when using a symmetric key.

public or asymmetric key uses the same signature method as JWT.

for both cases only 1 alg is used though per case.

everything in token us authenticated using the same method as tls, so cannot be tampered with.

simple to implement

### token anatomy

Local
version of paseto
purpose (local) which tells what type of key and signing to use

- local symmetric and authenticated encryption
payload (encrypted aka hex-encoded)\
- data
- nonce
- authentication tag
footer (optional)
- encoded extra data

public
version of paseto
purpose (local) which tells what type of key and signing to use

- asymmetric and digital synature
payload (encoded aka base64)\
- data
signature (encrypted - hex encoded)
- used by server to authenticate authenticity
footer (optional)
- encoded extra data

## Kubernetes

Worker nodes (machines) that run the containers
in each node is a kubelet agent that orchestrates pods running the containers

Container runtimes supported are Docker, containerd and CRI-O
Kubeproxy: maintain network rules, allows communication between pods

### Control plane

runs on the master node and manages worker nods and pods in the cluster.
also contains the API server managing the kubernetes API.
contains etcd, a persistent store for cluster data

#### API Server

Contains the following

connection with etcd, a persistent store for cluster data

scheduler (watches for new pods with no node, selects a node to run on)

controller manager
  node controller - to handle when nodes go down
  job controller - to watch for jobs/tasks and creates pods for them
  endpoint controller - populates endpoints object to join services (?)
  service account & token controller - creates default account and access tokens for new namespace for the kubernetes api

cloud controller manager - links the cluster in the cloud platform API (what makes kubernetes cloud agnostic)
  node controller - checking cloud provider to determine if non responding nodes have been deleted or not.
  route controller - setting up routes in the cloud ifx
  service controller - CRUD cloud load balancers

All of the worker nodes communicate with the API server for orchestration

### Kubectl

to get the eks cluster info into local config

```bash
aws eks update-kubeconfig --name simple-bank --region us-east-2
```

error

```bash
kubectl cluster-info

To further debug and diagnose cluster problems, use 'kubectl cluster-info dump'.
error: You must be logged in to the server (Unauthorized)
```

RBAC configuration not allowing your IAM entity to access the cluster.
Initally only the creator of the cluster has permissions to configure the cluster.

basically you need to be authenticated as the user you created the cluster with (in my case, my console login user). or add perms to the other user you want to use.

The solution is:

```bash
aws sts get-caller-identity # using github-ci user, which is not configured with the master creds
cat ~/.aws/credentials # github-ci user creds

```

create a new access key for your account in the console.
add it to .aws/credentials

to use the github profile

```bash
export AWS_PROFILE=github
```

more info <https://aws.amazon.com/premiumsupport/knowledge-center/amazon-eks-cluster-access/>

When starting the service in kubernetes, I was getting a timeout error connecting to postgres

I needed to configure the inbound rule for the rds security group.

<https://us-east-2.console.aws.amazon.com/ec2/v2/home?region=us-east-2#SecurityGroups:>

running the nginx ingress controller for aws

```bash
kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/controller-v1.3.0/deploy/static/provider/aws/deploy.yaml
```

kubectl exec -it -n ingress-nginx ingress-nginx-controller-54d587fbc6-nchxg -- cat /etc/nginx/nginx.conf

## DB docs

dbml2sql
dbdocs
npm global cli tools. pretty neat!

## GPRC

instructions for installing protoc for go <https://grpc.io/docs/languages/go/quickstart/>

remote procedure call framework

client can execute a function (procedure) on a server
protocol buffer handles the data structures of the calls. (request and response structure)

1. define api & data structure (protobuf)
2. generate gRPC stub code
3. implement the server code
4. implement the client code
5. profit

types

- unary - single request/response
- client streaming - client streams message and gets 1 response
- server streaming - client sends 1 request and server streams messages in response
- bidirectional - client and server stream to each other.

### grpc gateway

a protobug plugin that adds functionality for http json api as well
from proto file, you can generate grpc and json http code.

basically the client calls the gateway with http, then the gateway makes an rpc call from there

in process translation gateway calls grpc code directly (unary only), this does not require an extra network hop

separate proxy server - translated and forwarded to grpc via network call (unary and streaming)

followed instructions in <https://github.com/grpc-ecosystem/grpc-gateway> for installation.

Needed to clone the repo <https://github.com/googleapis/googleapis>

and copy over the following files

```sh
cp /home/okeefem355/Dev/repos/third-party/googleapis/google/api/annotations.proto ./proto/google/
cp /home/okeefem355/Dev/repos/third-party/googleapis/google/api/field_behavior.proto ./proto/google/
cp /home/okeefem355/Dev/repos/third-party/googleapis/google/api/http.proto ./proto/google/
cp /home/okeefem355/Dev/repos/third-party/googleapis/google/api/httpbody.proto ./proto/google/
```

got some errors with protoc command, I ended up removing the brew installed version, and using apt-get (had to use sudo though because my env is jacked up I think)

needed these files for openapi stuffs

```bash
cp /home/okeefem355/Dev/repos/third-party/grpc-gateway/protoc-gen-openapiv2/options/*.proto ./proto/protoc-gen-openapiv2/options/
```

## DNS STUFF

To set up routing, my domain is handled by netlify, I had to add a cname record that points to the nginx controller pod public IP
    a7e116e28080e4fa097945c3c60290c5-233bb7052b6ec35c.elb.us-east-2.amazonaws.com

Did NOT need to set anything up in aws for dns

## Serving Static files

the course uses a library called Statik which bundles the files in the binary rather than serving them from some other location on disk.

This is nice because everything will be loaded in memory (fast). But tradeoff is memory usage, best to use small static files, and I am not sure how well this scales for a lot of traffic - would this increase memory usage a lot.

Also check this out.
<https://dev.to/koddr/the-easiest-way-to-embed-static-files-into-a-binary-file-in-your-golang-app-no-external-dependencies-43pc>

I would be really curious to try this with a react app or something, and load test it.
