# Dockerfile
**Dockerfiles are the recipes for our images**; the define the base image, software to be installed and give us the capability to set the various structure that our application needs.

## Building application code for Docker
The build process is consists of two steps: compilation and linking. 
1. Compilation: We compile the program. A main.o file is produced: `go tool compile main.go` It is an intermediary archive file.
2. Linking: With go tool link, we produce the final executable. `go tool link -o server main.o`
3. We can examine the program with the file command. `file ./server`

We are going to execute a slightly different command for creating our files from the usual `go build`:
```go
CGO_ENABLED=0 GOOS=linux GOARCH=386 go build -a -installsuffix cgo -ldflags '-s -w' -o server
```
we are passing the argument -ldflags '-s', this argument passes the -s argument to the linker when we build the application and tells it to statically link all dependencies. This is very useful when we use the popular Scratch container as a base; Scratch is the lightest base you can get it has no application frameworks or applications this is opposed to Ubuntu.<br/>
The difference between Scratch and Ubuntu is that Scratch does not have access to the standard C library GLibC.<br/>
`-s` : Omit the symbol table and debug information.<br/>
`-w` : Omit the DWARF symbol table.<br/>
to strip DWARF, symbol table and debug info. Expect ~25% - ~40% binary size decrease and reduce the size of the resulting binary, you can strip off information not needed during execution.

### From [image-name:tag]
The FROM instruction **set the base image** for subsequent instructions. You can use any image that is either stored in a remote registry or locally on your Docker Engine. When you execute docker build, if you do not already have this image, then Docker will pull it from the registry as the first step of the build process.
* FROM image // assuming latest
* FROM image:tag // where you can specify a tag to use

```dockerfile
FROM scratch
```
`scratch`: When building Docker containers you define your **base image** in your dockerfile. The scratch image is the **smallest possible image** for docker. Actually, by itself it is empty (**in that it doesn’t contain any folders or files**) and is the starting point for building out images like(such as debian and busybox) or super minimal images (that contain only a single binary and whatever it requires, such as hello-world).<br/>
In order to run binary files on a scratch image, your executables need to be statically compiled and self-contained. This means there is no compiler in the image, so you’re left with just system calls. we need to run our Go application is the application itself then we can use scratch to produce the smallest possible image.<br/>
You also wouldn’t be able to login to the container either as there isn’t a shell unless you explicitly add one.

### MAINTAINER [Author name/Email]
The MAINTAINER instruction allows you to set the author of the generated image. This is an optional instruction; however, it can be good practice to include this even if you are not planning on publishing your image to the public registry.

### EXPOSE [port]
The EXPOSE instruction informs Docker that the container listens on the specified networks ports at runtime. Expose does not make the ports accessible to the host; this function still needs to be performed with the -p mapping.

### COPY [sourcefile]:[destinationfile]
The COPY instruction copies files from the source in the first part of this instruction to the destination specified in the second part:
* COPY <src> <dest>
* COPY ["<src">, "<dest>"] // useful when paths contain whitespace
The <src> in the COPY instruction may contain wildcards with the matching done using Go's filepath.Match rules.

Note:
* <src> must be part of the context for the build, you cannot specify relative folders such as ../;
* A root / specified in the <src> will be the root of the context
* A root / specified in the <dest> will map to the containers root file system
* Specifying a COPY instruction without a destination will copy the file or folder into the WORKDIR with the same name as the original

### ENTRYPOINT ["executable", "param1", "param2"]
An ENTRYPOINT allows you to configure the executable that you would like to run when your container starts. Using ENTRYPOINT makes it possible to specify arguments as part of the docker run command which is appended to the ENTRYPOINT.<br/>
ENTRYPOINT has two forms:
* ENTRYPOINT ["executable", "param1", "param2"] // preferred form
* ENTRYPOINT command param1 param2 //shell form

We can, however, pass additional arguments to the application via the docker run command arguments; these would then be appended to the ENTRYPOINT before the application is run. For example:
```dockerfile
docker run --rm helloworld --config=/configfile.json
```
it would be the equivalent of executing the following shell command: ` ./server --config=configfile.json`

### CMD
The CMD instruction has three forms:
* CMD ["executable", "param1", "param2"] // exec form
* CMD ["param1", "param2"] // append default parameters to ENTRYPOINT
* CMD command param1 param2 // shell form
When CMD is used to provide default arguments for the ENTRYPOINT instruction then both the CMD and ENTRYPOINT instructions should be specified using the JSON array format.<br/>
If we specify a default value for CMD, we can still override it by passing the command arguments to the docker run command.<br/>
**Only one CMD instruction is permitted in a Docker file.**

## Good practice for creating Dockerfiles
Taking all of this into account, we need to remember how the union file system works in Docker and how we can leverage it to create small and compact images. Every time we issue a command in the Dockerfile, Docker will create a new layer. When we mutate this command, the layer must be completely recreated and potentially all the following layers too, which can dramatically slow down your build. It is therefore recommended a good practice that you should attempt to group your commands as tightly as possible to reduce the possibility of this occurring.<br/>
Quite often, you will see Dockerfiles which instead of having a separate RUN command for every command we would like to execute, we chain these using standard bash formatting.<br/>
For example, consider the following, which would install software from a package manager.<br/><br/>
**Bad Practice:**
```
RUN apt-get update
RUN apt-get install -y wget
RUN apt-get install -y curl
RUN apt-get install -y nginx
```
<br/>

**Good Practice:**
```
RUN apt-get update && \
    apt-get install -y wget curl nginx
```

The second example would only create one layer, which in turn would create a much smaller and more compact image, it is also good practice to organize your COPY statements placing the statement which changes the least further up in the Dockerfile, this way you avoid invalidation of subsequent layers even if there are no changes to these layers.

## Building images from Dockerfiles
To build an image from our Dockerfile, we can execute a straightforward command:
```dockerfile
docker build -t testserver .
```
Breaking this down the `-t` argument is the tag we wish to give the container, this takes the form `name:tag`, If we omit the tag portion of the argument as we have in our example command, then the tag `latest` will be automatically assigned.

Now that we have our running container, let's test it out. Why not start a container from our newly built image and check the API by curling the endpoint:
```
$ docker run --rm -p 8080:8080 testserver
$ curl -XPOST localhost:8080/helloworld -d '{"name":"world"}'
```

