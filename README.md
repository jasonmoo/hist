#hist

cli tool to generate histograms

	jason@mba ~: go get github.com/jasonmoo/hist

	jason@mba ~: hist -?
	Description: hist reads lines from stdin and outputs a histogram next to each line
	  -?=false: output help
	  -c="#": char to use
	  -f=1: field to hist
	  -w=20: width of hist


	jason@mba ~: wc -l /usr/share/dict/* | hist
	                      :       39 /usr/share/dict/README
	                      :      150 /usr/share/dict/connectives
	                      :     1308 /usr/share/dict/propernames
	             ######## :   235886 /usr/share/dict/web2
	                   ## :    76205 /usr/share/dict/web2a
	             ######## :   235886 /usr/share/dict/words
	 #################### :   549474 total


[LICENSE](https://raw.github.com/jasonmoo/hist/master/LICENSE)