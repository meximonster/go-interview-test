## Go API test

Let's assume that you maintain an API, to keep a directory of all your friends! Each friend has an `id` and a `name`.
Your friends database is a `map[int]string` named `Friends`.

Currently, the API supports two endpoints:

- `GET /health`, which simply returns a 200 response and indicates that the API is up and running.
- `GET /friends`, which returns the full list of friends.

Your task is:

- Finish incomplete `POST /friends` endpoint, and make it add a new friend in the database
- Finish incomplete `GET /friends/{id}` endpoint, and make it return a single friend by his/her `id`.
- Filter `GET /friends` endpoint. Currently, it returns the full list of friends. Give it also the possibility to filter
  this list by adulthood: when requesting `/friends?isAdult=true`, it should only return friends who have age of 18 or
  above. If `false`, then it returns the non adults.
- Bonus: Currently, there is no database in the application, we keep everything in memory. Considering that many
  requests can happen at the same time, is there something dangerous with that? If you find something dangerous, please
  try to fix it!

### Notes:

- Some unit tests will execute and will check the functionality of the API. As soon as you see the "SUCCESS!" message,
  then congrats, your task is complete!
- No need to modify any code except for where it is explicitly indicated. If you think that this is necessary for your
  work though, you can do it!
- User input can be really evil. The more bad input we catch, the safer our API is. Please return appropriate status
  codes and messages if the user input is invalid.