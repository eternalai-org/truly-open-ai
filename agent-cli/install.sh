#!/bin/bash
yarn build

current_dir=$(pwd)
file="$current_dir/bin/eai.js"
alias=/usr/local/bin/eai
chmod +x $file
if [ -e "$alias" ]; then
    rm -f $alias
fi
ln -s $file $alias

echo "eai has been installed!!!"