## Go concurrency producer - consumer test

Let's assume that you want to do some statistical research on Twitter. a `Tweet` consists of a `Username` and
a `Content`.

Our final goal is to gather all usernames that are tweeting about "fortune".

There is an already existing `TweetStream` mechanism that needs 0.2 seconds to invoke a new tweet. Also notice that,
assuming that we use some sophisticated machine learning procedure, the calculation of if a tweet talks about fortune
takes 0.3 seconds.

There is a `Produce` function which reads the tweet stream, and a `Consume` function which does the calculation and
appends the appropriate tweets to the `result`.

Your task is to modify this code at your will, and use your Go concurrency knowledge to increase the throughput of this
program!

### Notes

- No need to modify any code except for `Consume()`, `Produce()` and `main()` functions! If you think that this is
  necessary for your work though, you can do it!