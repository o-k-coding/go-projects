# Http Client

note he keeps using the term SLA

this stands for Service level agreement
problems with the default client: no timeout set by default
so in the case of a microservice that is taking new connections and making requests, Go will continue to stack these up which could end up overloading your cpu
Probably just because Go will block each thread for each connection that makes an http request, and if they take along time, you have an issue

Go creates a new go routine per connection

(goroutine !== thread)

you can have a lot more go routines than you can threads. Go handles creating them and scheduling them

they exist only in the Go runtime virtual space, not on the OS directly like a thread does

<https://codeburst.io/why-goroutines-are-not-lightweight-threads-7c460c1f155f>

```quote
So, on startup, go runtime starts a number of goroutines for GC, scheduler and user code. An OS Thread is created to handle these goroutines. These threads can be at most equal to GOMAXPROCS.
```

```quote
If a goroutine blocks on system call, it blocks it’s running thread. But another thread is taken from the waiting queue of Scheduler (the Sched struct) and used for other runnable goroutines.

However, if you communicate using channels in go which exists only in virtual space, the OS doesn’t block the thread. Such goroutines simply go in the waiting state and other runnable goroutine (from the M struct) is scheduled in it’s place.
```

so multiple go routines can be concurrently run on a single thread available to the runtime, and if one thread is blocked, the other runnable goroutines can be scheduled to another available thread

## How does HTTP work

### Parts of an HTTP request

#### HTTP Method

#### URL

where to send the request

#### Request Headers

Used to configure the request and response

#### Request Body

Send data with the request, if you want to send json, you can set Content-Type application/json to let the server or client know the shape of the body, so it knows how to parse it.

You can also specify the Accept header to tell the server what you expect to get back
so if Content Type is json, but Accept is xml, the server would know to send xml back instead of Json
Wonder if frameworks usually handle this for you?? Or if you need to specify this? or if it is mainly just an accident

### Parts of an HTTP Response

#### Status Code

Tells the client if the request was succesful or not

#### Response Headers

Headers configuring the response

#### Response Body

Response data

## Go notes


