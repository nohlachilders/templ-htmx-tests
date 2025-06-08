{
    inputs = {
        nixpkgs.url = "github:NixOS/nixpkgs/nixpkgs-unstable";
        systems.url = "github:nix-systems/default";
        flake-utils = {
            url = "github:numtide/flake-utils";
            inputs.systems.follows = "systems";
        };

        gomod2nix = {
            url = "github:nix-community/gomod2nix";
            inputs.nixpkgs.follows = "nixpkgs";
        };
    };

    outputs = { self, nixpkgs, flake-utils, gomod2nix, ... } @ inputs:
        flake-utils.lib.eachDefaultSystem (system: let
            pkgs = import nixpkgs {
                inherit system;
                overlays = [
                    gomod2nix.overlays.default
                ];
            };
        in {
            packages = {
                gomod2nix = inputs.gomod2nix.default;
            };

            devShells.default = pkgs.mkShell {
                packages = with pkgs; [
                    go
                    gopls
                    go-tools
                    delve
                    gomod2nix.packages.${system}.default

                    templ
                    air
                ];
                hardeningDisable = ["all"];
            };
        }
        );
}
