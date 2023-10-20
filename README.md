This is a project to learn some Go, aka Golang. I'm following the youtube tutorial linked below. The app is a CLI (non-real world/DB-persisted) ticket booking app. In reality, you want a front-end UI and a DB to persist information.

The main idea we wanted to explore beyond the general syntax and structure of Go is concurrency. Concurrency is Go's major benefit. Go does concurrency with goroutines. Goroutines are "green threads" which is an abstraction of an operating system thread that a language like Java use. They are more memory efficient and quicker to create than threads managed by the kernel. You can run thousands or millions of threads without affecting performance. Go also has an easy way to communicate between threads via channels. Kubernetes and Docker were written in Go. 

One thing we didn't do in this app is share state between the treads. We'd like to be able to check the current state between multiple users at check-out to make sure there are tickets left when they complete their checkout. Instead in this app we pretend there are multiple ticket creation and email processes that take time, and we kick all of those processes off on different threads while the application continues to take in new customer information.

There are a lot of comments in the code that are a little verbose if you already know the language. A lot of the comments compare things to Python since that is my primary language.

https://www.youtube.com/watch?v=yyUHQIec83I



