References:-
http://stackoverflow.com/questions/17766770/socket-echo-server-in-go
http://golang.org/doc/effective_go.html


client
======
To download client package from github.com:-
	go get github.com/mgkanani/client

To create executable use below command(first make sure correctly setting of GOPATH and PATH env-variables):-
	go install
To run:-
	client 127.0.0.1:12345
Now to set key-value pair give input in below format:-
	set [key] [value]
To get value of specific key use:-
	get [key]
If the key does not exist it will return '\n'.
