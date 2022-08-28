what are channels?<br><br>
a channel is a technique which allows to let one goroutine to send data to another goroutine<br><br>
**channel work with two principal operations one is sending and another one is receiving**
<br>The zero value of a channel is nil. nil channels are not of any use and hence the channel has to be defined using make similar to maps and slices.

<br>syntax:
`a:=make(chan int)`
<br>
The syntax to send and receive data from a channel is given below,

```
data:=<-a  //read from channel a
a<-data //write to channel a
```
**VIP**
- when data is read from a channel, the read is blocked until some Goroutine writes data to that channel.<br>
- One important factor to consider while using channels is deadlock. If a Goroutine is sending data on a channel, then it is expected that some other Goroutine should be receiving the data. If this does not happen, then the program will panic at runtime with Deadlock.
- Similarly, if a Goroutine is waiting to receive data from a channel, then some other Goroutine is expected to write data on that channel, else the program will panic.

# Unidirectional channels
- All the channels we discussed so far are bidirectional channels, that is data can be both sent and received on them. It is also possible to create unidirectional channels, that is channels that only send or receive data.
- This is where channel conversion comes into use. It is possible to convert a bidirectional channel to a send only or receive only channel but not the vice versa.

# Closing channels and for range loops on channels

- Senders have the ability to close the channel to notify receivers that no more data will be sent on the channel.
- Receivers can use an additional variable while receiving data from the channel to check whether the channel has been closed.<br>
`v, ok := <- ch  
  `
- ok is true if the value was received by a successful send operation to a channel.if ok is false it means that we are reading from a closed channel.

in send side of channel int ,float,string,bool  is safe for send but pointers or reference like a slice or map is unsafe because the value of pointers or reference may change by sending goroutine or by the receiving goroutine at the same time and the result is unpredicted.<br>
`Mychannel <- element 
`
The above statement indicates that the data(element) send to the channel(Mychannel) with the help of a <- operator.<br>
<br><br>
close a channel with `close` 
and `ele, ok:= <- Mychannel
`
<br><br>
find the length of the channel using len()
<br><br>
find the capacity of the channel using cap()
all channel is basically unbuffered
`ch := make(chan type, capacity)  
`
**The capacity for an unbuffered channel is 0**


## **waitGroup**
A WaitGroup is used to wait for a collection of Goroutines to finish executing.