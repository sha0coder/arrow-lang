; print several hello worlds:

"hello world" -> print
"hello world\n" -> out
"hello" -> $h -> "$h world" -> print

; if the var is next to a space or the end of string use $var 
;but having the . next to the var use $var$
"world" -> $w -> "hello $w$." -> print


; let's ask user name
"enter your name: " -> out -> in -> $name
"welcome $name, hello world!!" -> print


; now on a function

* hello_world($name)
    $name -> print

"peter" -> hello_world

; now 3 times
(3) => 
    "hello world" -> print


; now n times
3 -> $n 
($n) =>
    "hello world" -> print


