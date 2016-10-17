package Nonogram

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
)

func getLocalAddress() string {
	var localaddress string

	ifaces, err := net.Interfaces()
	if err != nil {
		panic("init: failed to find network interfaces")
	}

	// find the first non-loop back interface with an IP address
	for _, elt := range ifaces {
		if elt.Flags&net.FlagLoopback == 0 && elt.Flags&net.FlagUp != 0 {
			addrs, err := elt.Addrs()
			if err != nil {
				panic("init: failed to get addresses for network interface")
			}

			for _, addr := range addrs {
				if ipnet, ok := addr.(*net.IPNet); ok {
					if ip4 := ipnet.IP.To4(); len(ip4) == net.IPv4len {
						localaddress = ip4.String()
						break
					}
				}
			}
		}
	}
	if localaddress == "" {
		panic("init: failed to find non-loopback interface with valid address on this node")
	}

	return localaddress
}

func call(address string, method string, request interface{}, reply interface{}) interface{} {
	node, derr := rpc.DialHTTP("tcp", string(address))
	if derr != nil {
		log.Fatalf("call: failed to dial: %s", address)
	}
	defer node.Close()

	nerr := node.Call(method, request, reply)
	if nerr != nil {
		log.Println("call:", address, method, request, reply)
		log.Println("call: return error:", nerr)
	}
	return nerr
}

func debugMessage(lvl int, message string) {
	if lvl <= DebugLevel {
		if lvl == 1 {
			log.Println(fmt.Sprintf("[Normal] %s", message))
		} else if lvl == 2 {
			log.Println(fmt.Sprintf("[Detailed] %s", message))
		} else if lvl == 3 {
			log.Println(fmt.Sprintf("[Verbose] %s", message))
		}
	}
}
