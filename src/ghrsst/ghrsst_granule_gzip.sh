#!/bin/sh

for dir in */
	do
	  base=$(basename "$dir")
	  tar -czvf "/nodc/projects/satdata/Granule_OneStop/GHRSST/xml/${base}.tar.gz" "$dir"
	done