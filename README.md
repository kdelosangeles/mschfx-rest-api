# mschfx-rest-api

To launch and test the provided Go file using serverless, please follow the steps below:

## Deployment and Updates

1\. Create an IAM user in the AWS console UI with programmatic access as credential type.

2\. If working with a new IAM account, attach existing policies directly and check the box for AdministratorAccess to grant the user administrative permissions.

3\. Run `aws configure` and enter the AWS Access Key ID, AWS Secret Access Key, us-east-1 for the region, and json for the output format.

4\. Make your change and updates here

5\. Run `go mod tidy` to resolve dependencies.

6\. Run `make` to build the Go binary.

7\. Run `serverless deploy` to deploy the function to AWS Lambda.

7\. Once the function has been deployed, you can test it by running serverless invoke -f <function-name> or by navigating to the AWS Lambda console and testing the function there.

<!-- In this case to test the GET request at getProducts use:  -->

```
serverless invoke -f getProducts
```

## Testing existing/deployed functions

#### via Serverless

```
serverless invoke -f getProducts
```

#### Locally via curl

```
curl https://rv3p1m2yoa.execute-api.us-east-1.amazonaws.com/getProducts
```

#### Unit Test

##### from within function directory

```
go test
```
