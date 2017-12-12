# ErrNegativeSqrt
Work for exercise about errors from "A Tour of Go" tutorial (https://tour.golang.org/methods/20).

Two things are important here:
- the type <code>ErrNegativeSqrt</code> is defined as <code>float24</code> (holds the cause of errorr - the bad number);
- a call to <code>fmt.Sprint(e)</code> inside the <code>Error</code> method will send the program into an infinite loop because the <code>fmt.Sprint()</code> looks for <code>e.Error()</code> to represent <code>e</code>.
