# Go Routines

routine is basically a thread I believe

When running a Go program, we create one Go routine, executing each line one by one
kinda like JS main thread

So things like http.Get are going to BLOCK the thread/go routine

New routines can be created with the `go` keyword which runs that code in that routine instead
the blocking call then blocks that routine instead of the main one

So while the new routine is blocked, control flow is passed back to the main routine which continues on.
The main routine then needs to make sure it is

Go has a scheduler, and Go is going to try and only use one CPU.
If a scheduler sees a routine is completed, or has a blocking call, that routine is "descheduled" for now and another is allowed to run
If multiple cores are available, routines can be run in parallel (rather than just concurrently)
