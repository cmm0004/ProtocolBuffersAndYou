# ProtocolBuffersAndYou

## 4- Looking at the differnece between this format and Json.

So, if all this is just an alternative to json, whats the point, why not just use json? Well, why are you using json instead of xml? Because Json is faster and smaller. And so it is with Protocol Buffers. They are faster and smaller than Json, and in this branch we will prove it!

Now, if you look at main.go, I have altered our handler. Now, the endpoind accepts a query parameter `?response=json` to return the response in json, and will return protocol buffed data if left blank. I've also added some timings, and a line to print the size of the data to the terminal window. So with each request, we will now have the size and the time if takes to serialize being printed out to the terminal window that is running main. 

One thing to keep in mind, the first time you run through each code path (json vs protobuf) will take longer than normal, thats expected, just call the endpoint 5 or 6 times to get an idea of the actual normal time it takes to serialize.

So, run the program by executing `./main`, and go to postman and GET localhost:8080/jobs/12345 5 or 6 times, and then GET localhost:8080/jobs/12345?response=json 5 or 6 times, and then go look in your terminal window at the output we created, should look something like this:
```
time to serialize to protobufs: 10.045µs
size of proto: 51
time to serialize to protobufs: 9.361µs
size of proto: 51
time to serialize to protobufs: 8.034µs
size of proto: 51
time to serialize to json: 12.575µs
size of json: 89
time to serialize to json: 12.803µs
size of json: 89
time to serialize to json: 10.915µs
size of json: 89
```

The size is reduce by almost half! and the time to serialize is reduced (on average) by about 30%, and while you might be saying, who cares about a difference of *microseconds* keep in mind this is a very simple message definition, a real job object has hundreds of fields, some quite deeply nested. Yours are probably very complex as well.

I think this is all very convincing.
