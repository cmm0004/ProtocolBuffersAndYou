# ProtocolBuffersAndYou
## Initial Running Webservice

This is a hello world for a go restful web service. Its going to be our starting point for returning Protocol Buffers to clients rather than json. Neat!

I've got the source in main.go, and have already compiled it for you as 'main'.
main is an executable, so, in your terminal, navigate to where you downloaded this repo and type `./main` and this will run the file. 


This will start a webserver on localhost:8080 and have /ping and /job/:jobdid endpoints. now navigate to it with your favorite browser, postman, or `curl localhost:8080/jobs/123456` in a new terminal window if you are feeling frisky.

If you get back something like "Ive got job 123456" then you are good to go! Lets move on to branch 2.
