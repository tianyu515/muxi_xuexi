package main
import"fmt"

type Student struct {
    Name   string
    Scores []int // 记录若干项成绩
}
// func (s Student) Average() float64{
// 	if len(s.Scores) == 0 {
// 		return 0
// 	}
// 	sum := 0
// 	for i := 0;i < len(s.Scores);i++ {
// 		sum += s.Scores[i]
// 	}
// 	return  float64(sum) / float64(len(s.Scores))
// }



// func (s Student) Addscore() []int {
// 	scores := append(s.Scores,88)
// 	return scores
// }



// func (s Student) Top() []int {
// 	if len(s.Scores) <= 0 {
// 		return nil
// 	}
// 	a := make([]int,len(s.Scores))
// 	copy(a,s.Scores)
// 	for i := 0;i < len(a);i++ {
// 		for j :=0;j < len(a)-1-i;j++ {
// 			if a[j] < a[j+1] {
// 				temp := a[j+1]
// 				a[j+1] = a[j]
// 				a[j] = temp
// 			}
// 		}
// 	}
// 	var n int
// 	fmt.Scanf("%d",&n)
// 	if n >= len(a) {
// 		for i := 0;i < len(a);i++ {
// 			fmt.Println(a[i])
// 		}
// 	} else {
// 		for i := 0;i < n;i++ {
// 			fmt.Println(a[i])
// 		}
// 	}
// 	return a

// }


// func (s Student) Merge() Student {
// 	student1 := Student{
// 	 	Name : "小王",
// 		Scores : []int{85,90,94},
// 	}
// 	student2 := Student{
// 	 	Name : "小张",
// 		Scores : []int{88,100,75,99},
// 	}
// 	b := len(student1.Scores)
// 	c := len(student2.Scores)
// 	a := make([]int,b+c)
// 	for i := 0;i < b ;i++ {
// 		a[i] = student1.Scores[i]
// 	}
// 	for j :=0 ;j <c;j++ {
// 		a[b+j] = student2.Scores[j]

// 	}
// 	s.Name = "s"
// 	for i :=0; i < b+c;i++ {
// 		s.Scores[i] = a[i]
// 	}
// 	return s

// }



// func main() {
// 	// student := Student{
// 	// 	Name : "小明",
// 	// 	Scores : []int{85,90,94,88,100,75,99},
// 	// }
// 	// //fmt.Println(student.Scores)
// 	// //fmt.Println(student.Average())
// 	// for i := 0;i < 7; i++ {
// 	// fmt.Println(student.Scores[i])
// 	// }
// 	//fmt.Println(student.Top())
// 	var student3 Student
// 	fmt.Println(student3.Merge())



//}



func (s Student) Merge(other Student) Student {

    // 计算两个切片的长度
    b := len(s.Scores)
    c := len(other.Scores)

    // 创建一个全新的切片
    a := make([]int, b+c)

    // 把第一个人的成绩放进去
    for i := 0; i < b; i++ {
        a[i] = s.Scores[i]
    }

    // 把第二个人的成绩接在后面
    for j := 0; j < c; j++ {
        a[b+j] = other.Scores[j]
    }

    // 返回一个新的 Student
    result := Student{
        Name:   s.Name,
        Scores: a,
    }

    return result
}

func main() {

    student1 := Student{
        Name:   "小王",
        Scores: []int{85, 90, 94},
    }

    student2 := Student{
        Name:   "小张",
        Scores: []int{88, 100, 75, 99},
    }

    merged := student1.Merge(student2)

    fmt.Println(merged)
}
