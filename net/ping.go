package toolkits
import (
	"os/exec"
	"fmt"
	"bytes"
)

// 探测给定的ip是否连通
func Ping(ipaddr string) string{
	cmd := exec.Command("/sbin/ping", "-c", "3", "-t","3", ipaddr)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	ping := ""
	if err != nil{
		fmt.Println(err)
		ping = "no"
	}else{
		ping = "yes"
	}
	fmt.Println(&out)
	return ping

}
