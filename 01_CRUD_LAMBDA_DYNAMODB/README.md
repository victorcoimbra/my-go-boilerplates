# A simple CRUD boilerplate using AWS Lambda/DynamoDB, and the Serveless framework for CI/CD pipeline



## Prerequisites

- [Node.js & NPM](https://github.com/creationix/nvm)
- [Serverless framework](https://serverless.com/framework/docs/providers/aws/guide/installation/)
- [Go](https://golang.org/dl/)
- [dep](https://github.com/golang/dep)
- [AWS CLI](https://aws.amazon.com/pt/cli/): `Make sure to setup all AWS and Serverless credentials`

## Quick Start

```
1. Install Go dependencies

```
dep ensure
```

2. Compile functions as individual binaries for deployment package:

```
./scripts/build.sh
```
3. Deploy!

```
serverless deploy
```
