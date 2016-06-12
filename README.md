# ProtocolBuffersAndYou

## 5 - Calling with another program, and deserializing into something that program can understand

Now that we can create and send protobuffed data, your consumers are going to need to deserialize it so they can do the work they need to do on it. with json, since its so well known, most every language has a library for deserializing json, and people dont need to think about it too much. In our since we are using something that isnt the industry standard yet, you will need to either provide you clients with a copy of the definition file compiled for the language they are using (they will need to be using a supported language, im afraid) or you will need to create a wrapper for them to use in their language that also contains the message definition. You should use the wrapper method. But thats a discussion for another time.

Here, I've done something nasty and created a client program in the same repository as the service its calling. Sorry. but it allows us to see the whole process from start to finish.

So, you will need two terminal windows open to the root of this project. In the first run `./main` and then in the second, run `./main-client`.

What main-client.go is doing is the same thing that we have been doing this whole time, its making an http get request to localhost:8080/jobs/1234 BUT! since its going to be doing work on the response, it needs to be able to deserialize the response it gets back into a struct (or whatever your language uses for object-like structures). So thats exactly what it does, and then i have it print the go struct to the terminal to show that the info was correctly parse and put in the struct and is ready to have other work done on it.

For simplicity, I wrote both the client and the main service in Go, but just like with json, the service could be in Go, and the client could be in any supported language, such as java, python, ruby, c#, whatever. you would simply have a version of your definition .proto file compiled to Go in your service, and (if your client was written in C#) a copy of it compiled to c# for your client.
