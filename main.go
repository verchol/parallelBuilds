package main

import (
    "fmt"
	"math/rand" 
	"sort"
	"time"
	"flag"
)

type Build struct {
	 start int
	 end   int
	 isEnd bool
}

func GenerateBuild(num int)   map[int]Build{
	 builds:= map[int]Build{}
	 rand.Seed(time.Now().UnixNano())
	for i := 0; i < num; i++ {
		b := Build{}
	 
		b.start = rand.Intn(80) + 1
		b.end =   b.start + rand.Intn(20)  + 1
		builds[b.start] = b;
		builds[b.end] = Build{b.start, b.end, true}
		fmt.Printf("generated - [%v-%v]\n", b.start, b.end)
	}
	
	return builds
	
}
 
func Format(n int, c rune) string{
  str:= ""
  for i:=0; i <n ; i++ {
	  //s := fmt.Sprintf("%s", c)
	  str = str + string(c)
	}

  return str;
}
func Draw(keys []int, builds map[int]Build){
	for _, start:= range  keys{
	
	   b:= builds[start]
	   if (b.isEnd){
		continue
	   } 
	   buildLen := b.end - b.start
	   leftPad := Format(b.start , '.');
	   drawBuild := Format(buildLen, 'X');
       rightPad := Format(100 - buildLen - b.start, '.');

	   fmt.Printf("%v\n", leftPad + drawBuild+ rightPad)
	  
	  }
}
func main() {
	numOfBuilds := flag.Int("builds", 100, "number of build to generate")
	toDraw := flag.Bool("draw", true, "to draw builds ")
	flag.Parse();
	builds:= GenerateBuild(*numOfBuilds)
	
	play(builds, *toDraw)
}
type Builds []int

func (a Builds) Len() int           { return len(a) }
func (a Builds) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a Builds) Less(i, j int) bool { return a[i]  < a[j]  }
func endBuilds(builds []Build, start int, max *int) {
	 
}
func play(builds map[int]Build, toDraw bool){
 
  //sort.Sort(Builds(builds))
  var max int = 0
  result := []int{}
  keys := []int{}
  for k, _:= range  builds{
	 
		keys= append(keys, k)
 
  }
  if (toDraw){
  Draw(keys, builds)
  }
  sort.Ints(keys);
  for _, start:= range  keys{
 
	b:= builds[start]
 
	//fmt.Printf("[%v-%v]\n", b.start, b.end)
   if (b.isEnd){
	result = append(result , max);
	max--
   }else {
	   max++
   }
  }
  sort.Sort(Builds(result));
  fmt.Printf("[%v]\n", result)
  fmt.Printf("max paralell is [%d]", result[len(result) - 1])

}