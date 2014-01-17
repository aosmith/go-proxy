TLDR; A simple websocket / tcp proxy written in go.

I use this to work around upstream restrictions.  It allows you to proxy http traffic through a websocket (eg, javascript client).

You'll need my weird go alias:

```
alias go='GOPATH=`pwd`:$GOPATH go'
```

You'll also need to copy & configure src/config/config.go


Todo:
-----

  * Consider using channels instead of simple byte buffers.
  * HTTPS?
