# localactions
local GitHub Actions

https://nektosact.com

install

curl --proto '=https' --tlsv1.2 -sSf https://raw.githubusercontent.com/nektos/act/master/install.sh | sudo bash

act --pull=false


if workflow more then one  

$ act -l

-j job id

$ act -j check -P ubuntu-latest=node:20-bullseye --pull=false

