package client

import (
    "bufio"
    "fmt"
    "net"
    "os"
    "strings"
)

func main() {

if len(os.Args)!=2{
	fmt.Println("please give addr:port");
	return
}

    stdin := bufio.NewReader(os.Stdin)

    conn, err := net.Dial("tcp", os.Args[1])
    if err != nil {
            fmt.Print(err)
            return
    }
    defer conn.Close()

    for {
            fmt.Print("Enter message to transmit: ")
            msg, err := stdin.ReadString('\n')
            if err != nil {
                    fmt.Print(err)
                    return
            }

            msg = msg[:len(msg)-1]+" "
            if (strings.ToLower(msg) == "quit") || (strings.ToLower(msg) == "exit") {
                    fmt.Println("bye")
                    return
            }



            n, err := conn.Write([]byte(msg))
            if err != nil {
                    fmt.Print(err)
                    return
            }
            fmt.Printf("CLIENT: sent %v bytes\n", n)

           str := strings.Fields(string(msg));
           if str[0]=="get"{
		    temp := make([]byte, 1000)
	            n, err = conn.Read(temp)
        	    if err != nil {
	                    fmt.Print(err)
        	            return
	            }
        	    fmt.Printf("CLIENT: received %v bytes\n", n)

	            fmt.Println("Received message:", string(temp));
	  }
    }
}
