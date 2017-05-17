package main

import "fmt"

type data struct {
	Start int64
	End   int64
	Delta int64
}

func main() {
	var repr, N int
	fmt.Scanf("%d\n", &repr)
	fmt.Scanf("%d\n", &N)

	var times []data
	for i := 0; i < N; i++ {
		var start, end int64
		fmt.Scanf("%d %d\n", &start, &end)

		d := data{
			Start: start,
			End:   end,
			Delta: end - start,
		}

		times = append(times, d)
	}

	var overlaps int
	for i := 0; i < len(times); i++ {
		for j := 1; j < len(times); j++ {
			if i == j {
				continue
			}
			if tt := times[j].End - times[i].Delta; tt > times[i].Start && tt < times[i].End {
				overlaps++
			}
		}
	}

	fmt.Println(overlaps - repr + 1)

}
