#!/bin/sh

for dir in */
	do
	  base=$(basename "$dir")
	  tar -czvf "/nodc/projects/metadata/GHRSST/xml/${base}.tar.gz" "$dir"
	done