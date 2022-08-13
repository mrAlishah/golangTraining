package main

import "fmt"

const (
	online      = 0
	offline     = 1
	maintenance = 2
	retired     = 3
)

func printserverstatus(servers map[string]int) {
	fmt.Println("\nthere are", len(servers), "servers")
	stats := make(map[int]int)
	for _, status := range servers {
		switch status {
		case online:
			stats[online] += 1
		case offline:
			stats[online] += 1
		case maintenance:
			stats[online] += 1
		case retired:
			stats[online] += 1
		default:
			panic("unhandled server status")

		}
	}
	fmt.Println(stats[online])
	fmt.Println(stats[offline])
	fmt.Println(stats[maintenance])
	fmt.Println(stats[retired])
}
func main() {
	servers := []string{"web", "page", "docker", "flask", "django"}
	serverstatus := make(map[string]int)
	for _, servers := range servers {
		serverstatus[servers] = online

	}
	printserverstatus(serverstatus)
	serverstatus["darkweb"] = retired
	serverstatus["google drive"] = offline
	printserverstatus(serverstatus)
	for server, _ := range serverstatus {
		serverstatus[server] = maintenance

	}
	printserverstatus(serverstatus)
}
