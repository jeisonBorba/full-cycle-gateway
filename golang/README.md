<p align="center">
	![Golang](img/golang.svg)
</p>

## Description

Golang application responsible for processing transactions

## Getting Started

### Configure /etc/hosts

The communication between applications takes place directly through the machine's network. For this it is necessary to configure an address that all Docker containers can access.

Add in the hosts file: (Windows -> C:\Windows\system32\drivers\etc\hosts | Linux -> /etc/hosts):

`127.0.0.1 host.docker.internal`


```bash
NOTES:
* Make sure that both docker and docker compose are installed in your machine in order to run the application.
* The kafka application must be started.
* On all operating systems it is necessary to open the program to edit the hosts as the machine's Administrator or root.
```

### Run application

Enter the /golang folder and execute the following command:

> docker-compose up -d (the -d is a optional flag to run the container in detached mode)

To teste the application, use the Apache Kafka to produce and consume messages.