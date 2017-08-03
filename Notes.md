### ADC Notes

#### TODO

-   [x] When building the project, a number of Go references  point to their own repository.  
        **Conclusion**: These now point to our forked repo.

-   [x] Review all dependencies and decide which ones we may want to copy/fork.  
        **Conclusion**: All 3rd-party Go package references are committed to the "vendor" directory.
        The only other dependencies I have found are related to the user "farazdagi" and described
        below.

        About the **Go dependencies**: It is not fully clear whether the build uses the online
        versions or the committed versions, but it seems that it does use the committed versions (the
        build output indicates so much). OTOH, Go's dependency listing prints out the package names as
        shown below. However, these could just be the package names and not necessarily their locations.
        To make sure, we'd have to analyze the `xgo` build configurations in more detail. (See also the
        next item.) In any case, I think we can leave this as-is; we definitely have the packages
        committed, and even if the current build uses online sources, we can make it work with the
        committed ones if needed.

        3rd party package dependencies:
        -   https://github.com/btcsuite
            -   github.com/btcsuite/btcd/btcec
            -   github.com/btcsuite/btcd/chaincfg
            -   github.com/btcsuite/btcd/chaincfg/chainhash
            -   github.com/btcsuite/btcutil
            -   github.com/btcsuite/btcutil/base58
        -   https://github.com/eapache
            -   github.com/eapache/go-resiliency/semaphore
        -   https://github.com/ethereum/go-ethereum
            -   github.com/ethereum/go-ethereum/*
        -   https://github.com/robertkrimen/otto
            -   github.com/robertkrimen/otto
        -   https://github.com/stretchr/testify
            -   github.com/stretchr/testify/*
        -   https://godoc.org/-/subrepo
            -   golang.org/x/crypto/pbkdf2
            -   golang.org/x/text/unicode/norm
        -   http://gopkg.in/urfave/cli.v1 -> https://github.com/urfave/cli
            -   gopkg.in/urfave/cli.v1

-   [ ] A number of external dependencies are committed to to the `status-go` repository.
        However, the Go files in those dependencies often import the original location. (For
        example, a number of `btcsuite` projects are committed in
        status-go/vendor/github.com/btcsuite, but btcsuite/btcd/chaincfg/chainhash/hashfuncs.go
        contains `import "github.com/btcsuite/fastsha256"`.) I don't know for sure whether this
        `import` statement references the file as indicated by the URL path or the one that's
        committed to the repository. It looks to me as if it referenced the original file on
        Github. If this is correct, cleaning this up would be a significant amount of work.

-   [x] Find all references of `farazdagi` and copy the referenced objects (S3 files, gists,
        Docker images, ...) or at least make sure we reference something that's guaranteed
        to be persisting (Docker images with tag?).  
        **Note**: With the proposed changes below, we now have several places that we need to
        check for upstream updates when we update the repository.

    -   [x] The [Dockerfile](Dockerfile) creates an image with a `statusd` executable. It uses the
            branch "feature/statusd-replaces-geth-on-cluster" of a
            [fork of status-go](https://github.com/farazdagi/status-go). This branch doesn't exist
            in this fork (and neither in the original). It seems there isn't much we can do at this
            point; the Dockerfile doesn't seem to work as-is. We can clone the `farazdagi/status-go`
            repository, but it wouldn't help us to get this going.
            **Conclusion**: Do nothing.
    -   [x] The build uses two Docker images: [farazdagi/xgo](https://hub.docker.com/r/farazdagi/xgo/)
            and [farazdagi/xgo-ios-simulator](https://hub.docker.com/r/farazdagi/xgo-ios-simulator/).
            **Conclusion**: We submit these images as they are to the Docker hub with our own
            path/tag and then reference these images. The new paths are
            [algebraixendurance/status-xgo](https://hub.docker.com/r/algebraixendurance/status-xgo/)
            and
            [algebraixendurance/status-xgo-ios-simulator](https://hub.docker.com/r/algebraixendurance/status-xgo-ios-simulator/).
    -   [x] The iOS simulator build uses a copy of the iPhone simulator SDK that they have made available through a
            [public S3 location](https://s3.amazonaws.com/farazdagi/status-im/iPhoneSimulator9.3.sdk.tar.gz). I'm not
            sure this it is legal to publish a copy of the SDK. This is something that we can address if we should have to.
            **Conclusion**: Do nothing, but keep a private copy of the SDK at
            [s3://alice-status/iPhoneSimulator9.3.sdk.tar.gz](https://s3.amazonaws.com/alice-status/iPhoneSimulator9.3.sdk.tar.gz).
            When/if we need to do something here, we probably shouldn't make the SDK publicly available
            without checking the license first.
    -   [x] Several mentions in the [package.json](package.json) file. This file seems to be an `npm` package
            specification and looks like a left-over from some time ago. It also references the 
            [farazdagi/status-go](https://github.com/farazdagi/status-go) repository even though it is in the 
            [status-im/status-go](https://github.com/status-im/status-go) repository. To me it seems this file is not
            used anymore and the references in it can be ignored.  
            **Conclusion**: Do nothing.
    -   [x] The code uses two `CHTRootConfigURL`s in the form of gists
            (`cht.json`
            [farazdagi/a8d36e2818b3b2b6074d691da63a0c36](https://gist.githubusercontent.com/farazdagi/a8d36e2818b3b2b6074d691da63a0c36)
            and `cht-sandbox.json`
            [farazdagi/3d05d1d3bfa36db7b650c955e23fd7ae](https://gist.githubusercontent.com/farazdagi/3d05d1d3bfa36db7b650c955e23fd7ae)).
            These are marked with "TODO remove this hack, once CHT sync is implemented on LES side".  
            **Conclusion**: Fork the gists and reference our gists. Replacements:
            -   `cht.json`: [GFiedler-ADC/082d351623c7cf80d9b31eebed457087](https://gist.github.com/GFiedler-ADC/082d351623c7cf80d9b31eebed457087) replaces [farazdagi/a8d36e2818b3b2b6074d691da63a0c36](https://gist.githubusercontent.com/farazdagi/a8d36e2818b3b2b6074d691da63a0c36).
            -   `cht-sandbox.json`: [GFiedler-ADC/ccd7b962775bded311d4ae500bfbc27a](https://gist.github.com/GFiedler-ADC/ccd7b962775bded311d4ae500bfbc27a) replaces [farazdagi/3d05d1d3bfa36db7b650c955e23fd7ae](https://gist.githubusercontent.com/farazdagi/3d05d1d3bfa36db7b650c955e23fd7ae).
        
-   [ ] Understand what the different OS builds need and make sure it's deployed to the Github pages.
        Ben is working on this.
-   [ ] Re-enable the test and analyze the failures. (Consider that the original repository also has failures.)
-   [ ] Create an organization on Docker hub (rather than continue to use Endurance's personal account), move
        the two Docker images that are used in the build to the organization and set up the integration with Github
        to automatically build them.


#### Observations

*   [Dockerfile](Dockerfile): 
    *   References the branch `feature/statusd-replaces-geth-on-cluster` in the fork 
        [farazdagi/status-go](https://github.com/farazdagi/status-go). The branch doesn't seem to exist.
*   [Makefile](Makefile): 
    *   Uses the docker image [algebraixendurance/status-xgo](https://hub.docker.com/r/algebraixendurance/status-xgo/)
        for the normal builds. It is based `xgo-1.7.1` (a cross-compiler for Go) and adds `build.sh` as custom
        build script. This script sets up rather complex build environments and performs the actual build. The iOS
        simulator build is done with another docker image
        [algebraixendurance/status-xgo-ios-simulator](https://hub.docker.com/r/algebraixendurance/status-xgo-ios-simulator/)
        that is derived by what appears to be replacing of the iOS 9.3 SDK with the iPhoneSimulator 9.3 SDK.
        (The code for these docker images is in the directory "xgo".)
    *   Because the builds are run in docker images, typical `make` dependency management doesn't work; all
        builds are full builds, starting from scratch. (`make` dependencies are used to create dependencies
        between the different `make` targets.)
    *   `make statusgo-android`:
        *   Make result is the file `build/bin/statusgo-android-16.aar`.
    *   `make statusgo-ios`:
        *   Make result is the directory `build/bin/statusgo-ios-9.3-framework`.
        *   The console output contains a number of what appear to be asserts like this:
            ```
            ldid.cpp(602): _assert(): Swap(mach_header_->filetype) == MH_EXECUTE || Swap(mach_header_->filetype) == MH_DYLIB || Swap(mach_header_->filetype) == MH_BUNDLE
            ```
            This doesn't seem to prevent the build from succeeding.
    *   `make statusgo-ios-simulator`:
        *   Result looks similar to `make statusgo-ios`, but build process is different.
            Specifically, it uses a different `xgo` docker image.

