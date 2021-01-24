# tube

`tube` is a Youtube-like (_without censorship and features you don't need!_)
Video Sharing App written in Go which also supports automatic transcoding to
MP4 H.265 AAC, multiple collections and RSS feed.

**this is a fork stripped out of features I don't need, the original project is located [here](https://github.com/prologic/tube)**


## Features

- Easy to add videos (just move a file into the folder)
- Builtin ffmpeg-based Transcoder that automatically converts your uploaded content to MP4 H.264 / AAC
- Builtin automatic thumbnail generator
- No database (video info pulled from file metadata)
- No JavaScript (the player UI is entirely HTML)
- Easy to customize CSS and HTML template
- Clean, simple, familiar UI


## Configuration

`tube` can be configured to suit your particular needs and comes by default with
a sensbile set of defaults. There is also a default configuration at the
top-level [config.json](/config.json) that you can use as a start point and
modify to suite your needs.

To Run `tube` with a provided configuration just pass the `-c /path/to/config`
option; for example:

```#!sh
$ tube -c config.json
```

Everything in the configuration is optional as the builtin defaults are used
if you do not supply anything, omit some sections or values or the configuration
is invalid. Refer to the [default config.json](/config.json) for the builtin
defaults (_this files matches the builtin defaults_).

Here are some documentation on key configuration items:

### Library Options and Upload / Video Paths(s)

```#!json
{
    "library": [
        {
            "path": "videos",
            "prefix": ""
        }
    ],
}
```

Set `path` to the value of the path where you want to store videos and where
`tube` will look for new videos.

### Server Options / Upload Path and Max Upload Size

```#!json
{
    "server": {
        "host": "0.0.0.0",
        "port": 8000,
        "store_path": "tube.db"
    }
}
```

- Set `host` to the interface you wish to bind to. If you want to only bind
  your local machine (_ie: localhost_) set this to `127.0.0.1`.
- Set `port` to any port you wish to bind the listening socket of the server
  to. It doesn't matter what it is as long as there it doesn't collide with
  a port already in use on your system.
- Set `store_path` to a directory where `tube` will store statistics on videos
  viewed.

### Thumbnailer / Transcoder Timeouts

```#!json
{
    "thumbnailer": {
        "timeout": 60
    },
    "transcoder": {
        "timeout": 300,
        "sizes": null
    }
}
```

- Set `timeout` to the no. of seconds to permit for thumbnail generation and
  video transcoding. This value has to be large enough for thumbnail generation
  and transcoding to take place depending on the `max_upload_size` permitted.
  These values also depend on the underlying performance of the machine Tube
  runs on. Use sensible values for your `max_upload_size` + system performance.
  This is a safety measure to ensure background processed do not run away
  and/or hog system resources. The thumbnailer and transcoder processes will
  be killed if their execution time exceeds these values.

- Set `sizes` to an map of `size` => `suffix` that you wish to support for
  transcoding videos to lower quality on Upload/Import. This is especially
  useful for serving up videos to users that have poor bandwidth or where
  data charges are high for them. The following is a valid map:

```#!json
{
    "transcoder": {
        "sizes": {
          "hd720": "720p",
          "hd480": "480p",
          "nhd":   "360p",
          "film":  "240p"
        }
    }
}
```

## License

tube source code is available under the MIT [License](/LICENSE).
