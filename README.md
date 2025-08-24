# Website for the Konnekt, the non-profit for upcoming artists.

[Konnekt](https://knnkt.dk) is a non-profit youth-driven organization helping young musicians get a platform on the Danish music scene.

We are run by volunteers, and backed by several local sponsors, giving us the possibility to give a stage to the upcoming musicians while giving them the pay they deserve.


# Building the website

The website is consisting of a backend written in [Golang](https://go.dev) and a frontend in [React](https://react.dev), powered by [Vite](https://vite.dev).

The frontend and backend is proxied behind a [Caddy](https://caddyserver.com) HTTPS reverse proxy. All of this is containerized and run, preferably with [Docker Compose](https://docs.docker.com/compose). 

It is prefered to set up [Docker Stack](https://docs.docker.com/reference/cli/docker/stack/) between development machine and the server hosting the website, creating a 
[Docker Context](https://docs.docker.com/engine/manage-resources/contexts/) and deploying directly from the development machine.


## CI/CD Pipeline

The repository makes use of [Github Actions](https://github.com/features/actions) to push tested containers to the repository's Container Registry (see [GitHub Docs](https://docs.github.com/en/packages/working-with-a-github-packages-registry/working-with-the-container-registry) for more information).

Before a pull request is made from `dev` to `main`, both backend and frontend code is tested before being built into their respective images and pushed to the registry. Hereafter it can be deployed using Docker Stack (see below).


## Docker Stack Deployment

```
# Creating the Docker context.
$ docker context create <context_name> --docker "host=ssh://<remote_username>@<domain>"

# Using the newly created Docker context. This makes all docker commands be run remotely. All of the following commands are now run on the remote Docker instance.
$ docker context use <context_name> 

# Initialising and enabling swarm mode on remote.
$ docker swarm init

# Deploying the services with the input compose file.
$ docker stack deploy -c <compose_file_path> <stack_name> --with-registry-auth
```
