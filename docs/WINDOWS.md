# Virtual FIDO for Windows x64 (beta)

Extract the entire ZIP to a writable folder. Keep the `usbip` folder beside
`virtual-fido.exe`. Go is not required.

## Driver prerequisite

This demo requires the USB/IP VHCI driver. The bundled legacy drivers are test
signed; they are not a production driver installer. Driver installation requires
administrator rights and may require test-signing configuration and a reboot.
Review `usbip/README.md` for the original installation requirements and procedure.
Use a test machine or VM; this release does not change Windows security settings
or install drivers automatically. A test signing private key is not distributed.

Once the driver prerequisites are satisfied, run PowerShell as administrator:

```powershell
cd .\usbip
.\usbip.exe install
cd ..
```

## Run

From PowerShell in the extracted folder:

```powershell
.\virtual-fido.exe --help
.\virtual-fido.exe --vault .\vault.json --passphrase 'choose-your-own-passphrase' start
```

Keep the terminal open and approve registration or login requests by entering `y`.
To list saved credentials:

```powershell
.\virtual-fido.exe --vault .\vault.json --passphrase 'choose-your-own-passphrase' list
```

The vault path is relative to your terminal's current folder. Back up your vault
and remember its passphrase. Command-line passphrases can appear in shell history.
This is the upstream experimental demo, with demo attestation and key settings;
use it for testing, not as your only means of accessing important accounts.

## Integrity and licenses

Compare `Get-FileHash .\virtual-fido-windows-amd64.zip -Algorithm SHA256` with
`SHA256SUMS.txt`. The application is not Authenticode signed. Virtual FIDO's MIT
license is in `LICENSE`; the USB/IP license and original instructions are in
`usbip/LICENSE` and `usbip/README.md`.
