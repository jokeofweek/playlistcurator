# playlistcurator

Given a local music library, the Playlist Curator will create a playlist of similar artists based on a seed artist and threshold. Artist similarity is currently obtained through the [Music Map](http://www.music-map.com/).

The playlist curator supports:
- Arbitrary sources of music through a LibraryProvider interface. There is currently only one implementation which works for the Banshee music player.
- Arbitrary playlist output format. There is currently one implementation which builds an M3U playlist.

## Use
The main package is used as follows:

```
./main "seed artist"
```

A playlist will be created based on the artists similar to the seed artist and output to STDOUT.
