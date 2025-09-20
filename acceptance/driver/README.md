# Drivers

[Protocol Drivers](https://continuous-delivery.co.uk/downloads/ATDD%20Guide%2026-03-21.pdf) are advocated by Dave Farley and others, as a common abstraction to represent our access to the system under test. 

They allow us to reason about the system without thinking about whether we're accessing it via a UI or an API layer such as HTTP or gRPC, and allow us to carry out the same tests against different layers of the system if necessary.

i would prefer these of these drivers to 'live' with their implementation packages but `go`'s but `go`'s import rules would not allow this without polluting the exported symbols for the package. I prefer this separation of concerns.