# Back-end challenge

## how to run?

I solved every challenge in it's package and included them all in main.go for testing.
you can run tasks from A to C by passing it's character as a command-line argument. like:

```shell-session
foo@bar:~$ go run . a
```

you can write a or A . and replace it with the letter numbering of the challenge (A,B,C or D).

## unit-tests

I added unit-tests for challenges a,b and c. you can run them like:

```shell-session
foo@bar:~$ go test ./c -v

=== RUN   TestWithOdd
    c_test.go:13: tests with odd number passed.
--- PASS: TestWithOdd (0.00s)
=== RUN   TestWithEven
    c_test.go:21: tests with even number passed.
--- PASS: TestWithEven (0.00s)
PASS
ok  	github.com/MohabMohamed/challenge/c	0.002s

```

## About the extra challenge in question B

there's nothing we would change in the code, as every unicode contains some Ascii characters and we aren't considered with anything unless the resulting ascii codes, did we see them in the second trying or not. and I added unit-tests with unicode for proving that. 