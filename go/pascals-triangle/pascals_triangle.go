package pascal

func Triangle(n int) [][]int {

	// The numbers in this exercise differ from the mathematical definition of the triangle as far as i know.
	// The first row here counts as 1, i.e. Triangle(1) is {1}, whereas normally this would be {1},{1,1}

	if n<=0 {
		return [][]int{{}}
	}
	tri := make([][]int,n)
	tri[0] = []int{1}
	if n == 1 {
		return tri
	}
	
	for i:=1;i<n;i++ {
		tri[i] = make([]int,i+1)
		tri[i][0] = 1
		tri[i][i] = 1
		for j:=1;j<i;j++ {
			tri[i][j] = tri[i-1][j-1] + tri[i-1][j]
		}
	}
	return tri
}

