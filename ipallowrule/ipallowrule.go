package ipallowrule

import (
	"log"
	//	"github.com/coreos/go-iptables/iptables"
	"github.com/abdfnx/gosh"
)

/*
2. Write a script (bash/python/whatever you're comfortable with) that will be called with the following parameters:
./script_location NAME IP_ADDRESS ACTION
NAME             - string
IP_ADDRESS   - IPv4 IP address
ACTION          - connect/disconnect

Based on the inputs, the script will ADD / DELETE an iptables rule:
2.1 ADD an allow rule FROM the IP_ADDRESS to ANY, on ACTION=connect
2.2 DELETE a previously ALLOW rule FROM the IP_ADDRESS to ANY on ACTION=disconnect.


---------------------NOTES---------------------
I wanted to use the iptables library but it has poor documentation and as i just should have 4 hours for this work im going just to use the cmd library
*/

//on Action-> Connect i will run this funktion
//the Name is probably some notes what the rule actually does
//Ip Adresses i will try to convert it as strings not very sure if it possible but i think it should

func add_iptables_rule(name string, ip_adress string) {
	log.Printf("started adding %v with the adress %v", name, ip_adress)

	cmd := "sudo iptables -A INPUT -s" + ip_adress + " -j ACCEPT -m comment --comment '" + name + "'"
	gosh.ShellCommand(cmd)

}

func delete_iptables_rule(name string, ip_adress string) {
	log.Printf("started deleting %v with the adress %v", name, ip_adress)

	cmd := "sudo iptables -D INPUT -s" + ip_adress + " -j ACCEPT -m comment --comment '" + name + "'"
	gosh.ShellCommand(cmd)

}
