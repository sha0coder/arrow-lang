; This is the guess number game

; random number
rand 1 100 -> $guessme

; infinite loop
=>

    ; print and read keyboard in one line, in for stdin, and int for integer.
    "say a number: " -> out -> in -> int -> $num

    [$num == $guessme]
        "contrats you got it, the number is $guessme$." -> print -> end

    [$num < $guessme]
        "the number is bigger" -> print
    
    ; else
    []
        "the number is lower" -> print


; you can use end, continue/cont, break/brk and return
