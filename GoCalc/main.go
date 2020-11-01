package main
import (
"bufio"
"fmt"
"os"
"strconv"
)
var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var st = new(Stack)
type Stack s t r u c t {
i i n t
data [10]i n t
}
func (s *Stack) push(k i n t) {
i f s.i+1 > 9 {
re tu rn
}
s.data[s.i] = k
s.i++
}
func (s *Stack) pop() (ret i n t) {
s.i--
i f s.i < 0 {
s.i = 0
re tu rn
}
ret = s.data[s.i]
re tu rn
}
func main() {
fo r {
s, err := reader.ReadString('\n')
var token s t ri n g
i f err != nil {
re tu rn
}
fo r _, c := range s {
switch {
case c >= '0' && c <= '9':
token = token + s t ri n g(c)
case c == ' ':
r, _ := strconv.Atoi(token)
st.push(r)
token = ""
case c == '+':
fmt.Printf("%d\n", st.pop()+
st.pop())
case c == '*':
fmt.Printf("%d\n", st.pop()*
st.pop())
case c == '-':