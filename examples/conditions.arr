; conditions 

1 -> $a -> $b

; lists without spaces on commas
list "a","b","c" -> $list
"this is a test" -> $txt



[1 == 1 && $a == $b && 'c' in $list && "test" in $txt]
    "it\'s logic, isn\'t it?" -> print


"11111111-C" -> $id

; perl style
[$id =~ "[0-9]{8}-[A-Za-z]"]
    "sytactically correct" -> print
[]
    "syntactically incorrect" -> print


; the [] is an else
