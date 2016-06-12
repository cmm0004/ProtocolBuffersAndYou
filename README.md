# ProtocolBuffersAndYou

## 3 - Responding to requests with protobuffed data

This one is pretty self-explanitory, for the purpose of simplicity, I've just initalized a struct literal with some values and serialized them into (in go terms) a slice of bytes based on the message we defined in the second branch, and struct that was created for us after compiling the .proto file. Then I'm adding the string representation of that byte slice to the response writer object and when the function ends, that is what is returned to the client.

If you run main, and call the 'localhost:8080/jobs/123353whatever' endpoint with postman, you should get my values I've set, plus some weird characters that postman doesnt quite know what to do with on either end. That's okay. This data, unlike json, isnt really human friendly, and doesnt always translate to text perfectly as is.

Now lets move on to the next branch.
