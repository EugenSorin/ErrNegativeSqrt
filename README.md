# ErrNegativeSqrt
Work for exercise about errors from "A Tour of Go" tutorial (https://tour.golang.org/methods/20).
Two things are important here:
- the type <b>ErrNegativeSqrt</b> is defined as <b>float24</b> (holds the cause of errorr);
- a call to fmt.Sprint(e) inside the Error method will send the program into an infinite loop because the fmt.Sprint() looks for e.Error() to represent e.
