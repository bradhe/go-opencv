# Go bindings for OpenCV

This is a mirror of the [original OpenCV bindings for
Go](https://code.google.com/p/go-opencv/) project hosted on Google Code. I
wanted to contribute to it and believe that Github's model for
contributing/sharing code is better so I mirrored it here.

## Installing OpenCV from Git

Someone should really make a script out of this.

1. `git clone git://github.com/Itseez/opencv.git`
1. `git checkout 2.4.5`
    1. **NOTE:** You may want a different tag...
1. `cd opencv`
1. `mkdir build && cd build`
1. `cmake ..`
1. `make`
1. `sudo make install`

## Contributing

I can't promise that any contributions here will make it back upstream as I
have never contacted the original authors, but who knows! Additionally, there
doesn't seem to be a strong emphasis on testing in the original repo so I'm
going to keep that up here.

1. Fork it
1. Commit your changes
1. Submit a pull request
