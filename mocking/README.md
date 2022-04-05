## Go Mocking and Unit Tests

Let's assume that you have a `SimpleApp` that does one simple job; it fetches a list of users from an external
API (`UserFetcher`) and filters out the ones that their company catchphrase contains one or more given keywords.

This `SimpleApp` needs to be tested; there is already a unit test there. But it has a drawback. It is still making
external requests to the user API! This is not something that we want.

Your task is to modify the code and the unit test in a way that it avoids any external connections.

### Notes

- You can modify the code in both `users.go` and `users_test.go` files as much as you want!
- Use Go `interface` magic!