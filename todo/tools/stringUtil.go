package tools

func BuildStringFromArray(input []string) string{
	output := ""
	
	for _,s := range input{
		if s != ""{
			output += s + "\n"
		}
	}
	return output
}