### go run main.go 3
####Case 1:
``````
Delay: 3
f1(): 2019-05-17 10:29:39.5637578 +0700 +07 m=+3.006711801
f2(): context deadline exceeded
f3(): 2019-05-17 10:29:45.5670935 +0700 +07 m=+9.010047501
``````
#####What causes the error?
`context deadline exceeded` error occurs when using `context.WithDeadline()` and the deadline has expired.
#####How can I fix it?
In most cases this isn’t an error you would fix. This is simple a signal that your code should stop and gracefully exit what its doing.

####Case 2:
``````
Delay: 3
f1(): 2019-05-17 10:29:51.619131 +0700 +07 m=+3.006000001
f2(): 2019-05-17 10:29:54.6217923 +0700 +07 m=+6.008661401
f3(): 2019-05-17 10:29:57.6219789 +0700 +07 m=+9.008847901
``````