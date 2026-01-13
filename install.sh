GEVM_ARCH=$(uname -m)
GEVM_OS=$(uname -s)
GEVM_FILENAME=""

if [[ $GEVM_OS == "Darwin" ]]; then
    if [[ $GEVM_ARCH == "x86_64" ]]; then
        GEVM_FILENAME="gevm-darwin-amd64.zip"
    elif [[ $GEVM_ARCH == "arm64" ]]; then
        GEVM_FILENAME="gevm-darwin-arm64.zip"
    fi
elif [[ $GEVM_OS == "Linux" ]]; then
    if [[ $GEVM_ARCH == "x86_64" ]]; then
        GEVM_FILENAME="gevm-linux-amd64.zip"
    elif [[ $GEVM_ARCH == "arm64" ]]; then
        GEVM_FILENAME="gevm-linux-arm64.zip"
    fi
fi

if [[ -z $GEVM_FILENAME ]]; then
    echo "$GEVM_ARCH on $GEVM_OS not supported"
    exit 1
fi

curl -LOs https://github.com/bashmills/gevm/releases/latest/download/$GEVM_FILENAME
unzip -qo $GEVM_FILENAME
rm $GEVM_FILENAME
mkdir -p ~/.local/bin
mv gevm ~/.local/bin
