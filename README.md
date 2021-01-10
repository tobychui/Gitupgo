# GitUpGo

A git based update utile written in Go

## Build

```
cd Gitupgo
go build
```

## Usage

### Setup the Config file

The config file is a simple json file that allows you to set the basic git clone / pull operation of the update utility. Here is an example

```
{
    "gitrepo":"https://github.com/tobychui/arozos",
    "pre-script": "pre.bat",
    "post-script":"post.bat",
    "folder":"./arozos",
    "interval":259200
}
```



You can also change the pre and post script to the file you want. The interval is update interval in seconds (Leave empty for update every 24 hours) and the folder is the target folder for git to clone / pull to.

### Startup Flags

You can also type ./gitupgo -h for showing this message

```
  -b    Enable update on application startup (default true)
  -config string
        The system configuration file (default "./config.json")
```



Yup, basically there is nothing more on in this very simple project.



## License

MIT License

Copyright 2021 tobychui

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the  "Software"), to deal in the Software without restriction, including  without limitation the rights to use, copy, modify, merge, publish,  distribute, sublicense, and/or sell copies of the Software, and to  permit persons to whom the Software is furnished to do so, subject to  the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,  EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF  MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.  IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY  CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT,  TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE  SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.