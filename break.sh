# This should only run on cronjobs
if [ "pull_request" != $TRAVIS_EVENT_TYPE ]; then
   exit 0
fi

# get kubernetes
go get k8s.io/kubernetes
cd $GOPATH/src/k8s.io/kubernetes/vendor

#replace vendored SDK in kubernetes with the latest version
files=()

while IFS= read -d $'\0' -r files
do files+=("$file")
done < <(find github.com/Azure/azure-sdk-for-go '*.go' -print0)

for i in "${files[@]}"
do
    rm $i
    cp $GOPATH/src/$i $(dirname $i)
done

# try to build
cd $GOPATH/src/k8s.io/kubernetes
make
exit $?