# kingpin

[TOC]

https://github.com/alecthomas/kingpin

```
package main

import (
	"fmt"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	debug=kingpin.Flag("debug","Enable debug mode.").Bool()
	timeout=kingpin.Flag("timeout","Timeout waiting for ping.").Default("5s").OverrideDefaultFromEnvar("PING_TIMEOUT").Short('t').Duration()
	ip=kingpin.Arg("ip","IP address to ping.").Required().IP()
	count=kingpin.Arg("count","Number of packets to send").Int()
)

func main()  {
	kingpin.Version("0.0.1")
	kingpin.Parse()
	fmt.Printf("Would ping: %s with timeout %s and count %d\n",*ip,*timeout,*count)
}
//
//
$ ping --help
usage: ping [<flags>] <ip> [<count>]

Flags:
  --debug            Enable debug mode.
  --help             Show help.
  -t, --timeout=5s   Timeout waiting for ping.

Args:
  <ip>        IP address to ping.
  [<count>]   Number of packets to send
$ ping 1.2.3.4 5
Would ping: 1.2.3.4 with timeout 5s and count 5
```


```
package main

import (
	"gopkg.in/alecthomas/kingpin.v2"
	"os"
	"strings"
)

var (
	app		=kingpin.New("chat","A command-line chat application")
	debug	=app.Flag("debug","Enable debug mode.").Bool()
	serverIP	=app.Flag("server","Server address.").Default("127.0.0.1").IP()

	register	=app.Command("register","Register a new user")
	registerNick=register.Arg("nick","Nickname for user.").Required().String()
	registerName=register.Arg("name","Name of user.").Required().String()

	post		=app.Command("post","Post a message to a channel.")
	postImage	=post.Flag("image","Image to post.").File()
	postChannel	=post.Arg("channel","Channel to post to.").Required().String()
	postText	=post.Arg("text","Text to post.").Strings()
)

func main()  {
	switch kingpin.MustParse(app.Parse(os.Args[1:])){
	case register.FullCommand():
		println(*registerNick)
	case post.FullCommand():
		if *postImage!=nil{		}
		text:=strings.Join(*postText," ")
		println("Post:",text)
	}
}
```
