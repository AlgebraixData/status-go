#### ADC Notes

##### TODO

-   [x] When building the project, a number of references (look like `go` dependencies) point to their 
        own repository. These should point into our repo.
-   [ ] Find all references of `farazdagi` and copy the referenced objects (S3 files, gists, Docker images, ...) 
        or at least make sure we reference something that's guaranteed to be persisting (Docker images with tag?). 

##### Observations

*   [Dockerfile](Dockerfile): 
    *   References the branch `feature/statusd-replaces-geth-on-cluster` in the fork 
        [farazdagi/status-go](https://github.com/farazdagi/status-go). The branch doesn't seem to exist.
*   [Makefile](Makefile): 
    *   Uses the docker image [farazdagi/xgo](https://hub.docker.com/r/farazdagi/xgo/) for most build goals.
    *   Seems to have only partial dependency management implemented. Looks like it's mostly a collection of script
        snippets (goals) with dependency management between them.
    *   `make statusgo-android`:
        *   Make result is the file `build/bin/statusgo-android-16.aar`.
    *   `make statusgo-ios`:
        *   Make result is the directory `build/bin/statusgo-ios-9.3-framework`.
        *   The console output contains a number of what appear to be asserts like this:
            ```
            ldid.cpp(602): _assert(): Swap(mach_header_->filetype) == MH_EXECUTE || Swap(mach_header_->filetype) == MH_DYLIB || Swap(mach_header_->filetype) == MH_BUNDLE
            ```
            This doesn't seem to prevent the build from succeeding.

