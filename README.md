# flutter_go

An experimental flutter plugin which let you write backend code in Go for your flutter apps.

## How It Works

This plugin acts as a thin layer between the Dart runtime and Golang runtime.
Messages sent from one runtime will be forwarded to the other.

JSON-RPC based rpc is provided so that Dart code could call Golang functions using rpc.

When you wants to expose some native functions to the frontend,
now you could implement those functions in Go instead of writing both java and obj-c code for
Android and iOS.

## Getting Started

Unlike most flutter plugins, you have to include the source code of this repo into your
flutter project in order to add your own gomobile generated module.

Add flutter_go using git submodule:

```
git submodule add https://github.com/adieu/flutter_go.git
```

Update your project dependency, include flutter_go as a local module.

```
dependencies:
  flutter:
    sdk: flutter

  # The following adds the Cupertino Icons font to your application.
  # Use with the CupertinoIcons class for iOS style icons.
  cupertino_icons: ^0.1.0

  flutter_go:
    path: ./flutter_go
```

Put another copy of flutter_go in your GOPATH

```
go get github.com/adieu/flutter_go/go
```

Generate your go modules using gomobile

```
# Generate Plugin.framework for ios platform
cd flutter_go/ios
gomobile bind --target=ios github.com/adieu/flutter_go/go/plugin github.com/adieu/flutter_go/go/channel github.com/adieu/flutter_go/go/channel/types

# Generate plugin.aar for android platform
cd flutter_go/android
mkdir libs
cd libs
gomobile bind --target=android github.com/adieu/flutter_go/go/plugin github.com/adieu/flutter_go/go/channel github.com/adieu/flutter_go/go/channel/types
```

Build and test if it works. If everything goes well, you could call go functions in your flutter app.

Whenever you modified the go code, you have to regenerate the native module with gomobile again.

## How to organize your Go code

Since the flutter frontend interacts with the Golang backend via rpc,
you have to register exported function in the `net/rpc` package.
Your Go code is essentially a RPC server.
You can expose rpc handlers just like you did in a normal http server.
The only difference is that you don't call
`http.Serve` and let the plugin handle request/response for you.

Next you have to create a new go package for flutter_go with these requirements:

- The package name should be plugin
- The package expose a function called Bootstrap
- The package doesn't contain code which gomobile cannot process

We suggest to keep this module minimal. You could include other Go modules
or doing some initialization in the Bootstrap function.

Your business logic should be exposed through the rpc package.

We included a demo plugin module in the go/ folder as a reference.
