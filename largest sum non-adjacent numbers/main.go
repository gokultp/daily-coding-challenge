package main

func main(){

}

func largestSumNonAdjascent(data []int) int{
	with := data[0]
	without := 0
	for i := 1; i< len(data); i++ {
		tmp := without + data[i]
		without = with
		if tmp 	> with {
			with = tmp
		}
	}
	return with
}

