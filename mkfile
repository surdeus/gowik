MKSHELL = rc
TGT = gowikd
all:QV: build 
	echo -n
build:QV: `{echo build-^($TGT)}
	echo -n
clean:QV:`{echo clean-^($TGT)}
	echo -n
install:QV:`{echo install-^($TGT)}
	echo -n
build-&:V:
	go build -o bin/$stem ./src/cmd/$stem
install-&:V: build-&
	go install ./src/cmd/$stem
clean-&:V:
	go clean ./src/cmd/$stem	
