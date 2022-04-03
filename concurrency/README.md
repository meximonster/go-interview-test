## Go concurrency test

Let's assume that there is a `task`, which takes 0.5 seconds to complete. There are 10 tasks with 10 different outputs 
that need to be executed.

Those tasks, if ran in a sequential order, take -obviously- 5 seconds to complete.

Please use your Go concurrency knowledge and run those 10 tasks in less than 3 seconds!

Notes:

- No need to modify any code except for `RunTasks()` function! If you think that this is necessary for your work though, 
you can do it!
- Every time a task is fulfilled, it is appended in a slice. This will help the unit tests recognize that all of them 
are fulfilled.
- No need to worry about appending concurrently to the `completedTasks` slice. It's protected with `mutex`.