#!/bin/sh

start=$(date +%s.%N)

#Find all WOA13 .nc files and copy them into input folder of nc2iso script
# find /nodc/web/data.nodc/htdocs/nodc/archive/data/0114815/public -name \*.nc -print0 | xargs -I{} -0 cp -v {} /nodc/users/tjaensch/git_repo/gostuff/src/boulder_demo/

#Write input.txt file for nc2iso script based on .nc files in input folder
find . -type f -name "*.nc" -printf "%f \n" > input.txt

FILENAME=input.txt
FILESIZE=$(stat -c%s "$FILENAME")
echo "$FILESIZE"
	if (($FILESIZE > 0)); then
		while read line
		do
			bash WOA_nc2iso_thomas_script_blossom.sh $line  ;
		done < input.txt;

	else
	echo "No .nc files found in input folder!"
	fi

echo "All done!"

end=$(date +%s.%N)
runtime=$(python -c "print(${end} - ${start})")

echo "Runtime was $runtime seconds."
