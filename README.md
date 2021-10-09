# Appointy Task

![badge](https://img.shields.io/badge/Appointy-Technical%20Task-%23099DFD)

![go-gopher](https://user-images.githubusercontent.com/59786899/136669624-9007d601-5c78-4807-925e-4f4adc7c91e0.png)


## Requirements

* Golang SDK. See the [Golang SDK](https://go.dev/) installation instructions.
* An IDE that supports Golang. You can install **Android Studio, IntelliJ IDEA, or Visual Studio Code** and install the Golang plugins to enable language support and tools for refactoring, running and debugging your backend.

## Commands to run

* To start a local project using Golang
```
go mod init <project name>
```
* To add packages to our project
```
go get <package name>
```


## Steps to run

Run the following commands to run the project

```
go build main.go
go run main.go
```


## Disclaimer
I have not used a .env file for the MongoDB Url now for ease of evaluation. I can do the same while working on real life projects.


## Routes
* /users - to get all users (GET)
* /users - to create a user (POST)
* /users/{id} - to get an existing user (GET)
* /posts - to get all posts (GET)
* /posts - to create a post (POST)
* /posts/{id} - to get a exisitng post (GET)

## Sample Tests

Get an existing user             |  Create a new user
:-------------------------:|:-------------------------:
![image](https://user-images.githubusercontent.com/59786899/136669809-73ee4559-ebf9-4565-812f-30b9566b7e88.png)  |  ![image](https://user-images.githubusercontent.com/59786899/136669761-77146bd8-786b-409d-947a-accee3252c59.png)


<br><br>
## Thank you for reading this!

![thank-you-gif-animation](https://user-images.githubusercontent.com/59786899/136669879-db0af519-5ed2-4457-8ea1-b9f76b28a110.gif)
