# ProtocolBuffersAndYou
## Step 2 define your messages

Defining messages is creating a definition of what the data your app returns looks like. we need this because any clients consuming this data will need to know exacly what to expect, so that the data can be deserialized on thier end. To make this easier for your clients, you should provide them with either a wrapper for your services or a copy of the .proto file that they can compile with the protoc command line tool into a supported language of thier choice.

[Here is where you can learn more about the syntax of .proto files](https://developers.google.com/protocol-buffers/docs/proto3)

[Here is where you can see the supported languages](https://developers.google.com/protocol-buffers/) Though, I would google for implementations in other languages if you dont see yours there, its likely someone has built something for your language of choice.

[Here is where you can download the protoc cli so you can compile your .proto files into files your language of choice can understand](https://github.com/google/protobuf)  Also, another list of supported languages is there in the readme, halfway down.

Now that you have all the components, you can see here in protos/job.proto that I've defined a simple message of Data that contains an Error and a Job object. I've also compiled that into a go file using the protoc tool so that my main.go file can import the struct (object) definitions and do things with them. In this case the command (from the project directory) would look like this : `protoc -I=./proto --go_out=./proto/ ./proto/job.proto`

You should really consider the structure of your messages and read through the docs i've linked, while protobufs make it easy to update definitions, you should still come up with a plan for how you will do that. Also theres lots of nifty things you can do with them that i'm not including in this very simple hello world, so go read the docs.
