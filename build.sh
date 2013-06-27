sh mkroute.sh 
cp  /home/rainkid/ago/src/github.com/rainkid/dogo/* /home/rainkid/dogo/
export GOPATH=/home/rainkid/ago
sh /home/rainkid/ago/src/github.com/rainkid/dogo/build.sh

export GOPATH=$PWD
cd $GOPATH
go build && ./ago
