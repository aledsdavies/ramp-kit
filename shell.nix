{ pkgs ? import <nixpkgs> { } }:

pkgs.mkShell {
  buildInputs = [
    pkgs.air
    pkgs.go
    pkgs.git
    pkgs.templ
    pkgs.bun
  ];

  shellHook = ''
    export GOPATH=$HOME/go
    echo "Go version: $(go version)"
  '';
}

