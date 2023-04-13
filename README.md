# dockerstopper

Have you ever been annoyed when you try to stop and remove docker containers because you must run `docker ps` to get the container IDs? If so, this tool might help you! This is my first foray into Golang and it was fantastically simple.

## Installation
Clone this repo and you will have the `dockerstopper` binary. All you need to do is add it to your `~/.bashrc` or `~/.zshrc` dotfile by adding

    export PATH=$PATH:~/your/path/to/dockerstopper

or by symbolically linking it to `/usr/local/bin`. Maybe I'll make this easier in the future, idk. 

## Usage
Say you have a few docker containers running locally. You want to inspect them so you run `docker ps` and see the following:

![image](https://user-images.githubusercontent.com/50613550/231899357-bd87f7ec-29fb-4575-bb1c-b1eab3f8eacb.png)

If I want to stop and remove either of these containers, I need to either copy the CONTAINER ID to my clipboard or type it out when I write 

    docker rm -f 49e7b2250bc1`. With `dockerstopper

all I need to write in the terminal is:

    dockerstopper 1

If I want to stop and remove multiple containers, all I need to do is pass `dockerstopper` a list of integers that represent the index of the containers you want to stop. For example, if I want to stop and remove both of these containers, all i need to run is:

   dockerstopper 1,2

That is all. 
