{
  inputs = {
    # nixpkgs.url = "github:NixOS/nixpkgs/nixpkgs-unstable";
    systems.url = "github:nix-systems/default";
  };

  outputs =
    {
      systems,
      nixpkgs,
      ...
    }@inputs:
    let
      eachSystem = f: nixpkgs.lib.genAttrs (import systems) (system: f nixpkgs.legacyPackages.${system});
    in
    {
      devShells = eachSystem (pkgs: {
        default =
          with pkgs;
          mkShell {
            JAVA_HOME = jdk17.home;
            packages =
              [
                go
                jdk17
                clang
              ]
              ++ (
                if stdenv.isLinux then
                  [
                    vulkan-headers
                    libxkbcommon
                    wayland
                    xorg.libX11
                    xorg.libXcursor
                    xorg.libXfixes
                    libGL
                    pkg-config
                  ]
                else
                  [ ]
              );
          }
          // (
            if stdenv.isLinux then
              {
                LD_LIBRARY_PATH = "${vulkan-loader}/lib";
              }
            else
              { }
          );
      });
    };
}
