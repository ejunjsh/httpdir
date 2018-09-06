# httpdir

[![Build Status](https://travis-ci.org/ejunjsh/httpdir.svg?branch=master)](https://travis-ci.org/ejunjsh/httpdir)

a commandline to make your directory become a static file server

## install

    #  go get github.com/ejunjsh/httpdir

## usage

    # httpdir -help
    Usage of httpdir:
      -addr string
            listen address (default ":6789")
      -dir string
            directory (default ".")
      -help
            usage

## example

    # httpdir

then open browser and navigate to `http://localhost:6789`, you will see your files in the page.