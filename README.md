# Inventory Dispusipda

A stand alone web app for managing assets for Dinas Perpustakaan Jawa Barat (Dispusipda Jabar).
React Js was used for the frontend and Golang for backend. Golang choosed for the backend because it can be compiled to single binary, so it is easier for distribution and opeartion. Besides running in local, it also can run on a server.

*NB The latest version was not ready to deployed on server, the auth is still not good.

## Forntend
React Js used to make the app feels fast using the SPA architecture. For fast development, ant design library was used to compose the pages. Context was used to manage global state.

## Backend 
REST API was used for client-server communication to make the app more decoupled and easy to extend. The backend implement a Hex architecture and Domain Driven Design. For authentication & authorization using a bearer token. 

## Release
Git hash was used for versioning the release.