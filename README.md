# go-cobra
CLI tool built over go using cobra.

# Basic information
```
$ ./go-cobra 
A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.

Usage:
  go-cobra [command]

Available Commands:
  completion  generate the autocompletion script for the specified shell
  get         Command to get the desired Cat
  help        Help about any command

Flags:
      --config string   config file (default is $HOME/.go-cobra.yaml)
  -h, --help            help for go-cobra
  -t, --toggle          Help message for toggle

Use "go-cobra [command] --help" for more information about a command.
```

# TO download random image of a cute cat
```
$ ./go-cobra get cute
Trying to get the image of a cute cat
Pictue of a cute cat is saved at file cute.png
```

# To download random image of a baby cat
```
$ ./go-cobra get baby
Trying to get the image of a baby cat
Pictue of a baby cat is saved at file baby.png
```
