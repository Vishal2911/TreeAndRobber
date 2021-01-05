package houserobber

func MaxedSumAcquired(input []int) int{
	var oddsum,evensum int
		for i,v:=range input{
			if len(input) % 2 != 0 && i+1 == len(input) {
				break
			}
			if i%2==0{
				evensum+=v
	
			} else {
				oddsum+=v
			}
		}
	if oddsum > evensum {
		return oddsum
	}else {
		return evensum
	}
}