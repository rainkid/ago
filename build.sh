basedir=$PWD

cp /home/rainkid/ago/src/github.com/rainkid/dogo/* /home/rainkid/dogo/
export GOPATH=/home/rainkid/ago
cd /home/rainkid/ago/src/github.com/rainkid/dogo/
go build

export GOPATH=$basedir
cd $GOPATH
go build && ./ago
