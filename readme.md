# Smart Home Server API

A lightweight HTTP server for controlling Tuya smart devices (protocol version 3.3) over your local network.

> Some Taschibra devices are under the hood tuya devices 


## Prerequisites

Before getting started, you'll need the following information for each device:
- Device IP address
- Device key
- Device ID

## Setup

1. **Configure your devices**
   ```bash
   cp devices.example.json devices.json
   ```
   Edit `devices.json` with your device information.

2. **Set up scenes**
   ```bash
   cp scenes.example.json scenes.json
   ```
   Customize `scenes.json` to define your automation scenarios.

## Running the Server

Start the server:
```bash
go mod download
go run .
```

The server will start on `localhost:3010` by default.

## Usage

Apply a scene using the REST API:
```bash
curl -X POST "http://localhost:3010/api/apply-scene?scene=turn-on"
```

Replace `turn-on` with any scene name defined in your `scenes.json` file.

## API Endpoints

- `POST /api/apply-scene?scene={scene-name}` - Apply a predefined scene

## Goal 
I built this aiming to have simple and lightweight server able to control some led devices
over my LAN using apple shortcuts
