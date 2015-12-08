package main

type route func([]string) error

func applyRoute(routes map[string][]route, input []string) {
	//if val, ok := routes[command]; ok {

	//} else {
	// Did you forget to include
	//}
	// Sanity check
	if len(input) == 0 {
		for _, v := range routes["status"] {
			v(input)
		}
	} else if len(input) < 2 {
		command := input[0]
		args := []string{}
		for _, v := range routes[command] {
			v(args)
		}
	} else {
		// Shift
		command, args := input[0], input[1:]
		for _, v := range routes[command] {
			v(args)
		}
	}
}
