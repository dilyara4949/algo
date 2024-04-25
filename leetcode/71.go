func simplifyPath(path string) string {
	res := []string{}
	p := strings.Split(path, "/")
	fmt.Println(p)
	for _, s := range p {
		if s == "" || s == "." {
			continue
		} else if s == ".." {
			if len(res) > 0 {
				res = res[:len(res)-1]
				if len(res) > 0 {
					res = res[:len(res)-1]
				}

			}
		} else {
			res = append(res, "/")
			res = append(res, s)
		}
	}
	if len(res) == 0 {
		res = append(res, "/")
	}
	return strings.Join(res[:], "")
}