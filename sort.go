package sort

import(
	"math/rand"
)

func Bubble(a []int)[]int{
	tmp := 0

	for i:=0; i<len(a); i++{
		for j:=0; j<len(a)-i-1; j++{
			if a[j] > a[j+1]{
				tmp = a[j]
				a[j] = a[j+1]
				a[j+1] = tmp
			}
		}
	}
	return a
}

func Select(a []int)[]int{
	tmp := 0

	for i:=0; i<len(a); i++{
		for j:=0; j<len(a); j++{
			if (a[i] < a[j]){
				tmp = a[j];
				a[j] = a[i];
				a[i] = tmp;
			}
		}
	}
	return a
}

func Insert(a []int)[]int{
	for i:=1; i<len(a); i++{
		tmp := a[i]
		j:=i-1
		for ; j > -1 && a[j] > tmp; j--{
			a[j+1] = a[j]
		}
		a[j+1] = tmp
	}
	return a
}

func Merge(a []int)[]int{
	if len(a) < 2 {
		return a
	}
	mid := len(a)/2
	
	return MergeSort(Merge(a[:mid]), Merge(a[mid:]))
}

func MergeSort(left, right []int)[]int{
	result := make([]int, len(left)+len(right))
	i := 0

	for len(left) > 0 && len(right) > 0{
		if left[0] < right[0]{
			result[i] = left[0]
			left = left[1:]
		} else {
			result[i] = right[0]
			right = right[1:]
		}
		i++
	}

	for j:=0; j < len(left); j++{
		result[i] = left[j]
		i++
	}
	for j:=0; j<len(right); j++{
		result[i] = right[j]
		i++
	}
	return result
}

func Quick(a []int)[]int{
	if len(a)<2{
		return a
	}

	left, right := 0 , len(a)-1
	pivot := rand.Int() % len(a)

	a[pivot], a[right] = a[right], a[pivot]
      
    for i, _ := range a {
        if a[i] < a[right] {
            a[left], a[i] = a[i], a[left]
            left++
        }
    }
      
    a[left], a[right] = a[right], a[left]
      
    Quick(a[:left])
    Quick(a[left+1:])
      
    return a
}