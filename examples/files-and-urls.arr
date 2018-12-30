; on arrow language urls and files are a type of data:


; is productive isn't it?
http://google.com -> $html -> ./myfile.txt

; or even
http://google.com -> ./myfile.txt

http://google.com -> append ./myfile.txt

; here is not an url type, its an string:
"http://google.com" -> $url -> get -> $html -> pritn


; post
"login=admin&pwd=admin" -> post http://lalala.com/login.go -> $html
[$html =~ "granted"]
    "yeah" -> print



./urls_list.txt -> split '\n' -> $urls
$urls =>
    -> $url -> append 'html.log'

