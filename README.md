Here, Gin Framework is used. The CRUD operations are to create a product, fetch a product, fetch all products and delete a product. Made use of NoSQL DB (Mongo).

Test the API's using :-

1. Create a user - http://localhost:8000/users
    Send user object in payload. Example - {id: 1, name: "Nishant", age: 24, address: "Thane, Mumbai"}

2. Get a single user - http://localhost:8000/users/1

3. Get all users - http://localhost:8000/users

4. Delete a user - http://localhost:8000/users/1 
   here the request type is DELETE