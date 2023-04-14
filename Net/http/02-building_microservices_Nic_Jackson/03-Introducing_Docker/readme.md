#Docker

## What is docker?
>Docker is the silver bullet that solved the problem with software containers and virtualization once and for all. Yes, that’s a strong claim! Other products had attempted to deal with these problems, but Docker’s fresh approach and ecosystem has wiped the competition off the map. This guide will help you understand the basic concepts of Docker.

- Docker is a technology that allows you to build, run, test, and deploy distributed applications. It uses operating-system-level virtualization to deliver software in packages called containers.
- The way Docker does this is by packaging an application and its dependencies in a virtual container that can run on any computer. This containerization allows for much better portability and efficiency when compared to virtual machines.
- These containers are isolated from each other and bundle their own tools, libraries, and configuration files. They can communicate with each other through well-defined channels. All containers are run by a single operating system kernel, and therefore use few resources.
- They’re portable and can run on any computer that has a Docker runtime environment.
- They’re isolated from each other and can run different versions of the same software without affecting each other.
- They’re extremely lightweight, so they can start up faster and use fewer resources.

![01.jpg](Img/01.jpg)<br/>

## Docker Components and Tools
Docker consists of three major components:
- the **Docker Engine**, a runtime environment for containers
- the **Docker command line** client, used to interact with the Docker Engine
- the **Docker Hub**, a cloud service that provides registry and repository services for Docker images

n addition to these core components, there’s also a number of other tools that work with Docker, including:
- **Swarm**, a clustering and scheduling tool for dockerized applications
- **Docker Desktop**, successor of Docker Machine, and the fastest way to containerize applications
- **Docker Compose**, a tool for defining and running multi-container Docker applications
- **Docker Registry**, an on-premises registry service for storing and managing Docker images
- **Kubernetes**, a container orchestration tool that can be used with Docker
- **Rancher**, a container management platform for delivering Kubernetes-as-a-Service

