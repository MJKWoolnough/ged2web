#!/bin/bash

cd "$(dirname "$0")";

if [ ! -e gedcom.js ] || [ $(stat -c %s gedcom.js) -ne 41 ]; then
	echo "export const people = [], families = [];" > gedcom.js;
fi;

jslib="$(realpath "../jslib/")";
$jslib/html.sh "$($jslib/requiredHTML.sh ged2web.js)" lib/html.js;

(
	echo "package main";
	echo;
	echo "// File automatically generated with ./embed.sh";
	echo;
	echo "const (";
	echo "	htmlStart = \"<html lang=\\\"en\\\"><head><title>Ged2Web</title><meta charset=\\\"UTF-8\\\" /><script type=\\\"module\\\">\"";
	echo -n "	htmlEnd   = \"</script><style type=\\\"text/css\\\">";
	uglifycss style.css | tr -d "\n"; 
	echo "</style></head><body></body></html>\"";
	echo "	modStart  = \"export const people = [\"";
	echo "	modMid    = \"], families = [\"";
	echo "	modEnd    = \"]\"";

	echo -n "	jsStart   = \"";

	jspacker -i /ged2web.js -n -x | terser -m --module --compress pure_getters,passes=3 --ecma 6 | tr -d '\n' | sed -e 's/\\/\\\\/g' -e 's/"/\\\"/g' -e 's/\(.*=\[\)\(\],[^=]*=\[\)\]/\1\"'"\n"'	jsMid     = \"\2\"'"\n"'	jsEnd     = \"\]/';

	echo "\"";

	echo ")";
) > parts.go
