Project Slartibartfast
======================

My game development project, now in Go. Currently only runs on OSX.

http://jasonroelofs.com/2013/07/01/third-times-a-charm/

### Installation

    $ git clone https://github.com/jasonroelofs/slartibartfast.git
    $ git submodule init
    $ git submodule update
    $ source ./environment

    ## on OSX with homebrew
    # Install FreeImage
    $ brew install FreeImage

    # Install GLFW3
    $ brew tap homebrew/versions
    $ brew install glfw3

    # Install glew
    $ brew install glew

    # Build and run the tests
    $ rake

    # Run the app
    $ rake run
