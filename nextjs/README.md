<p align="center">
	<img src="../img/nextjs.png" width="72"/>
</p>

This is a [Next.js](https://nextjs.org/) project bootstrapped with [`create-next-app`](https://github.com/vercel/next.js/tree/canary/packages/create-next-app).

## Description

Nextjs application responsible for provide a dashboard in order to track the transactions

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

Enter the /nextjs folder and execute the following command:

> docker-compose up -d (the -d is a optional flag to run the container in detached mode)

Open [http://localhost:3000](http://localhost:3000) with your browser to see the result.

The `pages/api` directory is mapped to `/api/*`. Files in this directory are treated as [API routes](https://nextjs.org/docs/api-routes/introduction) instead of React pages.
