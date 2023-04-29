## MGO
To install mgo
```
go get gopkg.in/mgo.v2 
```
More details: [here](https://github.com/tobyzxj/mgo)

## Cucumber
### Install Godog
To install Cucumber (godog) package
```
go get github.com/DATA-DOG/godog/cmd/godog
```

### Set environment
Export PATH point to GOLANG environment
```
#=============== Linux ===================#
# the executable is here after installation
# $GOPATH/bin/godog
export PATH=$PATH:$GOPATH

#=============== Window ===================#
set PATH=%PATH%;%GOPATH%

```
### Folder contain godog library
To check godog.exe is exist in your environment
```
# godog.exe is in directory C:\Go\bin
C:\Go\bin\godog.exe
```
More details: [here](https://techblog.fexcofts.com/2019/08/09/go-and-test-cucumber/)


### Issue knowledge when download
#### Issue 1: wrong link to download
> go get github.com/DATA-DOG/godog/cmd/godog
```
module declares its path as: github.com/cucumber/godog
                but was required as: github.com/DATA-DOG/godog
```
To fix it.
> go get github.com/cucumber/godog/cmd/godog

More details: [here](https://github.com/cucumber/godog/issues/211)

#### Issue 2: Wrong path to working godog

**Behavior:** When you run command cucumber in Makefile
```
make cucumber
```

**Error:** You will see error below:
```
- Container chapter4-mongodb-1  Running                                                                                                                                                                                                0.0s
cd features && godog ./
failed to compile tested package: ....\features, reason: exit status 1, output: WORK=....\go-build013581706
# docker-compose/features
search_test.go:14:2: import "github.com/cucumber/godog/cmd/godog" is a program, not an importable package
FAIL    docker-compose/features [setup failed]

make: *** [cucumber] Error 1
```
**Solution:**
```
Change command from Makefile.
From:
    cd features && godog ./
To: 
    cd godog ./
```

**Error:** You will see error below:
```
failed to compile testmain package: exit status 1 - output: compile: 
/var/folders/n9/m54nxwb5133cvnr0289zcb9h0000gn/T/go-build1244285593/b001/importcfg.link:185:
unknown directive "modinfo""
```

**Solution:**
```
go install github.com/cucumber/godog/cmd/godog@latest
```

### Run godog
To run this godog for this project
```
godog ./
```

To run with docker-compose
```
make cucumber
```

