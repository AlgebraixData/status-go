### ADC Notes

#### TODO

-   [x] When building the project, a number of Go references  point to their own repository.  
        **Conclusion**: These now point to our forked repo.
-   [x] Review all dependencies and decide which ones we may want to copy/fork.  
        **Conclusion**: All 3rd-party Go package references seem to be committed to the "vendor" directory.
        It is not fully clear whether the build uses the online versions or the committed versions,
        but it seems that it does use the committed versions. To make sure we'd have to analyze the
        `xgo` build configurations. See also my 
        [Slack post](https://algebraixdata.slack.com/files/gerhard/F6HBHNDGA/Go_dependencies_in_status-go)
        with the dependencies.
-   [ ] A number of external dependencies are committed to to the `status-go` repository.
        However, the Go files in those dependencies often import the original location. (For
        example, a number of `btcsuite` projects are committed in
        status-go/vendor/github.com/btcsuite, but btcsuite/btcd/chaincfg/chainhash/hashfuncs.go
        contains `import "github.com/btcsuite/fastsha256"`.) I don't know for sure whether this
        `import` statement references the file as indicated by the URL path or the one that's
        committed to the repository. It looks to me as if it referenced the original file on
        Github. If this is correct, cleaning this up would be a significant amount of work.
-   [ ] Find all references of `farazdagi` and copy the referenced objects (S3 files, gists,
        Docker images, ...) or at least make sure we reference something that's guaranteed
        to be persisting (Docker images with tag?).  
        **Note**: With the proposed changes below, we now have several places that we need to check for upstream 
        updates when we update the repository.

    -   [ ] The [Dockerfile](Dockerfile) creates an image with a `statusd` executable. It uses the branch
            "feature/statusd-replaces-geth-on-cluster" of a [fork of status-go](https://github.com/farazdagi/status-go).
            This branch doesn't exist in this fork (and neither in the original). It seems there isn't much we can do at 
            this point; the Dockerfile doesn't seem to work as-is. We can clone the `farazdagi/status-go` repository, 
            but it wouldn't help us to get this going.  
            **Tentative conclusion**: Do nothing. 
    -   [ ] The build uses two Docker images: [farazdagi/xgo](https://hub.docker.com/r/farazdagi/xgo/) and 
            [farazdagi/xgo-ios-simulator](https://hub.docker.com/r/farazdagi/xgo-ios-simulator/).  
            **Tentative conclusion**: We probably can submit these images as they are to the Docker hub with our own 
            path/tag and then reference these images.
    -   [ ] The iOS simulator build uses a copy of the iPhone simulator SDK that they have made available through a 
            [public S3 location](https://s3.amazonaws.com/farazdagi/status-im/iPhoneSimulator9.3.sdk.tar.gz). I'm not
            sure this is legal. This is probably something that we can get when we need it -- if we should need it.  
            **Tentative conclusion**: Do nothing. When/if we need to do something here, we probably shouldn't make the  
            SDK publicly available without checking the license first.
    -   [ ] Several mentions in the [package.json](package.json) file. This file seems to be an `npm` package 
            specification and looks like a left-over from some time ago. It also references the 
            [farazdagi/status-go](https://github.com/farazdagi/status-go) repository even though it is in the 
            [status-im/status-go](https://github.com/status-im/status-go) repository. To me it seems this file is not
            used anymore and the references in it can be ignored.  
            **Tentative conclusion**: Do nothing.
    -   [ ] The code uses two `CHTRootConfigURL`s in the form of gists
            ([farazdagi/a8d36e2818b3b2b6074d691da63a0c36](https://gist.githubusercontent.com/farazdagi/a8d36e2818b3b2b6074d691da63a0c36/raw/)
            and [farazdagi/3d05d1d3bfa36db7b650c955e23fd7ae](https://gist.githubusercontent.com/farazdagi/3d05d1d3bfa36db7b650c955e23fd7ae/raw/)).
            These are marked with "TODO remove this hack, once CHT sync is implemented on LES side".  
            **Tentative conclusion**: Clone the gists and reference our gists.
        
-   [ ] Understand what the different OS builds need and make sure it's deployed to the Github pages.
-   [ ] Re-enable the test and analyze the failures. (Consider that the original repository also has failures.)


#### Observations

*   [Dockerfile](Dockerfile): 
    *   References the branch `feature/statusd-replaces-geth-on-cluster` in the fork 
        [farazdagi/status-go](https://github.com/farazdagi/status-go). The branch doesn't seem to exist.
*   [Makefile](Makefile): 
    *   Uses the docker image [farazdagi/xgo](https://hub.docker.com/r/farazdagi/xgo/) (or a derivation of it)
        for the builds. It is based `xgo-1.7.1` (a cross-compiler for Go) and adds `build.sh` as custom build
        script. This script sets up rather complex build environments and performs the actual build. The iOS
        simulator build is done with another docker image that is derived by what appears to be replacing of
        the iOS 9.3 SDK with the iPhoneSimulator 9.3 SDK. (The code for these docker images is in the directory
        "xgo".)
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

