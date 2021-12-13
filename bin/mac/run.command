#!/bin/bash


#dirname
project_path=$(cd `dirname $0`; pwd)
project_name="${project_path##*/}"
echo $project_path
echo $project_name

#runexe
cd $project_path
./$project_name

chmod +x ./easyhttpd
./easyhttpd