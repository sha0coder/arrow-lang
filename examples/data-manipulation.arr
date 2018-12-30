; manipulating data

; a simple list
list 1,2,3 -> $l


; push and pop
4 -> push $l
$l -> pop $l -> $last_element
 


; filter arry with grep+regex like perl
./somefile.txt -> split '\n' -> grep '[a-zA-Z0-9]' -> $remaining_lines -> join '\n' -> ./file2.txt


; a simple recursive crawler in three lines
* crawl_url(url)
    ; make an http get, extract urls and store on $urls array
    $url -> get -> extract 'src="([^"])"' -> $urls
    ; iterate $urls array
    $urls =>
        ; print each url and send to crawl_url to get more urls
        -> print -> crawl_url

"http://soundcloud.com" -> crawl_url

; we can use any python method for example upper:
"test".upper() -> print


; grep filter an array and create a new array
; extract filter text and create an array