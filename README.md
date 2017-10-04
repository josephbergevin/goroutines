# goroutines
Goroutines at Nozzle
Basic and Advanced Usage

History behind tiJoe

NOTE: figure out how to illustrate all of these points throughout the presentation
http://divan.github.io/posts/go_concurrency_visualize/

Concurrency vs Parallelism
(see youtube video)

Talk about my "concurrency" in Excel with 6 computers and then 8 cores

Beginner's guide to Goroutines
- Common gotcha's for beginners
- What makes GR's so light?
https://tour.golang.org/concurrency/1
https://tour.golang.org/concurrency/2
https://tour.golang.org/concurrency/3
https://tour.golang.org/concurrency/4
https://gobyexample.com/goroutines

Blocking
  - What is blocking?
  - How can I block in Go?
    - Block with Channel (receiving or signaling channel)
    - Block with Mutex
    - Block with timer (uses a channel)
    - WaitGroup
    - Throttler 

Channels in Go
  - closing channels
    - a closed channel never blocks
    - https://golang.org/ref/spec#Close
    - a closed channel cannot be sent to
    - https://gobyexample.com/closing-channels
  - using a channel for signaling
    - one-time signaling shiftStatusDoneChan for shiftManager
    - multi-signaling (counter) clockOutChan for mozWorker
    - why use `chan struct{}` for these?
  - using a channel for data
    - data chan 
  - buffered channels
    - block when full
    

Goroutines: Worker Pools
https://gobyexample.com/worker-pools
Demo Moz Workers/Supervisors
- Common gotcha's for advanced programmers

Examples:
https://www.reddit.com/r/golang/comments/6ziix2/question_about_ranging_over_channels/
https://play.golang.org/p/U4x02kxvba