package client

import (
	contract "01-Introduction_to_Microservices/08-RPC_HTTP/1-contract"
	"fmt"
	"log"
	"net/rpc"
)

const port = 1234

func CreateClient() *rpc.Client {
	/*
			With gobs, the source and destination values and types do not need to correspond exactly,
		    when you send struct, if a field is in the source but not in the receiving struct,
		    then the decoder will ignore this field and the processing will continue without error.
		    If a field is present in the destination that is not in the source, then again the decoder
		    will ignore this field and will successfully process the rest of the message.
	*/
	/*#### NEW ::: FOR RPC as a HTTP Transport protocol ####*/
	client, err := rpc.DialHTTP("tcp", fmt.Sprintf("localhost:%v", port))
	if err != nil {
		log.Fatal("dialing:", err)
	}

	return client
}

func PerformRequest(client *rpc.Client) contract.HelloWorldResponse {
	req := &contract.HelloWorldRequest{Name: "Alishah"}
	var res contract.HelloWorldResponse

	// After connect to server successfully, when we use client.Call(), it can automatically run Dial() or something else
	// Then we can make a request to the Server by inputting parameters into client.Call().
	err := client.Call("HelloWorldHandler.HelloWorld", req, &res)
	if err != nil {
		log.Fatal("error:", err)
	}

	return res
}
