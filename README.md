This is a project to learn some Go, aka Golang. I'm following the youtube tutorial linked below. 

The main idea we wanted to explore beyond the general syntax and structure was concurrency. We do that with goroutines, but we didn't share state. The app is a ticket booking app. We'd like to be able to check the current state between multiple users at check-out to make sure there are tickets left when they complete their checkout. Instead in this app we pretend there are multiple ticket creation and email processes that take time, and we kick all of those processes off on different threads while the application continues to take in new customer information.

There are a lot of comments in the code that are a little verbose if you already know the language. A lot of the comments compare things to Python since that is my primary language.

https://www.youtube.com/watch?v=yyUHQIec83I



