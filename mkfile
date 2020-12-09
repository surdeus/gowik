MKSHELL = rc
TGT = gowikd
all:V: build 
build:V: `{echo build-^($TGT)}
clean:V:`{echo clean-^($TGT)}
install:V:`{echo install-^($TGT)}
build-&:V:
	go build -o bin/$stem ./src/cmd/$stem
install-&:V: build-&
	go install ./src/cmd/$stem
test:V: build-$TGT
	./bin/$TGT
clean-&:V:
	go clean ./src/cmd/$stem	
