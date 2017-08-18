#### Repository forked from [status-im/status-go](https://github.com/status-im/status-go)

See (and maintain) our [notes](Notes.md).

Build results are accessible through S3. 
* The bucket is "adc-status-go". It is followed by a path that consists of the branch name 
  (typically "master" or "develop"), followed by a dash ("-"), followed by a timestamp in the format
  `%Y-%m-%d-%H-%M-%S`. The branch name may contain forward slashes ("/"); this works normally.
* For write access, you need ADC AWS keys. We provide write access to the Travis CI script 
  [.travis.yml](.travis.yml) through encrypted keys in the file. They only work in this project.
* For read access, no keys are necessary. A URL in the form 
  `https://s3.amazonaws.com/adc-status-go/<branch>-<timestamp>/<filename>` works.  
 

[![Master Build Status](https://img.shields.io/travis/AlgebraixData/status-go/master.svg?label=build/master)](https://github.com/AlgebraixData/status-go/tree/master)
[![Develop Build Status](https://img.shields.io/travis/AlgebraixData/status-go/develop.svg?label=build/develop)](https://github.com/AlgebraixData/status-go/tree/develop)
[![TravisCI Builds](https://img.shields.io/badge/TravisCI-URL-yellowgreen.svg?link=https://travis-ci.org/AlgebraixData/status-go)](https://travis-ci.org/AlgebraixData/status-go)

Original content below.

----

# Status bindings for go-ethereum [![Master Build Status](https://img.shields.io/travis/status-im/status-go/master.svg?label=build/master)](https://github.com/status-im/status-go/tree/master) [![Develop Build Status](https://img.shields.io/travis/status-im/status-go/develop.svg?label=build/develop)](https://github.com/status-im/status-go/tree/develop)

[![TravisCI Builds](https://img.shields.io/badge/TravisCI-URL-yellowgreen.svg?link=https://travis-ci.org/status-im/status-go)](https://travis-ci.org/status-im/status-go)
[![GoDoc](https://godoc.org/github.com/status-im/status-go?status.svg)](https://godoc.org/github.com/status-im/status-go)




# Docs

- [How To Build](https://github.com/status-im/status-go/wiki/Build-Process-Explained)
- [Notes on Bindings](https://github.com/status-im/status-go/wiki/Notes-on-Bindings)
- [Developers' Guide](https://github.com/status-im/status-react/wiki/Developers'-Guide)
