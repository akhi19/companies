# companies

## Steps to run
1. Install docker by following steps mentioned here: https://docs.docker.com/engine/install/
2. cd to the repository and run the following command

```
docker-compose up --build
```

## Steps to import postman collection
1. Download postman app: https://www.postman.com/downloads/
2. Open app and click on Import.
3. Select Link and copy the following url in the box: https://www.getpostman.com/collections/c4059465335c7130d324
4. Press Continue
5. You should now see collection name: 'Company Collection'
6. Click Import
7. You can now start hitting API's from the imported collection

## Sample test case file
### ./pkg/company/internal/services/company_command_service_test.go

## Constraints
1. Only soft delete support.
2. User email and name should be unique.
3. Company name should be unique
4. For easy usage while testing have used name for patching, getting and deleting company/user. However the correct way is to use ID for such operations

## Testing
1. Create users using add user API.
2. Once user created, you can create/patch/delete companies using admin role users token.
3. A dummy admin user is created in the beginning automtically so as to access Admin API's with following credentials : 
```
Email : test@test.com
Password : 1234
```
4. To get a token, use GetTokenAPI with appropriate payload.
5. Once token received add it to request header with `Token` key and value as received token.
5. Get Companies by Name API can be used by both admin and viewer users.
6. You can change user name, role, password by using update user API.
