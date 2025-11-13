package arcgis

type CodeWrapper string

func (c *CodeWrapper) UnmarshalJSON(data []byte) (err error) {
	// Does it look like a string?
	var content string
	if len(data) > 1 && data[0] == '"' && data[len(data)-1] == '"' {
		content = string(data[1 : len(data)-1])
	} else {
		if data[0] == 0 {
			content = "0"
		} else {
			content = "1"
		}
	}

	*c = CodeWrapper(string(content))
	return nil
}
