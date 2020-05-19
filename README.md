# dockerhorn

`dockerhorn` is an easy way to manage your Docker process, with simple actions on your browser.

## What does `dockerhorn` provide?

`dockerhorn` will make your Docker management more easily. You can execute following commands on your browser, with simple actions (like click). 

- Docker
    - `docker info`
- Images
    - `docker images`
    - `docker pull`
    - `docker rmi`
- Containers
    - `docker ps -a`
    - `docker start`
    - `docker commit`
    - `docker stop`
    - `docker rm`

## How to use dockerhorn?

You can start `dockerhorn` on environments that already installed Docker.

```bash
git clone https://github.com/prs-watch/dockerhorn.git
cd dockerhorn
docker-compose up
```

Then, you can open `dockerhorn` on `127.0.0.1:3000`.