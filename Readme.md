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
    # install FreeImage
    $ brew install FreeImage

    # install GLFW
    $ brew install GLFW

    # install   glew
    $ brew install  glew

    # Run the tests
    $ rake

    # Run the app
    $ rake run
