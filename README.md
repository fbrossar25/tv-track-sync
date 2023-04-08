# Installation

## Tautulli

Dans Webhook Settings :
    - URL : http://<SERVER_IP>:8090/tautulli/webhook
    - Method : POST
    - Triggers : Watched
    - Condition : Username is <username>
    - Data > Watched > JSON Data
```json
{
    "action": "{action}",
    "media_type": "{media_type}",
    "user": "{user}",
    "season_num": "{season_num}",
    "episode_num": "{episode_num}",
    "title": "{title}",
    "plex_id": "{plex_id}",
    "imdb_id": "{imdb_id}",
    "thetvdb_id": "{thetvdb_id}",
    "themoviedb_id": "{themoviedb_id}",
    "trakt_url": "{trakt_url}"
}
```

## MongoDB

Dans le dossier du fichier `docker-compose.yml`
```shell
docker-compose up
 ```

Création de l'utilisateur et de la base de données pour tv-track-sync:

1. Ouvrir un bash sur le conainer mongo
    ```shell
    docker exec -it mongo /bin/bash
    ```
2. Ouvrir mongosh
    ```shell
    mongosh
    ```
3. Utiliser la base de données avec le nom dbName du fichier de config `tv-tack-sync.yml`
    ```mongsh
    use tvtracksync 
    ```
4. Créer l'utilisateur pour tv-track-sync correspondant au fichier de config `tv-track-sync.yml`
    ```mongosh
    db.createUser({
      user: "tvtracksync",
      pwd: "tvtracksync",
      roles:[{role: "readWrite", db:"tvtracksync"}]
    })
    ```
