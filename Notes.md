### ADC Notes

#### Observations

-   Test accounts: The tests use several accounts (listed below). I know that at least one of them 
    is used on the Ropsten network; I'm not sure whether other networks are used in the tests. These
    accounts need to have sufficient funding. (At least the first one needs >= 100 Ropsten ETH.)
    
    The accounts are (see also 
    [static/config/test-data.json](https://github.com/AlgebraixData/status-go/blob/develop/static/config/test-data.json)):
    -   `0xadaf150b905cf5e6a778e553e15a139b6618bbb7`: See 
        [static/keys/test-account1.pk](https://github.com/AlgebraixData/status-go/blob/develop/static/keys/test-account1.pk).
    -   `0x65c01586aa0ce152835c788ace665e91ab3527b8`: See 
        [static/keys/test-account2.pk](https://github.com/AlgebraixData/status-go/blob/develop/static/keys/test-account2.pk).
    
-   Build: 
    -   Summary:
        -   The targets that start with `statusgo` in [Makefile](Makefile) contain the build commands.
        -   They run in an environment set up by [build/env.sh](build/env.sh). 
        -   They run `xgo` cross-compilation commands (or a `go build` command for the local build).
        -   The local build (the `statusd` executable) is build from the code in cmd/statusd/.
        -   The `xgo` cross-compilation commands use Docker images derived from the xgo image. They
            also use the complex build configuration script [xgo/base/build.sh](xgo/base/build.sh).
    -   [Makefile](Makefile):
        -   Because the builds are run in docker images, typical `make` dependency management 
            doesn't work; all builds are full builds, starting from scratch. (`make` dependencies 
            are used to create dependencies between the different `make` targets.)
        -   Targets:
            -   `make statusgo-android`:
                -   Make result is the file `build/bin/statusgo-android-16.aar`.
            -   `make statusgo-ios`:
                -   Make result is the directory `build/bin/statusgo-ios-9.3-framework`.
                -   The console output contains a number of what appear to be asserts like this:
                    ```
                    ldid.cpp(602): _assert(): Swap(mach_header_->filetype) == MH_EXECUTE || Swap(mach_header_->filetype) == MH_DYLIB || Swap(mach_header_->filetype) == MH_BUNDLE
                    ```
                    This doesn't seem to prevent the build from succeeding.
            -   `make statusgo-ios-simulator`:
                -   Result looks similar to `make statusgo-ios`, but uses a different `xgo` docker 
                    image (see below).
    -   Go cross-compiler [karalabe/xgo](https://github.com/karalabe/xgo): Uses two Docker
        images (the code for these images is in the directory "xgo"):
        -   The Docker image 
            [algebraixendurance/status-xgo](https://hub.docker.com/r/algebraixendurance/status-xgo/)
            is used for the normal builds. It is based on `xgo-1.7.1` and adds `build.sh` as custom 
            build script. This script sets up rather complex build environments and performs the 
            actual build. 
        -   The Docker image
            [algebraixendurance/status-xgo-ios-simulator](https://hub.docker.com/r/algebraixendurance/status-xgo-ios-simulator/)
            is used for the iOS simulator build. It is derived from the previous image by what 
            appears to be replacing of the iOS 9.3 SDK (from the normal image) with the 
            iPhoneSimulator 9.3 SDK.
    -   The build runs in a directory structure under build/_workspace that is set up by 
        [build/env.sh](build/env.sh).
    -   A central piece is the file [xgo/base/build.sh](xgo/base/build.sh). It seems to configure
        `xgo` for the different targets.

-   Comments to individual files:
    -   [Dockerfile](Dockerfile): References the branch `feature/statusd-replaces-geth-on-cluster` 
        in the fork [farazdagi/status-go](https://github.com/farazdagi/status-go). The branch 
        doesn't seem to exist. The file doesn't work as-is.
    -   [package.json](package.json): This file seems to be an `npm` package specification and looks 
        like a left-over from some time ago that isn't functional anymore.

#### TODO

-   Understand **what the different OS builds need** and make sure it's deployed to the Github 
    pages.  
    **Ben et al are working on this.**

-   **Create an organization on Docker hub** (rather than continue to use Endurance's personal 
    account), move the two Docker images that are used in the build to the organization and set up 
    the integration with Github to automatically build them. Or decide what else we want to do with 
    our public Docker images.
    
-   Possibly **replace their test accounts** (Ropsten only?) with our own accounts.

-   **Build/execution environment**:
    -   Fully understand their build/execution environment (in build/_workspace; see 
        [build/env.sh](build/env.sh) and [xgo/base/build.sh](xgo/base/build.sh)). 
    -   Compare the committed 3rd party packages with their originals so that we know whether they 
        are original (and which version) or modified. 
    -   Possibly change the build/execution environment so that Go 3rd party dependency management 
        works as intended. (It doesn't seem to work with the current setup.)

-   Create a **procedure for pulling updates from the upstream repository**. Some initial thoughts 
    (very likely not complete):
    -   Pull the upstream code changes into a branch and closely review. Make it a PR in our repo.
    -   Pull any changes to the two gists.
    -   Pull any changes to the two Docker images. (After we have the automatic Docker build 
        enabled, this would be a verification step only.)
    -   Propagate any changes to the Travis CI configuration to ours.
    -   Handle any changes that would require structural changes in our setup.
     
##### Done

-   When building the project, a number of Go references  point to their own repository.  
    **Conclusion**: These now point to our forked repo.

-   Review all dependencies and decide which ones we may want to copy/fork.  
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

-   Find all references of `farazdagi` and copy the referenced objects (S3 files, gists,
    Docker images, ...) or at least make sure we reference something that's guaranteed
    to be persisting (Docker images with tag?).  
    **Note**: With the proposed changes below, we now have several places that we need to
    check for upstream updates when we update the repository.

    -   The [Dockerfile](Dockerfile) creates an image with a `statusd` executable. It uses the
        branch "feature/statusd-replaces-geth-on-cluster" of a
        [fork of status-go](https://github.com/farazdagi/status-go). This branch doesn't exist
        in this fork (and neither in the original). It seems there isn't much we can do at this
        point; the Dockerfile doesn't seem to work as-is. We can clone the `farazdagi/status-go`
        repository, but it wouldn't help us to get this going.  
        **Conclusion**: Do nothing.
    -   The build uses two Docker images: [farazdagi/xgo](https://hub.docker.com/r/farazdagi/xgo/)
        and [farazdagi/xgo-ios-simulator](https://hub.docker.com/r/farazdagi/xgo-ios-simulator/).  
        **Conclusion**: We submit these images as they are to the Docker hub with our own
        path/tag and then reference these images. The new paths are
        [algebraixendurance/status-xgo](https://hub.docker.com/r/algebraixendurance/status-xgo/)
        and
        [algebraixendurance/status-xgo-ios-simulator](https://hub.docker.com/r/algebraixendurance/status-xgo-ios-simulator/).
    -   The iOS simulator build uses a copy of the iPhone simulator SDK that they have made 
        available through a
        [public S3 location](https://s3.amazonaws.com/farazdagi/status-im/iPhoneSimulator9.3.sdk.tar.gz). 
        I'm not sure this it is legal to publish a copy of the SDK. This is something that we can 
        address if we should have to.  
        **Conclusion**: Do nothing, but keep a private copy of the SDK at
        [s3://alice-status/iPhoneSimulator9.3.sdk.tar.gz](https://s3.amazonaws.com/alice-status/iPhoneSimulator9.3.sdk.tar.gz).
        When/if we need to do something here, we probably shouldn't make the SDK publicly available
        without checking the license first.
    -   Several mentions in the [package.json](package.json) file. This file seems to be an 
        `npm` package specification and looks like a left-over from some time ago. It also references 
        the [farazdagi/status-go](https://github.com/farazdagi/status-go) repository even though it 
        is in the [status-im/status-go](https://github.com/status-im/status-go) repository. To me it 
        seems this file is not used anymore and the references in it can be ignored.  
        **Conclusion**: Do nothing.
    -   The code uses two `CHTRootConfigURL`s in the form of gists (`cht.json`
        [farazdagi/a8d36e2818b3b2b6074d691da63a0c36](https://gist.githubusercontent.com/farazdagi/a8d36e2818b3b2b6074d691da63a0c36)
        and `cht-sandbox.json`
        [farazdagi/3d05d1d3bfa36db7b650c955e23fd7ae](https://gist.githubusercontent.com/farazdagi/3d05d1d3bfa36db7b650c955e23fd7ae)).
        These are marked with "TODO remove this hack, once CHT sync is implemented on LES side".  
        **Conclusion**: Fork the gists and reference our gists. Replacements:
        -   `cht.json`: [GFiedler-ADC/082d351623c7cf80d9b31eebed457087](https://gist.github.com/GFiedler-ADC/082d351623c7cf80d9b31eebed457087) replaces [farazdagi/a8d36e2818b3b2b6074d691da63a0c36](https://gist.githubusercontent.com/farazdagi/a8d36e2818b3b2b6074d691da63a0c36).
        -   `cht-sandbox.json`: [GFiedler-ADC/ccd7b962775bded311d4ae500bfbc27a](https://gist.github.com/GFiedler-ADC/ccd7b962775bded311d4ae500bfbc27a) replaces [farazdagi/3d05d1d3bfa36db7b650c955e23fd7ae](https://gist.githubusercontent.com/farazdagi/3d05d1d3bfa36db7b650c955e23fd7ae).
        
-   Re-enable the test and analyze the failures. (Consider that the original repository also 
    has failures.)  
    **Conclusion**: The initial failures were due to low balance on the test accounts. Adding 
    Ropsten ETH fixed this and made the tests run on my Macbook. Some tests still fail on Travis CI
    because of what appear to be race conditions (see below).

