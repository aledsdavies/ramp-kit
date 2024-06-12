{ pkgs ? import <nixpkgs> { } }:

pkgs.mkShell {
  buildInputs = [
    pkgs.air
    pkgs.go
    pkgs.git
    pkgs.kratos
    pkgs.templ
  ];

  shellHook = ''
    export GOPATH=$HOME/go
    export PATH=$GOPATH/bin:$PATH:$kratos/bin
    echo "Go version: $(go version)"
    echo "Kratos version: $(kratos version)"
  '';
}

