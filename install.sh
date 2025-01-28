#!/bin/bash
current_dir=$(pwd)
eai_file=eai.sh
file="$current_dir/$eai_file"
alias=/usr/local/bin/eai
chmod +x $file
if [ -e "$alias" ]; then
    rm -f $alias
fi
ln -s $file $alias

echo "eai has been installed!!!!"