# Leago

___

This project comes from some challenges I faced while developing hobby projects using the Riot API, the goal here is to provide a working client for it's access.

Leago is a wrapper for the Riot APIs, providing clean API access to the Riot API and Data Dragon.
It works with multiple client instances, with each client being coupled to it's region or platform.

The goal here is to add reliable completion and separation between the clients, that way a Platform Client (NA1 for example) can't be used with Region specific endpoints.

One of the features that I wish to implement too is the support for a RateLimiting interface, handling both token limits and method limits.
___
