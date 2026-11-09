Windows x64 beta package for the Virtual FIDO command-line demo.

- Includes virtual-fido.exe and the existing USB/IP helper and VHCI driver files.
- Runs without installing Go; USB/IP helpers resolve beside the executable.
- Includes Windows instructions, licenses, and SHA-256 checksum.
- Built and tested on Windows using GitHub Actions.

Extract the entire ZIP and read README.md before running. The legacy USB/IP
drivers require separate administrator setup and test-signing prerequisites.
No driver installation or Windows security changes are performed automatically.
The application is unsigned. Full USB attachment and browser WebAuthN flows have
not been tested by this release pipeline. This is experimental demo software.
