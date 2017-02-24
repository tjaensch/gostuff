#!/bin/bash

#Function for generic error checking
function error_exit
{
	echo "$1" 1>&2
	exit 1
}

#WOA13 collection metadata template file
isocofile="/nodc/web/data.nodc/htdocs/nodc/archive/metadata/approved/iso/0114815.xml"

#Loop through script arguments and process one by one
for filename in "$@"
	do
		if [ -z "$filename" ]; then
			error_exit "File not available, program exiting."
		fi

		#WOA13 file path on network
		if datapath=$(find /nodc/web/data.nodc/htdocs/nodc/archive/data/0114815/public -type f -name $filename); then
			#Shorten path to make it start with WOA13/DATA/...
			datapathandfile=${datapath#/nodc/web/woa.data.nodc/}
			#Add "woa/" at the beginning and remove filename from data path at the end
			datapathonly=woa/${datapathandfile%$filename}
		else
			error_exit "Something went wrong with retrieving the data path, program exiting."
		fi

		#Determine .nc file size in MB
		filesize=$(( $( stat -c '%s' $filename) / 1024 / 1024 ))

		#Remove ".nc" from $filename
		filename=${filename%.nc}

		#nccopy/ncdump on blossom
		if	ncdump -x $filename.nc > $filename.ncml; then
			true
		else
			echo "$filename " >> errors.txt
			ncdump -x $filename.nc 2>&1 | tee -a errors.txt
			sed -i '$ a ++++++++++++++++: ' errors.txt
			error_exit "Something went wrong with nccopy/ncdump, program exiting."
		fi

		#Replace the second line in the ncml file with <netcdf>
		if sed -i '2 c\ \<netcdf\>' $filename.ncml;
		   #Insert variables into ncml file to be used for extra service links in XSLT file
		   sed -i '3 a <path>'$datapathonly'</path>' $filename.ncml;
		   sed -i '4 a <title>'$filename'</title>' $filename.ncml;
		   sed -i '5 a <filesize>'$filesize'</filesize>' $filename.ncml; then
			#Apply modified UnidataDD2MI XSL to work with WOA data
			xsltproc /nodc/users/tjaensch/xsl.git/boulder_demo/bash_process/XSL/ncml2iso_modified_from_UnidataDD2MI_demo_WOA_Thomas_edits.xsl $filename.ncml > $filename.xml
			#Apply WOA collection metadata
			xsltproc --stringparam collFile $isocofile /nodc/users/tjaensch/xsl.git/boulder_demo/bash_process/XSL/granule.xsl $filename.xml > output/$filename.xml
		else
			error_exit "Something went wrong with xsltproc, program exiting."
		fi

		rm *.ncml
		rm $filename.xml

		echo "$filename.xml successfully written to output directory."
	done
	#end of script argument loop

exitcode=$?
	exit $exitcode
