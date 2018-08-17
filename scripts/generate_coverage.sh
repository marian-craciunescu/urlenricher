cd $GOPATH/src/github.com/marian-craciunescu/urlenricher

rm -rf ./cov
mkdir cov

i=0
for dir in $(find . -maxdepth 10 -not -path './.git*' -not -path '*/_test.go' -type d);
do
    if ls ${dir}/*.go &> /dev/null; then
        GO_TEST_DISABLED=true go test -v -covermode=atomic -coverprofile=./cov/$i.out ./${dir}
        i=$((i+1))
    fi
done

gocovmerge ./cov/*.out > full_cov.out
rm -rf ./cov