## Installing Docker
[Take a look](https://docs.docker.com/engine/install/)


## Docker Containers

![Docker Containerized vs Virtualization](Img/02-Docker-containerized-vd-vm-transparent.jpg)<br/>

- Containers are often compared to virtual machines, but there are some important differences between the two. Virtual machines run a full copy of an operating system, whereas containers share the host kernel with other containers. This makes containers much more lightweight and efficient than virtual machines.
- A container image is a lightweight, stand-alone, executable package of a piece of software that includes everything needed to run it: code, runtime, system tools, system libraries, settings.
- Containers isolate software from its surroundings, for example, dif erences between development and staging environments and help reduce conflicts between teams running dif erent software on the same infrastructure.
- Docker containers are built from images, which are read-only template with all the dependencies and configurations required to run an application.
- 

A container, in fact, is a runtime instance of an image — what the image becomes in memory when actually executed. It runs completely isolated from the host environment by default, only accessing host files and ports if configured to do so. As such, containers have their own networking, storage, and process space; and this isolation makes it easy to move containers between hosts without having to worry about compatibility issues.<br/>  
Where they benefit application development is that we can take advantage of this when deploying these applications as it allows us to pack them closer together, saving on hardware resources. <br/>
From a development and test lifecycle, containers give us the capability to run production code on our development machines with no complicated setup; it also allows us to create that **Clean Room environment** without having different instances of the same database installed to trial new software. <br/>
Containers have become the primary choice for packaging microservices, Containers work by isolating processes and filesystems from each other. Unless explicitly specified, containers cannot access each other's file systems. They also cannot interact with one another via TCP or UDP sockets unless again specified.

- the **Dockerfile**, used to build the image.
- the **image** itself, a read-only template with instructions for creating a Docker container
- the **container**, a runnable instance created from an image (you can create, start, stop, move or delete a container using the Docker API or CLI)

## How to Run a Container?

## `docker run [ImageName]`
```dockerfile
 docker run --rm hello-world
```
When you execute a docker run the first thing the engine does is check to see if you have the image installed locally. If it doesn't then it connects to the default registry,in this case, https://hub.docker.com/ to retrieve it.
Once the image has been downloaded, the daemon can create a container from the downloaded image, all the output is streamed to the output on your terminal:

```
Unable to find image 'hello-world:latest' locally
latest: Pulling from library/hello-world
c04b14da8d14: Pull complete
Digest: sha256:0256e8a36e2070f7bf2d0b0763dbabdd67798512411de4cdcf9431a1feb60fd9
Status: Downloaded newer image for hello-world:latest

Hello from Docker!
This message shows that your installation appears to be working correctly.
```
### `--rm`
The `--rm` flag tells the Docker engine to remove the container and delete any resources such as volumes it was using on exit. Unless we would like to re-start a container at some point it is good practice to use the `--rm` flag to **keep our filesystem clean**, otherwise, all of the temporary volumes which are created will sit around and consume space.


```dockerfile
docker run -it --rm alpine:latest sh
```
Alpine is a lightweight version of Linux and is perfect for running Go applications. The `-it` flags stand for interactive terminal it maps the standard in from your terminal to the input of the running container. The sh statement after the name of the image we want to run is the name of the command we would like to execute in the container when it starts.

### `-it`
The `-i` flag stands for **interactive** and is used to keep stdin open even if not attached. The `-t` flag allocates a pseudo TTY device. Together, these two flags allow us to attach directly to our running container and give us an interactive shell session.
If we want to exit this shell, we can simply type `exit`.

>Containers are **immutable** instances of images, and the data volumes are by default **non-persistent**

after `touch mytestfile.txt` and restart container, you can see  file will not exist.

## `docker ps -a`
```dockerfile
docker ps -a
```
The `docker ps` command queries the engine and returns a list of the containers, by default this only shows the running containers, however, if we add the `-a` flag we can also see stopped containers.

## `docker start -it [container_id] sh`
```dockerfile
docker start -it [container_id] sh
```
This is the one we want to restart, so grab the ID of the container and execute the following command,So we can restart our container.

## `docker rm`
If you want to remove all the stopped containers you can use the following command:
```dockerfile
docker rm -v $(docker ps -a -q)
```
The `docker ps -a -q` the `-a` flag will list all the containers including the stopped ones, `-q` will return a list of the container IDs rather than the full details. We are passing this as a parameter list to docker rm, which will remove all the containers in the list.<br/>
To avoid having to remove a container we can use the `--rm` flag when starting a new container. This flag tells Docker to remove the container when it stops.

```dockerfile
docker rm -v [container_id/container_name]
```
Containers that start with a name parameter are not automatically removed even if you specify the `--rm` argument. <br/>
To remove a container started in this way, we must manually use the docker rm command. If we append the `-v` option to the command, we can also remove the volumes that are associated with it. 

## Docker volumes
We have seen how Docker containers are immutable; however, there are some instances when you may wish to write some files to a disk or when you want to read data from a disk such as in a development setup. Docker has the concept of volumes, which can be mounted either from the host running the Docker machine or from another Docker container.

### `-v hostfolder:destinationfolder`
```dockerfile
docker run -it -v $(pwd):/host alpine:latest /bin/sh
```
The `-v, or --volume` parameter allows you to specify a pair of values corresponding to the file system you wish to **mount** on the host and the path where you would like to mount the volume inside the container.<br/>
If you change into the host folder, you will see that there is access to the same folder from where you ran the docker run command. The syntax for the values for `-v is hostfolder:destinationfolder`. <br/>
one thing I think is **important** to point out is that these **_paths need to be absolute_**, ~~and you cannot use a relative path like ./ or ../foldername~~. <br/>
The volume you have just mounted has read/write access, any changes you make will be synchronized to the folder on the host so be careful to not go running rm -rf *. Creating Volumes on a production environment should be used very sparingly, I would advise that where possible you avoid doing it all together as in a production environment there is no guarantee if a container dies and is re-created that it will be replaced on the same host where it was previously. This means that any changes you have made to the volume will be lost.

## Docker ports
When running web applications inside a container, it is quite common that we will need to expose some ports to the outside world.By default, a Docker container is completely isolated, and if you start a server running on port 8080 inside your container unless you explicitly specify that port is accessible from the outside, it will not be accessible.Mapping ports is a good thing from a security perspective as we are operating on a principle of no trust.

```dockerfile
docker run -it --rm -v $(pwd):/src -p 8080:8080 -w /src golang:alpine /bin/sh
```
### `-w [Path]`
The `-w` flag we are passing is to set the working directory that means that any command we run in the container will be run inside this folder. When we start the shell, you will see that rather than having to change into the folder we specify in the second part of the volume mounting we are already in that folder and can run our application.

### golang:alpine
We are also using a slightly different image this time. We are not using alpine:latest, which is a lightweight version of Linux, we are using golang:alpine, which is a version of Alpine with the most recent Go tools installed.

## Docker networking
![04-Network.png](Img/04-Network.png)<br/>
Docker networking is an interesting topic, and by default, Docker supports the following network modes:
* bridge
* host
* none
* overlay
* macvlan

![03-Network.png](Img/03-Network.png)

### Bridge networking
![05-Network bridg.png](Img/05-Network_bridg.png) <br/>
The default network driver. If you don’t specify a driver, this is the type of network you are creating. Bridge networks are usually used when your applications run in standalone containers that need to communicate.<br/>
For the containers on bridge network to communicate or be reachable from the outside world, port mapping needs to be configured.<br/>
When the Docker engine starts, it creates the docker0 virtual interface on the host machine. The docker0 interface is a virtual Ethernet bridge that automatically forwards packets between any other network interfaces that are attached to it. When a container starts it creates a veth pair, it gives one to the container, which becomes its eth0, and the other connects to the docker0 bridge.

* It is a private default network created on the host
* Containers linked to this network have an internal IP address through which they communicate with each other easily
* The Docker server (daemon) creates a virtual ethernet bridge docker0 that operates automatically, by delivering packets among various network interfaces
* These are widely used when applications are executed in a standalone container 
* Default subnet is 172.20.0.0

### Host networking
![06-Network_Host.png](Img%2F06-Network_Host.png) <br/>
For standalone containers, remove network isolation between the container and the Docker host, and use the host’s networking directly.<br/>
The host network is essentially the same network that the Docker engine is running on. When you connect a container to the host network all of the ports that are exposed by the container are automatically mapped to the hosts, it also shares the IP address of the host.<br/>
The host network can also pose a security risk to your container as it is no longer protected by the principle of no trust and you no longer have the ability to explicitly control if a port is exposed or not.<br/>
* It is a public network
* It utilizes the host’s IP address and TCP port space to display the services running inside the container
* It effectively disables network isolation between the docker host and the docker containers, which means using this network driver a user will be unable to run multiple containers on the same host

### None network
For this container, disable all networking. Usually used in conjunction with a custom network driver. none is not available for swarm services.<br/>
Removing your container from any network might in some instances be something you wish to do. Consider the situation where you have an application that only processes data stored in a file.<br/>
* In this network driver, the Docker containers will neither have any access to external networks nor will it be able to communicate with other containers
* This option is used when a user wants to disable the networking access to a container
* In simple terms, None is called a loopback interface, which means it has no external network interfaces 

### Overlay network
![07-Network Overlay.png](Img/07-Network_Overlay.png) <br/>
Overlay networks connect multiple Docker daemons together and enable swarm services to communicate with each other. You can also use overlay networks to facilitate communication between a swarm service and a standalone container, or between two standalone containers on different Docker daemons. This strategy removes the need to do OS-level routing between these containers. <br/>
An overlay network uses software virtualization to create additional layers of network abstraction running on top of a physical network. In Docker, an overlay network driver is used for multi-host network communication. This driver utilizes Virtual Extensible LAN (VXLAN) technology which provide portability between cloud, on-premise and virtual environments. VXLAN solves common portability limitations by extending layer 2 subnets across layer 3 network boundaries, hence containers can run on foreign IP subnets.<br/>
The Docker overlay network is a unique Docker network that is used to connect containers running on separate hosts to one another.The Docker overlay network solves this problem, it is in effect a network tunnel between machines which passes the traffic unmodified over the physical network.<br/>
* This is utilized for creating an internal private network to the Docker nodes in the Docker swarm cluster
* Note: Docker Swarm is a service for containers which facilitates developer teams to build and manage a cluster of swarm nodes within the Docker platform
* It is an important network driver in Docker networking. It helps in providing the interaction between the stand-alone container and the Docker swarm service

### macvlan network
![08-Network_MacvLan.png](Img%2F08-Network_MacvLan.png) <br/>
Macvlan networks allow you to assign a MAC address to a container, making it appear as a physical device on your network. The Docker daemon routes traffic to containers by their MAC addresses. Using the macvlan driver is sometimes the best choice when dealing with legacy applications that expect to be directly connected to the physical network, rather than routed through the Docker host’s network stack.<br/>
The macvlan driver is used to connect Docker containers directly to the host network interfaces through layer 2 segmentation. No use of port mapping or network address translation (NAT) is needed and containers can be assigned a public IP address which is accessible from the outside world. Latency in macvlan networks is low since packets are routed directly from Docker host network interface controller (NIC) to the containers.<br/>
Note that macvlan has to be configured per host, and has support for physical NIC, sub-interface, network bonded interfaces and even teamed interfaces. Traffic is explicitly filtered by the host kernel modules for isolation and security. <br/>
* It simplifies the communication process between containers
* This network assigns a MAC address to the Docker container. With this Mac address, the Docker server (daemon) routes the network traffic to a router
* Note: Docker Daemon is a server which interacts with the operating system and performs all kind of services
* It is suitable when a user wants to directly connect the container to the physical network rather than the Docker host

### ipvlan network
IPvlan networks give users total control over both IPv4 and IPv6 addressing. The VLAN driver builds on top of that in giving operators complete control of layer 2 VLAN tagging and even IPvlan L3 routing for users interested in underlay network integration.<br/>

### Custom network drivers
You can install and use third-party network plugins with Docker. These plugins are available from Docker Hub or from third-party vendors. See the vendor’s documentation for installing and using a given network plugin.<br/>
Docker also supports plugins for networking, based around its open source libnetwork project, you can write custom networking plugins that can replace the networking subsystem of the Docker engine. They also give the capability for you to connect non-Docker applications to your container network such as a physical database server.<br/>
Likes:
* Weaveworks : https://www.weave.works/
* Project Calico : https://www.tigera.io/project-calico/

## `docker network ls`
To see the currently running networks on your Docker engine, we can execute the following command:
```dockerfile
docker network ls
```
## Creating a bridge network
```dockerfile
docker network create testnetwork
```

## Connecting containers to a custom network
after Create your network, Execute :
```dockerfile
docker run -it --rm -v $(pwd):/src -w /src --name server --network=testnetwork golang:alpine go run main.go
```
`--network=testnetwork` To connect a container to a custom network.

for running client request CURL, execute this code on your custom network:
```dockerfile
docker run --rm --network=testnetwork appropriate/curl:latest curl -i -XPOST server:8080/helloworld -d '{"name":"World"}'
```

at first time ,Downloaded curl and then again execute it.

## Delete network
```dockerfile
docker network rm testnetwork
```
