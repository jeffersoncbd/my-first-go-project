# GOnotepad

![demonstration](https://raw.githubusercontent.com/jeffersoncbd/my-first-go-project/main/demo.png)

## Use (linux only)

1. clone this repository:
   ```bash
   git clone https://github.com/jeffersoncbd/my-first-go-project.git gonotepad
   ```
2. run build command (you need [go](https://go.dev/dl/) installed):
   ```bash
   cd gonotepad
   go build .
   ```
3. run binary file:
   ```bash
   ./gonotepad
   ```
4. access [interface](http://localhost:8080)

## Observations

This application does not communicate with a server, your annotations are stored in the "notes" folder in the same directory as the executable. It uses the "titles.yml" file to associate each annotation with its ID.
