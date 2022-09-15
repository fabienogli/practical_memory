# Practical Memory 
See the original with better documentation, better code and instructions
[github link](https://github.com/ardanlabs/gotraining/tree/master/topics/go/profiling/memcpu)
# First cmd that show bench and allocation	
```bash
go test -bench . -benchmem
```
## export memory profile to p.out
```bash
go test -bench . -benchmem -memprofile p.out 
```
Now we have this
```bash
practical_memory ➤ ls
cmd.md  
go.mod  
p.out					#
practical_memory.test	# test binary compiler built for this test
stream.go  				# program file
stream_test.go			# benchmark test file
```


We want to explore our mem
```bash
go tool pprof p.out 
```
### inside pprof
(pprof) list algOne
--> allow to see prog
2 column -> 1 is flat and the other is cumulative (cum):
ROUTINE ======================== github.com/fabienogli/practical_memory.algOne in /home/fabien.ogli/gophercon/practical_memory/stream.go
   14.50MB   153.51MB (flat, cum) 98.40% of Total
            .          .     57:   fmt.Printf("Matched: %v\nInp: [%s]\nExp: [%s]\nGot: [%s]\n", matched == 0, in, out, output.Bytes())

flat 	--> allocation is happening because of code within function
cum 	--> happenning inside the called path

You can see a web view of it by doing
(pprof) weblist algOne
What's weird is 1 allocation is not showned cumulative but flat
# Go deeper
Ok but we want to know what the compiler is doing, we need to see escape analysis report
```bash
go test -bench . -benchmem -memprofile p.out -gcflags -m=2 
```
The following allows you to see the options
(pprof) o

And we see the following
noinlines                 = false

inlining --> the most important compiler optimisation -> we are not doing this function call, we are gonne take the code out of this function and inline it within the funciton

if we set it to true, now the terminal view is the same than the webview
You wan also run pprof directly with noinline option set to true
```bash
go tool pprof -noinlines p.out 
```
## CPU Profiling
If you want to see the where in the programing, it can be optimized
```bash
go test -bench . -benchmem  -cpuprofile p.out
```
