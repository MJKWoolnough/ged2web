#!/bin/bash
(
	echo "package main";
	echo;
	echo "const (";
	echo "	htmlStart = \"<html><head><title>Ged2Web</title><script type=\\\"module\\\">\"";
	echo -n "	htmlEnd   = \"</script><style type=\\\"text/css\\\">";
	uglifycss style.css | tr -d "\n"; 
	echo "</style></head><body></body></html>\"";
	echo "	modStart  = \"export const people = [\"";
	echo "	modMid    = \"], families = [\"";
	echo "	modEnd    = \"]\"";

	names=( "jsStart  " "jsMid    " "jsEnd    " )
	pos=0;

	while read part; do
		echo "	${names[$pos]} = \""$part"\"";
		let "pos++";
	done < <(jslib -i /ged2web.js -n -x | tail -n+2 | sed -e 's/pageLoad/(document.readyState == "complete" ? Promise.resolve() : new Promise(successFn => globalThis.addEventListener("load", successFn, {once: true})))/' | terser -m  --module --compress pure_getters,passes=3 --ecma 6 | sed -e 's/\(.*=\[\)\(\],[^=]*=\[\)/\1'"\n"'\2'"\n"'/' -e 's/"/\\\\\"/g');

	echo ")";
) > parts.go
