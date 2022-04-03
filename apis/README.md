## Go API test

Let's assume that you maintain an API, to keep a directory of all your friends! Each friend has an `id` and a `name`. 
Your friends database is a slice named `Friends`. 

Currently, the API supports two endpoints:
- `GET /health`, which simply returns a 200 response and indicates that the API is up and running.
- `GET /friends`, which returns the full list of friends.

Your task is to fulfill two missing endpoints:
- `POST /friends`, which will add a new friend in the database
- `GET /friends/{id}`, which will return a single friend by his/her `id`.

### Notes:
- Please notice that the `main()` function runs the API in the background, via goroutine. Then, it performs some checks
concerning the correct functionality of the API. As soon as you see the "SUCCESS!" message, then congrats, your task 
is complete!
- No need to modify any code except for where it is explicitly indicated. If you think that this is necessary for your 
work though, you can do it!
- Unfortunately this platform does NOT let us use external libraries, such as `gorilla/mux`. We ask for your patience
to only work with native Go `http` library. 