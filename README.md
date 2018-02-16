# hypothesisbot

Hypothesisbot (AKA StarkRavingMadBot) is a Discord bot written in Go. 
The goal for the project is to focus on statisitics and other fun things you can do with the data available from chat logs (like Markov chains).

## Running

I highly recommened running this project inside docker with mongo. To do so:

```
docker pull 7thfox/hypothesisbot
```

Next we'll create the configs:

```mkdir config
touch token
touch config.json
```

Within `token` you should place your token. `config.json` can have whatever setting you wish
in addition to a few things that are specific for this setup:

```
{
  "global": {
    "token": "/data/token",
    ...
    "db": {
      "type": "mongo",
      "host": "hypo-mongo",
      ...
    }
    ...
  }
  ...
}
```

Now that the config is setup, we need to make a container for MongoDB as well.

```
docker pull mongo
mkdir mongo-data
```

Finally, we need to setup a network for these containers to communicate on.

```
docker network create hypo-net
```

Now we can run stuff:
```
# Setup mongo db with external data
docker run -dit -v $(pwd)/mongo-data:/data/db --name hypo-mongo --network=hypo-net --net-alias=hypo-net mongo
# Setup bot with external config and web interface exposed on port 42968
docker run -dit -v $(pwd)/config:/data --network=hypo-net --net-alias=bot --name hypobot --publish 42968:8080 --restart always 7thfox/hypothesisbot
```

## Upgrading

Comming soon:tm:

TL;DR: rebuild the image, which in turn, pulls the latest from this github repo.
