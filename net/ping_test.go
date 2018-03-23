package toolkits

import (
	"fmt"
	"testing"
)

func Test_Ping(t *testing.T){
	ipaddr := "127.0.0.1"
	ok := Ping(ipaddr)
	if ok == "no"{
		fmt.Printf("Connect to %s is faild", ipaddr)
	}else{
		fmt.Printf("Connect to %s is sucess", ipaddr)
	}
}
