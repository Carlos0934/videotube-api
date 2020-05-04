# Videotuve-API Project

## What is About?

This project is about a simple RestAPI that you may login and upload videos, with sure way and a authorization system made on JWT.


## What You Will Need?

You will need a EC (Elliptc Cuve) private key and public key pairs, with the filenames **key.pem** and **public.pem**. Also you must create 3 directorys, **/data** inside of data **/covers** and **/videos**.

## Routes

#### Users  **/users** GET, POST

#### User **/users/{user-id}*** GET, PUT, DELETE

#### Get New Token **/auth** Any

#### Videos **/users/{user-id}/videos** GET, POST

#### Video **/users/{user-id}/videos{video-id}** GET, PUT, DELETE

#### Upload Media Content **/upload{user-id}/{video-id}** POST

### Get Media Content **/statict/{content}/{fileanme}** GET


### Note
Only there 2 kind of contents, cover and video.
