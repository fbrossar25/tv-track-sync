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