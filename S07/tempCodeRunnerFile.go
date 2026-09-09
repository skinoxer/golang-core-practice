 i, item := range arrinput {
		// go
		// is
		for j, _ := range arrinput {
			if item == arrinput[i] {
				flag = flag + 1
				object[arrinput[j]] = flag
			}
		}
